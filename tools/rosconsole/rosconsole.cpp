/*
 * File      : ExampleSubscriber.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#include <unistd.h>
#include <stdio.h>
#include <time.h>
#include <sys/socket.h>   
#include <netinet/in.h>   
#include <arpa/inet.h> 
#include <sys/stat.h> 
#include "tiny_ros/ros.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "spdlog/spdlog.h"
#include "spdlog/sinks/stdout_sinks.h"
#include "spdlog/sinks/basic_file_sink.h"
#include "spdlog/sinks/rotating_file_sink.h"

#define LOG_ROTATION_SIZE    (10*1024*1024) //10MB
#define LOG_ROTATION_FILES   (5)

enum { LEVEL_DEBUG = 0 };
enum { LEVEL_INFO = 1 };
enum { LEVEL_WARN = 2 };
enum { LEVEL_ERROR = 3 };
enum { LEVEL_FATAL = 4 };
static int level_ = LEVEL_DEBUG;

enum { OPTION_P = 0 };
enum { OPTION_F = 1 };
enum { OPTION_D = 2 };
static int option_ = OPTION_P;

static std::string param_ = "";

static std::string dir_ = "";


static void print_usage() {
  printf("\nUsage: tinyrosconsole [OPTION] [VAR=LOG_LEVEL]...\n");
  printf("\nOPTION:\n");
  printf(" -h : display this help and exit.\n");
  printf(" -p : output log to screen.\n");
  printf(" -f : output log to the specified file.\n");
  printf(" -d : output log to the specified directory.\n");
  printf("\nLOG_LEVEL:\n");
  printf(" 0 : log level is debug.\n");
  printf(" 1 : log level is info.\n");
  printf(" 2 : log level is warn.\n");
  printf(" 3 : log level is error.\n");
  printf(" 4 : log level is fatal.\n");
  printf("\nEXAMPLE:\n");
  printf(" tinyrosconsole -p 0 : output log to screen with debug level\n");
  printf(" tinyrosconsole -p 0 127.0.0.1: output log to screen with debug level & tinyrosdds address\n");
  printf(" tinyrosconsole -f 1 log.txt : output log to log.txt with info level\n");
  printf(" tinyrosconsole -f 1 log.txt 127.0.0.1: output log to log.txt with info level & tinyrosdds address\n");
  printf(" tinyrosconsole -d 2 tinyros_logs : output log to folder \"tinyros_logs\" with warn level\n");
  printf(" tinyrosconsole -d 2 tinyros_logs 127.0.0.1: output log to folder \"tinyros_logs\" with warn level & tinyrosdds address\n\n");
}


static int create_dir(const char *dirname) {
    char* dir = strdup(dirname);
    if (!dir) {
        return -1;
    }

    int i, ret = -1, len = strlen(dir);

    for(i=1; i < len; i++) {
        if(dir[i]=='/') {
            dir[i] = 0;
            if(access((const char*)dir, F_OK) != 0 ) {
                if(mkdir(dir, 0777) == -1) {
                    goto beach;
                }
            }
            dir[i] = '/';
        }
    }

    if(access((const char*)dir, F_OK) != 0 ) {
        if(mkdir(dir, 0777) == -1) {
            goto beach;
        }
    }

    ret = 0;

    beach:
    if (!dir) {
        free(dir);
    }
    return ret;
}

static void messageCb(const tinyros_msgs::Log& l) {
  if(l.level == tinyros_msgs::Log::ROSDEBUG && level_ <= LEVEL_DEBUG) {
    if (spdlog::get("logger")) spdlog::get("logger")->debug("{0}", l.msg);
  } else if(l.level == tinyros_msgs::Log::ROSINFO && level_ <= LEVEL_INFO) {
    if (spdlog::get("logger")) spdlog::get("logger")->info("{0}", l.msg); 
  } else if(l.level == tinyros_msgs::Log::ROSWARN && level_ <= LEVEL_WARN) {
    if (spdlog::get("logger")) spdlog::get("logger")->warn("{0}", l.msg);
  } else if(l.level == tinyros_msgs::Log::ROSERROR && level_ <= LEVEL_ERROR) {
    if (spdlog::get("logger")) spdlog::get("logger")->error("{0}", l.msg);
  } else if(l.level == tinyros_msgs::Log::ROSFATAL && level_ <= LEVEL_FATAL) {
    if (spdlog::get("logger")) spdlog::get("logger")->critical("{0}", l.msg);
  }
}

static tinyros::Subscriber<tinyros_msgs::Log> sub(TINYROS_LOG_TOPIC, messageCb);

static void init_log_environment() {
  std::shared_ptr<spdlog::sinks::sink> log_sink = nullptr;
  if (option_ == OPTION_P) {
    log_sink = std::make_shared<spdlog::sinks::stdout_sink_mt>();
  } else {
    if (option_ == OPTION_D) {
      log_sink = std::make_shared<spdlog::sinks::rotating_file_sink_mt> (
        param_ + "/tinyros_apps.log", LOG_ROTATION_SIZE, LOG_ROTATION_FILES);
    } else {
      log_sink = std::make_shared<spdlog::sinks::basic_file_sink_mt>(param_);
    }
  }

  if (log_sink != nullptr) {
    log_sink->set_level(spdlog::level::trace);
    auto logger = std::make_shared<spdlog::logger>("logger", log_sink);
    logger->set_level(spdlog::level::trace);
    logger->flush_on(spdlog::level::trace);
    logger->set_pattern("[%Y-%m-%d %H:%M:%S.%e] [%l] %v");
    spdlog::register_logger(logger);
  }
}

int main(int argc, char *argv[]) {
  std::string ip = "127.0.0.1";
  
  if(argc >= 2) {
    if(!strcmp(argv[1], "-h")) {
      print_usage();
      return 0;
    } else if (!strcmp(argv[1], "-p")) {
      if (argc < 3) {
        print_usage();
        return 0;
      } else {
        level_ = atoi(argv[2]);
        option_ = OPTION_P;
        if (argc >= 4) {
          ip = argv[3];
        }
      }
    } else if (!strcmp(argv[1], "-f")){
      if (argc < 4) {
        print_usage();
        return 0;
      } else {
        level_ = atoi(argv[2]);
        param_ = argv[3];
        option_ = OPTION_F;
        if (argc >= 5) {
          ip = argv[4];
        }
      }
    } else if(!strcmp(argv[1], "-d")){
      if (argc < 4) {
        print_usage();
        return 0;
      } else {
        level_ = atoi(argv[2]);
        param_ = argv[3];
        option_ = OPTION_D;
        if (argc >= 5) {
          ip = argv[4];
        }
      }
    } else {
      print_usage();
      return 0;
    }
  } else {
    level_ = LEVEL_DEBUG;
    option_ = OPTION_P;
  }

  if (level_ < LEVEL_DEBUG || level_ > LEVEL_FATAL) {
    printf("Invalid log level: %d\n", level_);
    return 0;
  }
  
  if (option_ == OPTION_F || option_ == OPTION_D) {
    if (*(param_.begin()) != '/') {
      char buffer[1024];
      char* ret = getcwd(buffer, sizeof(buffer));
      param_ = std::string(buffer) + "/" + param_;
    }
  }

  init_log_environment();
  tinyros::init("tinyrosconsole", ip);
  tinyros::nh()->subscribe(sub);
  while(1) sleep(10);
}
