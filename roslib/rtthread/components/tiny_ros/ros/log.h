/*
 * File      : log.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#ifndef _TINYROS_LOG_H_
#define _TINYROS_LOG_H_

#include "tiny_ros/ros/time.h"
#include "tiny_ros/tinyros_msgs/Log.h"

namespace tinyros
{
tinyros::string get_executable_name();
void mtrace(int level, const char *chfr, ...);

#define log_info(format, args...) tinyros::mtrace(tinyros_msgs::Log::ROSINFO, format, ##args)
#define log_warn(format, args...) tinyros::mtrace(tinyros_msgs::Log::ROSWARN, format, ##args)
#define log_error(format, args...) tinyros::mtrace(tinyros_msgs::Log::ROSERROR, format, ##args)
#define log_debug(format, args...) tinyros::mtrace(tinyros_msgs::Log::ROSDEBUG, format, ##args)
#define log_diag(format, args...) tinyros::mtrace(tinyros_msgs::Log::ROSDIAG, format, ##args)

#define LOG_ONCE_TIME_THROTTLE  1
#define log_once_info(format, args...) { \
  static double __once_time = -1.0f; \
  if(__once_time < 0.0f) \
    __once_time = tinyros::Time::now().toSec(); \
  double __once_tem_sec =(tinyros::Time::now().toSec() - __once_time); \
  if((__once_tem_sec > LOG_ONCE_TIME_THROTTLE) || (__once_tem_sec < 0.0f)) { \
    __once_time = tinyros::Time::now().toSec(); \
    tinyros::mtrace(tinyros_msgs::Log::ROSINFO, format, ##args); \
  } \
}
#define log_once_warn(format, args...) { \
  static double __once_time = -1.0f; \
  if(__once_time < 0.0f) \
    __once_time = tinyros::Time::now().toSec(); \
  double __once_tem_sec =(tinyros::Time::now().toSec() - __once_time); \
  if((__once_tem_sec > LOG_ONCE_TIME_THROTTLE) || (__once_tem_sec < 0.0f)) { \
    __once_time = tinyros::Time::now().toSec(); \
    tinyros::mtrace(tinyros_msgs::Log::ROSWARN, format, ##args); \
  } \
}
#define log_once_error(format, args...) { \
  static double __once_time = -1.0f; \
  if(__once_time < 0.0f) \
    __once_time = tinyros::Time::now().toSec(); \
  double __once_tem_sec =(tinyros::Time::now().toSec() - __once_time); \
  if((__once_tem_sec > LOG_ONCE_TIME_THROTTLE) || (__once_tem_sec < 0.0f)) { \
    __once_time = tinyros::Time::now().toSec(); \
    tinyros::mtrace(tinyros_msgs::Log::ROSERROR, format, ##args); \
  } \
}
#define log_once_debug(format, args...) { \
  static double __once_time = -1.0f; \
  if(__once_time < 0.0f) \
    __once_time = tinyros::Time::now().toSec(); \
  double __once_tem_sec =(tinyros::Time::now().toSec() - __once_time); \
  if((__once_tem_sec > LOG_ONCE_TIME_THROTTLE) || (__once_tem_sec < 0.0f)) { \
    __once_time = tinyros::Time::now().toSec(); \
    tinyros::mtrace(tinyros_msgs::Log::ROSDEBUG, format, ##args); \
  } \
}

}  // namespace tinyros

#endif


