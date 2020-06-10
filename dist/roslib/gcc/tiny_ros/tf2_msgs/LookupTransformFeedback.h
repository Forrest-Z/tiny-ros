#ifndef _TINYROS_tf2_msgs_LookupTransformFeedback_h
#define _TINYROS_tf2_msgs_LookupTransformFeedback_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace tf2_msgs
{

  class LookupTransformFeedback : public tinyros::Msg
  {
    public:

    LookupTransformFeedback()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tf2_msgs/LookupTransformFeedback"; }
    virtual std::string getMD5(){ return "e6217f8a8e77aa218a8d6f594d08ba08"; }

  };

}
#endif
