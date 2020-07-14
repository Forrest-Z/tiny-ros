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
std::string get_executable_name() {
#ifdef WIN32
  int i = 0;
  TCHAR exeFullPath[MAX_PATH];
  char actualpath[MAX_PATH];
  GetModuleFileName(NULL, exeFullPath, MAX_PATH);
  for(; exeFullPath[i] != 0; i++) {
    actualpath[i]=exeFullPath[i];
  }
  actualpath[i] = '\0';
  
  std::string processname = actualpath;
  std::size_t pos = processname.find_last_of('\\', processname.length());
  if (pos != std::string::npos) {
    return processname.substr(pos+1);
  }
#else
  char processdir[PATH_MAX];
  memset(processdir, 0, sizeof(processdir));
  if(readlink("/proc/self/exe", processdir,sizeof(processdir)) <= 0) {
    return "";
  }
  std::string processname = processdir;
  std::size_t pos = processname.find_last_of('/', processname.length());
  if (pos != std::string::npos) {
    return processname.substr(pos+1);
  }
#endif
  return "";
}

void mtrace(int level, const char *chfr, ...) {
  static std::string processname = "";
  if (processname.empty()) {
    processname = get_executable_name();
    processname = std::string("[") + processname + std::string("] ");
  }

  char buffer[TINYROS_LOG_MAX_SIZE] = {0};

  if (!processname.empty()) {
#ifdef WIN32
    _snprintf(buffer, sizeof(buffer) - 1, "%s", processname.c_str());
#else
    snprintf(buffer, sizeof(buffer) - 1, "%s", processname.c_str());
#endif
  }

  va_list ap;
  va_start(ap, chfr);
#ifdef WIN32
  _vsnprintf(buffer + processname.size(), sizeof(buffer) - (processname.size() + 1), chfr, ap);
#else
  vsnprintf(buffer + processname.size(), sizeof(buffer) - (processname.size() + 1), chfr, ap);
#endif
  va_end(ap);

  if(level == tinyros_msgs::Log::ROSDEBUG) tinyros::nh()->logdebug(buffer);
  else if(level == tinyros_msgs::Log::ROSINFO) tinyros::nh()->loginfo(buffer);
  else if(level == tinyros_msgs::Log::ROSWARN) tinyros::nh()->logwarn(buffer);
  else if(level == tinyros_msgs::Log::ROSERROR) tinyros::nh()->logerror(buffer);
  else if(level == tinyros_msgs::Log::ROSFATAL) tinyros::nh()->logfatal(buffer);
}
}

