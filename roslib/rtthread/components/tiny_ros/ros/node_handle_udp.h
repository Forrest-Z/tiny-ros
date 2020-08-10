#ifndef TINYROS_NODE_HANDLE_UDP_H_
#define TINYROS_NODE_HANDLE_UDP_H_
#include <map>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/ros/hardware_udp.h"

namespace tinyros {
class NodeHandleUdp: public NodeHandleBase_
{
private:
  HardwareUdp hardware_;
  uint8_t message_in[INPUT_SIZE];
  uint8_t message_out[OUTPUT_SIZE];

  bool spin_;

  Publisher * publishers[MAX_PUBLISHERS];
  Subscriber_ * subscribers[MAX_SUBSCRIBERS];

  uint32_t generate_id() {
    uint32_t h = 0;
    for (int i = 0; i < 64; i++) {
      srand(rt_tick_get());
      unsigned char rc = rand();
      h = 31 * h + rc;
    }
    return h;
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

  static void keepalive(void *pthis) {
    NodeHandleUdp *nh = (NodeHandleUdp *)pthis;
    while(nh->negotiate_keepalive_) {
      nh->negotiateTopics();
      rt_thread_delay(1000);
    }
  }

  NodeHandleUdp(): negotiate_keepalive_(false) {
    rt_mutex_init(&mutex_, "nhudp", RT_IPC_FLAG_FIFO);
    rt_mutex_init(&srv_id_mutex_, "srvudp", RT_IPC_FLAG_FIFO);
    rt_mutex_init(&main_loop_mutex_, "mludp", RT_IPC_FLAG_FIFO);

    for (unsigned int i = 0; i < MAX_PUBLISHERS; i++)
      publishers[i] = NULL;

    for (unsigned int i = 0; i < MAX_SUBSCRIBERS; i++)
      subscribers[i] = NULL;
  }

public:
  bool negotiate_keepalive_;
  char negotiate_keepalive_stack_[THREAD_NEGOTIATE_KEEPALIVE_STACK_SIZE];
  struct rt_thread negotiate_keepalive_thread_;

  static NodeHandleUdp* getInstance() {
    static NodeHandleUdp g_nh;
    return &g_nh;
  }

  ~NodeHandleUdp() {
    exit();

    rt_mutex_detach(&mutex_);
    rt_mutex_detach(&srv_id_mutex_);
    rt_mutex_detach(&main_loop_mutex_);
  }

  void negotiateTopics() {
    int i;
    for (i = 0; i < MAX_PUBLISHERS; i++) {
      rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
      if (publishers[i] != NULL) { // non-empty slot
        rt_mutex_release(&mutex_);
        negotiateTopics(publishers[i]);
      } else {
        rt_mutex_release(&mutex_);
      }
    }
    for (i = 0; i < MAX_SUBSCRIBERS; i++) {
      rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
      if (subscribers[i] != NULL) { // non-empty slot
        rt_mutex_release(&mutex_);
        negotiateTopics(subscribers[i]);
      } else {
        rt_mutex_release(&mutex_);
      }
    }
  }

  virtual bool initNode(tinyros::string node_name, tinyros::string ip_addr) {
    ip_addr_ = ip_addr;
    node_name_ = node_name;
    if(!negotiate_keepalive_) {
      rt_err_t result = rt_thread_init(&negotiate_keepalive_thread_, "nego", &NodeHandleUdp::keepalive, this,
          &negotiate_keepalive_stack_[0], sizeof(negotiate_keepalive_stack_),
          THREAD_NEGOTIATE_KEEPALIVE_PRIORITY, THREAD_NEGOTIATE_KEEPALIVE_TICK);
      RT_ASSERT(result == RT_EOK);
      result = rt_thread_startup(&negotiate_keepalive_thread_);
      RT_ASSERT(result == RT_EOK);
      negotiate_keepalive_ = true;
    }

    return hardware_.init(ip_addr_);
  }

  virtual void exit() {
    spin_ = false;
    negotiate_keepalive_ = false;
    rt_thread_detach(&negotiate_keepalive_thread_);
    hardware_.close();
  }

  virtual bool ok() {
    return hardware_.connected();
  }

  /* Register a new publisher */
  bool advertise(Publisher & p) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (int i = 0; i < MAX_PUBLISHERS; i++) {
      if (publishers[i] == NULL) { // empty slot
        p.id_ = generate_id();
        p.nh_ = this;
        publishers[i] = &p;
        rt_mutex_release(&mutex_);
        negotiateTopics(publishers[i]);
        return true;
      }
    }
    rt_mutex_release(&mutex_);
    return false;
  }

  /* Register a new subscriber */
  template<typename SubscriberT>
  bool subscribe(SubscriberT& s) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == NULL) {// empty slot
        s.id_ = generate_id();
        subscribers[i] = &s;
        rt_mutex_release(&mutex_);
        negotiateTopics(subscribers[i]);
        return true;
      }
    }
    rt_mutex_release(&mutex_);
    return false;
  }

  virtual int spin() {
    spin_ = true;
    while (spin_ && ok()) {
      if (subscribers[0] == NULL) {
        rt_thread_delay(1000);
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
            for (int32_t j = 0; j < MAX_SUBSCRIBERS; j++) {
              if (subscribers[j] != NULL && subscribers[j]->id_ == topic) {
                subscribers[j]->callback(message_in + index);
                break;
              }
            }
          }
        } while(0);
      }
    }

    return true;
  }

  virtual int publish(uint32_t id, const Msg * msg, bool islog = false) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
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
      rt_mutex_release(&mutex_);
      return l;
    } else {
      rt_mutex_release(&mutex_);
      return -2;
    }
  }
};

NodeHandleUdp* udp();
}

#endif //TINYROS_NODE_HANDLE_UDP_H_
