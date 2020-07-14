#ifndef _TINYROS_diagnostic_msgs_DiagnosticStatus_h
#define _TINYROS_diagnostic_msgs_DiagnosticStatus_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/diagnostic_msgs/KeyValue.h"

namespace diagnostic_msgs
{

  class DiagnosticStatus : public tinyros::Msg
  {
    public:
      typedef int8_t _level_type;
      _level_type level;
      typedef std::string _name_type;
      _name_type name;
      typedef std::string _message_type;
      _message_type message;
      typedef std::string _hardware_id_type;
      _hardware_id_type hardware_id;
      uint32_t values_length;
      typedef diagnostic_msgs::KeyValue _values_type;
      _values_type st_values;
      _values_type * values;
      enum { OK = 0 };
      enum { WARN = 1 };
      enum { ERROR = 2 };
      enum { STALE = 3 };

    DiagnosticStatus():
      level(0),
      name(""),
      message(""),
      hardware_id(""),
      values_length(0), values(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      union {
        int8_t real;
        uint8_t base;
      } u_level;
      u_level.real = this->level;
      *(outbuffer + offset + 0) = (u_level.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->level);
      uint32_t length_name = this->name.size();
      varToArr(outbuffer + offset, length_name);
      offset += 4;
      memcpy(outbuffer + offset, this->name.c_str(), length_name);
      offset += length_name;
      uint32_t length_message = this->message.size();
      varToArr(outbuffer + offset, length_message);
      offset += 4;
      memcpy(outbuffer + offset, this->message.c_str(), length_message);
      offset += length_message;
      uint32_t length_hardware_id = this->hardware_id.size();
      varToArr(outbuffer + offset, length_hardware_id);
      offset += 4;
      memcpy(outbuffer + offset, this->hardware_id.c_str(), length_hardware_id);
      offset += length_hardware_id;
      *(outbuffer + offset + 0) = (this->values_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->values_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->values_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->values_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->values_length);
      for( uint32_t i = 0; i < values_length; i++) {
        offset += this->values[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      union {
        int8_t real;
        uint8_t base;
      } u_level;
      u_level.base = 0;
      u_level.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->level = u_level.real;
      offset += sizeof(this->level);
      uint32_t length_name;
      arrToVar(length_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_name; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_name-1]=0;
      this->name = (char *)(inbuffer + offset-1);
      offset += length_name;
      uint32_t length_message;
      arrToVar(length_message, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_message; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_message-1]=0;
      this->message = (char *)(inbuffer + offset-1);
      offset += length_message;
      uint32_t length_hardware_id;
      arrToVar(length_hardware_id, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_hardware_id; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_hardware_id-1]=0;
      this->hardware_id = (char *)(inbuffer + offset-1);
      offset += length_hardware_id;
      uint32_t values_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      values_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      values_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      values_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->values_length);
      if(values_lengthT > values_length)
        this->values = (diagnostic_msgs::KeyValue*)realloc(this->values, values_lengthT * sizeof(diagnostic_msgs::KeyValue));
      values_length = values_lengthT;
      for( uint32_t i = 0; i < values_length; i++) {
        offset += this->st_values.deserialize(inbuffer + offset);
        memcpy( &(this->values[i]), &(this->st_values), sizeof(diagnostic_msgs::KeyValue));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->level);
      uint32_t length_name = this->name.size();
      length += 4;
      length += length_name;
      uint32_t length_message = this->message.size();
      length += 4;
      length += length_message;
      uint32_t length_hardware_id = this->hardware_id.size();
      length += 4;
      length += length_hardware_id;
      length += sizeof(this->values_length);
      for( uint32_t i = 0; i < values_length; i++) {
        length += this->values[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_level; ss_level << "\"level\":" << (int16_t)level <<",";
      string_echo += ss_level.str();
      std::size_t name_pos = name.find("\"");
      while(name_pos != std::string::npos){
        name.replace(name_pos, 1,"\\\"");
        name_pos = name.find("\"", name_pos+2);
      }
      string_echo += "\"name\":\"";
      string_echo += name;
      string_echo += "\",";
      std::size_t message_pos = message.find("\"");
      while(message_pos != std::string::npos){
        message.replace(message_pos, 1,"\\\"");
        message_pos = message.find("\"", message_pos+2);
      }
      string_echo += "\"message\":\"";
      string_echo += message;
      string_echo += "\",";
      std::size_t hardware_id_pos = hardware_id.find("\"");
      while(hardware_id_pos != std::string::npos){
        hardware_id.replace(hardware_id_pos, 1,"\\\"");
        hardware_id_pos = hardware_id.find("\"", hardware_id_pos+2);
      }
      string_echo += "\"hardware_id\":\"";
      string_echo += hardware_id;
      string_echo += "\",";
      string_echo += "values:[";
      for( uint32_t i = 0; i < values_length; i++) {
        if( i == (values_length - 1)) {
          string_echo += this->values[i].echo();
          string_echo += "";
        } else {
          string_echo += this->values[i].echo();
          string_echo += ",";
        }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "diagnostic_msgs/DiagnosticStatus"; }
    virtual std::string getMD5(){ return "9ec892d2145f478061efd60bb1762361"; }

  };

}
#endif
