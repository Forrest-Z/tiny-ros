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
#include <math.h>
#include <queue>
#include <rtthread.h>
#include "tiny_ros/ros/time.h"
#include "tiny_ros/ros/log.h"
#include "tiny_ros/roslib_msgs/Time.h"
#include "tiny_ros/roslib_msgs/String.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"
#include "tiny_ros/ros/service_server.h"
#include "tiny_ros/ros/service_client.h"
#include "tiny_ros/ros/hardware.h"

namespace tinyros {
using tinyros_msgs::TopicInfo;

class NodeHandle : public NodeHandleBase_
{
public:
  tinyros::string port_name_;
  Hardware hardware_;

  Hardware loghd_;
  bool loghd_keepalive_;
  char loghd_keepalive_stack_[THREAD_LOG_KEEPALIVE_STACK_SIZE];
  struct rt_thread loghd_keepalive_thread_;

  uint8_t message_in[INPUT_SIZE];
  uint8_t message_tmp[INPUT_SIZE];
  uint8_t message_out[OUTPUT_SIZE];

  Publisher * publishers[MAX_PUBLISHERS];
  Subscriber_ * subscribers[MAX_SUBSCRIBERS];

  static NodeHandle* getInstance() {
    static NodeHandle g_nh;
    return &g_nh;
  }
private:
  NodeHandle():
    loghd_keepalive_(false),
    topic_list(""),
    service_list("") {
    rt_mutex_init(&mutex_, "nh", RT_IPC_FLAG_FIFO);
    rt_mutex_init(&srv_id_mutex_, "srv", RT_IPC_FLAG_FIFO);
    rt_mutex_init(&main_loop_mutex_, "ml", RT_IPC_FLAG_FIFO);

    for (unsigned int i = 0; i < MAX_PUBLISHERS; i++)
      publishers[i] = NULL;

    for (unsigned int i = 0; i < MAX_SUBSCRIBERS; i++)
      subscribers[i] = NULL;
  }

public:
  ~NodeHandle() {
    exit();

    rt_mutex_detach(&mutex_);
    rt_mutex_detach(&srv_id_mutex_);
    rt_mutex_detach(&main_loop_mutex_);
  }

  virtual bool initNode(tinyros::string portName = "127.0.0.1") {
    bytes_ = 0;
    index_ = 0;
    topic_ = 0;
    spin_ = true;
    mode_ = MODE_FIRST_FF;
    port_name_ = portName;

    roslib_msgs::String msg;
    if (hardware_.init(port_name_)) {
      msg.data = get_executable_name();
      publish(TopicInfo::ID_SESSION_ID, &msg);
    }

    if (!loghd_keepalive_) {
      rt_err_t result = RT_EOK;
      result = rt_thread_init(&loghd_keepalive_thread_, "loghd", &NodeHandle::keepalive, this, &loghd_keepalive_stack_[0],
          sizeof(loghd_keepalive_stack_), THREAD_LOG_KEEPALIVE_PRIORITY, THREAD_LOG_KEEPALIVE_TICK);
      RT_ASSERT(result == RT_EOK);
      result = rt_thread_startup(&loghd_keepalive_thread_);
      RT_ASSERT(result == RT_EOK);
      loghd_keepalive_ = true;
    }

    return hardware_.connected();
  }

  static void keepalive(void *pthis) {
    uint8_t in[200];
    NodeHandle *nh = (NodeHandle *)pthis;
    while(nh->loghd_keepalive_) {
      if (!nh->loghd_.connected()) {
        if (nh->loghd_.init(nh->port_name_)) {
          roslib_msgs::String msg;
          tinyros::string process = get_executable_name();
          msg.data = process + "_log";
          nh->publish(TopicInfo::ID_SESSION_ID, &msg, true);
        }
        rt_thread_delay(1000);
      } else {
        nh->loghd_.read(in, sizeof(in));
      }
    }
  }

  virtual void exit() {
    spin_ = false;
    loghd_keepalive_ = false;
    rt_thread_detach(&loghd_keepalive_thread_);

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

    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (unsigned int i = 0; i < MAX_PUBLISHERS; i++) {
      if (publishers[i] != NULL)
        publishers[i]->negotiated_ = false;
    }
    for (unsigned int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] != NULL)
        subscribers[i]->negotiated_ = false;
    }
    rt_mutex_release(&mutex_);

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
            roslib_msgs::String msg;
            msg.deserialize(message_in);
            topic_list = msg.data;
            topic_list_recieved = true;
          } else if (topic_ == TopicInfo::ID_ROSSERVICE_REQUEST) {
            roslib_msgs::String msg;
            msg.deserialize(message_in);
            service_list = msg.data;
            service_list_recieved = true;
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
            if ((((int)(topic_-100)) >= 0) && subscribers[(int)(topic_-100)]) {
              subscribers[(int)(topic_-100)]->callback(message_in);
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
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (int i = 0; i < MAX_PUBLISHERS; i++) {
      if (publishers[i] == NULL) { // empty slot
        p.id_ = i + 100 + MAX_SUBSCRIBERS;
        p.nh_ = this;
        publishers[i] = &p;
        rt_mutex_release(&mutex_);
        negotiateTopics(publishers[i]);
        log_debug("Publishers[%d] topic_id: %d, topic_name: %s", i, p.id_, p.topic_.c_str());
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
      if (subscribers[i] == 0) {// empty slot
        s.id_ = i + 100;
        subscribers[i] = &s;
        rt_mutex_release(&mutex_);
        negotiateTopics(subscribers[i]);
        log_debug("Subscribers[%d] topic_id: %d, topic_name: %s", i, s.id_, s.topic_.c_str());
        return true;
      }
    }
    rt_mutex_release(&mutex_);
    return false;
  }

  /* Register a new Service Server */
  template<typename MReq, typename MRes, typename ObjT>
  bool advertiseService(ServiceServer<MReq, MRes, ObjT>& srv) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == NULL) { // empty slot
        subscribers[i] = &srv;
        srv.id_ = i + 100;
        for (int j = 0; j < MAX_PUBLISHERS; j++) {
          if (publishers[j] == NULL) { // empty slot
            publishers[j] = &srv.pub;
            srv.pub.id_ = srv.id_;
            srv.pub.nh_ = this;
            rt_mutex_release(&mutex_);
            negotiateTopics(subscribers[i]);
            negotiateTopics(publishers[j]);
            log_debug("advertiseService Subscribers[%d] topic_id: %d, topic_name: %s | Publishers[%d] topic_id: %d, topic_name: %s", 
                i, srv.id_, srv.topic_.c_str(), j, srv.pub.id_, srv.pub.topic_.c_str());
            return true;
          }
        }
      }
    }
    rt_mutex_release(&mutex_);
    return false;
  }

  /* Register a new Service Client */
  template<typename MReq, typename MRes>
  bool serviceClient(ServiceClient<MReq, MRes>& srv) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
    for (int i = 0; i < MAX_SUBSCRIBERS; i++) {
      if (subscribers[i] == NULL) { // empty slot
        subscribers[i] = &srv;
        srv.id_ = i + 100;
        for (int j = 0; j < MAX_PUBLISHERS; j++) {
          if (publishers[j] == NULL) { // empty slot
            publishers[j] = &srv.pub;
            srv.pub.id_ = srv.id_;
            srv.pub.nh_ = this;
            rt_mutex_release(&mutex_);
            negotiateTopics(subscribers[i]);
            negotiateTopics(publishers[j]);
          }
        }
      }
    }
    rt_mutex_release(&mutex_);
    return false;
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

  virtual int publish(uint32_t id, const Msg * msg, bool islog = false) {
    rt_mutex_take(&mutex_, RT_WAITING_FOREVER);
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
      rt_mutex_release(&mutex_);
      return l;
    } else {
      rt_mutex_release(&mutex_);
      return -2;
    }
  }

private:
  void log(char byte, const char* msg) {
    if (loghd_.connected()) {
      tinyros_msgs::Log l;
      l.level = byte;
      l.msg = msg;
      publish(tinyros_msgs::TopicInfo::ID_LOG, &l, true);
    }
  }

public:
  virtual void logdebug(const char* msg) { log(tinyros_msgs::Log::ROSDEBUG, msg); }
  virtual void loginfo(const char* msg) { log(tinyros_msgs::Log::ROSINFO, msg); }
  virtual void logwarn(const char* msg) { log(tinyros_msgs::Log::ROSWARN, msg); }
  virtual void logerror(const char* msg) { log(tinyros_msgs::Log::ROSERROR, msg); }
  virtual void logfatal(const char* msg) { log(tinyros_msgs::Log::ROSFATAL, msg); }

private:
  bool topic_list_recieved;
  tinyros::string topic_list;

  bool service_list_recieved;
  tinyros::string service_list;

public:
  tinyros::string getTopicList(int timeout = 1000)
  {
    roslib_msgs::String msg;
    topic_list_recieved = false;
    publish(TopicInfo::ID_ROSTOPIC_REQUEST, &msg);
    int64_t to = (int64_t)(tinyros::Time().now().toMSec() + timeout);
    while (!topic_list_recieved)
    {
      int64_t now = (int64_t)tinyros::Time().now().toMSec();
      if (now > to ) {
        rt_kprintf("Failed to get getTopicList: timeout expired.");
        return "";
      }
    }
    return topic_list;
  }

  tinyros::string getServiceList(int timeout = 1000)
  {
    roslib_msgs::String msg;
    service_list_recieved = false;
    publish(TopicInfo::ID_ROSSERVICE_REQUEST, &msg);
    int64_t to = (int64_t)(tinyros::Time().now().toMSec() + timeout);
    while (!service_list_recieved)
    {
      int64_t now = (int64_t)tinyros::Time().now().toMSec();
      if (now > to) {
        rt_kprintf("Failed to get getServiceList: timeout expired.");
        return "";
      }
    }
    return service_list;
  }
};

NodeHandle* nh();

}

#endif

