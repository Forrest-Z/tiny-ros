#ifndef _TINYROS_map_msgs_ProjectedMapInfo_h
#define _TINYROS_map_msgs_ProjectedMapInfo_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace map_msgs
{

  class ProjectedMapInfo : public tinyros::Msg
  {
    public:
      typedef std::string _frame_id_type;
      _frame_id_type frame_id;
      typedef double _x_type;
      _x_type x;
      typedef double _y_type;
      _y_type y;
      typedef double _width_type;
      _width_type width;
      typedef double _height_type;
      _height_type height;
      typedef double _min_z_type;
      _min_z_type min_z;
      typedef double _max_z_type;
      _max_z_type max_z;

    ProjectedMapInfo():
      frame_id(""),
      x(0),
      y(0),
      width(0),
      height(0),
      min_z(0),
      max_z(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      uint32_t length_frame_id = this->frame_id.size();
      varToArr(outbuffer + offset, length_frame_id);
      offset += 4;
      memcpy(outbuffer + offset, this->frame_id.c_str(), length_frame_id);
      offset += length_frame_id;
      union {
        double real;
        uint64_t base;
      } u_x;
      u_x.real = this->x;
      *(outbuffer + offset + 0) = (u_x.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_x.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_x.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_x.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_x.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_x.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_x.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_x.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->x);
      union {
        double real;
        uint64_t base;
      } u_y;
      u_y.real = this->y;
      *(outbuffer + offset + 0) = (u_y.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_y.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_y.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_y.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_y.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_y.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_y.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_y.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->y);
      union {
        double real;
        uint64_t base;
      } u_width;
      u_width.real = this->width;
      *(outbuffer + offset + 0) = (u_width.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_width.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_width.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_width.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_width.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_width.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_width.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_width.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->width);
      union {
        double real;
        uint64_t base;
      } u_height;
      u_height.real = this->height;
      *(outbuffer + offset + 0) = (u_height.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_height.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_height.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_height.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_height.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_height.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_height.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_height.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->height);
      union {
        double real;
        uint64_t base;
      } u_min_z;
      u_min_z.real = this->min_z;
      *(outbuffer + offset + 0) = (u_min_z.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_min_z.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_min_z.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_min_z.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_min_z.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_min_z.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_min_z.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_min_z.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->min_z);
      union {
        double real;
        uint64_t base;
      } u_max_z;
      u_max_z.real = this->max_z;
      *(outbuffer + offset + 0) = (u_max_z.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_max_z.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_max_z.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_max_z.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_max_z.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_max_z.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_max_z.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_max_z.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->max_z);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t length_frame_id;
      arrToVar(length_frame_id, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_frame_id; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_frame_id-1]=0;
      this->frame_id = (char *)(inbuffer + offset-1);
      offset += length_frame_id;
      union {
        double real;
        uint64_t base;
      } u_x;
      u_x.base = 0;
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_x.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->x = u_x.real;
      offset += sizeof(this->x);
      union {
        double real;
        uint64_t base;
      } u_y;
      u_y.base = 0;
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_y.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->y = u_y.real;
      offset += sizeof(this->y);
      union {
        double real;
        uint64_t base;
      } u_width;
      u_width.base = 0;
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_width.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->width = u_width.real;
      offset += sizeof(this->width);
      union {
        double real;
        uint64_t base;
      } u_height;
      u_height.base = 0;
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_height.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->height = u_height.real;
      offset += sizeof(this->height);
      union {
        double real;
        uint64_t base;
      } u_min_z;
      u_min_z.base = 0;
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_min_z.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->min_z = u_min_z.real;
      offset += sizeof(this->min_z);
      union {
        double real;
        uint64_t base;
      } u_max_z;
      u_max_z.base = 0;
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_max_z.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->max_z = u_max_z.real;
      offset += sizeof(this->max_z);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_frame_id = this->frame_id.size();
      length += 4;
      length += length_frame_id;
      length += sizeof(this->x);
      length += sizeof(this->y);
      length += sizeof(this->width);
      length += sizeof(this->height);
      length += sizeof(this->min_z);
      length += sizeof(this->max_z);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"frame_id\": \"";
      string_echo += frame_id;
      string_echo += "\", ";
      std::stringstream ss_x; ss_x << "\"x\": " << x <<", ";
      string_echo += ss_x.str();
      std::stringstream ss_y; ss_y << "\"y\": " << y <<", ";
      string_echo += ss_y.str();
      std::stringstream ss_width; ss_width << "\"width\": " << width <<", ";
      string_echo += ss_width.str();
      std::stringstream ss_height; ss_height << "\"height\": " << height <<", ";
      string_echo += ss_height.str();
      std::stringstream ss_min_z; ss_min_z << "\"min_z\": " << min_z <<", ";
      string_echo += ss_min_z.str();
      std::stringstream ss_max_z; ss_max_z << "\"max_z\": " << max_z <<"";
      string_echo += ss_max_z.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "map_msgs/ProjectedMapInfo"; }
    virtual std::string getMD5(){ return "f661365637fb759e63cb5d179a4461e1"; }

  };

}
#endif
