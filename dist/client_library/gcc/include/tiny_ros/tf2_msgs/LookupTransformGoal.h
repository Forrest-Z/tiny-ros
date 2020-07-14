#ifndef _TINYROS_tf2_msgs_LookupTransformGoal_h
#define _TINYROS_tf2_msgs_LookupTransformGoal_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/ros/time.h"
#include "tiny_ros/ros/duration.h"

namespace tf2_msgs
{

  class LookupTransformGoal : public tinyros::Msg
  {
    public:
      typedef std::string _target_frame_type;
      _target_frame_type target_frame;
      typedef std::string _source_frame_type;
      _source_frame_type source_frame;
      typedef tinyros::Time _source_time_type;
      _source_time_type source_time;
      typedef tinyros::Duration _timeout_type;
      _timeout_type timeout;
      typedef tinyros::Time _target_time_type;
      _target_time_type target_time;
      typedef std::string _fixed_frame_type;
      _fixed_frame_type fixed_frame;
      typedef bool _advanced_type;
      _advanced_type advanced;

    LookupTransformGoal():
      target_frame(""),
      source_frame(""),
      source_time(),
      timeout(),
      target_time(),
      fixed_frame(""),
      advanced(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      uint32_t length_target_frame = this->target_frame.size();
      varToArr(outbuffer + offset, length_target_frame);
      offset += 4;
      memcpy(outbuffer + offset, this->target_frame.c_str(), length_target_frame);
      offset += length_target_frame;
      uint32_t length_source_frame = this->source_frame.size();
      varToArr(outbuffer + offset, length_source_frame);
      offset += 4;
      memcpy(outbuffer + offset, this->source_frame.c_str(), length_source_frame);
      offset += length_source_frame;
      *(outbuffer + offset + 0) = (this->source_time.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->source_time.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->source_time.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->source_time.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->source_time.sec);
      *(outbuffer + offset + 0) = (this->source_time.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->source_time.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->source_time.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->source_time.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->source_time.nsec);
      *(outbuffer + offset + 0) = (this->timeout.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->timeout.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->timeout.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->timeout.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->timeout.sec);
      *(outbuffer + offset + 0) = (this->timeout.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->timeout.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->timeout.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->timeout.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->timeout.nsec);
      *(outbuffer + offset + 0) = (this->target_time.sec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->target_time.sec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->target_time.sec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->target_time.sec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->target_time.sec);
      *(outbuffer + offset + 0) = (this->target_time.nsec >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->target_time.nsec >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->target_time.nsec >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->target_time.nsec >> (8 * 3)) & 0xFF;
      offset += sizeof(this->target_time.nsec);
      uint32_t length_fixed_frame = this->fixed_frame.size();
      varToArr(outbuffer + offset, length_fixed_frame);
      offset += 4;
      memcpy(outbuffer + offset, this->fixed_frame.c_str(), length_fixed_frame);
      offset += length_fixed_frame;
      union {
        bool real;
        uint8_t base;
      } u_advanced;
      u_advanced.real = this->advanced;
      *(outbuffer + offset + 0) = (u_advanced.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->advanced);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t length_target_frame;
      arrToVar(length_target_frame, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_target_frame; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_target_frame-1]=0;
      this->target_frame = (char *)(inbuffer + offset-1);
      offset += length_target_frame;
      uint32_t length_source_frame;
      arrToVar(length_source_frame, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_source_frame; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_source_frame-1]=0;
      this->source_frame = (char *)(inbuffer + offset-1);
      offset += length_source_frame;
      this->source_time.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->source_time.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->source_time.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->source_time.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->source_time.sec);
      this->source_time.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->source_time.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->source_time.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->source_time.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->source_time.nsec);
      this->timeout.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->timeout.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->timeout.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->timeout.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->timeout.sec);
      this->timeout.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->timeout.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->timeout.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->timeout.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->timeout.nsec);
      this->target_time.sec =  ((uint32_t) (*(inbuffer + offset)));
      this->target_time.sec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->target_time.sec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->target_time.sec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->target_time.sec);
      this->target_time.nsec =  ((uint32_t) (*(inbuffer + offset)));
      this->target_time.nsec |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->target_time.nsec |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->target_time.nsec |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->target_time.nsec);
      uint32_t length_fixed_frame;
      arrToVar(length_fixed_frame, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_fixed_frame; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_fixed_frame-1]=0;
      this->fixed_frame = (char *)(inbuffer + offset-1);
      offset += length_fixed_frame;
      union {
        bool real;
        uint8_t base;
      } u_advanced;
      u_advanced.base = 0;
      u_advanced.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->advanced = u_advanced.real;
      offset += sizeof(this->advanced);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_target_frame = this->target_frame.size();
      length += 4;
      length += length_target_frame;
      uint32_t length_source_frame = this->source_frame.size();
      length += 4;
      length += length_source_frame;
      length += sizeof(this->source_time.sec);
      length += sizeof(this->source_time.nsec);
      length += sizeof(this->timeout.sec);
      length += sizeof(this->timeout.nsec);
      length += sizeof(this->target_time.sec);
      length += sizeof(this->target_time.nsec);
      uint32_t length_fixed_frame = this->fixed_frame.size();
      length += 4;
      length += length_fixed_frame;
      length += sizeof(this->advanced);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::size_t target_frame_pos = target_frame.find("\"");
      while(target_frame_pos != std::string::npos){
        target_frame.replace(target_frame_pos, 1,"\\\"");
        target_frame_pos = target_frame.find("\"", target_frame_pos+2);
      }
      string_echo += "\"target_frame\":\"";
      string_echo += target_frame;
      string_echo += "\",";
      std::size_t source_frame_pos = source_frame.find("\"");
      while(source_frame_pos != std::string::npos){
        source_frame.replace(source_frame_pos, 1,"\\\"");
        source_frame_pos = source_frame.find("\"", source_frame_pos+2);
      }
      string_echo += "\"source_frame\":\"";
      string_echo += source_frame;
      string_echo += "\",";
      std::stringstream ss_source_time;
      ss_source_time << "\"source_time\":{\"sec\":" << source_time.sec;
      ss_source_time << ",\"nsec\":" << source_time.nsec << "},";
      string_echo += ss_source_time.str();
      std::stringstream ss_timeout;
      ss_timeout << "\"timeout\":{\"sec\":" << timeout.sec;
      ss_timeout << ",\"nsec\":" << timeout.nsec << "},";
      string_echo += ss_timeout.str();
      std::stringstream ss_target_time;
      ss_target_time << "\"target_time\":{\"sec\":" << target_time.sec;
      ss_target_time << ",\"nsec\":" << target_time.nsec << "},";
      string_echo += ss_target_time.str();
      std::size_t fixed_frame_pos = fixed_frame.find("\"");
      while(fixed_frame_pos != std::string::npos){
        fixed_frame.replace(fixed_frame_pos, 1,"\\\"");
        fixed_frame_pos = fixed_frame.find("\"", fixed_frame_pos+2);
      }
      string_echo += "\"fixed_frame\":\"";
      string_echo += fixed_frame;
      string_echo += "\",";
      std::stringstream ss_advanced; ss_advanced << "\"advanced\":" << advanced <<"";
      string_echo += ss_advanced.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tf2_msgs/LookupTransformGoal"; }
    virtual std::string getMD5(){ return "677c84a912b788ccaaea5fbc38570182"; }

  };

}
#endif
