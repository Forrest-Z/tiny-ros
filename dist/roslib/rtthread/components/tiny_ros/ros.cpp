/*
 * File      : ros.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros/node_handle.h"
#include "tiny_ros/ros/node_handle_udp.h"
namespace tinyros {

static tinyros::string ip_addr_ = "127.0.0.1";

//////////////////////////////////////////
static struct rt_thread main_loop_thread_;
static char main_loop_stack_[THREAD_MAIN_LOOP_STACK_SIZE];

static struct rt_thread main_loop_udp_thread_;
static char main_loop_udp_stack_[THREAD_MAIN_LOOP_STACK_SIZE];

static void tinyros_main_loop(void* param) {
  NodeHandleBase_ *pnh = (NodeHandleBase_*) param;
  retry:
  pnh->initNode(ip_addr_);
  while (pnh->ok()) {
    pnh->spin();
    rt_thread_delay(1000);
  }
  rt_thread_delay(1000);
  goto retry;
}

NodeHandle* nh(){
  NodeHandle* nh = NodeHandle::getInstance();
  if (!nh->main_loop_init_){
    rt_mutex_take(&nh->main_loop_mutex_, RT_WAITING_FOREVER);
    if (!nh->main_loop_init_) {
      rt_err_t result = RT_EOK;
      result = rt_thread_init(&main_loop_thread_, "ml", tinyros_main_loop, nh,
          &main_loop_stack_[0], sizeof(main_loop_stack_), THREAD_MAIN_LOOP_PRIORITY, THREAD_MAIN_LOOP_TICK);
      RT_ASSERT(result == RT_EOK);
      result = rt_thread_startup(&main_loop_thread_);
      RT_ASSERT(result == RT_EOK);
      nh->main_loop_init_ = true;
    }
    rt_mutex_release(&nh->main_loop_mutex_);
  }
  return nh;
}

NodeHandleUdp* udp() {
  NodeHandleUdp* nh = NodeHandleUdp::getInstance();
  if (!nh->main_loop_init_){
    rt_mutex_take(&nh->main_loop_mutex_, RT_WAITING_FOREVER);
    if (!nh->main_loop_init_) {
      rt_err_t result = RT_EOK;
      result = rt_thread_init(&main_loop_udp_thread_, "mludp", tinyros_main_loop, nh,
          &main_loop_udp_stack_[0], sizeof(main_loop_udp_stack_), THREAD_MAIN_LOOP_PRIORITY, THREAD_MAIN_LOOP_TICK);
      RT_ASSERT(result == RT_EOK);
      result = rt_thread_startup(&main_loop_udp_thread_);
      RT_ASSERT(result == RT_EOK);
      nh->main_loop_init_ = true;
    }
    rt_mutex_release(&nh->main_loop_mutex_);
  }
  return nh;
}
}

