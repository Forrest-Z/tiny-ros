#ifndef _TINYROS_roslib_msgs_UInt8_h
#define _TINYROS_roslib_msgs_UInt8_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace roslib_msgs
{

  class UInt8 : public tinyros::Msg
  {
    public:
      typedef uint8_t _data_type;
      _data_type data;

    UInt8():
      data(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->data >> (8 * 0)) & 0xFF;
      offset += sizeof(this->data);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->data =  ((uint8_t) (*(inbuffer + offset)));
      offset += sizeof(this->data);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->data);
      return length;
    }

    virtual std::string getType(){ return "roslib_msgs/UInt8"; }
    virtual std::string getMD5(){ return "b8c6e02cc18bddcbace85025e42b55e5"; }

  };

}
#endif
