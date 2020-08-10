#include "tiny_ros/ros.h"
#include <stdio.h>
#include <thread>

#define USAGE "\n\nUSAGE:\n" \
              "       tinyrosservice          //ros_ip: 127.0.0.1, no arguments\n" \
              "       tinyrosservice <ros_ip> //ros_ip: tiny-ros ip address\n"

static void threadCb(std::string ros_ip) { 
  while(!tinyros::nh()->initNode("tinyrosservice", ros_ip)) {
    sleep(1);
  }
  
  std::string srvlist = tinyros::nh()->getServiceList(3000);
  printf(USAGE);
  printf("%s\n", srvlist.c_str());
  tinyros::nh()->exit();
  exit(0);
}

int main(int argc, char **argv) {
  char *ros_ip = (char*)"127.0.0.1";
  if (argc >= 2) {
    ros_ip = (char*)argv[1];
  }
  
  new std::thread(std::bind(threadCb, ros_ip));
  
  while (true) {
    tinyros::nh()->spin();
    sleep(1);
  }
  return 0;
}
