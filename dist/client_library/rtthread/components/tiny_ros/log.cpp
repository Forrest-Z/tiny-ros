/*
 * File      : log.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#include <stdarg.h>
#include "tiny_ros/ros.h"
#include "tiny_ros/ros/log.h"

#define TINYROS_LOG_MAX_SIZE 100

namespace tinyros {
tinyros::string get_executable_name() {
  return "rtthread";
}

void mtrace(int level, const char *chfr, ...) {
  static tinyros::string processname = "";
  if (processname.empty()) {
    processname = get_executable_name();
    processname = tinyros::string("[") + processname + tinyros::string("] ");
  }

  char buffer[TINYROS_LOG_MAX_SIZE] = {0};

  if (!processname.empty()) {
    snprintf(buffer, sizeof(buffer) - 1, "%s", processname.c_str());
  }

  va_list ap;
  va_start(ap, chfr);
  vsnprintf(buffer + processname.size(), sizeof(buffer) - (processname.size() + 1), chfr, ap);
  va_end(ap);

  if(level == tinyros_msgs::Log::ROSDEBUG) tinyros::nh()->logdebug(buffer);
  else if(level == tinyros_msgs::Log::ROSINFO) tinyros::nh()->loginfo(buffer);
  else if(level == tinyros_msgs::Log::ROSWARN) tinyros::nh()->logwarn(buffer);
  else if(level == tinyros_msgs::Log::ROSERROR) tinyros::nh()->logerror(buffer);
  else if(level == tinyros_msgs::Log::ROSFATAL) tinyros::nh()->logfatal(buffer);
}
}

