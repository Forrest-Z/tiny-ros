#ifndef _TINYROS_tf2_msgs_TF2Error_h
#define _TINYROS_tf2_msgs_TF2Error_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace tf2_msgs
{

  class TF2Error : public tinyros::Msg
  {
    public:
      typedef uint8_t _error_type;
      _error_type error;
      typedef tinyros::string _error_string_type;
      _error_string_type error_string;
      enum { NO_ERROR =  0 };
      enum { LOOKUP_ERROR =  1 };
      enum { CONNECTIVITY_ERROR =  2 };
      enum { EXTRAPOLATION_ERROR =  3 };
      enum { INVALID_ARGUMENT_ERROR =  4 };
      enum { TIMEOUT_ERROR =  5 };
      enum { TRANSFORM_ERROR =  6 };

    TF2Error():
      error(0),
      error_string("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->error >> (8 * 0)) & 0xFF;
      offset += sizeof(this->error);
      uint32_t length_error_string = this->error_string.size();
      varToArr(outbuffer + offset, length_error_string);
      offset += 4;
      memcpy(outbuffer + offset, this->error_string.c_str(), length_error_string);
      offset += length_error_string;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->error =  ((uint8_t) (*(inbuffer + offset)));
      offset += sizeof(this->error);
      uint32_t length_error_string;
      arrToVar(length_error_string, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_error_string; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_error_string-1]=0;
      this->error_string = (char *)(inbuffer + offset-1);
      offset += length_error_string;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->error);
      uint32_t length_error_string = this->error_string.size();
      length += 4;
      length += length_error_string;
      return length;
    }

    virtual tinyros::string getType(){ return "tf2_msgs/TF2Error"; }
    virtual tinyros::string getMD5(){ return "ed32adf5a372962d977aea0e5630d1d6"; }

  };

}
#endif
