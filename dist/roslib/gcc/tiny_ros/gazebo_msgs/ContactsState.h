#ifndef _TINYROS_gazebo_msgs_ContactsState_h
#define _TINYROS_gazebo_msgs_ContactsState_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/gazebo_msgs/ContactState.h"

namespace gazebo_msgs
{

  class ContactsState : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t states_length;
      typedef gazebo_msgs::ContactState _states_type;
      _states_type st_states;
      _states_type * states;

    ContactsState():
      header(),
      states_length(0), states(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->states_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->states_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->states_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->states_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->states_length);
      for( uint32_t i = 0; i < states_length; i++){
      offset += this->states[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t states_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      states_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      states_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      states_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->states_length);
      if(states_lengthT > states_length)
        this->states = (gazebo_msgs::ContactState*)realloc(this->states, states_lengthT * sizeof(gazebo_msgs::ContactState));
      states_length = states_lengthT;
      for( uint32_t i = 0; i < states_length; i++){
      offset += this->st_states.deserialize(inbuffer + offset);
        memcpy( &(this->states[i]), &(this->st_states), sizeof(gazebo_msgs::ContactState));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->states_length);
      for( uint32_t i = 0; i < states_length; i++){
      length += this->states[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\": {";
      string_echo += this->header.echo();
      string_echo += "}, ";
      string_echo += "states: [";
      for( uint32_t i = 0; i < states_length; i++){
      if( i == (states_length - 1)) {
      std::stringstream ss_statesi; ss_statesi << "{\"states" << i <<"\": {";
      string_echo += ss_statesi.str();
      string_echo += this->states[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_statesi; ss_statesi << "{\"states" << i <<"\": {";
      string_echo += ss_statesi.str();
      string_echo += this->states[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "gazebo_msgs/ContactsState"; }
    virtual std::string getMD5(){ return "d19cd2a086cbd43da4252eb8d5cc64f5"; }

  };

}
#endif
