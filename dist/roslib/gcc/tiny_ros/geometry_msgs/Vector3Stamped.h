#ifndef _TINYROS_geometry_msgs_Vector3Stamped_h
#define _TINYROS_geometry_msgs_Vector3Stamped_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/geometry_msgs/Vector3.h"

namespace geometry_msgs
{

  class Vector3Stamped : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef geometry_msgs::Vector3 _vector_type;
      _vector_type vector;

    Vector3Stamped():
      header(),
      vector()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->vector.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->vector.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->vector.serializedLength();
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\": {";
      string_echo += this->header.echo();
      string_echo += "}, ";
      string_echo += "\"vector\": {";
      string_echo += this->vector.echo();
      string_echo += "}";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "geometry_msgs/Vector3Stamped"; }
    virtual std::string getMD5(){ return "4b85025eb6f70f6b1e0cefbb75f69ac2"; }

  };

}
#endif
