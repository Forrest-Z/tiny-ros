/*
 * File      : rostopic.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include "tiny_ros/ros.h"
#include <stdio.h>
#include <thread>

#define USAGE "\n\nUSAGE:\n" \
              "       rostopic          //ros_ip: 127.0.0.1, no arguments\n" \
              "       rostopic <ros_ip> //ros_ip: tiny-ros ip address\n"

static void threadCb(std::string ros_ip) { 
  while(!tinyros::nh()->initNode(ros_ip)) {
    sleep(1);
  }
  
  std::string topiclist = tinyros::nh()->getTopicList(3000);
  printf(USAGE);
  printf("%s\n", topiclist.c_str());
  tinyros::nh()->exit();
  exit(0);
}

int main(int argc, char **argv) {
  char *ros_ip = (char*)"127.0.0.1";
  if (argc >= 2) {
    ros_ip = (char*)argv[1];
  }
  
  new std::thread(std::bind(threadCb, ros_ip));
  
  while(true) {
    tinyros::nh()->spin();
    sleep(1);
  }
  return 0;
}

