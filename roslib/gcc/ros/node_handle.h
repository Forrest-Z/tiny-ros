/*
 * File      : node_handle.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#ifndef TINYROS_NODE_HANDLE_H_
#define TINYROS_NODE_HANDLE_H_

#include <stdint.h>
#include <mutex>
#include <memory>
#include "tiny_ros/ros/log.h"
#include "tiny_ros/ros/threadpool.h"
#include "tiny_ros/std_msgs/String.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/ros/hardware_tcp.h"
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"
#include "tiny_ros/ros/service_server.h"
#include "tiny_ros/ros/service_client.h"

#define TINYROS_LOG_TOPIC "tinyros_log_11315"

namespace tinyros
{
using tinyros_msgs::TopicInfo;

class NodeHandle : public NodeHandleBase_
{
protected:
  HardwareTcp hardware_;
  
  HardwareTcp loghd_;
  bool loghd_keepalive_;
  ThreadPool loghd_thread_pool_;
  
  ThreadPool spin_thread_pool_;
  ThreadPool spin_log_thread_pool_;
  ThreadPool spin_srv_thread_pool_;
  
  std::mutex mutex_;

  uint8_t message_in[INPUT_SIZE];
  uint8_t message_tmp[INPUT_SIZE];
  uint8_t message_out[OUTPUT_SIZE];

  Publisher * publishers[MAX_PUBLISHERS];
  Subscriber_ * subscribers[MAX_SUBSCRIBERS];

private:
  
  virtual void spin_task(const std::shared_ptr<SpinObject> obj) {
      if((obj != nullptr) && subscribers[obj->id] && obj->message_in){
        int64_t time_start = (int64_t)tinyros::Time().now().toMSec();
        int64_t timeout_time = time_start + 1000;
        
        subscribers[obj->id]->callback( obj->message_in );
        
        int64_t time_end = (int64_t)tinyros::Time().now().toMSec();
        if (time_end > timeout_time) {
          log_warn("subscriber topic: %s, time escape: %lld(ms)", subscribers[obj->id]->topic_.c_str(), (time_end - time_start));
        }
      }
  }
  
  virtual void keepalive() {
    uint8_t in[200];
    while(loghd_keepalive_) {
      if (!loghd_.connected()) {
        if (loghd_.init(ip_addr_)) {
          std_msgs::String msg;
          msg.data = node_name_ + "_log";
          publish(TopicInfo::ID_SESSION_ID, &msg, true);
        }
#ifdef WIN32
        Sleep(1000);
#else
        sleep(1);
#endif
      } else {
        loghd_.read(in, sizeof(in));
      }
    }
  }
  
public:
  NodeHandle()
    : loghd_keepalive_(false)
    , loghd_thread_pool_(1)
    , spin_thread_pool_(3)
    , spin_log_thread_pool_(1)
    , spin_srv_thread_pool_(3)
    , topic_list("")
    , service_list("") {

    for (unsigned int i = 0; i < MAX_PUBLISHERS; i++)
      publishers[i] = NULL;

    for (unsigned int i = 0; i < MAX_SUBSCRIBERS; i++)
      subscribers[i] = NULL;
  }

  ~NodeHandle() {
    exit();
  }

  /* Start a named port, which may be network server IP, initialize buffers */
  virtual bool initNode(std::string node_name, std::string ip_addr) {
    bytes_ = 0;
    index_ = 0;
    topic_ = 0;
    spin_ = true;
    mode_ = MODE_FIRST_FF;
    ip_addr_ = ip_addr;
    node_name_ = node_name;

    std_msgs::String msg;
    if (hardware_.init(ip_addr_)) {
      msg.data = node_name_;
      publish(TopicInfo::ID_SESSION_ID, &msg);
    }

    if(!loghd_keepalive_) {
      loghd_keepalive_ = true;
      loghd_thread_pool_.schedule(std::bind(&NodeHandleBase_::keepalive, this));
    }
    
    return hardware_.connected();
  }

  virtual void exit() {
    spin_ = false;
    loghd_keepalive_ = false;
    spin_thread_pool_.shutdown();
    spin_log_thread_pool_.shutdown();
    spin_srv_thread_pool_.shutdown();
    loghd_thread_pool_.shutdown();
    
    loghd_.close();
    hardware_.close();
  }

protected:
  //State machine variables for spin
  uint32_t topic_;
  int mode_;
  int bytes_;
  int index_;
  int checksum_;
  bool spin_;
  int total_bytes_;

public:
  virtual int spin() {
    int i, rv, len = 1;
    
    mode_ = MODE_FIRST_FF;

    {
      std::unique_lock<std::mutex> lock(mutex_);
      for (unsigned int i = 0; i < MAX_PUBLISHERS; i++) {
        if (publishers[i] != NULL)
          publishers[i]->negotiated_ = false;
      }
      for (unsigned int i = 0; i < MAX_SUBSCRIBERS; i++) {
        if (subscribers[i] != NULL)
          subscribers[i]->negotiated_ = false;
      }
    }
    
    if (!hardware_.connected()) {
        return SPIN_ERR;
    } else {
        negotiateTopics();
    }

    /* while available buffer, read data */
    while (spin_ && hardware_.connected()) {
      if (len > INPUT_SIZE) {
        log_error("Input overflow(%d>%d)", len, INPUT_SIZE);
        len = 1;
        mode_ = MODE_FIRST_FF;
        continue;
      }
      
      rv = hardware_.read(message_tmp, len);
      if (rv < 0) {
        mode_ = MODE_FIRST_FF;
        return SPIN_ERR;
      }
      
      
      for (i = 0; i < rv; i++) {
        checksum_ += message_tmp[i];
      }
      
      if (mode_ == MODE_MESSAGE) {
        for (i = 0; i < rv; i++) {
          message_in[index_++] = message_tmp[i];
          bytes_--;
        }

        if (bytes_ == 0) {
          len = 1;
          mode_ = MODE_MSG_CHECKSUM;
        } else {
          len = bytes_;
        }
      } else if (mode_ == MODE_FIRST_FF) {
        if (message_tmp[0] == 0xff) {
          mode_++;
        }
      } else if (mode_ == MODE_PROTOCOL_VER) {
        if (message_tmp[0] == PROTOCOL_VER) {
          mode_++;
        } else {
          mode_ = MODE_FIRST_FF;
        }
      } else if (mode_ == MODE_SIZE_L) {
        bytes_ = message_tmp[0];
        index_ = 0;
        mode_++;
        checksum_ = message_tmp[0];
      } else if (mode_ == MODE_SIZE_L1) {
        bytes_ += message_tmp[0] << 8;
        mode_++;
      } else if (mode_ == MODE_SIZE_H) {
        bytes_ += message_tmp[0] << 16;
        mode_++;
      } else if (mode_ == MODE_SIZE_H1) {
        bytes_ += message_tmp[0] << 24;
        total_bytes_ = bytes_ > 0 ? bytes_ : 1;
        mode_++;
      } else if (mode_ == MODE_SIZE_CHECKSUM) {
        if ((checksum_ % 256) == 255)
          mode_++;
        else
          mode_ = MODE_FIRST_FF;
      } else if (mode_ == MODE_TOPIC_L) {
        topic_ = message_tmp[0];
        mode_++;
        checksum_ = message_tmp[0];
      } else if (mode_ == MODE_TOPIC_L1) {
        topic_ += message_tmp[0] << 8;
        mode_++;
      } else if (mode_ == MODE_TOPIC_H) {
        topic_ += message_tmp[0] << 16;
        mode_++;
      } else if (mode_ == MODE_TOPIC_H1) {
        topic_ += message_tmp[0] << 24;
        mode_ = MODE_MESSAGE;
        if (bytes_ == 0)
          mode_ = MODE_MSG_CHECKSUM;
        else
          len = bytes_;
      } else if (mode_ == MODE_MSG_CHECKSUM) {
        mode_ = MODE_FIRST_FF;
        if ((checksum_ % 256) == 255) {
          if (topic_ == TopicInfo::ID_PUBLISHER) {
            negotiateTopics();
          } else if (topic_ == TopicInfo::ID_ROSTOPIC_REQUEST) {
            std_msgs::String msg;
            msg.deserialize(message_in);
            topic_list = msg.data;
            topic_list_recieved = true;
          } else if (topic_ == TopicInfo::ID_ROSSERVICE_REQUEST) {
            std_msgs::String msg;
            msg.deserialize(message_in);
            service_list = msg.data;
            service_list_recieved = true;
          } else if (topic_ == TopicInfo::ID_TIME) {
            sync_time(message_in);
          } else if (topic_ == TopicInfo::ID_NEGOTIATED) {
            tinyros_msgs::TopicInfo ti;
            ti.deserialize(message_in);
            for (int i = 0; i < MAX_PUBLISHERS; i++) {
              if (publishers[i] != NULL && publishers[i]->id_ == ti.topic_id) {
                publishers[i]->negotiated_ = ti.negotiated;
              }
            }
            
            for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
              if (subscribers[i] != NULL && subscribers[i]->id_ == ti.topic_id) {
                subscribers[i]->negotiated_ = ti.negotiated;
              }
            }
          } else {
            if ((topic_-100) >= 0) {
              if(subscribers[topic_-100]) {
                std::shared_ptr<SpinObject> obj = std::shared_ptr<SpinObject> (new SpinObject());
                obj->id = topic_-100;
                obj->message_in = (uint8_t*)calloc(total_bytes_, sizeof(uint8_t));
                memcpy(obj->message_in, message_in, total_bytes_);
                if (subscribers[topic_-100]->topic_ == TINYROS_LOG_TOPIC) {
                  spin_log_thread_pool_.schedule(std::bind(&NodeHandleBase_::spin_task, this, obj));
                } else {
                  if (subscribers[topic_-100]->srv_flag_) {
                    spin_srv_thread_pool_.schedule(std::bind(&NodeHandleBase_::spin_task, this, obj));
                  } else {
                    spin_thread_pool_.schedule(std::bind(&NodeHandleBase_::spin_task, this, obj));
                  }
                }
              }
            }
          }
        }
      }
    }

    return SPIN_OK;
  }


  /* Are we connected to the PC? */
  virtual bool ok() {
    return hardware_.connected();
  }

  /* Register a new publisher */
  bool advertise(Publisher & p) {
    std::unique_lock<std::mutex> lock(mutex_);
    for (int i = 0; i < MAX_PUBLISHERS; i++) {
      if (publishers[i] == NULL) { // empty slot
        p.id_ = i + 100 + MAX_SUBSCRIBERS;
        p.nh_ = this;
        publishers[i] = &p;
        lock.unlock();
        negotiateTopics(publishers[i]);
        log_debug("Publishers[%d] topic_id: %d, topic_name: %s", i, p.id_, p.topic_.c_str());
        return true;
      }
    }
    return false;
  }

  /* Register a new subscriber */
  template<typename SubscriberT>
  bool subscribe(SubscriberT& s) {
    char buffer[512];
    std::unique_lock<std::mutex> lock(mutex_);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == 0) {// empty slot
        s.id_ = i + 100;
        subscribers[i] = &s;
        lock.unlock();
        negotiateTopics(subscribers[i]);
        log_debug("Subscribers[%d] topic_id: %d, topic_name: %s", i, s.id_, s.topic_.c_str());
        return true;
      }
    }
    return false;
  }

  /* Register a new Service Server */
  template<typename MReq, typename MRes, typename ObjT>
  bool advertiseService(ServiceServer<MReq, MRes, ObjT>& srv) {
    char buffer[512];
    std::unique_lock<std::mutex> lock(mutex_);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == NULL) { // empty slot
        subscribers[i] = &srv;
        srv.id_ = i + 100;
        for (int j = 0; j < MAX_PUBLISHERS; j++) {
          if (publishers[j] == NULL) { // empty slot
            publishers[j] = &srv.pub;
            srv.pub.id_ = srv.id_;
            srv.pub.nh_ = this;
            lock.unlock();
            negotiateTopics(subscribers[i]);
            negotiateTopics(publishers[j]);
            log_debug("advertiseService Subscribers[%d] topic_id: %d, topic_name: %s | Publishers[%d] topic_id: %d, topic_name: %s",
              i, srv.id_, srv.topic_.c_str(), j, srv.pub.id_, srv.pub.topic_.c_str());
            return true;
          }
        }
      }
    }
    return false;
  }

  /* Register a new Service Client */
  template<typename MReq, typename MRes>
  bool serviceClient(ServiceClient<MReq, MRes>& srv) {
    char buffer[512];
    std::unique_lock<std::mutex> lock(mutex_);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == NULL) { // empty slot
        subscribers[i] = &srv;
        srv.id_ = i + 100;
        for (int j = 0; j < MAX_PUBLISHERS; j++) {
          if (publishers[j] == NULL) { // empty slot
            publishers[j] = &srv.pub;
            srv.pub.id_ = srv.id_;
            srv.pub.nh_ = this;
            lock.unlock();
            negotiateTopics(subscribers[i]);
            negotiateTopics(publishers[j]);
            log_debug("serviceClient Subscribers[%d] topic_id: %d, topic_name: %s | Publishers[%d] topic_id: %d, topic_name: %s", 
              i, srv.id_, srv.topic_.c_str(), j, srv.pub.id_, srv.pub.topic_.c_str());
            return true;
          }
        }
      }
    }
    return false;
  }

  void negotiateTopics(Publisher * p) {
    tinyros_msgs::TopicInfo ti;
    ti.topic_id = p->id_;
    ti.topic_name = p->topic_;
    ti.message_type = p->msg_->getType();
    ti.md5sum = p->msg_->getMD5();
    ti.buffer_size = OUTPUT_SIZE;
    ti.node = node_name_;
    publish(p->getEndpointType(), &ti);
  }
  
  void negotiateTopics(Subscriber_ * s) {
    tinyros_msgs::TopicInfo ti;
    ti.topic_id = s->id_;
    ti.topic_name = s->topic_;
    ti.message_type = s->getMsgType();
    ti.md5sum = s->getMsgMD5();
    ti.buffer_size = INPUT_SIZE;
    ti.node = node_name_;
    publish(s->getEndpointType(), &ti);
  }

  void negotiateTopics() {
    int i;
    for (i = 0; i < MAX_PUBLISHERS; i++) {
      std::unique_lock<std::mutex> lock(mutex_);
      if (publishers[i] != NULL) { // non-empty slot
        lock.unlock();
        negotiateTopics(publishers[i]);
      }
    }
    for (i = 0; i < MAX_SUBSCRIBERS; i++) {
      std::unique_lock<std::mutex> lock(mutex_);
      if (subscribers[i] != NULL) { // non-empty slot
        lock.unlock();
        negotiateTopics(subscribers[i]);
      }
    }
  }

  virtual int publish(uint32_t id, const Msg * msg, bool islog = false) {
    std::unique_lock<std::mutex> lock(mutex_);
    /* serialize message */
    int l = msg->serialize(message_out + 11);

    /* setup the header */
    message_out[0] = 0xff;
    message_out[1] = PROTOCOL_VER;
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
      if (!islog) {
        l = hardware_.write(message_out, l) ? l : -1;
      } else {
        l = loghd_.write(message_out, l) ? l : -1;
      }
      lock.unlock();
      return l;
    } else {
      lock.unlock();
      return -2;
    }
  }

  void log(char byte, std::string msg) {
    if (loghd_.connected()) {
      tinyros_msgs::Log l;
      l.level = byte;
      l.msg = std::string("[") + node_name_ + std::string("] ") + msg;
      publish(tinyros_msgs::TopicInfo::ID_LOG, &l, true);
    }
  }

  /*********************************************************************/
  private:
    bool topic_list_recieved;
    std::string topic_list;

    bool service_list_recieved;
    std::string service_list;

  public:
    std::string getTopicList(int timeout = 1000)
    {
      std_msgs::String msg;
      topic_list_recieved = false;
      publish(TopicInfo::ID_ROSTOPIC_REQUEST, &msg);
      int64_t to = (int64_t)(tinyros::Time().now().toMSec() + timeout);
      while (!topic_list_recieved)
      {
        int64_t now = (int64_t)tinyros::Time().now().toMSec();
        if (now > to ) {
          printf("Failed to get getTopicList: timeout expired\n");
          return "";
        }
#ifdef WIN32
        Sleep(100);
#else
        usleep(100*1000);
#endif
      }
      return topic_list;
    }

    std::string getServiceList(int timeout = 1000)
    {
      std_msgs::String msg;
      service_list_recieved = false;
      publish(TopicInfo::ID_ROSSERVICE_REQUEST, &msg);
      int64_t to = (int64_t)(tinyros::Time().now().toMSec() + timeout);
      while (!service_list_recieved)
      {
        int64_t now = (int64_t)tinyros::Time().now().toMSec();
        if (now > to) {
          printf("Failed to get getServiceList: timeout expired\n");
          return "";
        }
#ifdef WIN32
        Sleep(100);
#else
        usleep(100*1000);
#endif
      }
      return service_list;
    }
};

NodeHandle* nh();

}

#endif

