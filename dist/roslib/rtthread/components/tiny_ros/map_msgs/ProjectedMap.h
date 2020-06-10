#ifndef _TINYROS_map_msgs_ProjectedMap_h
#define _TINYROS_map_msgs_ProjectedMap_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/nav_msgs/OccupancyGrid.h"

namespace map_msgs
{

  class ProjectedMap : public tinyros::Msg
  {
    public:
      typedef nav_msgs::OccupancyGrid _map_type;
      _map_type map;
      typedef double _min_z_type;
      _min_z_type min_z;
      typedef double _max_z_type;
      _max_z_type max_z;

    ProjectedMap():
      map(),
      min_z(0),
      max_z(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->map.serialize(outbuffer + offset);
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
      offset += this->map.deserialize(inbuffer + offset);
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
      length += this->map.serializedLength();
      length += sizeof(this->min_z);
      length += sizeof(this->max_z);
      return length;
    }

    virtual tinyros::string getType(){ return "map_msgs/ProjectedMap"; }
    virtual tinyros::string getMD5(){ return "cbd5598c259cc16f5aa07335587a7367"; }

  };

}
#endif
