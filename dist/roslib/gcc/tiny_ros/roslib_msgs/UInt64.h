#ifndef _TINYROS_roslib_msgs_UInt64_h
#define _TINYROS_roslib_msgs_UInt64_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace roslib_msgs
{

  class UInt64 : public tinyros::Msg
  {
    public:
      typedef uint64_t _data_type;
      _data_type data;

    UInt64():
      data(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->data >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->data >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->data >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (this->data >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (this->data >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (this->data >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (this->data >> (8 * 7)) & 0xFF;
      offset += sizeof(this->data);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->data =  ((uint64_t) (*(inbuffer + offset)));
      this->data |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->data |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->data |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      this->data |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      this->data |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      this->data |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      this->data |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      offset += sizeof(this->data);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->data);
      return length;
    }

    virtual std::string getType(){ return "roslib_msgs/UInt64"; }
    virtual std::string getMD5(){ return "5743ecd192777ee751f903f18c34befc"; }

  };

}
#endif
