#ifndef _TINYROS_ROSTOPIC_SUBSCRIBERS_h
#define _TINYROS_ROSTOPIC_SUBSCRIBERS_h

#include "tiny_ros/ros.h"
#include "tiny_ros/roslib_msgs/Bool.h"
#include "tiny_ros/roslib_msgs/Byte.h"
#include "tiny_ros/roslib_msgs/Char.h"
#include "tiny_ros/roslib_msgs/Duration.h"
#include "tiny_ros/roslib_msgs/Empty.h"
#include "tiny_ros/roslib_msgs/Float32.h"
#include "tiny_ros/roslib_msgs/Float64.h"
#include "tiny_ros/roslib_msgs/Header.h"
#include "tiny_ros/roslib_msgs/Int16.h"
#include "tiny_ros/roslib_msgs/Int32.h"
#include "tiny_ros/roslib_msgs/Int64.h"
#include "tiny_ros/roslib_msgs/Int8.h"
#include "tiny_ros/roslib_msgs/String.h"
#include "tiny_ros/roslib_msgs/Time.h"
#include "tiny_ros/roslib_msgs/UInt16.h"
#include "tiny_ros/roslib_msgs/UInt32.h"
#include "tiny_ros/roslib_msgs/UInt64.h"
#include "tiny_ros/roslib_msgs/UInt8.h"
#include "tiny_ros/tinyros_hello/TinyrosHello.h"
#include "tiny_ros/tinyros_msgs/Log.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"

namespace tinyros
{
template<typename MsgT>
class EchoSubscriber: public tinyros::Subscriber_
{
public:
MsgT msg;
  virtual void callback(unsigned char* data)
  {
    MsgT tmsg;
    tmsg.deserialize(data);
    spdlog::get("logger")->info("{0}[{1}]->>{2}", topic_, getMsgType(), tmsg.echo());
  }
  virtual std::string getMsgType()
  {
    return this->msg.getType();
  }
  virtual std::string getMsgMD5()
  {
    return this->msg.getMD5();
  }
  virtual int getEndpointType()
  {
    return tinyros_msgs::TopicInfo::ID_SUBSCRIBER;
  }
};

static std::map<std::string, tinyros::Subscriber_*> rostopic_subscribers = {
    {"roslib_msgs/Bool", new EchoSubscriber<roslib_msgs::Bool>()},
    {"roslib_msgs/Byte", new EchoSubscriber<roslib_msgs::Byte>()},
    {"roslib_msgs/Char", new EchoSubscriber<roslib_msgs::Char>()},
    {"roslib_msgs/Duration", new EchoSubscriber<roslib_msgs::Duration>()},
    {"roslib_msgs/Empty", new EchoSubscriber<roslib_msgs::Empty>()},
    {"roslib_msgs/Float32", new EchoSubscriber<roslib_msgs::Float32>()},
    {"roslib_msgs/Float64", new EchoSubscriber<roslib_msgs::Float64>()},
    {"roslib_msgs/Header", new EchoSubscriber<roslib_msgs::Header>()},
    {"roslib_msgs/Int16", new EchoSubscriber<roslib_msgs::Int16>()},
    {"roslib_msgs/Int32", new EchoSubscriber<roslib_msgs::Int32>()},
    {"roslib_msgs/Int64", new EchoSubscriber<roslib_msgs::Int64>()},
    {"roslib_msgs/Int8", new EchoSubscriber<roslib_msgs::Int8>()},
    {"roslib_msgs/String", new EchoSubscriber<roslib_msgs::String>()},
    {"roslib_msgs/Time", new EchoSubscriber<roslib_msgs::Time>()},
    {"roslib_msgs/UInt16", new EchoSubscriber<roslib_msgs::UInt16>()},
    {"roslib_msgs/UInt32", new EchoSubscriber<roslib_msgs::UInt32>()},
    {"roslib_msgs/UInt64", new EchoSubscriber<roslib_msgs::UInt64>()},
    {"roslib_msgs/UInt8", new EchoSubscriber<roslib_msgs::UInt8>()},
    {"tinyros_hello/TinyrosHello", new EchoSubscriber<tinyros_hello::TinyrosHello>()},
    {"tinyros_msgs/Log", new EchoSubscriber<tinyros_msgs::Log>()},
    {"tinyros_msgs/TopicInfo", new EchoSubscriber<tinyros_msgs::TopicInfo>()},
};

}
#endif

