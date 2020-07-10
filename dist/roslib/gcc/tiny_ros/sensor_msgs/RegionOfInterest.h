#ifndef _TINYROS_sensor_msgs_RegionOfInterest_h
#define _TINYROS_sensor_msgs_RegionOfInterest_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace sensor_msgs
{

  class RegionOfInterest : public tinyros::Msg
  {
    public:
      typedef uint32_t _x_offset_type;
      _x_offset_type x_offset;
      typedef uint32_t _y_offset_type;
      _y_offset_type y_offset;
      typedef uint32_t _height_type;
      _height_type height;
      typedef uint32_t _width_type;
      _width_type width;
      typedef bool _do_rectify_type;
      _do_rectify_type do_rectify;

    RegionOfInterest():
      x_offset(0),
      y_offset(0),
      height(0),
      width(0),
      do_rectify(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->x_offset >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->x_offset >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->x_offset >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->x_offset >> (8 * 3)) & 0xFF;
      offset += sizeof(this->x_offset);
      *(outbuffer + offset + 0) = (this->y_offset >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->y_offset >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->y_offset >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->y_offset >> (8 * 3)) & 0xFF;
      offset += sizeof(this->y_offset);
      *(outbuffer + offset + 0) = (this->height >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->height >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->height >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->height >> (8 * 3)) & 0xFF;
      offset += sizeof(this->height);
      *(outbuffer + offset + 0) = (this->width >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->width >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->width >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->width >> (8 * 3)) & 0xFF;
      offset += sizeof(this->width);
      union {
        bool real;
        uint8_t base;
      } u_do_rectify;
      u_do_rectify.real = this->do_rectify;
      *(outbuffer + offset + 0) = (u_do_rectify.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->do_rectify);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->x_offset =  ((uint32_t) (*(inbuffer + offset)));
      this->x_offset |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->x_offset |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->x_offset |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->x_offset);
      this->y_offset =  ((uint32_t) (*(inbuffer + offset)));
      this->y_offset |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->y_offset |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->y_offset |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->y_offset);
      this->height =  ((uint32_t) (*(inbuffer + offset)));
      this->height |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->height |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->height |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->height);
      this->width =  ((uint32_t) (*(inbuffer + offset)));
      this->width |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->width |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->width |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->width);
      union {
        bool real;
        uint8_t base;
      } u_do_rectify;
      u_do_rectify.base = 0;
      u_do_rectify.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->do_rectify = u_do_rectify.real;
      offset += sizeof(this->do_rectify);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->x_offset);
      length += sizeof(this->y_offset);
      length += sizeof(this->height);
      length += sizeof(this->width);
      length += sizeof(this->do_rectify);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_x_offset; ss_x_offset << "\"x_offset\":" << x_offset <<",";
      string_echo += ss_x_offset.str();
      std::stringstream ss_y_offset; ss_y_offset << "\"y_offset\":" << y_offset <<",";
      string_echo += ss_y_offset.str();
      std::stringstream ss_height; ss_height << "\"height\":" << height <<",";
      string_echo += ss_height.str();
      std::stringstream ss_width; ss_width << "\"width\":" << width <<",";
      string_echo += ss_width.str();
      std::stringstream ss_do_rectify; ss_do_rectify << "\"do_rectify\":" << do_rectify <<"";
      string_echo += ss_do_rectify.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "sensor_msgs/RegionOfInterest"; }
    virtual std::string getMD5(){ return "8370dc286f915405c906299aef5bb442"; }

  };

}
#endif
