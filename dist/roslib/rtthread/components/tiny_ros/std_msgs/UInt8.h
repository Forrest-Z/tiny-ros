#ifndef _TINYROS_std_msgs_UInt8_h
#define _TINYROS_std_msgs_UInt8_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace std_msgs
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

    virtual tinyros::string getType(){ return "std_msgs/UInt8"; }
    virtual tinyros::string getMD5(){ return "6f90555707d539e508484b884b2acc65"; }

  };

}
#endif
