/*
 * File      : ExamplePublisher.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_hello/TinyrosHello.h"

int main (int argc, char *argv[]) {
  tinyros::Publisher hello_pub ("tinyros_hello", new tinyros_hello::TinyrosHello());
  tinyros::nh()->advertise(hello_pub);
  while (true) {
    tinyros_hello::TinyrosHello msg;
    msg.hello = "Hello, tiny-ros ^_^ ";
    if (hello_pub.negotiated()) {
      hello_pub.publish (&msg);
    }
#ifdef WIN32
    Sleep(1000);
#else
    sleep(1);
#endif
  }
  return 0;
}

