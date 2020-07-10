#ifndef _TINYROS_std_msgs_MultiArrayLayout_h
#define _TINYROS_std_msgs_MultiArrayLayout_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/MultiArrayDimension.h"

namespace std_msgs
{

  class MultiArrayLayout : public tinyros::Msg
  {
    public:
      uint32_t dim_length;
      typedef std_msgs::MultiArrayDimension _dim_type;
      _dim_type st_dim;
      _dim_type * dim;
      typedef uint32_t _data_offset_type;
      _data_offset_type data_offset;

    MultiArrayLayout():
      dim_length(0), dim(NULL),
      data_offset(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->dim_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->dim_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->dim_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->dim_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->dim_length);
      for( uint32_t i = 0; i < dim_length; i++) {
        offset += this->dim[i].serialize(outbuffer + offset);
      }
      *(outbuffer + offset + 0) = (this->data_offset >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data_offset >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->data_offset >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->data_offset >> (8 * 3)) & 0xFF;
      offset += sizeof(this->data_offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t dim_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      dim_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      dim_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      dim_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->dim_length);
      if(dim_lengthT > dim_length)
        this->dim = (std_msgs::MultiArrayDimension*)realloc(this->dim, dim_lengthT * sizeof(std_msgs::MultiArrayDimension));
      dim_length = dim_lengthT;
      for( uint32_t i = 0; i < dim_length; i++) {
        offset += this->st_dim.deserialize(inbuffer + offset);
        memcpy( &(this->dim[i]), &(this->st_dim), sizeof(std_msgs::MultiArrayDimension));
      }
      this->data_offset =  ((uint32_t) (*(inbuffer + offset)));
      this->data_offset |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->data_offset |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->data_offset |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->data_offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->dim_length);
      for( uint32_t i = 0; i < dim_length; i++) {
        length += this->dim[i].serializedLength();
      }
      length += sizeof(this->data_offset);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "dim:[";
      for( uint32_t i = 0; i < dim_length; i++) {
        if( i == (dim_length - 1)) {
          string_echo += this->dim[i].echo();
          string_echo += "";
        } else {
          string_echo += this->dim[i].echo();
          string_echo += ",";
        }
      }
      string_echo += "],";
      std::stringstream ss_data_offset; ss_data_offset << "\"data_offset\":" << data_offset <<"";
      string_echo += ss_data_offset.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "std_msgs/MultiArrayLayout"; }
    virtual std::string getMD5(){ return "f40f0b5b285a93ca167c98c1012a989a"; }

  };

}
#endif
