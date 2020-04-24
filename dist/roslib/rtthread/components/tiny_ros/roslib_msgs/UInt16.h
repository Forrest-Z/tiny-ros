#ifndef _TINYROS_roslib_msgs_UInt16_h
#define _TINYROS_roslib_msgs_UInt16_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace roslib_msgs
{

  class UInt16 : public tinyros::Msg
  {
    public:
      typedef uint16_t _data_type;
      _data_type data;

    UInt16():
      data(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->data >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data >> (8 * 1)) & 0xFF;
      offset += sizeof(this->data);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->data =  ((uint16_t) (*(inbuffer + offset)));
      this->data |= ((uint16_t) (*(inbuffer + offset + 1))) << (8 * 1);
      offset += sizeof(this->data);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->data);
      return length;
    }

    virtual tinyros::string getType(){ return "roslib_msgs/UInt16"; }
    virtual tinyros::string getMD5(){ return "f7fc05ba72b1e0b94d9852ae4a5c54cb"; }

  };

}
#endif
