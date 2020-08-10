#ifndef _TINYROS_SERVICE_CLIENT_H_
#define _TINYROS_SERVICE_CLIENT_H_
#include <stdint.h>
#include <mutex>
#include <chrono>
#include "tiny_ros/ros/log.h"
#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/ros/publisher.h"
#include "tiny_ros/ros/subscriber.h"

namespace tinyros
{

template<typename MReq , typename MRes>
class ServiceClient : public Subscriber_
{
public:
  ServiceClient(std::string topic_name) :
    pub(topic_name, &req, tinyros_msgs::TopicInfo::ID_SERVICE_CLIENT + tinyros_msgs::TopicInfo::ID_PUBLISHER)
  {
    this->negotiated_ = false;
    this->srv_flag_ = true;
    this->topic_ = topic_name;
    this->call_resp = NULL;
    this->call_req = NULL;
  }
  virtual bool call(MReq & request, MRes & response, int duration = 3)
  {
    std::unique_lock<std::mutex> g_lock(g_mutex_);
    std::unique_lock<std::mutex> lock(mutex_);
    if (!pub.nh_->ok()) {
      call_req = NULL;
      call_resp = NULL;
      return false;
    }
    call_req = &request;
    call_resp = &response;

    {
      std::unique_lock<std::mutex> gg_lock(gg_mutex_);
      call_req->setID(gg_id_++);
    }
    
    if (pub.publish(&request) <= 0) {
      call_req = NULL; call_resp = NULL;
      return false;
    }
    if (cond_.wait_until(lock, std::chrono::system_clock::now() +
      std::chrono::milliseconds(duration * 1000)) == std::cv_status::timeout) {
      log_warn("Service[%s] call_req.id: %u, call timeout", this->topic_.c_str(), call_req->getID());
      call_req = NULL; call_resp = NULL;
      return false;
    }
    call_req = NULL; call_resp = NULL;
    return true;
  }

  // these refer to the subscriber
  virtual void callback(unsigned char *data)
  {
    if (call_resp && call_req) {
      std::unique_lock<std::mutex> lock(mutex_);
      if (call_resp && call_req) {
        uint32_t req_id  = call_req->getID();
        uint32_t resp_id =  ((uint32_t) (*(data + 0)));
        resp_id |= ((uint32_t) (*(data + 1))) << (8 * 1);
        resp_id |= ((uint32_t) (*(data + 2))) << (8 * 2);
        resp_id |= ((uint32_t) (*(data + 3))) << (8 * 3);

        if (req_id == resp_id) {
          call_resp->deserialize(data);
          cond_.notify_all();
        }
      }
    }
  }
  virtual std::string getMsgType()
  {
    return this->resp.getType();
  }
  virtual std::string getMsgMD5()
  {
    return this->resp.getMD5();
  }
  virtual int getEndpointType()
  {
    return tinyros_msgs::TopicInfo::ID_SERVICE_CLIENT + tinyros_msgs::TopicInfo::ID_SUBSCRIBER;
  }
  
  virtual bool negotiated()
  { 
    return (negotiated_ && pub.negotiated_); 
  }

  MReq req;
  MRes resp;
  MReq * call_req;
  MRes * call_resp;
  Publisher pub;
  std::mutex mutex_;
  std::mutex g_mutex_;
  std::condition_variable cond_;
  static uint32_t gg_id_;
  static std::mutex gg_mutex_;
};

template<typename MReq , typename MRes>
uint32_t ServiceClient<MReq , MRes>::gg_id_ = 1;

template<typename MReq , typename MRes>
std::mutex ServiceClient<MReq , MRes>::gg_mutex_;
}

#endif
