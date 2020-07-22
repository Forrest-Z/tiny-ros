/*
 * File      : node_handle_base.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_NODE_HANDLE_BASE_H_
#define TINYROS_NODE_HANDLE_BASE_H_

#include <stdint.h>
#include <mutex>
#include <memory>
#include "tiny_ros/ros/time.h"
#include "tiny_ros/tinyros_msgs/SyncTime.h"

namespace tinyros {
const int MAX_SUBSCRIBERS = 100;
const int MAX_PUBLISHERS = 100;
const int INPUT_SIZE = 64*1024; // bytes
const int OUTPUT_SIZE = 64*1024; // bytes

const uint8_t MODE_FIRST_FF = 0;
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

class SpinObject {
public:
  uint32_t id;
  uint8_t *message_in;
  SpinObject() { message_in = NULL; }
  ~SpinObject() { if(message_in) free((void*)message_in); }
};

class NodeHandleBase_
{
protected:
  std::string ip_addr_;
  std::string node_name_;

public:
  virtual bool initNode(std::string node_name, std::string ip_addr) { return false; }
  virtual int publish(uint32_t id, const Msg* msg, bool islog = false) { return 0; }
  virtual int spin() { return -1; }
  virtual void exit() {}
  virtual bool ok() { return false; }
  virtual void spin_task(const std::shared_ptr<SpinObject> obj) {}
  virtual void keepalive() {}
  virtual void sync_time(unsigned char* data) {
     tinyros_msgs::SyncTime t;
     t.deserialize(data);
     int64_t now = (int64_t)(Time::now().toMSec());
     std::unique_lock<std::mutex> lock(Time::mutex_);
     int64_t scope = now - Time::time_last_ - t.tick;
     if ((Time::time_start_ == 0) || (scope >= 0 && scope <= SYNC_TIME_SCOPE)) {
        Time::time_dds_ = (int64_t)(t.data.toMSec());
        Time::time_start_ = now;
     }
     Time::time_last_ = now;
  }
};

void init(std::string node_name, std::string ip_addr = "127.0.0.1");
void logdebug(std::string msg);
void loginfo(std::string msg);
void logwarn(std::string msg);
void logerror(std::string msg);
void logfatal(std::string msg);

}
#endif

