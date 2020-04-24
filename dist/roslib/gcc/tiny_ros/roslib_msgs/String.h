#ifndef _TINYROS_roslib_msgs_String_h
#define _TINYROS_roslib_msgs_String_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace roslib_msgs
{

  class String : public tinyros::Msg
  {
    public:
      typedef std::string _data_type;
      _data_type data;

    String():
      data("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      uint32_t length_data = this->data.size();
      varToArr(outbuffer + offset, length_data);
      offset += 4;
      memcpy(outbuffer + offset, this->data.c_str(), length_data);
      offset += length_data;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t length_data;
      arrToVar(length_data, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_data; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_data-1]=0;
      this->data = (char *)(inbuffer + offset-1);
      offset += length_data;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_data = this->data.size();
      length += 4;
      length += length_data;
      return length;
    }

    virtual std::string getType(){ return "roslib_msgs/String"; }
    virtual std::string getMD5(){ return "5803b93aeb3e71858cec3497cca77163"; }

  };

}
#endif
