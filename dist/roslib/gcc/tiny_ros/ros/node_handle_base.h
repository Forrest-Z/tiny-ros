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
}
#endif

