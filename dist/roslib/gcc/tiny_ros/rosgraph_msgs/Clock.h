#ifndef _TINYROS_rosgraph_msgs_Clock_h
#define _TINYROS_rosgraph_msgs_Clock_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/time.h"

namespace rosgraph_msgs
{

  class Clock : public tinyros::Msg
  {
    public:
      typedef tinyros::Time _clock_type;
      _clock_type clock;

    Clock():
      clock()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->clock.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->clock.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->clock.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->clock.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->clock.sec);
      *(outbuffer + offset + 0) = (this->clock.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->clock.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->clock.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->clock.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->clock.nsec);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->clock.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->clock.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->clock.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->clock.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->clock.sec);
      this->clock.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->clock.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->clock.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->clock.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->clock.nsec);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->clock.sec);
      length += sizeof(this->clock.nsec);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_clock;
      ss_clock << "\"clock.sec\": " << clock.sec;
      ss_clock << ", \"clock.nsec\": " << clock.nsec << "";
      string_echo += ss_clock.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "rosgraph_msgs/Clock"; }
    virtual std::string getMD5(){ return "d3bedbe03b904b8181e3fef4bbe0a73e"; }

  };

}
#endif
