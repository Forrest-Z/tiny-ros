#ifndef _TINYROS_diagnostic_msgs_DiagnosticArray_h
#define _TINYROS_diagnostic_msgs_DiagnosticArray_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/diagnostic_msgs/DiagnosticStatus.h"

namespace diagnostic_msgs
{

  class DiagnosticArray : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t status_length;
      typedef diagnostic_msgs::DiagnosticStatus _status_type;
      _status_type st_status;
      _status_type * status;

    DiagnosticArray():
      header(),
      status_length(0), status(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->status_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->status_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->status_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->status_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->status_length);
      for( uint32_t i = 0; i < status_length; i++){
      offset += this->status[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t status_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      status_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->status_length);
      if(status_lengthT > status_length)
        this->status = (diagnostic_msgs::DiagnosticStatus*)realloc(this->status, status_lengthT * sizeof(diagnostic_msgs::DiagnosticStatus));
      status_length = status_lengthT;
      for( uint32_t i = 0; i < status_length; i++){
      offset += this->st_status.deserialize(inbuffer + offset);
        memcpy( &(this->status[i]), &(this->st_status), sizeof(diagnostic_msgs::DiagnosticStatus));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->status_length);
      for( uint32_t i = 0; i < status_length; i++){
      length += this->status[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\": {";
      string_echo += this->header.echo();
      string_echo += "}, ";
      string_echo += "status: [";
      for( uint32_t i = 0; i < status_length; i++){
      if( i == (status_length - 1)) {
      std::stringstream ss_statusi; ss_statusi << "{\"status" << i <<"\": {";
      string_echo += ss_statusi.str();
      string_echo += this->status[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_statusi; ss_statusi << "{\"status" << i <<"\": {";
      string_echo += ss_statusi.str();
      string_echo += this->status[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "diagnostic_msgs/DiagnosticArray"; }
    virtual std::string getMD5(){ return "79a87210f85eb6afbd600eb2ba49dd85"; }

  };

}
#endif
