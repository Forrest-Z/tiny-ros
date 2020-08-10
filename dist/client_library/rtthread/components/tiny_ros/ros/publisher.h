#ifndef _TINYROS_PUBLISHER_H_
#define _TINYROS_PUBLISHER_H_

#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/ros/node_handle_base.h"

namespace tinyros
{

/* Generic Publisher */
class Publisher
{
public:
  Publisher(tinyros::string topic_name, Msg * msg, int endpoint = tinyros_msgs::TopicInfo::ID_PUBLISHER) :
    topic_(topic_name),
    msg_(msg),
    nh_(NULL),
    negotiated_(false),
    endpoint_(endpoint) { }

  int publish(const Msg * msg, bool islog = false)
  {
    if (nh_ != NULL) {
      return nh_->publish(id_, msg, islog);
    } else {
      return -1;
    }
  }

  int getEndpointType()
  {
    return endpoint_;
  }

  bool negotiated()
  {
    return negotiated_;
  }

  tinyros::string topic_;
  Msg *msg_;
  // id_ and no_ are set by NodeHandle when we advertise
  uint32_t id_;
  NodeHandleBase_* nh_;

  // negotiated_ is set by NodeHandle when we negotiateTopics
  bool negotiated_;

private:
  int endpoint_;
};

}

#endif

