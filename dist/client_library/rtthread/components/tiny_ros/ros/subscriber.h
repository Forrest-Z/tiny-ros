#ifndef TINYROS_SUBSCRIBER_H_
#define TINYROS_SUBSCRIBER_H_

#include "tiny_ros/tinyros_msgs/TopicInfo.h"

namespace tinyros
{

/* Base class for objects subscribers. */
class Subscriber_
{
public:
  virtual void callback(unsigned char *data) = 0;
  virtual int getEndpointType() = 0;

  // id_ is set by NodeHandle when we advertise
  uint32_t id_;

  virtual tinyros::string getMsgType() = 0;
  virtual tinyros::string getMsgMD5() = 0;
  virtual bool negotiated() { return negotiated_; }
  tinyros::string topic_;

  // negotiated_ is set by NodeHandle when we negotiateTopics
  bool negotiated_;
  bool srv_flag_;
};

/* Bound function subscriber. */
template<typename MsgT, typename ObjT = void>
class Subscriber: public Subscriber_
{
public:
  typedef void(ObjT::*CallbackT)(const MsgT&);
  MsgT msg;

  Subscriber(tinyros::string topic_name, CallbackT cb, ObjT* obj, int endpoint = tinyros_msgs::TopicInfo::ID_SUBSCRIBER) :
    cb_(cb),
    obj_(obj),
    endpoint_(endpoint) {
    topic_ = topic_name;
    negotiated_ = false;
    srv_flag_ = false;
  }

  virtual void callback(unsigned char* data)
  {
    MsgT tmsg;
    tmsg.deserialize(data);
    (obj_->*cb_)(tmsg);
  }

  virtual tinyros::string getMsgType()
  {
    return this->msg.getType();
  }
  virtual tinyros::string getMsgMD5()
  {
    return this->msg.getMD5();
  }
  virtual int getEndpointType()
  {
    return endpoint_;
  }

private:
  CallbackT cb_;
  ObjT* obj_;
  int endpoint_;
};

/* Standalone function subscriber. */
template<typename MsgT>
class Subscriber<MsgT, void>: public Subscriber_
{
public:
  typedef void(*CallbackT)(const MsgT&);
  MsgT msg;

  Subscriber(tinyros::string topic_name, CallbackT cb, int endpoint = tinyros_msgs::TopicInfo::ID_SUBSCRIBER) :
    cb_(cb),
    endpoint_(endpoint) {
    topic_ = topic_name;
    negotiated_ = false;
    srv_flag_ = false;
  }

  virtual void callback(unsigned char* data)
  {
    MsgT tmsg;
    tmsg.deserialize(data);
    this->cb_(tmsg);
  }

  virtual tinyros::string getMsgType()
  {
    return this->msg.getType();
  }
  virtual tinyros::string getMsgMD5()
  {
    return this->msg.getMD5();
  }
  virtual int getEndpointType()
  {
    return endpoint_;
  }

private:
  CallbackT cb_;
  int endpoint_;
};

}

#endif

