#ifndef _TINYROS_tinyros_msgs_Log_h
#define _TINYROS_tinyros_msgs_Log_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace tinyros_msgs
{

  class Log : public tinyros::Msg
  {
    public:
      typedef uint8_t _level_type;
      _level_type level;
      typedef tinyros::string _msg_type;
      _msg_type msg;
      enum { ROSDEBUG = 0 };
      enum { ROSINFO = 1 };
      enum { ROSWARN = 2 };
      enum { ROSERROR = 3 };
      enum { ROSFATAL = 4 };

    Log():
      level(0),
      msg("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->level >> (8 * 0)) & 0xFF;
      offset += sizeof(this->level);
      uint32_t length_msg = this->msg.size();
      varToArr(outbuffer + offset, length_msg);
      offset += 4;
      memcpy(outbuffer + offset, this->msg.c_str(), length_msg);
      offset += length_msg;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->level =  ((uint8_t) (*(inbuffer + offset)));
      offset += sizeof(this->level);
      uint32_t length_msg;
      arrToVar(length_msg, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_msg; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_msg-1]=0;
      this->msg = (char *)(inbuffer + offset-1);
      offset += length_msg;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->level);
      uint32_t length_msg = this->msg.size();
      length += 4;
      length += length_msg;
      return length;
    }

    virtual tinyros::string getType(){ return "tinyros_msgs/Log"; }
    virtual tinyros::string getMD5(){ return "0bd74339b4d77cb15766d831a3d15eeb"; }

  };

}
#endif
