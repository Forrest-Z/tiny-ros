/*
 * File      : topic_handlers.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_TOPIC_HANDLERS_H
#define TINY_ROS_TOPIC_HANDLERS_H

#include <boost/signals2.hpp>
#include <boost/thread.hpp>
#include "tiny_ros/tinyros_msgs/TopicInfo.h"

namespace tinyros
{
class Rostopic;
class Publisher;
class Subscriber;
class ServiceServer;
class ServiceClient;

typedef  boost::shared_ptr<Rostopic> RostopicPtr;
typedef  boost::shared_ptr<Publisher> PublisherPtr;
typedef  boost::shared_ptr<Subscriber> SubscriberPtr;
typedef  boost::shared_ptr<ServiceServer> ServiceServerPtr;
typedef  boost::shared_ptr<ServiceClient> ServiceClientPtr;

struct RostopicConnection {
  RostopicPtr rostopic_;
  boost::signals2::connection connection_;
};

class Rostopic {
public:
  Rostopic(const tinyros_msgs::TopicInfo& topic_info) {
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
    ref_count_ = 0;
    signal_ = boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> >
      (new boost::signals2::signal<void(std::vector<uint8_t>&)>);
  }

public:
  static std::map<std::string, RostopicPtr> topics_;
  static boost::mutex topics_mutex_;

public:
  boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> > signal_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
  int32_t ref_count_;
};
boost::mutex Rostopic::topics_mutex_;
std::map<std::string, RostopicPtr> Rostopic::topics_;

class Publisher {
public:
  Publisher(const tinyros_msgs::TopicInfo& topic_info) {
    topic_id_ = topic_info.topic_id;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
  }

  void handle(tinyros::serialization::IStream stream) {
    uint32_t length = stream.getLength();
    std::vector<uint8_t> message(length);
    memcpy(&message[0], stream.getData(), length);
    if (Rostopic::topics_.count(topic_name_)) {
      (*(Rostopic::topics_[topic_name_]->signal_))(message);
    }
  }

  std::string get_topic() {
    return topic_name_;
  }

private:
  uint16_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
};

class Subscriber {
public:
  Subscriber(tinyros_msgs::TopicInfo& topic_info,
      boost::function<void(std::vector<uint8_t>& buffer)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = topic_info.topic_id;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
  }

  std::string get_topic() {
    return topic_name_;
  }

private:
  void handle(std::vector<uint8_t>& message) {
    write_fn_(message);
  }

  boost::function<void(std::vector<uint8_t>& buffer)> write_fn_;
  uint16_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
};

class ServiceServer {
public:
  ServiceServer(tinyros_msgs::TopicInfo& topic_info,
      boost::function<void(std::vector<uint8_t>& buffer, const uint16_t topic_id)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = -1;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
    signal_ = boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> >
      (new boost::signals2::signal<void(std::vector<uint8_t>&)>);
    destroy_signal_ = boost::shared_ptr<boost::signals2::signal<void(std::string&)> >
      (new boost::signals2::signal<void(std::string&)>);
  }
  void setTopicId(uint16_t topic_id) {
    topic_id_ = topic_id;
  }

  void handle(tinyros::serialization::IStream stream) {
    uint32_t length = stream.getLength();
    std::vector<uint8_t> message(length);
    memcpy(&message[0], stream.getData(), length);
    (*signal_)(message);
  }

  void callback(std::vector<uint8_t>& message) {
    write_fn_(message, topic_id_);
  }

public:
  boost::function<void(std::vector<uint8_t>& buffer, const uint16_t topic_id)> write_fn_;
  boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> > signal_;
  boost::shared_ptr<boost::signals2::signal<void(std::string&)> > destroy_signal_;
  static std::map<std::string, ServiceServerPtr> services_;
  static boost::mutex services_mutex_;
  uint16_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
};
boost::mutex ServiceServer::services_mutex_;
std::map<std::string, ServiceServerPtr> ServiceServer::services_;

class ServiceClient {
public:
  ServiceClient(tinyros_msgs::TopicInfo& topic_info,
      boost::function<void(std::vector<uint8_t>& buffer, const uint16_t topic_id)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = -1;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
    signal_ = boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> >
      (new boost::signals2::signal<void(std::vector<uint8_t>&)>);
  }

  void setTopicId(uint16_t topic_id) {
    topic_id_ = topic_id;
  }

  void handle(tinyros::serialization::IStream stream) {
    uint32_t length = stream.getLength();
    std::vector<uint8_t> message(length);
    memcpy(&message[0], stream.getData(), length);
    (*signal_)(message);
  }

  void callback(std::vector<uint8_t>& message) {
    write_fn_(message, topic_id_);
  }

public:
  boost::function<void(std::vector<uint8_t>& buffer, const uint16_t topic_id)> write_fn_;
  boost::shared_ptr<boost::signals2::signal<void(std::vector<uint8_t>&)> > signal_;
  boost::signals2::connection client_connection_;
  boost::signals2::connection service_connection_;
  boost::signals2::connection destroy_connection_;
  uint16_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
};

}  // namespace

#endif  // TINY_ROS_TOPIC_HANDLERS_H
