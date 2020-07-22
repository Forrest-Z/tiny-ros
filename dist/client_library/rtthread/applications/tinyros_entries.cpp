/*
 * File      : tinyros_entries.cpp
 * This file is part of tinyros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-01-22     Pinkie.Fu    initial version
 */
#include "tinyros_entries.h"
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_hello/Test.h"
#include "tiny_ros/tinyros_hello/TinyrosHello.h"

//////////////////////////////////////////////////////////
void tinyros_example_publisher(void* parameter) {
  tinyros::init("RT-Thread", "192.168.8.1");
  
  tinyros::Publisher hello_pub ("tinyros_hello", new tinyros_hello::TinyrosHello());
#if 1
  tinyros::nh()->advertise(hello_pub);
#else
  tinyros::udp()->advertise(hello_pub);
#endif
  while (true) {
    tinyros_hello::TinyrosHello msg;
    msg.hello = "Hello, tiny-ros ^_^ ";
    hello_pub.publish (&msg);
    rt_thread_delay(1000);
  }
}


//////////////////////////////////////////////////////////
static void subscriber_cb(const tinyros_hello::TinyrosHello& received_msg) {
  rt_kprintf("%s\n", received_msg.hello.c_str());
}
extern "C" void tinyros_example_subscriber(void* parameter) {
  tinyros::init("RT-Thread", "192.168.8.1");
  
  tinyros::Subscriber<tinyros_hello::TinyrosHello> sub("tinyros_hello", subscriber_cb);
#if 1
  tinyros::nh()->subscribe(sub);
#else
  tinyros::udp()->subscribe(sub);
#endif
  while(true) {
    rt_thread_delay(10*1000);
  }
}

//////////////////////////////////////////////////////////
static void service_cb(const tinyros_hello::Test::Request & req, tinyros_hello::Test::Response & res) {
  res.output = "Hello, tiny-ros ^_^";
}
void tinyros_example_service(void* parameter) {
  tinyros::init("RT-Thread", "192.168.8.1");
  
  tinyros::ServiceServer<tinyros_hello::Test::Request, tinyros_hello::Test::Response> server("test_srv", &service_cb);
  tinyros::nh()->advertiseService(server);
  while(true) {
    rt_thread_delay(10*1000);
  }
}


//////////////////////////////////////////////////////////
extern "C" void tinyros_example_service_client(void* parameter) {
  tinyros::init("RT-Thread", "192.168.8.1");
  
  tinyros::ServiceClient<tinyros_hello::Test::Request, tinyros_hello::Test::Response> client("test_srv");
  tinyros::nh()->serviceClient(client);
  while (true) {
    tinyros_hello::Test::Request req;
    tinyros_hello::Test::Response res;
    req.input = "hello world!";
    if (client.call(req, res)) {
      rt_kprintf("Service responsed with \"%s\"\n", res.output.c_str());
    } else {
      rt_kprintf("Service call failed.\n");
    }
    rt_thread_delay(1000);
  }
}
//////////////////////////////////////////////////////////
