#ifndef _TINYROS_actionlib_msgs_GoalStatusArray_h
#define _TINYROS_actionlib_msgs_GoalStatusArray_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/actionlib_msgs/GoalStatus.h"

namespace actionlib_msgs
{

  class GoalStatusArray : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t status_list_length;
      typedef actionlib_msgs::GoalStatus _status_list_type;
      _status_list_type st_status_list;
      _status_list_type * status_list;

    GoalStatusArray():
      header(),
      status_list_length(0), status_list(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->status_list_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->status_list_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->status_list_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->status_list_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->status_list_length);
      for( uint32_t i = 0; i < status_list_length; i++) {
        offset += this->status_list[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t status_list_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      status_list_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      status_list_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      status_list_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->status_list_length);
      if(status_list_lengthT > status_list_length)
        this->status_list = (actionlib_msgs::GoalStatus*)realloc(this->status_list, status_list_lengthT * sizeof(actionlib_msgs::GoalStatus));
      status_list_length = status_list_lengthT;
      for( uint32_t i = 0; i < status_list_length; i++) {
        offset += this->st_status_list.deserialize(inbuffer + offset);
        memcpy( &(this->status_list[i]), &(this->st_status_list), sizeof(actionlib_msgs::GoalStatus));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->status_list_length);
      for( uint32_t i = 0; i < status_list_length; i++) {
        length += this->status_list[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\":";
      string_echo += this->header.echo();
      string_echo += ",";
      string_echo += "status_list:[";
      for( uint32_t i = 0; i < status_list_length; i++) {
        if( i == (status_list_length - 1)) {
          string_echo += this->status_list[i].echo();
          string_echo += "";
        } else {
          string_echo += this->status_list[i].echo();
          string_echo += ",";
        }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "actionlib_msgs/GoalStatusArray"; }
    virtual std::string getMD5(){ return "53f6501f7c14f5f3963638de4bbe3a71"; }

  };

}
#endif
