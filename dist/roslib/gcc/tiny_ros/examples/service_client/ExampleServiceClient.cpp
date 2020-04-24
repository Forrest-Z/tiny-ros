/*
 * File      : ExampleServiceClient.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_hello/Test.h"

int main() {
  tinyros::ServiceClient<tinyros_hello::Test::Request, tinyros_hello::Test::Response> client("test_srv");
  tinyros::nh()->serviceClient(client);
  while (true) {
    tinyros_hello::Test::Request req;
    tinyros_hello::Test::Response res;
    req.input = "hello world!";
    if (client.call(req, res)) {
       printf("Service responsed with \"%s\"\n", res.output.c_str());
    } else {
       printf("Service call failed.\n");
    }
#ifdef WIN32
    Sleep(1000);
#else
    sleep(1);
#endif
  }
  
  return 0;
}

