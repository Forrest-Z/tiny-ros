/*
 * File      : node_handle_udp.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#ifndef TINYROS_NODE_HANDLE_UDP_H_
#define TINYROS_NODE_HANDLE_UDP_H_
#include <map>
#include <sstream>
#include <random>
#include <thread>
#include "tiny_ros/ros/threadpool.h"
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/ros/hardware_linux_udp.h"
#include "tiny_ros/ros/hardware_windows_udp.h"

namespace tinyros {
template<class Hardware,
         int INPUT_SIZE = 65*1024,
         int OUTPUT_SIZE = 65*1024>
class NodeHandleUdp_: public NodeHandleBase_
{
private:
  Hardware hardware_;
  std::mutex mutex_;
  uint8_t message_in[INPUT_SIZE];
  uint8_t message_out[OUTPUT_SIZE];

  bool spin_;
  ThreadPool spin_thread_pool_;
  
  bool negotiate_keepalive_;
  ThreadPool negotiate_thread_pool_;
  
  std::map<uint32_t, Publisher*> publishers_;
  std::map<uint32_t, Subscriber_ *> subscribers_;
  
  int random_char() {
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(0, 255);
    return dis(gen);
  }

  std::string generate_uuid() {
    std::stringstream ss;
    for (int i = 0; i < 16; i++) {
      int rc = random_char();
      std::stringstream hexstream;
      hexstream << std::hex << rc;
      std::string hex = hexstream.str();
      ss << (hex.length() < 2 ? std::string("0") + hex : hex);
    }
    return ss.str();
  }
  
  uint32_t generate_id() {
    uint32_t h = 0;
    std::string uuid = generate_uuid();
    const char *p = uuid.c_str();
    for(; *p; p++) {
      h = 31 * h + *p;
    }
    return h;
  }
  
  virtual void spin_task(const std::shared_ptr<SpinObject> obj) {
    if((obj != nullptr) && obj->message_in){
      int64_t time_start = (int64_t)tinyros::Time().now().toMSec();
      int64_t timeout_time = time_start + 1000;

      subscribers_[obj->id]->callback(obj->message_in );
      
      int64_t time_end = (int64_t)tinyros::Time().now().toMSec();
      if (time_end > timeout_time) {
        log_warn("subscriber topic: %s, time escape: %lld(ms)", subscribers_[obj->id]->topic_.c_str(), (time_end - time_start));
      }
    }
  }

  void negotiateTopics(Publisher * p) {
    tinyros_msgs::TopicInfo ti;
    ti.topic_id = p->id_;
    ti.topic_name = p->topic_;
    ti.message_type = p->msg_->getType();
    ti.md5sum = p->msg_->getMD5();
    ti.buffer_size = OUTPUT_SIZE;
    publish(p->getEndpointType(), &ti);
  }
  
  void negotiateTopics(Subscriber_ * s) {
    tinyros_msgs::TopicInfo ti;
    ti.topic_id = s->id_;
    ti.topic_name = s->topic_;
    ti.message_type = s->getMsgType();
    ti.md5sum = s->getMsgMD5();
    ti.buffer_size = INPUT_SIZE;
    publish(s->getEndpointType(), &ti);
  }

  void negotiateTopics() {
    std::map<uint32_t, Publisher*>::iterator pit;
    std::map<uint32_t, Subscriber_*>::iterator sit;
    std::unique_lock<std::mutex> lock(mutex_);
    for(pit = publishers_.begin(); pit != publishers_.end(); pit++) {
      Publisher* p = pit->second;
      lock.unlock();
      negotiateTopics(p);
      lock.lock();
    }
    for(sit = subscribers_.begin(); sit != subscribers_.end(); sit++) {
      Subscriber_* s = sit->second;
      lock.unlock();
      negotiateTopics(s);
      lock.lock();
    }
    lock.unlock();
  }

  virtual void keepalive() {
    while(negotiate_keepalive_) {
      negotiateTopics();
#ifdef WIN32
      Sleep(1000);
#else
      sleep(1);
#endif
    }
  }

public:
  NodeHandleUdp_()
    : spin_thread_pool_(3)
    , negotiate_thread_pool_(1) {
  }
  
  ~NodeHandleUdp_() {
    exit();
  }
   
  virtual bool initNode(std::string ipaddr = "127.0.0.1") {
    if(!negotiate_keepalive_) {
      negotiate_keepalive_ = true;
      negotiate_thread_pool_.schedule(std::bind(&NodeHandleBase_::keepalive, this));
    }
    
    return hardware_.init(ipaddr);
  }
  
  virtual void exit() {
    spin_ = false;
    spin_thread_pool_.shutdown();
    hardware_.close();
  }
    
  virtual bool ok() {
    return hardware_.connected();
  }
  
   /* Register a new publisher */
  bool advertise(Publisher & p) {
    std::unique_lock<std::mutex> lock(mutex_);
    p.id_ = generate_id();
    p.nh_ = this;
    p.negotiated_ = true;
    publishers_[p.id_] = &p;
    lock.unlock();
    negotiateTopics(&p);
    return true;
  }

  /* Register a new subscriber */
  template<typename SubscriberT>
  bool subscribe(SubscriberT& s) {
    std::unique_lock<std::mutex> lock(mutex_);
    s.id_ = generate_id();
    s.negotiated_ = true;
    subscribers_[s.id_] = &s;
    lock.unlock();
    negotiateTopics(&s);
    return true;
  }

  virtual int spin() {
    spin_ = true;
    while (spin_ && ok()) {
      if (subscribers_.size() <= 0) {
#ifdef WIN32
        Sleep(1000);
#else
        sleep(1);
#endif
        continue;
      }
      
      int32_t rv = hardware_.read(message_in, INPUT_SIZE);
      if (INPUT_SIZE >= rv && rv > 0) {
        uint32_t topic = 0;
        int bytes = 0, index= 0, checksum = 0;
        do {
          index = 0;
          if (message_in[index++] != 0xff) {
            break;
          }

          if (message_in[index++] != 0xb9) {
            break;
          }

          bytes = message_in[index];
          bytes += message_in[index + 1] << 8;
          bytes += message_in[index + 2] << 16;
          bytes += message_in[index + 3] << 24;
          checksum = message_in[index];
          checksum += message_in[index + 1];
          checksum += message_in[index + 2];
          checksum += message_in[index + 3];
          checksum += message_in[index + 4];
          index += 5;
          
          if((checksum % 256) != 255) {
            break;
          }

          topic = message_in[index];
          topic += message_in[index + 1] << 8;
          topic += message_in[index + 2] << 16;
          topic += message_in[index + 3] << 24;
          checksum = message_in[index];
          checksum += message_in[index + 1];
          checksum += message_in[index + 2];
          checksum += message_in[index + 3];
          index += 4;

          if (INPUT_SIZE < (index + bytes + 1)) {
            break;
          }

          if(bytes > 0) {
            for (int32_t i=0; i < bytes + 1; i++) {
              checksum += message_in[index + i];
            }
          } else {
            checksum += message_in[index];
          }

          if ((checksum % 256) == 255) {
            if(subscribers_.count(topic) == 1) {
              std::shared_ptr<SpinObject> obj = std::shared_ptr<SpinObject> (new SpinObject());
              int total_bytes = bytes > 0 ? bytes : 1;
              obj->id = topic;
              obj->message_in = (uint8_t*)calloc(total_bytes, sizeof(uint8_t));
              memcpy(obj->message_in, message_in + index, total_bytes);
              spin_thread_pool_.schedule(std::bind(&NodeHandleBase_::spin_task, this, obj));
            }
          }
        } while(0);
      }
    }

    return true;
  }
  
  virtual int publish(uint32_t id, const Msg * msg, bool islog = false) {
      std::unique_lock<std::mutex> lock(mutex_);
      /* serialize message */
      int l = msg->serialize(message_out + 11);
  
      /* setup the header */
      message_out[0] = 0xff;
      message_out[1] = 0xb9;
      message_out[2] = (uint8_t)((uint32_t)l & 0xFF);
      message_out[3] = (uint8_t)((uint32_t)((l >> 8) & 0xFF));
      message_out[4] = (uint8_t)((uint32_t)((l >> 16) & 0xFF));
      message_out[5] = (uint8_t)((uint32_t)((l >> 24) & 0xFF));
      message_out[6] = 255 - ((message_out[2] + message_out[3] + message_out[4] + message_out[5]) % 256);
      message_out[7] = (uint8_t)((uint32_t)id & 0xFF);
      message_out[8] = (uint8_t)((uint32_t)((id >> 8) & 0xFF));
      message_out[9] = (uint8_t)((uint32_t)((id >> 16) & 0xFF));
      message_out[10] = (uint8_t)((uint32_t)((id >> 24) & 0xFF));
  
      /* calculate checksum */
      int chk = 0;
      for (int i = 7; i < l + 11; i++)
        chk += message_out[i];
      l += 11;
      message_out[l++] = 255 - (chk % 256);
  
      if (l <= OUTPUT_SIZE) {
        l = hardware_.write(message_out, l) ? l : -1;
        return l;
      } else {
        return -2;
      }
    }

};

#ifdef WIN32
typedef NodeHandleUdp_<HardwareWindowsUdp> NodeHandleUdp;
#else
typedef NodeHandleUdp_<HardwareLinuxUdp> NodeHandleUdp;
#endif

NodeHandleUdp* udp();
}

#endif //TINYROS_NODE_HANDLE_UDP_H_

