#ifndef _TINYROS_std_msgs_MultiArrayDimension_h
#define _TINYROS_std_msgs_MultiArrayDimension_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace std_msgs
{

  class MultiArrayDimension : public tinyros::Msg
  {
    public:
      typedef tinyros::string _label_type;
      _label_type label;
      typedef uint32_t _size_type;
      _size_type size;
      typedef uint32_t _stride_type;
      _stride_type stride;

    MultiArrayDimension():
      label(""),
      size(0),
      stride(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      uint32_t length_label = this->label.size();
      varToArr(outbuffer + offset, length_label);
      offset += 4;
      memcpy(outbuffer + offset, this->label.c_str(), length_label);
      offset += length_label;
      *(outbuffer + offset + 0) = (this->size >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->size >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->size >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->size >> (8 * 3)) & 0xFF;
      offset += sizeof(this->size);
      *(outbuffer + offset + 0) = (this->stride >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->stride >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->stride >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->stride >> (8 * 3)) & 0xFF;
      offset += sizeof(this->stride);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t length_label;
      arrToVar(length_label, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_label; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_label-1]=0;
      this->label = (char *)(inbuffer + offset-1);
      offset += length_label;
      this->size =  ((uint32_t) (*(inbuffer + offset)));
      this->size |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->size |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->size |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->size);
      this->stride =  ((uint32_t) (*(inbuffer + offset)));
      this->stride |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->stride |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->stride |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->stride);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_label = this->label.size();
      length += 4;
      length += length_label;
      length += sizeof(this->size);
      length += sizeof(this->stride);
      return length;
    }

    virtual tinyros::string getType(){ return "std_msgs/MultiArrayDimension"; }
    virtual tinyros::string getMD5(){ return "c2aacf83d49c7aa4a8504bd32158e990"; }

  };

}
#endif
