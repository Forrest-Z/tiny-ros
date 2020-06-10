#ifndef _TINYROS_tf2_msgs_TFMessage_h
#define _TINYROS_tf2_msgs_TFMessage_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/TransformStamped.h"

namespace tf2_msgs
{

  class TFMessage : public tinyros::Msg
  {
    public:
      uint32_t transforms_length;
      typedef geometry_msgs::TransformStamped _transforms_type;
      _transforms_type st_transforms;
      _transforms_type * transforms;

    TFMessage():
      transforms_length(0), transforms(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->transforms_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->transforms_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->transforms_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->transforms_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->transforms_length);
      for( uint32_t i = 0; i < transforms_length; i++){
      offset += this->transforms[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t transforms_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->transforms_length);
      if(transforms_lengthT > transforms_length)
        this->transforms = (geometry_msgs::TransformStamped*)realloc(this->transforms, transforms_lengthT * sizeof(geometry_msgs::TransformStamped));
      transforms_length = transforms_lengthT;
      for( uint32_t i = 0; i < transforms_length; i++){
      offset += this->st_transforms.deserialize(inbuffer + offset);
        memcpy( &(this->transforms[i]), &(this->st_transforms), sizeof(geometry_msgs::TransformStamped));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->transforms_length);
      for( uint32_t i = 0; i < transforms_length; i++){
      length += this->transforms[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "transforms: [";
      for( uint32_t i = 0; i < transforms_length; i++){
      if( i == (transforms_length - 1)) {
      std::stringstream ss_transformsi; ss_transformsi << "{\"transforms" << i <<"\": {";
      string_echo += ss_transformsi.str();
      string_echo += this->transforms[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_transformsi; ss_transformsi << "{\"transforms" << i <<"\": {";
      string_echo += ss_transformsi.str();
      string_echo += this->transforms[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tf2_msgs/TFMessage"; }
    virtual std::string getMD5(){ return "cb93cfe6a141f8d8af7cc34997ec99fe"; }

  };

}
#endif
