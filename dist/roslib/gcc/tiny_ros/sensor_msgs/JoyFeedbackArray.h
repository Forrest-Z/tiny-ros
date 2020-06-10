#ifndef _TINYROS_sensor_msgs_JoyFeedbackArray_h
#define _TINYROS_sensor_msgs_JoyFeedbackArray_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/sensor_msgs/JoyFeedback.h"

namespace sensor_msgs
{

  class JoyFeedbackArray : public tinyros::Msg
  {
    public:
      uint32_t array_length;
      typedef sensor_msgs::JoyFeedback _array_type;
      _array_type st_array;
      _array_type * array;

    JoyFeedbackArray():
      array_length(0), array(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->array_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->array_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->array_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->array_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->array_length);
      for( uint32_t i = 0; i < array_length; i++){
      offset += this->array[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t array_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      array_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      array_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      array_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->array_length);
      if(array_lengthT > array_length)
        this->array = (sensor_msgs::JoyFeedback*)realloc(this->array, array_lengthT * sizeof(sensor_msgs::JoyFeedback));
      array_length = array_lengthT;
      for( uint32_t i = 0; i < array_length; i++){
      offset += this->st_array.deserialize(inbuffer + offset);
        memcpy( &(this->array[i]), &(this->st_array), sizeof(sensor_msgs::JoyFeedback));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->array_length);
      for( uint32_t i = 0; i < array_length; i++){
      length += this->array[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "array: [";
      for( uint32_t i = 0; i < array_length; i++){
      if( i == (array_length - 1)) {
      std::stringstream ss_arrayi; ss_arrayi << "{\"array" << i <<"\": {";
      string_echo += ss_arrayi.str();
      string_echo += this->array[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_arrayi; ss_arrayi << "{\"array" << i <<"\": {";
      string_echo += ss_arrayi.str();
      string_echo += this->array[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "sensor_msgs/JoyFeedbackArray"; }
    virtual std::string getMD5(){ return "45361e458d526d5670706a9f083819b6"; }

  };

}
#endif
