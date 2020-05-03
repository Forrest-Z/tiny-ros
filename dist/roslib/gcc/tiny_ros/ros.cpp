/*
 * File      : ros.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#include <thread>
#include "tiny_ros/ros/node_handle.h"
#include "tiny_ros/ros/node_handle_udp.h"

namespace tinyros {
static std::string ip_addr_ = "127.0.0.1";
  
//////////////////////////////////////////
static bool main_loop_init_ = false;
static std::mutex main_loop_mutex_;

static bool main_loop_udp_init_ = false;
static std::mutex main_loop_udp_mutex_;

static void tinyros_main_loop(NodeHandleBase_ * p) {
  NodeHandleBase_ *pnh = p;
retry:
  pnh->initNode(ip_addr_);
  while (pnh->ok()) {
    pnh->spin();
#ifdef WIN32
    Sleep(1000);
  } 
  Sleep(1000);
#else
    sleep(1);
  } 
  sleep(1);
#endif
  goto retry;
}

NodeHandle* nh(){
  static NodeHandle g_nh;
  if (!main_loop_init_){
    std::unique_lock<std::mutex> lock(main_loop_mutex_);
    if (!main_loop_init_) {
      main_loop_init_ = true;
      std::thread tid(std::bind(tinyros_main_loop, &g_nh));
      tid.detach();
    }
  }
  return &g_nh;
}

NodeHandleUdp* udp() {
  static NodeHandleUdp g_nh;
  if (!main_loop_udp_init_){
    std::unique_lock<std::mutex> lock(main_loop_udp_mutex_);
    if (!main_loop_udp_init_) {
      main_loop_udp_init_ = true;
      std::thread tid(std::bind(tinyros_main_loop, &g_nh));
      tid.detach();
    }
  }
  return &g_nh;
}

}

