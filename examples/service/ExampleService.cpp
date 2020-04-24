/*
 * File      : ExampleService.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_hello/Test.h"

static void service_cb(const tinyros_hello::Test::Request & req, tinyros_hello::Test::Response & res) {
  res.output = "Hello, tiny-ros ^_^";
}

int main() {
  tinyros::ServiceServer<tinyros_hello::Test::Request, tinyros_hello::Test::Response> server("test_srv", &service_cb);
  tinyros::nh()->advertiseService(server);
  while(true) {
#ifdef WIN32
    Sleep(10*1000);
#else
    sleep(10);
#endif
  }
  
  return 0;
}

