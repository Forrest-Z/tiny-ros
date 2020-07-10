#ifndef _TINYROS_diagnostic_msgs_KeyValue_h
#define _TINYROS_diagnostic_msgs_KeyValue_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace diagnostic_msgs
{

  class KeyValue : public tinyros::Msg
  {
    public:
      typedef std::string _key_type;
      _key_type key;
      typedef std::string _value_type;
      _value_type value;

    KeyValue():
      key(""),
      value("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      uint32_t length_key = this->key.size();
      varToArr(outbuffer + offset, length_key);
      offset += 4;
      memcpy(outbuffer + offset, this->key.c_str(), length_key);
      offset += length_key;
      uint32_t length_value = this->value.size();
      varToArr(outbuffer + offset, length_value);
      offset += 4;
      memcpy(outbuffer + offset, this->value.c_str(), length_value);
      offset += length_value;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t length_key;
      arrToVar(length_key, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_key; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_key-1]=0;
      this->key = (char *)(inbuffer + offset-1);
      offset += length_key;
      uint32_t length_value;
      arrToVar(length_value, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_value; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_value-1]=0;
      this->value = (char *)(inbuffer + offset-1);
      offset += length_value;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_key = this->key.size();
      length += 4;
      length += length_key;
      uint32_t length_value = this->value.size();
      length += 4;
      length += length_value;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::size_t key_pos = key.find("\"");
      while(key_pos != std::string::npos){
        key.replace(key_pos, 1,"\\\"");
        key_pos = key.find("\"", key_pos+2);
      }
      string_echo += "\"key\":\"";
      string_echo += key;
      string_echo += "\",";
      std::size_t value_pos = value.find("\"");
      while(value_pos != std::string::npos){
        value.replace(value_pos, 1,"\\\"");
        value_pos = value.find("\"", value_pos+2);
      }
      string_echo += "\"value\":\"";
      string_echo += value;
      string_echo += "\"";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "diagnostic_msgs/KeyValue"; }
    virtual std::string getMD5(){ return "1baa904b80c685c77d1a42a872ca1d07"; }

  };

}
#endif
