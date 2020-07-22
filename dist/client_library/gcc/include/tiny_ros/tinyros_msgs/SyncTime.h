#ifndef _TINYROS_tinyros_msgs_SyncTime_h
#define _TINYROS_tinyros_msgs_SyncTime_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/time.h"

namespace tinyros_msgs
{

  class SyncTime : public tinyros::Msg
  {
    public:
      typedef uint32_t _tick_type;
      _tick_type tick;
      typedef tinyros::Time _data_type;
      _data_type data;

    SyncTime():
      tick(0),
      data()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->tick >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->tick >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->tick >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->tick >> (8 * 3)) & 0xFF;
      offset += sizeof(this->tick);
      *(outbuffer + offset + 0) = (this->data.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->data.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->data.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->data.sec);
      *(outbuffer + offset + 0) = (this->data.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->data.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->data.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->data.nsec);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->tick =  ((uint32_t) (*(inbuffer + offset)));
      this->tick |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->tick |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->tick |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->tick);
      this->data.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->data.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->data.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->data.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->data.sec);
      this->data.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->data.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->data.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->data.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->data.nsec);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->tick);
      length += sizeof(this->data.sec);
      length += sizeof(this->data.nsec);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_tick; ss_tick << "\"tick\":" << tick <<",";
      string_echo += ss_tick.str();
      std::stringstream ss_data;
      ss_data << "\"data\":{\"sec\":" << data.sec;
      ss_data << ",\"nsec\":" << data.nsec << "}";
      string_echo += ss_data.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tinyros_msgs/SyncTime"; }
    virtual std::string getMD5(){ return "45bf702585c65b1bb762993bdbb1de6f"; }

  };

}
#endif
