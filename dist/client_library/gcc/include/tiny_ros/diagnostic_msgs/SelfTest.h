#ifndef _TINYROS_SERVICE_SelfTest_h
#define _TINYROS_SERVICE_SelfTest_h
#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/diagnostic_msgs/DiagnosticStatus.h"

namespace diagnostic_msgs
{

static const char SELFTEST[] = "diagnostic_msgs/SelfTest";

  class SelfTestRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:

    SelfTestRequest()
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return SELFTEST; }
    virtual std::string getMD5(){ return "049f87742408b36b8ef5f7dd71e3ef5a"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SelfTestResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef std::string _id_type;
      _id_type id;
      typedef int8_t _passed_type;
      _passed_type passed;
      uint32_t status_length;
      typedef diagnostic_msgs::DiagnosticStatus _status_type;
      _status_type st_status;
      _status_type * status;

    SelfTestResponse():
      id(""),
      passed(0),
      status_length(0), status(NULL)
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      uint32_t length_id = this->id.size();
      varToArr(outbuffer + offset, length_id);
      offset += 4;
      memcpy(outbuffer + offset, this->id.c_str(), length_id);
      offset += length_id;
      union {
        int8_t real;
        uint8_t base;
      } u_passed;
      u_passed.real = this->passed;
      *(outbuffer + offset + 0) = (u_passed.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->passed);
      *(outbuffer + offset + 0) = (this->status_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->status_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->status_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->status_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->status_length);
      for( uint32_t i = 0; i < status_length; i++) {
        offset += this->status[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      uint32_t length_id;
      arrToVar(length_id, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_id; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_id-1]=0;
      this->id = (char *)(inbuffer + offset-1);
      offset += length_id;
      union {
        int8_t real;
        uint8_t base;
      } u_passed;
      u_passed.base = 0;
      u_passed.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->passed = u_passed.real;
      offset += sizeof(this->passed);
      uint32_t status_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->status_length);
      if(status_lengthT > status_length)
        this->status = (diagnostic_msgs::DiagnosticStatus*)realloc(this->status, status_lengthT * sizeof(diagnostic_msgs::DiagnosticStatus));
      status_length = status_lengthT;
      for( uint32_t i = 0; i < status_length; i++) {
        offset += this->st_status.deserialize(inbuffer + offset);
        memcpy( &(this->status[i]), &(this->st_status), sizeof(diagnostic_msgs::DiagnosticStatus));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_id = this->id.size();
      length += 4;
      length += length_id;
      length += sizeof(this->passed);
      length += sizeof(this->status_length);
      for( uint32_t i = 0; i < status_length; i++) {
        length += this->status[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::size_t id_pos = id.find("\"");
      while(id_pos != std::string::npos){
        id.replace(id_pos, 1,"\\\"");
        id_pos = id.find("\"", id_pos+2);
      }
      string_echo += "\"id\":\"";
      string_echo += id;
      string_echo += "\",";
      std::stringstream ss_passed; ss_passed << "\"passed\":" << (int16_t)passed <<",";
      string_echo += ss_passed.str();
      string_echo += "status:[";
      for( uint32_t i = 0; i < status_length; i++) {
        if( i == (status_length - 1)) {
          string_echo += this->status[i].echo();
          string_echo += "";
        } else {
          string_echo += this->status[i].echo();
          string_echo += ",";
        }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return SELFTEST; }
    virtual std::string getMD5(){ return "70aaf2a851ccb5e946b2d112ea26f7b9"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SelfTest {
    public:
    typedef SelfTestRequest Request;
    typedef SelfTestResponse Response;
  };

}
#endif
