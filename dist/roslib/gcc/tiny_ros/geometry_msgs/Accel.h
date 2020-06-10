#ifndef _TINYROS_geometry_msgs_Accel_h
#define _TINYROS_geometry_msgs_Accel_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Vector3.h"

namespace geometry_msgs
{

  class Accel : public tinyros::Msg
  {
    public:
      typedef geometry_msgs::Vector3 _linear_type;
      _linear_type linear;
      typedef geometry_msgs::Vector3 _angular_type;
      _angular_type angular;

    Accel():
      linear(),
      angular()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->linear.serialize(outbuffer + offset);
      offset += this->angular.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->linear.deserialize(inbuffer + offset);
      offset += this->angular.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->linear.serializedLength();
      length += this->angular.serializedLength();
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"linear\": {";
      string_echo += this->linear.echo();
      string_echo += "}, ";
      string_echo += "\"angular\": {";
      string_echo += this->angular.echo();
      string_echo += "}";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "geometry_msgs/Accel"; }
    virtual std::string getMD5(){ return "580cbad5f3bd2e9f0ca71e14b7ab1b0f"; }

  };

}
#endif
