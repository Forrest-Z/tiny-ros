#ifndef _TINYROS_sensor_msgs_TimeReference_h
#define _TINYROS_sensor_msgs_TimeReference_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/ros/time.h"

namespace sensor_msgs
{

  class TimeReference : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef tinyros::Time _time_ref_type;
      _time_ref_type time_ref;
      typedef tinyros::string _source_type;
      _source_type source;

    TimeReference():
      header(),
      time_ref(),
      source("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->time_ref.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->time_ref.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->time_ref.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->time_ref.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->time_ref.sec);
      *(outbuffer + offset + 0) = (this->time_ref.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->time_ref.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->time_ref.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->time_ref.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->time_ref.nsec);
      uint32_t length_source = this->source.size();
      varToArr(outbuffer + offset, length_source);
      offset += 4;
      memcpy(outbuffer + offset, this->source.c_str(), length_source);
      offset += length_source;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      this->time_ref.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->time_ref.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->time_ref.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->time_ref.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->time_ref.sec);
      this->time_ref.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->time_ref.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->time_ref.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->time_ref.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->time_ref.nsec);
      uint32_t length_source;
      arrToVar(length_source, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_source; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_source-1]=0;
      this->source = (char *)(inbuffer + offset-1);
      offset += length_source;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->time_ref.sec);
      length += sizeof(this->time_ref.nsec);
      uint32_t length_source = this->source.size();
      length += 4;
      length += length_source;
      return length;
    }

    virtual tinyros::string getType(){ return "sensor_msgs/TimeReference"; }
    virtual tinyros::string getMD5(){ return "8e1576e01de57cd0d55758112f0e84ec"; }

  };

}
#endif
