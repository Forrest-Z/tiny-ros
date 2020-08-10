#ifndef TINY_ROS_TOPIC_HANDLERS_H
#define TINY_ROS_TOPIC_HANDLERS_H
#include <thread>
#include <memory>
#include "signals.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"

namespace tinyros
{
class Rostopic;
class PublisherCore;
class SubscriberCore;
class ServiceServerCore;
class ServiceClientCore;

typedef  std::shared_ptr<Rostopic> RostopicPtr;
typedef  std::shared_ptr<PublisherCore> PublisherPtr;
typedef  std::shared_ptr<SubscriberCore> SubscriberPtr;
typedef  std::shared_ptr<ServiceServerCore> ServiceServerPtr;
typedef  std::shared_ptr<ServiceClientCore> ServiceClientPtr;

struct RostopicConnection {
  int id_;
  RostopicPtr rostopic_;
};

class Rostopic {
public:
  Rostopic(const tinyros_msgs::TopicInfo& topic_info) {
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    buffer_size_ = topic_info.buffer_size;
    ref_count_ = 0;
    signal_ = std::shared_ptr<Signal<tinyros::serialization::IStream&> >(new Signal<tinyros::serialization::IStream&>);
  }

public:
  static std::map<std::string, RostopicPtr> topics_;
  static std::mutex topics_mutex_;

public:
  std::shared_ptr<Signal<tinyros::serialization::IStream&> > signal_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  int32_t buffer_size_;
  int32_t ref_count_;
};
std::mutex Rostopic::topics_mutex_;
std::map<std::string, RostopicPtr> Rostopic::topics_;

class PublisherCore {
public:
  PublisherCore(const tinyros_msgs::TopicInfo& topic_info) {
    topic_id_ = topic_info.topic_id;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    node_name_ = topic_info.node;
    buffer_size_ = topic_info.buffer_size;
  }

  ~PublisherCore() {
    std::unique_lock<std::mutex> lock(Rostopic::topics_mutex_);
    connection_.rostopic_->ref_count_--;
    if (connection_.rostopic_->ref_count_ <=0) {
      Rostopic::topics_.erase(connection_.rostopic_->topic_name_);
    }
  }
  
  void handle(tinyros::serialization::IStream& stream) {
    if (Rostopic::topics_.count(topic_name_)) {
      Rostopic::topics_[topic_name_]->signal_->emit(stream);
    }
  }

  uint32_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  std::string node_name_;
  int32_t buffer_size_;
  uint64_t alive_time_;
  RostopicConnection connection_;
};

class SubscriberCore {
public:
  SubscriberCore(tinyros_msgs::TopicInfo& topic_info,
      std::function<void(tinyros::serialization::IStream& message)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = topic_info.topic_id;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    node_name_ = topic_info.node;
    buffer_size_ = topic_info.buffer_size;
  }

  ~SubscriberCore() {
    std::unique_lock<std::mutex> lock(Rostopic::topics_mutex_);
    connection_.rostopic_->signal_->disconnect(connection_.id_);
    connection_.rostopic_->ref_count_--;
    if (connection_.rostopic_->ref_count_ <=0) {
      Rostopic::topics_.erase(connection_.rostopic_->topic_name_);
    }
  }
  
  void handle(tinyros::serialization::IStream& message) {
    write_fn_(message);
  }

  std::function<void(tinyros::serialization::IStream& message)> write_fn_;
  uint32_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  std::string node_name_;
  int32_t buffer_size_;
  uint64_t alive_time_;
  RostopicConnection connection_;
};

class ServiceServerCore {
public:
  ServiceServerCore(tinyros_msgs::TopicInfo& topic_info,
      std::function<void(tinyros::serialization::IStream& message, const uint16_t topic_id)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = -1;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    node_name_ = topic_info.node;
    buffer_size_ = topic_info.buffer_size;
    signal_ = std::shared_ptr<Signal<tinyros::serialization::IStream&> >(new Signal<tinyros::serialization::IStream&>);
    destroy_signal_ = std::shared_ptr<Signal<std::string&> >(new Signal<std::string&>);
  }
  void setTopicId(uint32_t topic_id) {
    topic_id_ = topic_id;
  }

  void handle(tinyros::serialization::IStream& message) {
    signal_->emit(message);
  }

  void callback(tinyros::serialization::IStream& message) {
    write_fn_(message, topic_id_);
  }

public:
  std::function<void(tinyros::serialization::IStream& message, const uint16_t topic_id)> write_fn_;
  std::shared_ptr<Signal<tinyros::serialization::IStream&> > signal_;
  std::shared_ptr<Signal<std::string&> > destroy_signal_;
  static std::map<std::string, ServiceServerPtr> services_;
  static std::mutex services_mutex_;
  uint32_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  std::string node_name_;
  int32_t buffer_size_;
};
std::mutex ServiceServerCore::services_mutex_;
std::map<std::string, ServiceServerPtr> ServiceServerCore::services_;

class ServiceClientCore {
public:
  ServiceClientCore(tinyros_msgs::TopicInfo& topic_info,
      std::function<void(tinyros::serialization::IStream& message, const uint16_t topic_id)> write_fn)
    : write_fn_(write_fn) {
    topic_id_ = -1;
    topic_name_ = topic_info.topic_name;
    message_type_ = topic_info.message_type;
    md5sum_ = topic_info.md5sum;
    node_name_ = topic_info.node;
    buffer_size_ = topic_info.buffer_size;
    signal_ = std::shared_ptr<Signal<tinyros::serialization::IStream&> >(new Signal<tinyros::serialization::IStream&>);
  }

  void setTopicId(uint32_t topic_id) {
    topic_id_ = topic_id;
  }

  void handle(tinyros::serialization::IStream& message) {
    signal_->emit(message);
  }

  void callback(tinyros::serialization::IStream& message) {
    write_fn_(message, topic_id_);
  }

public:
  std::function<void(tinyros::serialization::IStream& message, const uint16_t topic_id)> write_fn_;
  std::shared_ptr<Signal<tinyros::serialization::IStream&> > signal_;
  int client_connection_;
  int service_connection_;
  int destroy_connection_;
  uint32_t topic_id_;
  std::string topic_name_;
  std::string message_type_;
  std::string md5sum_;
  std::string node_name_;
  int32_t buffer_size_;
};

}  // namespace

#endif  // TINY_ROS_TOPIC_HANDLERS_H
