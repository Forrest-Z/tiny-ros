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
#include <boost/log/trivial.hpp>
#include <boost/log/utility/setup/common_attributes.hpp>
#include <boost/log/utility/setup/formatter_parser.hpp>
#include <boost/log/utility/setup/filter_parser.hpp>
#include <boost/log/utility/setup/settings.hpp>
#include <boost/log/utility/setup/from_settings.hpp>


#define LOG_ROTATION_SIZE  10*1024*1024 //10MB

#define LOG_MAX_SIZE0       200*1024*1024 //200MB

#define LOG_MAX_SIZE1       50*1024*1024 //50MB

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
	printf(" tinyrosconsole -f 1 log.txt : output log to log.txt with info level\n");
	printf(" tinyrosconsole -d 2 /userdata/atris_app : output log to folder \"/userdata/atris_app\" with warn level\n\n");
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
  if(l.level == tinyros_msgs::Log::ROSDEBUG && level_ <= LEVEL_DEBUG) BOOST_LOG_TRIVIAL(debug) << l.msg;
  else if(l.level == tinyros_msgs::Log::ROSINFO && level_ <= LEVEL_INFO) BOOST_LOG_TRIVIAL(info) << l.msg;
  else if(l.level == tinyros_msgs::Log::ROSWARN && level_ <= LEVEL_WARN) BOOST_LOG_TRIVIAL(warning) << l.msg;
  else if(l.level == tinyros_msgs::Log::ROSERROR && level_ <= LEVEL_ERROR) BOOST_LOG_TRIVIAL(error) << l.msg;
  else if(l.level == tinyros_msgs::Log::ROSFATAL && level_ <= LEVEL_FATAL) BOOST_LOG_TRIVIAL(fatal) << l.msg;
}

static tinyros::Subscriber<tinyros_msgs::Log> sub(TINYROS_LOG_TOPIC, messageCb);

static void init_log_environment() {
  namespace logging = boost::log;
  using namespace logging::trivial;

  logging::settings setts;
  logging::add_common_attributes();
  logging::register_simple_formatter_factory<severity_level, char>("Severity");
  logging::register_simple_filter_factory<severity_level, char>("Severity");

  setts["Core"]["DisableLogging"] = false;
  if (level_ == LEVEL_DEBUG) {
    setts["Core"]["Filter"] = "%Severity% >= trace";
  } else if (level_ == LEVEL_INFO) {
    setts["Core"]["Filter"] = "%Severity% >= info";
  } else if (level_ == LEVEL_WARN) {
    setts["Core"]["Filter"] = "%Severity% >= warning";
  } else if (level_ == LEVEL_ERROR) {
    setts["Core"]["Filter"] = "%Severity% >= error";
  } else if (level_ == LEVEL_FATAL) {
    setts["Core"]["Filter"] = "%Severity% >= fatal";
  } else {
    setts["Core"]["Filter"] = "%Severity% >= trace";
  }

  if (option_ == OPTION_P) { // Sinks.console
    setts["Sinks.console.Format"] = "[%TimeStamp%] [%Severity%] %Message%";
    setts["Sinks.console.Destination"] = "Console";
    setts["Sinks.console.Filter"] = "%Severity% >= debug";
    setts["Sinks.console.Asynchronous"] = false;
    setts["Sinks.console.AutoFlush"] = true;
  } else { // Sinks.degug
    setts["Sinks.degug.Filter"] = "%Severity% >= debug";
    setts["Sinks.degug.Format"] = "[%TimeStamp%] [%Severity%] %Message%";
    setts["Sinks.degug.Destination"] = "TextFile";
    if (option_ == OPTION_D) {
      setts["Sinks.degug.FileName"] = param_ + "/log/%Y%m%d_%H%M%S.log";
      setts["Sinks.degug.Target"] = param_ + "/log";
    } else {
      setts["Sinks.degug.FileName"] = param_;
      setts["Sinks.degug.Target"] = param_.substr(0, param_.rfind("/"));
    }
    setts["Sinks.degug.Asynchronous"] = false;
    setts["Sinks.degug.AutoFlush"] = true;
    setts["Sinks.degug.RotationSize"] = LOG_ROTATION_SIZE;
    setts["Sinks.degug.MaxSize"] = LOG_MAX_SIZE0;
    setts["Sinks.degug.ScanForFiles"] = "All";
  }

  logging::init_from_settings(setts);
}

int main(int argc, char *argv[]) {
  if(argc >= 2){
		if(!strcmp(argv[1], "-h")) {
      print_usage();
			return 0;
		}else if (!strcmp(argv[1], "-p")) {
      if (argc != 3) {
        print_usage();
  			return 0;
      } else {
        level_ = atoi(argv[2]);
        option_ = OPTION_P;
      }
		} else if (!strcmp(argv[1], "-f")){
      if (argc != 4) {
        print_usage();
  			return 0;
      } else {
        level_ = atoi(argv[2]);
        param_ = argv[3];
        option_ = OPTION_F;
      }
		} else if(!strcmp(argv[1], "-d")){
      if (argc != 4) {
        print_usage();
  			return 0;
      } else {
        level_ = atoi(argv[2]);
        param_ = argv[3];
        option_ = OPTION_D;
      }
		} else {
      print_usage();
  	  return 0;
    }
	} else {
     print_usage();
  	 return 0;
  }

  if (level_ < LEVEL_DEBUG || level_ > LEVEL_FATAL) {
    printf("Invalid log level: %d\n", level_);
    return 0;
  }
  
  if (option_ == OPTION_F || option_ == OPTION_D) {
    if (*(param_.begin()) != '/') {
      char buffer[1024];
      getcwd(buffer, sizeof(buffer));
      param_ = std::string(buffer) + "/" + param_;
    }
  }

  init_log_environment();
  tinyros::nh()->subscribe(sub);
  while(1) sleep(10);
}
