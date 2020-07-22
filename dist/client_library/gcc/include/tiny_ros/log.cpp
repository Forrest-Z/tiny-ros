/*
 * File      : log.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#include <stdarg.h>
#include <limits.h>
#include <stdio.h>
#include <string>
#include "tiny_ros/ros.h"
#include "tiny_ros/ros/log.h"

#define TINYROS_LOG_MAX_SIZE 4096

namespace tinyros {
void mtrace(int level, const char *chfr, ...) {
  char buffer[TINYROS_LOG_MAX_SIZE] = {0};
  va_list ap;
  va_start(ap, chfr);
#ifdef WIN32
  _vsnprintf(buffer, sizeof(buffer) - 1, chfr, ap);
#else
  vsnprintf(buffer, sizeof(buffer) - 1, chfr, ap);
#endif
  va_end(ap);

  if(level == tinyros_msgs::Log::ROSDEBUG) tinyros::logdebug(buffer);
  else if(level == tinyros_msgs::Log::ROSINFO) tinyros::loginfo(buffer);
  else if(level == tinyros_msgs::Log::ROSWARN) tinyros::logwarn(buffer);
  else if(level == tinyros_msgs::Log::ROSERROR) tinyros::logerror(buffer);
  else if(level == tinyros_msgs::Log::ROSFATAL) tinyros::logfatal(buffer);
}
}

