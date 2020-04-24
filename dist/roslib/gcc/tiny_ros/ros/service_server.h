/*
 * File      : service_server.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */


#ifndef _TINYROS_SERVICE_SERVER_H_
#define _TINYROS_SERVICE_SERVER_H_

#include "tiny_ros/tinyros_msgs/TopicInfo.h"

#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"

namespace tinyros
{

template<typename MReq , typename MRes, typename ObjT = void>
class ServiceServer : public Subscriber_
{
public:
  typedef void(ObjT::*CallbackT)(const MReq&,  MRes&);

  ServiceServer(std::string topic_name, CallbackT cb, ObjT* obj) :
    pub(topic_name, &resp, tinyros_msgs::TopicInfo::ID_SERVICE_SERVER + tinyros_msgs::TopicInfo::ID_PUBLISHER),
    obj_(obj)
  {
    this->negotiated_ = false;
    this->srv_flag_ = true;
    this->topic_ = topic_name;
    this->cb_ = cb;
  }

  // these refer to the subscriber
  virtual void callback(unsigned char *data)
  {
    MReq treq; MRes tresp;
    treq.deserialize(data);
    (obj_->*cb_)(treq, tresp);
    tresp.setID(treq.getID());
    pub.publish(&tresp);
  }
  virtual std::string getMsgType()
  {
    return this->req.getType();
  }
  virtual std::string getMsgMD5()
  {
    return this->req.getMD5();
  }
  virtual int getEndpointType()
  {
    return tinyros_msgs::TopicInfo::ID_SERVICE_SERVER + tinyros_msgs::TopicInfo::ID_SUBSCRIBER;
  }
  
  virtual bool negotiated()
  { 
    return (negotiated_ && pub.negotiated_); 
  }

  MReq req;
  MRes resp;
  Publisher pub;
private:
  CallbackT cb_;
  ObjT* obj_;
};

template<typename MReq , typename MRes>
class ServiceServer<MReq, MRes, void> : public Subscriber_
{
public:
  typedef void(*CallbackT)(const MReq&,  MRes&);

  ServiceServer(std::string topic_name, CallbackT cb) :
    pub(topic_name, &resp, tinyros_msgs::TopicInfo::ID_SERVICE_SERVER + tinyros_msgs::TopicInfo::ID_PUBLISHER)
  {
    this->negotiated_ = false;
    this->srv_flag_ = true;
    this->topic_ = topic_name;
    this->cb_ = cb;
  }

  // these refer to the subscriber
  virtual void callback(unsigned char *data)
  {
    MReq treq; MRes tresp;
    treq.deserialize(data);
    cb_(treq, tresp);
    tresp.setID(treq.getID());
    pub.publish(&tresp);
  }
  virtual std::string getMsgType()
  {
    return this->req.getType();
  }
  virtual std::string getMsgMD5()
  {
    return this->req.getMD5();
  }
  virtual int getEndpointType()
  {
    return tinyros_msgs::TopicInfo::ID_SERVICE_SERVER + tinyros_msgs::TopicInfo::ID_SUBSCRIBER;
  }

  MReq req;
  MRes resp;
  Publisher pub;
private:
  CallbackT cb_;
};

}

#endif
