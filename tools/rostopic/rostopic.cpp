#include <stdio.h>
#include <thread>
#include "spdlog/spdlog.h"
#include "spdlog/sinks/stdout_sinks.h"
#include "spdlog/sinks/basic_file_sink.h"
#include "rostopic_subscribers.cpp"

static void print_usage() {
  printf("\n\nUsage:\n");
  printf(" tinyrostopic is a command-line tool for printing information about tinyros topics\n\n");
  printf("Commands:\n");
  printf(" tinyrostopic help : display this help usage\n");
  printf(" tinyrostopic list : list active topics\n");
  printf(" tinyrostopic echo /topic [options] : echo messages to screen\n");
  printf(" tinyrostopic echo -b [options] /topic [options] : echo messages to \".bag\" file\n\n");
  printf("Example:\n");
  printf(" tinyrostopic list : list active topics\n");
  printf(" tinyrostopic list 127.0.0.1: list active topics with tinyrosdds address\n");
  printf(" tinyrostopic echo /topic : echo messages to screen with topic\n");
  printf(" tinyrostopic echo /topic 127.0.0.1: echo messages to screen with topic & tinyrosdds address\n");
  printf(" tinyrostopic echo -b topic.bag /topic: echo messages to \".bag\" file with topic\n");
  printf(" tinyrostopic echo -b topic.bag /topic 127.0.0.1: echo messages to \".bag\" file with topic & tinyrosdds address\n\n");
}

static std::vector<std::string> string_split(const std::string& s, const std::string& delim="\n") {
  std::vector<std::string> elems;
  std::size_t pos = 0;
  std::size_t len = s.length();
  std::size_t delim_len = delim.length();
  if (delim_len == 0) return elems;
  while (pos < len) {
    std::size_t find_pos = s.find(delim, pos);
    if (find_pos < 0) {
      elems.push_back(s.substr(pos, len - pos));
      break;
    }
    elems.push_back(s.substr(pos, find_pos - pos));
    pos = find_pos + delim_len;
  }
  return elems;
}

static void init_log_environment(std::string bag) {
  std::shared_ptr<spdlog::sinks::sink> log_sink = nullptr;
  if (bag.empty()) {
    log_sink = std::make_shared<spdlog::sinks::stdout_sink_mt>();
  } else {
    std::string suffix = "";
    std::size_t suffix_pos = bag.rfind(".");
    if (suffix_pos != std::string::npos) {
      suffix = bag.substr(suffix_pos);
    }
    if (suffix != ".bag") {
      bag += ".bag";
    }
    
    log_sink = std::make_shared<spdlog::sinks::basic_file_sink_mt>(bag);
  }

  if (log_sink != nullptr) {
    log_sink->set_level(spdlog::level::trace);
    auto logger = std::make_shared<spdlog::logger>("logger", log_sink);
    logger->set_level(spdlog::level::trace);
    logger->flush_on(spdlog::level::trace);
    logger->set_pattern("[%Y-%m-%d %H:%M:%S.%e] %v");
    spdlog::register_logger(logger);
  }
}

static void rostopic_cmd_list() {
  while(!tinyros::nh()->ok()) {
#ifdef WIN32
    Sleep(1000);
#else
    sleep(1);
#endif
  }
  
  std::string topiclist = "";
  while (topiclist.empty()) {
    topiclist = tinyros::nh()->getTopicList(3000);
  }
  
  std::cout << topiclist.c_str() << std::endl;
  
  exit(0);
}

static void rostopic_cmd_echo(std::string topic) {
  while(!tinyros::nh()->ok()) {
#ifdef WIN32
    Sleep(1000);
#else
    sleep(1);
#endif
  }
  
  std::string topiclist = "";
  while (topiclist.empty()) {
    topiclist = tinyros::nh()->getTopicList(3000);
  }

  std::string type = "";
  std::string md5 = "";
  bool old = (topiclist.find("topic_list:") != std::string::npos);
  std::vector<std::string> topics = string_split(topiclist);
  for (std::size_t i = 0; i < topics.size(); i++) {
    std::string elem = topics[i];
    std::size_t pos = elem.find(topic);
    if (pos != std::string::npos) {
      if (!old) {
        std::size_t type_begin = elem.find("[type:");
        std::size_t type_end = elem.find(",");
        std::size_t md5_begin = elem.find("md5:");
        std::size_t md5_end = elem.find("]");
        std::string elem_topic = elem.substr(pos, type_begin - pos);
        elem_topic.erase(elem_topic.find_last_not_of(" \t\f\v\n\r") + 1);
        elem_topic.erase(0, elem_topic.find_first_not_of(" \t\f\v\n\r"));
        if (elem_topic == topic) {
          if (type_begin != std::string::npos && type_end != std::string::npos
            && md5_begin != std::string::npos && md5_end != std::string::npos) {
            type = elem.substr((type_begin + 6), (type_end - type_begin - 6));
            md5 = elem.substr((md5_begin + 4), (md5_end - md5_begin - 4));
            break;
          }
        }
      } else {
        std::size_t type_begin = elem.find("[");
        std::size_t type_end = elem.find("]");
        std::string elem_topic = elem.substr(pos, type_begin - pos);
        elem_topic.erase(elem_topic.find_last_not_of(" \t\f\v\n\r") + 1);
        elem_topic.erase(0, elem_topic.find_first_not_of(" \t\f\v\n\r"));
        if (elem_topic == topic) {
          if (type_begin != std::string::npos && type_end != std::string::npos) {
            type = elem.substr((type_begin + 1), (type_end - type_begin - 1));
            break;
          }
        }
      }
    }
  }
  
  type.erase(type.find_last_not_of(" \t\f\v\n\r") + 1);
  type.erase(0, type.find_first_not_of(" \t\f\v\n\r"));
  md5.erase(md5.find_last_not_of(" \t\f\v\n\r") + 1);
  md5.erase(0, md5.find_first_not_of(" \t\f\v\n\r"));

  if (type.empty()) {
    std::cout << "ERROR: Your topic[" << topic << "] seems invalid.\n" << std::endl;
    exit(0);
  }
  
  if (tinyros::rostopic_subscribers.count(type)) {
    tinyros::rostopic_subscribers[type]->topic_ = topic;
    tinyros::nh()->subscribe(*(tinyros::rostopic_subscribers[type]));
    
    std::string native_md5 = tinyros::rostopic_subscribers[type]->getMsgMD5();
    if (!md5.empty() && native_md5 != md5) {
      std::cout << "WARNING: Message MD5 does not match [remote:" << md5 << ", native: " << native_md5 << "]" << std::endl;
    }
  } else {
    std::cout << "ERROR: Cannot load message for [" << type << "]. Are your messages built?\n" << std::endl;
    exit(0);
  }
}

int main(int argc, char **argv) {
  std::string ip = "127.0.0.1";
  std::string file = "";
  
  if(argc >= 2) {
    if(!strcmp(argv[1], "help")) {
      print_usage();
      return 0;
    } else if (!strcmp(argv[1], "echo")) {
      if (argc < 3) {
        print_usage();
        return 0;
      } else {
        if (!strcmp(argv[2], "-b")) {
          if (argc < 5) {
            std::cout << "ERROR: Option is not valid with bag files\n" << std::endl;
            print_usage();
            return 0;
          }
          if (argc >= 6) {
            ip = argv[5];
          }

          file = argv[3];
          tinyros::init("tinyrostopic", ip);
          init_log_environment(file);
          rostopic_cmd_echo(argv[4]);
        } else {
          if (argc >= 4) {
            ip = argv[3];
          }
          tinyros::init("tinyrostopic", ip);
          init_log_environment(file);
          rostopic_cmd_echo(argv[2]);
        }
      }
    } else if (!strcmp(argv[1], "list")) {
      if (argc >= 3) {
        ip = argv[2];
      }
      tinyros::init("tinyrostopic", ip);
      init_log_environment(file);
      rostopic_cmd_list();
    } else {
      print_usage();
      return 0;
    }
  } else {
    print_usage();
    return 0;
  }

  while(1) sleep(10);
  
  return 0;
}

