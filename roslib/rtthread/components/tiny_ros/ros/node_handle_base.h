#ifndef TINYROS_NODE_HANDLE_BASE_H_
#define TINYROS_NODE_HANDLE_BASE_H_
#include <stdint.h>
#include <rtthread.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/time.h"
#include "tiny_ros/tinyros_msgs/SyncTime.h"

namespace tinyros {

#define TINYROS_LOG_TOPIC "/tinyrosout"

const int THREAD_MAIN_LOOP_PRIORITY   = 15;
const int THREAD_MAIN_LOOP_TICK       = 40;
const int THREAD_MAIN_LOOP_STACK_SIZE = 4096; // bytes

const int THREAD_MAIN_LOOP_UDP_PRIORITY   = 16;
const int THREAD_MAIN_LOOP_UDP_TICK       = 40;
const int THREAD_MAIN_LOOP_UDP_STACK_SIZE = 4096; // bytes

const int THREAD_LOG_KEEPALIVE_PRIORITY   = 17;
const int THREAD_LOG_KEEPALIVE_TICK       = 10;
const int THREAD_LOG_KEEPALIVE_STACK_SIZE = 4096; // bytes

const int THREAD_NEGOTIATE_KEEPALIVE_PRIORITY   = 18;
const int THREAD_NEGOTIATE_KEEPALIVE_TICK       = 10;
const int THREAD_NEGOTIATE_KEEPALIVE_STACK_SIZE = 4096; // bytes

const int MAX_SUBSCRIBERS = 20;
const int MAX_PUBLISHERS = 20;
const int INPUT_SIZE = 200; // bytes
const int OUTPUT_SIZE = 200; // bytes

const uint8_t MODE_FIRST_FF       = 0;
const uint8_t MODE_PROTOCOL_VER   = 1;
const uint8_t PROTOCOL_VER        = 0xb9;
const uint8_t MODE_SIZE_L         = 2;
const uint8_t MODE_SIZE_L1        = 3;
const uint8_t MODE_SIZE_H         = 4;
const uint8_t MODE_SIZE_H1        = 5;
const uint8_t MODE_SIZE_CHECKSUM  = 6;
const uint8_t MODE_TOPIC_L        = 7;
const uint8_t MODE_TOPIC_L1       = 8;
const uint8_t MODE_TOPIC_H        = 9;
const uint8_t MODE_TOPIC_H1       = 10;
const uint8_t MODE_MESSAGE        = 11;
const uint8_t MODE_MSG_CHECKSUM   = 12;

const int SPIN_OK = 0;
const int SPIN_ERR = -1;

const int SYNC_TIME_SCOPE = 10;  // milliseconds

class NodeHandleBase_
{
public:
  tinyros::string ip_addr_;
  tinyros::string node_name_;
  struct rt_mutex mutex_;
  struct rt_mutex srv_id_mutex_;
  struct rt_mutex main_loop_mutex_;
  struct rt_mutex sync_time_mutex_;
  bool main_loop_init_;

  NodeHandleBase_() {
    main_loop_init_ = false;
  }

  ~NodeHandleBase_() {
  }

  virtual bool initNode(tinyros::string node_name, tinyros::string ip_addr) { return false; }
  virtual int publish(uint32_t id, const Msg* msg, bool islog = false) { return -1; }
  virtual int spin() { return -1; }
  virtual void exit() {}
  virtual bool ok() { return false; }
  virtual void sync_time(unsigned char* data) {
     tinyros_msgs::SyncTime t;
     t.deserialize(data);
     int64_t now = (int64_t)(Time::now().toMSec());
     rt_mutex_take(&sync_time_mutex_, RT_WAITING_FOREVER);
     int64_t scope = now - Time::time_last_ - t.tick;
     if ((Time::time_start_ == 0) || (scope >= 0 && scope <= SYNC_TIME_SCOPE)) {
        Time::time_dds_ = (int64_t)(t.data.toMSec());
        Time::time_start_ = now;
     }
     Time::time_last_ = now;
     rt_mutex_release(&sync_time_mutex_);
  }
};

void init(tinyros::string node_name, tinyros::string ip_addr = "127.0.0.1");
void logdebug(tinyros::string msg);
void loginfo(tinyros::string msg);
void logwarn(tinyros::string msg);
void logerror(tinyros::string msg);
void logfatal(tinyros::string msg);
}

#endif /* TINYROS_NODE_HANDLE_BASE_H_ */
