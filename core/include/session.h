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
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include <boost/function.hpp>
#include <boost/log/trivial.hpp>
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/roslib_msgs/Time.h"
#include "tiny_ros/roslib_msgs/String.h"
#include "serialization.h"
#include "topic_handlers.h"
#include "serialization.h"

namespace tinyros
{
#define REQUEST_TOPICS_TIMER (1000*1000)

typedef std::vector<uint8_t> Buffer;
typedef boost::shared_ptr<Buffer> BufferPtr;
typedef std::deque<BufferPtr> AsyncWritebuffer;

template<typename Socket>
class Session : boost::noncopyable
{
public:
  Session(boost::asio::io_service& io_service, bool isudp = false)
    : socket_(io_service)
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
    callbacks_[TopicInfo::ID_PUBLISHER] = boost::bind(&Session::setup_publisher, this, _1);
    callbacks_[TopicInfo::ID_SUBSCRIBER] = boost::bind(&Session::setup_subscriber, this, _1);
    callbacks_[TopicInfo::ID_SERVICE_SERVER+TopicInfo::ID_PUBLISHER] = boost::bind(&Session::setup_service_server_publisher, this, _1);
    callbacks_[TopicInfo::ID_SERVICE_SERVER+TopicInfo::ID_SUBSCRIBER] = boost::bind(&Session::setup_service_server_subscriber, this, _1);
    callbacks_[TopicInfo::ID_SERVICE_CLIENT+TopicInfo::ID_PUBLISHER] = boost::bind(&Session::setup_service_client_publisher, this, _1);
    callbacks_[TopicInfo::ID_SERVICE_CLIENT+TopicInfo::ID_SUBSCRIBER] = boost::bind(&Session::setup_service_client_subscriber, this, _1);
    callbacks_[TopicInfo::ID_ROSTOPIC_REQUEST] = boost::bind(&Session::handle_rostopic_request, this, _1);
    callbacks_[TopicInfo::ID_ROSSERVICE_REQUEST] = boost::bind(&Session::handle_rosservice_request, this, _1);
    callbacks_[TopicInfo::ID_LOG] = boost::bind(&Session::handle_log, this, _1);
    callbacks_[TopicInfo::ID_TIME] = boost::bind(&Session::handle_time, this, _1);
    callbacks_[TopicInfo::ID_SESSION_ID] = boost::bind(&Session::handle_session_id, this, _1);

    active_ = true;
    
    require_check_running_ = true;
 
    message_write_thread_ = new boost::thread(boost::bind(&Session::write_completion_cb, this));
    if (isudp_) {
      message_in_thread_ = new boost::thread(boost::bind(&Session::read_message_sync_udp, this));
    } else {
      message_in_thread_ = new boost::thread(boost::bind(&Session::read_message_sync, this));
      require_check_thread_ = new boost::thread(boost::bind(&Session::request_topics, this));
    }
  }

  void stop()
  {
    {
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " begin.";
      boost::unique_lock<boost::mutex> lock(active_mutex_);
      if (!active_) {
        BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " end.";
        return;
      }
      if (require_check_thread_) {
        BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " require_check_thread interrupt begin.";
        require_check_running_ = false;
        require_check_thread_->interrupt();
        require_check_thread_->timed_join(boost::posix_time::milliseconds(500));
        delete require_check_thread_;
        require_check_thread_ = nullptr;
        BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " require_check_thread interrupt end.";
      }
      active_ = false;
    }
    
    {
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " topic_connections clear begin.";
      boost::unique_lock<boost::mutex> lock(Rostopic::topics_mutex_);
      for (uint32_t i=0; i<topic_connections_.size(); i++) {
        RostopicConnection connection = topic_connections_.at(i);
        connection.connection_.disconnect();
        connection.rostopic_->ref_count_--;

        if (connection.rostopic_->ref_count_ <=0) {
          Rostopic::topics_.erase(connection.rostopic_->topic_name_);
        }
      }
      topic_connections_.clear();
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " topic_connections clear end.";
    }

    {
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " services clear begin.";
      boost::unique_lock<boost::mutex> lock(ServiceServer::services_mutex_);
      for (uint32_t i=0; i<service_server_.size(); i++) {
        if (ServiceServer::services_.count(service_server_.at(i))) {
          ServiceServerPtr service = ServiceServer::services_[service_server_.at(i)];
          (*(service->destroy_signal_))(service_server_.at(i));
          ServiceServer::services_.erase(service_server_.at(i));
        }
      }
      
      std::map<uint32_t, ServiceClientPtr>::iterator it;
      for(it = services_client_.begin(); it != services_client_.end(); ) {
        ServiceClientPtr client = it->second;
        client->client_connection_.disconnect();
        client->service_connection_.disconnect();
        client->destroy_connection_.disconnect();
        it++;
      }
      service_server_.clear();
      services_client_.clear();
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " services clear end.";
    }

    // Reset the state of the session, dropping any publishers or subscribers
    // we currently know about from this client.
    callbacks_.clear();
    subscribers_.clear();
    publishers_.clear();
    
    // Close the socket.
    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " Close the socket begin.";
    socket_.close();
    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " Close the socket end.";

    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " clear async_write_buffers begin.";
    boost::unique_lock<boost::mutex> lock(async_write_mutex_);
    async_write_buffers_.clear();
    async_write_cond_.notify_all();
    lock.unlock();
    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " clear async_write_buffers end.";
    
    if (message_write_thread_) {
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " message_write_thread interrupt begin.";
      message_write_thread_->interrupt();
      message_write_thread_->timed_join(boost::posix_time::milliseconds(500));
      delete message_write_thread_;
      message_write_thread_ = nullptr;
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " message_write_thread interrupt end.";
    }

    if (message_in_thread_) {
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " message_in_thread interrupt begin.";
      message_in_thread_->interrupt();
      message_in_thread_->timed_join(boost::posix_time::milliseconds(500));
      delete message_in_thread_;
      message_in_thread_ = nullptr;
      BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " message_in_thread interrupt end.";
    }
    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ << " end.";
  }

  bool is_active()
  {
    return active_;
  }

private:
  void read_message_sync_udp() {
    while (is_active()) {
      uint8_t *message_in = (uint8_t*)calloc(buffer_max, sizeof(uint8_t));
      int32_t rv = socket_.read_some(boost::asio::buffer(message_in, buffer_max));
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
              BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ 
                << " Received message with unrecognized topicId ("<<topic<<").";
            }
          } else {
            BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ 
              << " Rejecting message on topicId="<<topic<<", length=" <<bytes<<" with bad checksum.";
          }
        } while(0);
      }
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
      
      boost::system::error_code ec;
      rv = boost::asio::read(socket_, boost::asio::buffer(message_tmp, len), ec);
      if (ec == boost::asio::error::eof) {
        BOOST_LOG_TRIVIAL(fatal) << "[" << session_id_<<"] "
          << __FUNCTION__ << " Socket asio error, closing socket: " << ec;
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
              BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ 
                << " Received message with unrecognized topicId ("<<topic<<").";
            }
          } else {
            BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< __FUNCTION__ 
              << " Rejecting message on topicId="<<topic<<", length=" <<total_bytes<<" with bad checksum.";
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
 
    boost::unique_lock<boost::mutex> lock(async_write_mutex_);
    async_write_buffers_.push_back(buffer_ptr);
    async_write_cond_.notify_one();
  }

  void write_completion_cb() {
    while (is_active()) {
      boost::unique_lock<boost::mutex> lock(async_write_mutex_);
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
        boost::system::error_code error;
        boost::asio::write(socket_, boost::asio::buffer(*buffer_ptr), error);
        if (error) {
          if (error == boost::system::errc::io_error) {
            BOOST_LOG_TRIVIAL(fatal) << "[" << session_id_<<"] "<< __FUNCTION__ << " Socket write operation returned IO error.";
          } else if (error == boost::system::errc::no_such_device) {
            BOOST_LOG_TRIVIAL(fatal) << "[" << session_id_<<"] "<< __FUNCTION__ << " Socket write operation returned no device.";
          } else {
            BOOST_LOG_TRIVIAL(fatal) << "[" << session_id_<<"] "<< __FUNCTION__ << " Unknown error returned during write operation: " << error;
          }
          if (is_active()) {
            boost::thread tid(boost::bind(&Session::stop, this));
            tid.detach();
            break;
          }
        }
      }
    }
  }

  //// HELPERS ////
  void request_topics() {
    while(require_check_running_) {
      std::vector<uint8_t> message(0);
      write_message(message, tinyros_msgs::TopicInfo::ID_PUBLISHER);
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
      BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] " << "setup_publisher(topic_id:" 
        << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
      PublisherPtr pub(new Publisher(topic_info));
      callbacks_[topic_info.topic_id] = boost::bind(&Publisher::handle, pub, _1);
      publishers_[topic_info.topic_id] = pub;

      boost::unique_lock<boost::mutex> lock(Rostopic::topics_mutex_);
      if (!Rostopic::topics_.count(topic_info.topic_name)) {
        Rostopic::topics_[topic_info.topic_name] = RostopicPtr(new Rostopic(topic_info));
      }

      // fake connection for stop
      RostopicConnection connection;
      connection.rostopic_ = Rostopic::topics_[topic_info.topic_name];
      connection.rostopic_->ref_count_++;
      topic_connections_.push_back(connection);
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_subscriber(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);
    if (!subscribers_.count(topic_info.topic_id)) {
      BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] "<< "setup_subscriber(topic_id:" 
        << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
      SubscriberPtr sub(new Subscriber(topic_info, boost::bind(&Session::write_message, this, _1, topic_info.topic_id)));
      subscribers_[topic_info.topic_id] = sub;

      boost::unique_lock<boost::mutex> lock(Rostopic::topics_mutex_);
      if (!Rostopic::topics_.count(topic_info.topic_name)) {
        Rostopic::topics_[topic_info.topic_name] = RostopicPtr(new Rostopic(topic_info));
      }

      RostopicConnection connection;
      connection.rostopic_ = Rostopic::topics_[topic_info.topic_name];
      connection.rostopic_->ref_count_++;
      connection.connection_ = connection.rostopic_->signal_->connect(boost::bind(&Subscriber::handle, sub, _1));
      topic_connections_.push_back(connection);
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_service_server_publisher(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    boost::unique_lock<boost::mutex> lock(ServiceServer::services_mutex_);
    if (!ServiceServer::services_.count(topic_info.topic_name)) {
      BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] "<< "setup_service_server_publisher(topic_id:" 
        << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
      ServiceServerPtr srv(new ServiceServer(topic_info, boost::bind(&Session::write_message, this, _1, _2)));
      callbacks_[topic_info.topic_id] = boost::bind(&ServiceServer::handle, srv, _1);
      ServiceServer::services_[topic_info.topic_name] = srv;
      ServiceServer::services_[topic_info.topic_name]->setTopicId(topic_info.topic_id);
      service_server_.push_back(topic_info.topic_name);
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_service_server_subscriber(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    boost::unique_lock<boost::mutex> lock(ServiceServer::services_mutex_);
    if (!ServiceServer::services_.count(topic_info.topic_name)) {
      BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] "<< "setup_service_server_subscriber(topic_id:"
        << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
      ServiceServerPtr srv(new ServiceServer(topic_info, boost::bind(&Session::write_message, this, _1, _2)));
      callbacks_[topic_info.topic_id] = boost::bind(&ServiceServer::handle, srv, _1);
      ServiceServer::services_[topic_info.topic_name] = srv;
      ServiceServer::services_[topic_info.topic_name]->setTopicId(topic_info.topic_id);
      service_server_.push_back(topic_info.topic_name);
    }
    
    topic_info.negotiated = true;
    
    handle_negotiated(topic_info);
  }

  void setup_service_client_publisher(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    boost::unique_lock<boost::mutex> lock(ServiceServer::services_mutex_);
    if (ServiceServer::services_.count(topic_info.topic_name)) {
      if (!services_client_.count(topic_info.topic_id)) {
        BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] " << "setup_service_client_publisher(topic_id:" 
          << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
        ServiceServerPtr service = ServiceServer::services_[topic_info.topic_name];
        ServiceClientPtr client(new ServiceClient(topic_info, boost::bind(&Session::write_message, this, _1, _2)));
        client->setTopicId(topic_info.topic_id);
        client->client_connection_ = client->signal_->connect(boost::bind(&ServiceServer::callback, service, _1));
        client->service_connection_ = service->signal_->connect(boost::bind(&ServiceClient::callback, client, _1));
        client->destroy_connection_ = service->destroy_signal_->connect(boost::bind(&Session::stop_service, this, _1));
        callbacks_[topic_info.topic_id] = boost::bind(&ServiceClient::handle, client, _1);
        services_client_[topic_info.topic_id] = client;
      }
      topic_info.negotiated = true;
    } else {
      topic_info.negotiated = false;
    }
    
    handle_negotiated(topic_info);
  }

  void setup_service_client_subscriber(tinyros::serialization::IStream& stream) {
    tinyros_msgs::TopicInfo topic_info;
    tinyros::serialization::Serializer<tinyros_msgs::TopicInfo>::read(stream, topic_info);

    boost::unique_lock<boost::mutex> lock(ServiceServer::services_mutex_);
    if (ServiceServer::services_.count(topic_info.topic_name)) {
      if (!services_client_.count(topic_info.topic_id)) {
        BOOST_LOG_TRIVIAL(info) << "[" << session_id_<<"] "<< "setup_service_client_subscriber(topic_id:" 
          << topic_info.topic_id << ", topic_name:" << topic_info.topic_name << ")";
        ServiceServerPtr service = ServiceServer::services_[topic_info.topic_name];
        ServiceClientPtr client(new ServiceClient(topic_info, boost::bind(&Session::write_message, this, _1, _2)));
        client->setTopicId(topic_info.topic_id);
        client->client_connection_ = client->signal_->connect(boost::bind(&ServiceServer::callback, service, _1));
        client->service_connection_ = service->signal_->connect(boost::bind(&ServiceClient::callback, client, _1));
        client->destroy_connection_ = service->destroy_signal_->connect(boost::bind(&Session::stop_service, this, _1));
        callbacks_[topic_info.topic_id] = boost::bind(&ServiceClient::handle, client, _1);
        services_client_[topic_info.topic_id] = client;
      }
      topic_info.negotiated = true;
    } else {
      topic_info.negotiated = false;
    }
      
    handle_negotiated(topic_info);
  }

  void stop_service(std::string &topic_name) {
    BOOST_LOG_TRIVIAL(warning) << "[" << session_id_<<"] "<< "stop_service topic_name: " << topic_name;
    std::map<uint32_t, ServiceClientPtr>::iterator it;
    for(it = services_client_.begin(); it != services_client_.end(); ) {
      ServiceClientPtr client = it->second;
      if (client->topic_name_ == topic_name) {
        client->client_connection_.disconnect();
        client->service_connection_.disconnect();
        client->destroy_connection_.disconnect();
        callbacks_.erase(it->first);
        services_client_.erase(it++);
      } else {
        it++;
      }
    }
  }

  void handle_log(tinyros::serialization::IStream& stream) {
    tinyros_msgs::Log l;
    tinyros::serialization::Serializer<tinyros_msgs::Log>::read(stream, l);
    if(l.level == tinyros_msgs::Log::ROSDEBUG) BOOST_LOG_TRIVIAL(debug) << l.msg;
    else if(l.level == tinyros_msgs::Log::ROSINFO) BOOST_LOG_TRIVIAL(info) << l.msg;
    else if(l.level == tinyros_msgs::Log::ROSWARN) BOOST_LOG_TRIVIAL(warning) << l.msg;
    else if(l.level == tinyros_msgs::Log::ROSERROR) BOOST_LOG_TRIVIAL(error) << l.msg;
    else if(l.level == tinyros_msgs::Log::ROSFATAL) BOOST_LOG_TRIVIAL(fatal) << l.msg;
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
    std::string topic_list = "\ntopic_list:\n";
    std::map<std::string, RostopicPtr>::iterator it;
    for(it = Rostopic::topics_.begin(); it != Rostopic::topics_.end(); ) {
      topic_list += "        " + it->first + " [" + it->second->message_type_ + "]\n";
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
    for(it = ServiceServer::services_.begin(); it != ServiceServer::services_.end(); ) {
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
    BOOST_LOG_TRIVIAL(warning) << "[" << session_info.data <<"] "<< "Starting session...";
    session_id_ = session_info.data;
  }

  boost::mutex async_write_mutex_;
  boost::condition_variable async_write_cond_;
  AsyncWritebuffer async_write_buffers_;

  Socket socket_;
  enum { buffer_max = 65*1024 };
  
  boost::mutex active_mutex_;
  bool active_;

  bool isudp_;

  std::map<uint32_t, boost::function<void(tinyros::serialization::IStream&)> > callbacks_;
  std::map<uint32_t, PublisherPtr> publishers_;
  std::map<uint32_t, SubscriberPtr> subscribers_;
  std::vector<RostopicConnection> topic_connections_;
  std::vector<std::string> service_server_;
  std::map<uint32_t, ServiceClientPtr> services_client_;

  std::string session_id_;

  boost::thread* message_in_thread_;

  boost::thread* message_write_thread_;

  boost::thread* require_check_thread_;
  bool require_check_running_;
};
}  // namespace

#endif  // TINY_ROS_SERVER_SESSION_H
