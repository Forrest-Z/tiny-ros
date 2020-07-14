#ifndef _TINYROS_std_msgs_String_h
#define _TINYROS_std_msgs_String_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace std_msgs
{

  class String : public tinyros::Msg
  {
    public:
      typedef tinyros::string _data_type;
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

    virtual tinyros::string getType(){ return "std_msgs/String"; }
    virtual tinyros::string getMD5(){ return "44b822292b4a9ed05e241aa225458f97"; }

  };

}
#endif
