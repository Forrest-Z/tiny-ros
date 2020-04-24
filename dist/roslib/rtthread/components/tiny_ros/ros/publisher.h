/*
 * File      : publisher.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */

#ifndef _TINYROS_PUBLISHER_H_
#define _TINYROS_PUBLISHER_H_

#include "tiny_ros/tinyros_msgs/TopicInfo.h"
#include "tiny_ros/ros/node_handle.h"

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
  int id_;
  NodeHandleBase_* nh_;

  // negotiated_ is set by NodeHandle when we negotiateTopics
  bool negotiated_;

private:
  int endpoint_;
};

}

#endif

