/*
 * File      : ExampleSubscriber.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_hello/TinyrosHello.h"

static void subscriber_cb(const tinyros_hello::TinyrosHello& received_msg) {
  printf("%s\n", received_msg.hello.c_str());
}

int main(void) {
  tinyros::Subscriber<tinyros_hello::TinyrosHello> sub("tinyros_hello", subscriber_cb);
  tinyros::nh()->subscribe(sub);
  while(true) {
#ifdef WIN32
    Sleep(10*1000);
#else
    sleep(10);
#endif
  }

  return 0;
}

