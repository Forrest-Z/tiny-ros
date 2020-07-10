#ifndef _TINYROS_geometry_msgs_Transform_h
#define _TINYROS_geometry_msgs_Transform_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Vector3.h"
#include "tiny_ros/geometry_msgs/Quaternion.h"

namespace geometry_msgs
{

  class Transform : public tinyros::Msg
  {
    public:
      typedef geometry_msgs::Vector3 _translation_type;
      _translation_type translation;
      typedef geometry_msgs::Quaternion _rotation_type;
      _rotation_type rotation;

    Transform():
      translation(),
      rotation()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->translation.serialize(outbuffer + offset);
      offset += this->rotation.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->translation.deserialize(inbuffer + offset);
      offset += this->rotation.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->translation.serializedLength();
      length += this->rotation.serializedLength();
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"translation\":";
      string_echo += this->translation.echo();
      string_echo += ",";
      string_echo += "\"rotation\":";
      string_echo += this->rotation.echo();
      string_echo += "";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "geometry_msgs/Transform"; }
    virtual std::string getMD5(){ return "2526ee1b1cc2e723e386c3c1b048ba72"; }

  };

}
#endif
