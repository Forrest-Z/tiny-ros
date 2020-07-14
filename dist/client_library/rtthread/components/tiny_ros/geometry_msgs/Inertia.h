#ifndef _TINYROS_geometry_msgs_Inertia_h
#define _TINYROS_geometry_msgs_Inertia_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Vector3.h"

namespace geometry_msgs
{

  class Inertia : public tinyros::Msg
  {
    public:
      typedef double _m_type;
      _m_type m;
      typedef geometry_msgs::Vector3 _com_type;
      _com_type com;
      typedef double _ixx_type;
      _ixx_type ixx;
      typedef double _ixy_type;
      _ixy_type ixy;
      typedef double _ixz_type;
      _ixz_type ixz;
      typedef double _iyy_type;
      _iyy_type iyy;
      typedef double _iyz_type;
      _iyz_type iyz;
      typedef double _izz_type;
      _izz_type izz;

    Inertia():
      m(0),
      com(),
      ixx(0),
      ixy(0),
      ixz(0),
      iyy(0),
      iyz(0),
      izz(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      union {
        double real;
        uint64_t base;
      } u_m;
      u_m.real = this->m;
      *(outbuffer + offset + 0) = (u_m.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_m.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_m.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_m.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_m.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_m.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_m.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_m.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->m);
      offset += this->com.serialize(outbuffer + offset);
      union {
        double real;
        uint64_t base;
      } u_ixx;
      u_ixx.real = this->ixx;
      *(outbuffer + offset + 0) = (u_ixx.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_ixx.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_ixx.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_ixx.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_ixx.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_ixx.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_ixx.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_ixx.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->ixx);
      union {
        double real;
        uint64_t base;
      } u_ixy;
      u_ixy.real = this->ixy;
      *(outbuffer + offset + 0) = (u_ixy.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_ixy.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_ixy.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_ixy.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_ixy.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_ixy.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_ixy.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_ixy.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->ixy);
      union {
        double real;
        uint64_t base;
      } u_ixz;
      u_ixz.real = this->ixz;
      *(outbuffer + offset + 0) = (u_ixz.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_ixz.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_ixz.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_ixz.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_ixz.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_ixz.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_ixz.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_ixz.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->ixz);
      union {
        double real;
        uint64_t base;
      } u_iyy;
      u_iyy.real = this->iyy;
      *(outbuffer + offset + 0) = (u_iyy.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_iyy.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_iyy.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_iyy.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_iyy.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_iyy.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_iyy.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_iyy.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->iyy);
      union {
        double real;
        uint64_t base;
      } u_iyz;
      u_iyz.real = this->iyz;
      *(outbuffer + offset + 0) = (u_iyz.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_iyz.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_iyz.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_iyz.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_iyz.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_iyz.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_iyz.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_iyz.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->iyz);
      union {
        double real;
        uint64_t base;
      } u_izz;
      u_izz.real = this->izz;
      *(outbuffer + offset + 0) = (u_izz.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_izz.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_izz.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_izz.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_izz.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_izz.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_izz.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_izz.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->izz);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      union {
        double real;
        uint64_t base;
      } u_m;
      u_m.base = 0;
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_m.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->m = u_m.real;
      offset += sizeof(this->m);
      offset += this->com.deserialize(inbuffer + offset);
      union {
        double real;
        uint64_t base;
      } u_ixx;
      u_ixx.base = 0;
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_ixx.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->ixx = u_ixx.real;
      offset += sizeof(this->ixx);
      union {
        double real;
        uint64_t base;
      } u_ixy;
      u_ixy.base = 0;
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_ixy.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->ixy = u_ixy.real;
      offset += sizeof(this->ixy);
      union {
        double real;
        uint64_t base;
      } u_ixz;
      u_ixz.base = 0;
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_ixz.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->ixz = u_ixz.real;
      offset += sizeof(this->ixz);
      union {
        double real;
        uint64_t base;
      } u_iyy;
      u_iyy.base = 0;
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_iyy.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->iyy = u_iyy.real;
      offset += sizeof(this->iyy);
      union {
        double real;
        uint64_t base;
      } u_iyz;
      u_iyz.base = 0;
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_iyz.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->iyz = u_iyz.real;
      offset += sizeof(this->iyz);
      union {
        double real;
        uint64_t base;
      } u_izz;
      u_izz.base = 0;
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_izz.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->izz = u_izz.real;
      offset += sizeof(this->izz);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->m);
      length += this->com.serializedLength();
      length += sizeof(this->ixx);
      length += sizeof(this->ixy);
      length += sizeof(this->ixz);
      length += sizeof(this->iyy);
      length += sizeof(this->iyz);
      length += sizeof(this->izz);
      return length;
    }

    virtual tinyros::string getType(){ return "geometry_msgs/Inertia"; }
    virtual tinyros::string getMD5(){ return "9116c935782bc29999dad1927624dff0"; }

  };

}
#endif
