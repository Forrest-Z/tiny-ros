/*
 * File      : session.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-03     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_SERVER_SESSION_H
#define TINY_ROS_SERVER_SESSION_H
#include <map>
#include <memory>
#include <thread>
#include <mutex>
#include <deque>
#include <chrono>
#include <functional>
#include <condition_variable>
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/roslib_msgs/Time.h"
#include "tiny_ros/roslib_msgs/String.h"
#include "serialization.h"
#include "topic_handlers.h"
#include "serialization.h"

namespace tinyros
{
#define TINYROS_LOG_TOPIC "tinyros_log_11315"

#define REQUEST_TOPICS_TIMER (1000*1000)

#define REQUEST_TOPICS_ALIVE_TIME (15) // seconds

typedef std::vector<uint8_t> Buffer;
typedef std::shared_ptr<Buffer> BufferPtr;
typedef std::deque<BufferPtr> AsyncWritebuffer;

template<typename Socket>
class Session
{
public:
  std::string session_id_;
  
  Session(Socket& socket, bool isudp = false)
    : socket_(socket)
    , active_(false)
    , isudp_(isudp)
    , message_in_thread_(nullptr)
    , message_write_thread_(nullptr)
  {
    if (isudp) {
      session_id_ = "session_udp";
    }
  }

  Socket& socket()
  {
    return socket_;
  }

  void start()
  {
    using namespace tinyros_msgs;
    callbacks_[TopicInfo::ID_PUBLISHER] = std::bind(&Session::setup_publisher, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SUBSCRIBER] = std::bind(&Session::setup_subscriber, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SERVICE_SERVER+TopicInfo::ID_PUBLISHER] = std::bind(&Session::setup_service_server, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SERVICE_SERVER+TopicInfo::ID_SUBSCRIBER] = std::bind(&Session::setup_service_server, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SERVICE_CLIENT+TopicInfo::ID_PUBLISHER] = std::bind(&Session::setup_service_client, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SERVICE_CLIENT+TopicInfo::ID_SUBSCRIBER] = std::bind(&Session::setup_service_client, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_ROSTOPIC_REQUEST] = std::bind(&Session::handle_rostopic_request, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_ROSSERVICE_REQUEST] = std::bind(&Session::handle_rosservice_request, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_LOG] = std::bind(&Session::handle_log, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_TIME] = std::bind(&Session::handle_time, this, std::placeholders::_1);
    callbacks_[TopicInfo::ID_SESSION_ID] = std::bind(&Session::handle_session_id, this, std::placeholders::_1);

    active_ = true;
    
    require_check_running_ = true;
 
    message_write_thread_ = new std::thread(std::bind(&Session::write_completion_cb, this));
    require_check_thread_ = new std::thread(std::bind(&Session::request_topics, this));
    if (isudp_) {
      message_in_thread_ = new std::thread(std::bind(&Session::read_message_sync_udp, this));
    } else {
      message_in_thread_ = new std::thread(std::bind(&Session::read_message_sync, this));
    }
  }

  void stop()
  {
    {
      std::unique_lock<std::mutex> lock(active_mutex_);
      if (!active_) {
        return;
      }
      
      active_ = false;
    
      spdlog_warn("[{0}] {1} begin.", session_id_.c_str(), __FUNCTION__);
      if (require_check_thread_) {
        spdlog_warn("[{0}] {1} require_check_thread interrupt begin.", session_id_.c_str(), __FUNCTION__);
        require_check_running_ = false;
        require_check_thread_->join();
        delete require_check_thread_;
        require_check_thread_ = nullptr;
        spdlog_warn("[{0}] {1} require_check_thread interrupt end.", session_id_.c_str(), __FUNCTION__);
      }

      if (message_in_thread_) {
        spdlog_warn("[{0}] {1} message_in_thread interrupt begin.", session_id_.c_str(), __FUNCTION__);
        message_in_thread_->join();
        delete message_in_thread_;
        message_in_thread_ = nullptr;
        spdlog_warn("[{0}] {1} message_in_thread interrupt end.", session_id_.c_str(), __FUNCTION__);
      }

      spdlog_warn("[{0}] {1} clear async_write_buffers begin.", session_id_.c_str(), __FUNCTION__);
      std::unique_lock<std::mutex> async_write_lock(async_write_mutex_);
      async_write_buffers_.clear();
      async_write_cond_.notify_all();
      async_write_lock.unlock();
      spdlog_warn("[{0}] {1} clear async_write_buffers end.", session_id_.c_str(), __FUNCTION__);
  
      if (message_write_thread_) {
        spdlog_warn("[{0}] {1} message_write_thread interrupt begin.", session_id_.c_str(), __FUNCTION__);
        message_write_thread_->join();
        delete message_write_thread_;
        message_write_thread_ = nullptr;
        spdlog_warn("[{0}] {1} message_write_thread interrupt end.", session_id_.c_str(), __FUNCTION__);
      }
    }

    {
      // Reset the state of the session, dropping any publishers or subscribers
      // we currently know about from this client.
      callbacks_.clear();    // 1
      subscribers_.clear();  // 2
      publishers_.clear();   // 3
    
      spdlog_warn("[{0}] {1} services clear begin.", session_id_.c_str(), __FUNCTION__);
      std::unique_lock<std::mutex> lock(ServiceServerCore::services_mutex_);
      for (uint32_t i=0; i<service_server_.size(); i++) {
        if (ServiceServerCore::services_.count(service_server_.at(i))) {
          ServiceServerPtr service = ServiceServerCore::services_[service_server_.at(i)];
          service->destroy_signal_->emit(service_server_.at(i));
          service->signal_->disconnect_all();
          service->destroy_signal_->disconnect_all();
          ServiceServerCore::services_.erase(service_server_.at(i));
        }
      }
      
      std::map<uint32_t, ServiceClientPtr>::iterator it;
      for(it = services_client_.begin(); it != services_client_.end(); ) {
        ServiceClientPtr client = it->second;
        if (ServiceServerCore::services_.count(client->topic_name_)) {
          ServiceServerPtr service = ServiceServerCore::services_[client->topic_name_];
          client->signal_->disconnect(client->client_connection_);
          service->signal_->disconnect(client->service_connection_);
          service->destroy_signal_->disconnect(client->destroy_connection_);
        }
        it++;
      }
      service_server_.clear();
      services_client_.clear();
      spdlog_warn("[{0}] {1} services clear end.", session_id_.c_str(), __FUNCTION__);
    }
    
    // Close the socket.
    spdlog_warn("[{0}] {1} Close the socket begin.", session_id_.c_str(), __FUNCTION__);
    socket_.close();
    spdlog_warn("[{0}] {1} Close the socket end.", session_id_.c_str(), __FUNCTION__);
    
    spdlog_warn("[{0}] {1} end.", session_id_.c_str(), __FUNCTION__);
  }

  bool is_active()
  {
    return active_;
  }

private:
  void read_message_sync_udp() {
    uint8_t *message_in = (uint8_t*)calloc(buffer_max, sizeof(uint8_t));
    while (is_active()) {
      int32_t rv = socket_.read_some(message_in, buffer_max, session_id_);
      if (buffer_max >= rv && rv > 0)  {
        uint32_t topic = 0;
        int bytes = 0, index= 0, checksum = 0;
        do {
          index = 0;
          if (message_in[index++] != 0xff) {
            break;
          }

          if (message_in[index++] != 0xb9) {
            break;
          }

          bytes = message_in[index];
          bytes += message_in[index + 1] << 8;
          bytes += message_in[index + 2] << 16;
          bytes += message_in[index + 3] << 24;
          checksum = message_in[index];
          checksum += message_in[index + 1];
          checksum += message_in[index + 2];
          checksum += message_in[index + 3];
          checksum += message_in[index + 4];
          index += 5;
          
          if((checksum % 256) != 255) {
            break;
          }

          topic = message_in[index];
          topic += message_in[index + 1] << 8;
          topic += message_in[index + 2] << 16;
          topic += message_in[index + 3] << 24;
          checksum = message_in[index];
          checksum += message_in[index + 1];
          checksum += message_in[index + 2];
          checksum += message_in[index + 3];
          index += 4;

          if (buffer_max < (index + bytes + 1)) {
            break;
          }

          if(bytes > 0) {
            for (int32_t i=0; i < bytes + 1; i++) {
              checksum += message_in[index + i];
            }
          } else {
            checksum += message_in[index];
          }

          if ((checksum % 256) == 255) {
            tinyros::serialization::IStream stream(message_in + index, bytes);
            if (callbacks_.count(topic) == 1) {
              try {
                callbacks_[topic](stream);
              } catch(tinyros::serialization::StreamOverrunException e) {
              }
            } else {
              spdlog_warn("[{0}] {1} Received message with unrecognized topicId ({2}).", session_id_.c_str(), __FUNCTION__, topic);
            }
          } else {
            spdlog_warn("[{0}] {1} Rejecting message on topicId({2}), bytes({3}) with bad checksum.", 
              session_id_.c_str(), __FUNCTION__, topic, bytes);
          }
        } while(0);
      }
    }

    if (message_in) {
      free(message_in);
    }
  }
  
  void read_message_sync() {
    const uint8_t MODE_FIRST_FF = 0;
    const uint8_t MODE_PROTOCOL_VER   = 1;
    const uint8_t PROTOCOL_VER        = 0xb9;
    const uint8_t MODE_SIZE_L         = 2;
    const uint8_t MODE_SIZE_L1        = 3;
    const uint8_t MODE_SIZE_H         = 4;
    const uint8_t MODE_SIZE_H1        = 5;
    const uint8_t MODE_SIZE_CHECKSUM  = 6;
    const uint8_t MODE_TOPIC_L        = 7;
    const uint8_t MODE_TOPIC_L1       = 8;
    const uint8_t MODE_TOPIC_H        = 9;
    const uint8_t MODE_TOPIC_H1       = 10;
    const uint8_t MODE_MESSAGE        = 11;
    const uint8_t MODE_MSG_CHECKSUM   = 12;
    
    uint8_t *message_in = (uint8_t*)calloc(buffer_max, sizeof(uint8_t));
    uint8_t *message_tmp = (uint8_t*)calloc(buffer_max, sizeof(uint8_t));
    uint32_t topic = 0;
    int mode = MODE_FIRST_FF;
    int total_bytes = 0;
    int bytes = 0;
    int index= 0;
    int checksum = 0;
    int i, rv, len = 1;
    
    while (message_tmp && message_in && is_active()) {
      if (len > buffer_max) {
        len = 1;
        mode = MODE_FIRST_FF;
        continue;
      }
      
      if ((rv = socket_.read_some(message_tmp, len, session_id_)) < 0) {
        if (is_active()) {
          std::thread tid(std::bind(&Session::stop, this));
          tid.detach();
        }
        break;
      }
      
      if (rv > 0) {
        if (mode != MODE_MESSAGE) {
          for (i = 0; i < rv; i++) {
            checksum += message_tmp[i];
          }
        }
        
        if (mode == MODE_MESSAGE) {
          for (i = 0; i < rv; i++) {
            checksum += message_tmp[i];
            message_in[index++] = message_tmp[i];
            bytes--;
          }

          if (bytes == 0) {
            len = 1;
            mode = MODE_MSG_CHECKSUM;
          } else {
            len = bytes;
          }
        } else if (mode == MODE_FIRST_FF) {
          if (message_tmp[0] == 0xff) {
            mode++;
          }
        } else if (mode == MODE_PROTOCOL_VER) {
          if (message_tmp[0] == PROTOCOL_VER) {
            mode++;
          } else {
            mode = MODE_FIRST_FF;
          }
        } else if (mode == MODE_SIZE_L) {
          bytes = message_tmp[0];
          index = 0;
          mode++;
          checksum = message_tmp[0];
        } else if (mode == MODE_SIZE_L1) {
          bytes += message_tmp[0] << 8;
          mode++;
        } else if (mode == MODE_SIZE_H) {
          bytes += message_tmp[0] << 16;
          mode++;
        } else if (mode == MODE_SIZE_H1) {
          bytes += message_tmp[0] << 24;
          mode++;
        } else if (mode == MODE_SIZE_CHECKSUM) {
          if ((checksum % 256) == 255) {
            mode++;
          } else {
            mode = MODE_FIRST_FF;
          }
        } else if (mode == MODE_TOPIC_L) {
          topic = message_tmp[0];
          mode++;
          checksum = message_tmp[0]; 
        } else if (mode == MODE_TOPIC_L1) {
          topic += message_tmp[0] << 8;
          mode++;
        } else if (mode == MODE_TOPIC_H) {
          topic += message_tmp[0] << 16;
          mode++;
        } else if (mode == MODE_TOPIC_H1) {
          topic += message_tmp[0] << 24;
          mode = MODE_MESSAGE;
          total_bytes = bytes;
          if (bytes == 0)
            mode = MODE_MSG_CHECKSUM;
          else
            len = bytes;
        } else if (mode == MODE_MSG_CHECKSUM) {
          mode = MODE_FIRST_FF;
          if ((checksum % 256) == 255) {
            tinyros::serialization::IStream stream(message_in, total_bytes);
            if (callbacks_.count(topic) == 1) {
              try {
                callbacks_[topic](stream);
              } catch(tinyros::serialization::StreamOverrunException e) {
              }
            } else {
              spdlog_warn("[{0}] {1} Received message with unrecognized topicId ({2}).", session_id_.c_str(), __FUNCTION__, topic);
            }
          } else {
            spdlog_warn("[{0}] {1} Rejecting message on topicId({2}), bytes({3}) with bad checksum.", 
              session_id_.c_str(), __FUNCTION__, topic, bytes);
          }
        }
      }
    }

    if (message_in) {
      free(message_in);
    }
    
    if (message_tmp) {
      free(message_tmp);
    }
  }

  //// SENDING MESSAGES ////

  void write_message(Buffer& message, const uint32_t topic_id) {
    if (!is_active()) return;
    
    uint8_t overhead_bytes = 12;
    uint32_t length = overhead_bytes + message.size();
    BufferPtr buffer_ptr(new Buffer(length));

    uint8_t msg_checksum;
    tinyros::serialization::IStream checksum_stream(message.size() > 0 ? &message[0] : NULL, message.size());

    tinyros::serialization::OStream stream(&buffer_ptr->at(0), buffer_ptr->size());
    uint8_t msg_len_checksum = 255 - checksum((uint32_t)message.size());

    stream << (uint16_t)0xb9ff << (uint32_t)message.size() << msg_len_checksum << topic_id;
    msg_checksum = 255 - (checksum(checksum_stream) + checksum(topic_id));

    memcpy(stream.advance(message.size()), &message[0], message.size());
    stream << msg_checksum;

    std::unique_lock<std::mutex> lock(async_write_mutex_);
    async_write_buffers_.push_back(buffer_ptr);
    async_write_cond_.notify_one();
  }

  void write_completion_cb() {
    while (is_active()) {
      std::unique_lock<std::mutex> lock(async_write_mutex_);
      if(async_write_buffers_.empty()) {
        async_write_cond_.wait(lock);
      }

      BufferPtr buffer_ptr = nullptr;
      if(is_active() && !async_write_buffers_.empty()) {
        buffer_ptr = async_write_buffers_.front();
        async_write_buffers_.pop_front();
      }
      lock.unlock();

      if (is_active() && buffer_ptr) {
        if ((socket_.write_some((uint8_t*)buffer_ptr->data(), (int)buffer_ptr->size(), session_id_)) < 0) {
          if (is_active()) {
            std::thread tid(std::bind(&Session::stop, this));
            tid.detach();
          }
          break;
        }
      }
    }
  }

  //// HELPERS ////
  void request_topics() {
    while(require_check_running_) {
      if (!isudp_) {
        std::vector<uint8_t> message(0);
        write_message(message, tinyros_msgs::TopicInfo::ID_PUBLISHER);
      } else {
        std::map<uint32_t, PublisherPtr>::iterator pit;
        for(pit = publishers_.begin(); pit != publishers_.end(); ) {
          PublisherPtr pub = pit->second;
          uint64_t now = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
          if ((now - pub->alive_time_) > REQUEST_TOPICS_ALIVE_TIME) {
            spdlog_info("[{0}] Publisher remove(topic_id: {1}, topic_name: {2})", session_id_.c_str(), pub->topic_id_, pub->topic_name_.c_str());
            callbacks_.erase(pit->first);
            publishers_.erase(pit++);
          } else {
            pit++;
          }
        }
        
        std::map<uint32_t, SubscriberPtr>::iterator sit;
        for(sit = subscribers_.begin(); sit != subscribers_.end(); ) {
          SubscriberPtr sub = sit->second;
          uint64_t now = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
          if ((now - sub->alive_time_) > REQUEST_TOPICS_ALIVE_TIME) {
            spdlog_info("[{0}] Subscriber remove(topic_id: {1}, topic_name: {2})", session_id_.c_str(), sub->topic_id_, sub->topic_name_.c_str());
            subscribers_.erase(sit++);
          } else {
            sit++;
          }
        }
      }
      
      usleep(REQUEST_TOPICS_TIMER);
    }
  }

  static uint8_t checksum(tinyros::serialization::IStream& stream) {
    uint8_t sum = 0;
    for (uint32_t i = 0; i < stream.getLength(); ++i) {
      sum += stream.getData()[i];
    }
    return sum;
  }

  static uint8_t checksum(uint16_t val) {
    return (val >> 8) + val;
  }

  static uint8_t checksum(uint32_t val) {
    uint8_t sum = 0;
    for (size_t i = 0; i < sizeof(val); i++) {
      sum += (uint8_t)((uint32_t)((val >> (8*i)) & 0xFF));
    }
    return sum;
  }

  //// RECEIVED MESSAGE HANDLERS ////
  void setup_publisher(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);
    if (!publishers_.count(topic_info.topic_id)) {
      spdlog_info("[{0}] setup_publisher(topic_id: {1}, topic_name: {2})", session_id_.c_str(), topic_info.topic_id, topic_info.topic_name.c_str());
      PublisherPtr pub(new PublisherCore(topic_info));
      pub->alive_time_ = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
      callbacks_[topic_info.topic_id] = std::bind(&PublisherCore::handle, pub, std::placeholders::_1);
      publishers_[topic_info.topic_id] = pub;

      std::unique_lock<std::mutex> lock(Rostopic::topics_mutex_);
      if (!Rostopic::topics_.count(topic_info.topic_name)) {
        Rostopic::topics_[topic_info.topic_name] = RostopicPtr(new Rostopic(topic_info));
      }

      // fake connection for stop
      RostopicConnection connection;
      connection.rostopic_ = Rostopic::topics_[topic_info.topic_name];
      connection.rostopic_->ref_count_++;
      publishers_[topic_info.topic_id]->connection_ = connection;
    } else {
      publishers_[topic_info.topic_id]->alive_time_ = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_subscriber(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);
    if (!subscribers_.count(topic_info.topic_id)) {
      spdlog_info("[{0}] setup_subscriber(topic_id: {1}, topic_name: {2})", session_id_.c_str(), topic_info.topic_id, topic_info.topic_name.c_str());
      SubscriberPtr sub(new SubscriberCore(topic_info, std::bind(&Session::write_message, this, std::placeholders::_1, topic_info.topic_id)));
      sub->alive_time_ = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
      subscribers_[topic_info.topic_id] = sub;

      std::unique_lock<std::mutex> lock(Rostopic::topics_mutex_);
      if (!Rostopic::topics_.count(topic_info.topic_name)) {
        Rostopic::topics_[topic_info.topic_name] = RostopicPtr(new Rostopic(topic_info));
      }

      RostopicConnection connection;
      connection.rostopic_ = Rostopic::topics_[topic_info.topic_name];
      connection.rostopic_->ref_count_++;
      connection.id_ = connection.rostopic_->signal_->connect(std::bind(&SubscriberCore::handle, sub.get(), std::placeholders::_1));
      subscribers_[topic_info.topic_id]->connection_ = connection;
    } else {
      subscribers_[topic_info.topic_id]->alive_time_ = std::chrono::system_clock::now().time_since_epoch().count() * 1e-9;
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_service_server(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    std::unique_lock<std::mutex> lock(ServiceServerCore::services_mutex_);
    if (!ServiceServerCore::services_.count(topic_info.topic_name)) {
      spdlog_info("[{0}] setup_service_server(topic_id: {1}, topic_name: {2})", session_id_.c_str(), topic_info.topic_id, topic_info.topic_name.c_str());
      ServiceServerPtr srv(new ServiceServerCore(topic_info, std::bind(&Session::write_message, this, std::placeholders::_1, std::placeholders::_2)));
      callbacks_[topic_info.topic_id] = std::bind(&ServiceServerCore::handle, srv, std::placeholders::_1);
      ServiceServerCore::services_[topic_info.topic_name] = srv;
      ServiceServerCore::services_[topic_info.topic_name]->setTopicId(topic_info.topic_id);
      service_server_.push_back(topic_info.topic_name);
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }
  
  void setup_service_client(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    std::unique_lock<std::mutex> lock(ServiceServerCore::services_mutex_);
    if (ServiceServerCore::services_.count(topic_info.topic_name)) {
      if (!services_client_.count(topic_info.topic_id)) {
        spdlog_info("[{0}] setup_service_client(topic_id: {1}, topic_name: {2})", session_id_.c_str(), topic_info.topic_id, topic_info.topic_name.c_str());
        ServiceServerPtr service = ServiceServerCore::services_[topic_info.topic_name];
        ServiceClientPtr client(new ServiceClientCore(topic_info, std::bind(&Session::write_message, this, std::placeholders::_1, std::placeholders::_2)));
        client->setTopicId(topic_info.topic_id);
        client->client_connection_ = client->signal_->connect(std::bind(&ServiceServerCore::callback, service, std::placeholders::_1));
        client->service_connection_ = service->signal_->connect(std::bind(&ServiceClientCore::callback, client, std::placeholders::_1));
        client->destroy_connection_ = service->destroy_signal_->connect(std::bind(&Session::stop_service, this, std::placeholders::_1));
        callbacks_[topic_info.topic_id] = std::bind(&ServiceClientCore::handle, client, std::placeholders::_1);
        services_client_[topic_info.topic_id] = client;
      }
      topic_info.negotiated = true;
    } else {
      topic_info.negotiated = false;
    }
    
    handle_negotiated(topic_info);
  }

  void stop_service(std::string &topic_name) {
    spdlog_warn("[{0}] stop_service(topic_name: {1})", session_id_.c_str(), topic_name.c_str());
    std::map<uint32_t, ServiceClientPtr>::iterator it;
    for(it = services_client_.begin(); it != services_client_.end(); ) {
      ServiceClientPtr client = it->second;
      if (client->topic_name_ == topic_name) {
        client->signal_->disconnect(client->client_connection_);
        callbacks_.erase(it->first);
        services_client_.erase(it++);
      } else {
        it++;
      }
    }
  }

  void handle_log(tinyros::serialization::IStream& stream) {
    uint32_t length = stream.getLength();
    std::vector<uint8_t> message(length);
    memcpy(&message[0], stream.getData(), length);
    if (Rostopic::topics_.count(TINYROS_LOG_TOPIC)) {
      Rostopic::topics_[TINYROS_LOG_TOPIC]->signal_->emit(message);
    }
  }

  void handle_time(tinyros::serialization::IStream& stream) {
    roslib_msgs::Time time;
    time.data = tinyros::Time::now();

    size_t length = tinyros::serialization::serializationLength(time);
    std::vector<uint8_t> message(length);

    tinyros::serialization::OStream ostream(&message[0], length);
    tinyros::serialization::Serializer<roslib_msgs::Time>::write(ostream, time);

    write_message(message, tinyros_msgs::TopicInfo::ID_TIME);
  }

  void handle_negotiated(const tinyros_msgs::TopicInfo& topic_info) {
    if (!isudp_) {
      size_t length = tinyros::serialization::serializationLength(topic_info);
      std::vector<uint8_t> message(length);

      tinyros::serialization::OStream ostream(&message[0], length);
      tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::write(ostream, topic_info);

      write_message(message, tinyros_msgs::TopicInfo::ID_NEGOTIATED);
    }
  }

  void handle_rostopic_request(tinyros::serialization::IStream& stream) {
    std::string topic_list = "\n";
    std::map<std::string, RostopicPtr>::iterator it;
    for(it = Rostopic::topics_.begin(); it != Rostopic::topics_.end(); ) {
      topic_list += it->first + " [type:" + it->second->message_type_ + ", md5:" + it->second->md5sum_ + "]\n";
      it++;
    }
    roslib_msgs::String msg;
    msg.data = topic_list.c_str();
    size_t length = tinyros::serialization::serializationLength(msg);
    std::vector<uint8_t> message(length);
    tinyros::serialization::OStream ostream(&message[0], length);
    tinyros::serialization::Serializer<roslib_msgs::String>::write(ostream, msg);
    write_message(message, tinyros_msgs::TopicInfo::ID_ROSTOPIC_REQUEST);
  }

  void handle_rosservice_request(tinyros::serialization::IStream& stream) {
    std::string service_list = "\n\nservice_list:\n";
    std::map<std::string, ServiceServerPtr>::iterator it;
    for(it = ServiceServerCore::services_.begin(); it != ServiceServerCore::services_.end(); ) {
      service_list += "          " + it->first + " [" + it->second->message_type_ + "]\n";
      it++;
    }
    roslib_msgs::String msg;
    msg.data = service_list.c_str();
    size_t length = tinyros::serialization::serializationLength(msg);
    std::vector<uint8_t> message(length);
    tinyros::serialization::OStream ostream(&message[0], length);
    tinyros::serialization::Serializer<roslib_msgs::String>::write(ostream, msg);
    write_message(message, tinyros_msgs::TopicInfo::ID_ROSSERVICE_REQUEST);
  }

  void handle_session_id(tinyros::serialization::IStream& stream) {
    roslib_msgs::String session_info;
    tinyros::serialization::Serializer<roslib_msgs::String>::read(stream, session_info);
    spdlog_info("[{0}] Starting session...", session_info.data.c_str());
    session_id_ = session_info.data;
  }

  std::mutex async_write_mutex_;
  std::condition_variable async_write_cond_;
  AsyncWritebuffer async_write_buffers_;

  Socket socket_;
  enum { buffer_max = 65*1024 };
  
  std::mutex active_mutex_;
  bool active_;

  bool isudp_;

  std::map<uint32_t, std::function<void(tinyros::serialization::IStream&)> > callbacks_;
  std::map<uint32_t, PublisherPtr> publishers_;
  std::map<uint32_t, SubscriberPtr> subscribers_;
  std::vector<std::string> service_server_;
  std::map<uint32_t, ServiceClientPtr> services_client_;

  std::thread* message_in_thread_;

  std::thread* message_write_thread_;

  std::thread* require_check_thread_;
  bool require_check_running_;
};
}  // namespace

#endif  // TINY_ROS_SERVER_SESSION_H
