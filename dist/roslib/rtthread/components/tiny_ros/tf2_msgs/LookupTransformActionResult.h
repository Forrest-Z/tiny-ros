#ifndef _TINYROS_tf2_msgs_LookupTransformActionResult_h
#define _TINYROS_tf2_msgs_LookupTransformActionResult_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/actionlib_msgs/GoalStatus.h"
#include "tiny_ros/tf2_msgs/LookupTransformResult.h"

namespace tf2_msgs
{

  class LookupTransformActionResult : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef actionlib_msgs::GoalStatus _status_type;
      _status_type status;
      typedef tf2_msgs::LookupTransformResult _result_type;
      _result_type result;

    LookupTransformActionResult():
      header(),
      status(),
      result()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->status.serialize(outbuffer + offset);
      offset += this->result.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->status.deserialize(inbuffer + offset);
      offset += this->result.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->status.serializedLength();
      length += this->result.serializedLength();
      return length;
    }

    virtual tinyros::string getType(){ return "tf2_msgs/LookupTransformActionResult"; }
    virtual tinyros::string getMD5(){ return "5a8abe079c2126ea9966563c5cae6d29"; }

  };

}
#endif
