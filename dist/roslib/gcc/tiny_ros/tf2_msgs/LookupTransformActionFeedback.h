#ifndef _TINYROS_tf2_msgs_LookupTransformActionFeedback_h
#define _TINYROS_tf2_msgs_LookupTransformActionFeedback_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/actionlib_msgs/GoalStatus.h"
#include "tiny_ros/tf2_msgs/LookupTransformFeedback.h"

namespace tf2_msgs
{

  class LookupTransformActionFeedback : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef actionlib_msgs::GoalStatus _status_type;
      _status_type status;
      typedef tf2_msgs::LookupTransformFeedback _feedback_type;
      _feedback_type feedback;

    LookupTransformActionFeedback():
      header(),
      status(),
      feedback()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->status.serialize(outbuffer + offset);
      offset += this->feedback.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->status.deserialize(inbuffer + offset);
      offset += this->feedback.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->status.serializedLength();
      length += this->feedback.serializedLength();
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\": {";
      string_echo += this->header.echo();
      string_echo += "}, ";
      string_echo += "\"status\": {";
      string_echo += this->status.echo();
      string_echo += "}, ";
      string_echo += "\"feedback\": {";
      string_echo += this->feedback.echo();
      string_echo += "}";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tf2_msgs/LookupTransformActionFeedback"; }
    virtual std::string getMD5(){ return "f432c3fa33a3ac9586ace248b606de3a"; }

  };

}
#endif
