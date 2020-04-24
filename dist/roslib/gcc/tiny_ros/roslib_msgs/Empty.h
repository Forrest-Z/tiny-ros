#ifndef _TINYROS_roslib_msgs_Empty_h
#define _TINYROS_roslib_msgs_Empty_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace roslib_msgs
{

  class Empty : public tinyros::Msg
  {
    public:

    Empty()
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

    virtual std::string getType(){ return "roslib_msgs/Empty"; }
    virtual std::string getMD5(){ return "a4758a8dc954913cf9a8cc8864ad5209"; }

  };

}
#endif
