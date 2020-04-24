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
namespace tinyros {
static std::string ip_addr_ = "127.0.0.1";
  
//////////////////////////////////////////
static bool main_loop_init_ = false;
static std::mutex main_loop_mutex_;
static std::thread* tinyros_main_loop_thread_ = NULL;

static void tinyros_main_loop() {
retry:
  tinyros::nh()->initNode(ip_addr_);
  while (tinyros::nh()->ok()) {
    tinyros::nh()->spin();
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
      tinyros_main_loop_thread_ =  new std::thread(std::bind(tinyros_main_loop));
      if (tinyros_main_loop_thread_ != NULL) {
        main_loop_init_ = true;
      }
    }
  }
  return &g_nh;
}
}

