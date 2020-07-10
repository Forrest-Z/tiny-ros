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

class SpinObject {
public:
  uint32_t id;
  uint8_t *message_in;
  SpinObject() { message_in = NULL; }
  ~SpinObject() { if(message_in) free((void*)message_in); }
};

class NodeHandleBase_
{
public:
  virtual bool initNode(std::string portName) { return false; }
  virtual int publish(uint32_t id, const Msg* msg, bool islog = false) { return 0; }
  virtual int spin() { return -1; }
  virtual void exit() {}
  virtual bool ok() { return false; }
  virtual void spin_task(const std::shared_ptr<SpinObject> obj) {}
  virtual void keepalive() {}
};

void init(std::string ip_addr = "127.0.0.1");
}
#endif

