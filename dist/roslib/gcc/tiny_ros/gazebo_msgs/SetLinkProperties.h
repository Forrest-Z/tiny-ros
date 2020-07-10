#ifndef _TINYROS_SERVICE_SetLinkProperties_h
#define _TINYROS_SERVICE_SetLinkProperties_h
#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Pose.h"

namespace gazebo_msgs
{

static const char SETLINKPROPERTIES[] = "gazebo_msgs/SetLinkProperties";

  class SetLinkPropertiesRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef std::string _link_name_type;
      _link_name_type link_name;
      typedef geometry_msgs::Pose _com_type;
      _com_type com;
      typedef bool _gravity_mode_type;
      _gravity_mode_type gravity_mode;
      typedef double _mass_type;
      _mass_type mass;
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

    SetLinkPropertiesRequest():
      link_name(""),
      com(),
      gravity_mode(0),
      mass(0),
      ixx(0),
      ixy(0),
      ixz(0),
      iyy(0),
      iyz(0),
      izz(0)
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      uint32_t length_link_name = this->link_name.size();
      varToArr(outbuffer + offset, length_link_name);
      offset += 4;
      memcpy(outbuffer + offset, this->link_name.c_str(), length_link_name);
      offset += length_link_name;
      offset += this->com.serialize(outbuffer + offset);
      union {
        bool real;
        uint8_t base;
      } u_gravity_mode;
      u_gravity_mode.real = this->gravity_mode;
      *(outbuffer + offset + 0) = (u_gravity_mode.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->gravity_mode);
      union {
        double real;
        uint64_t base;
      } u_mass;
      u_mass.real = this->mass;
      *(outbuffer + offset + 0) = (u_mass.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_mass.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_mass.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_mass.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_mass.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_mass.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_mass.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_mass.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->mass);
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
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      uint32_t length_link_name;
      arrToVar(length_link_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_link_name; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_link_name-1]=0;
      this->link_name = (char *)(inbuffer + offset-1);
      offset += length_link_name;
      offset += this->com.deserialize(inbuffer + offset);
      union {
        bool real;
        uint8_t base;
      } u_gravity_mode;
      u_gravity_mode.base = 0;
      u_gravity_mode.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->gravity_mode = u_gravity_mode.real;
      offset += sizeof(this->gravity_mode);
      union {
        double real;
        uint64_t base;
      } u_mass;
      u_mass.base = 0;
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_mass.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->mass = u_mass.real;
      offset += sizeof(this->mass);
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
      uint32_t length_link_name = this->link_name.size();
      length += 4;
      length += length_link_name;
      length += this->com.serializedLength();
      length += sizeof(this->gravity_mode);
      length += sizeof(this->mass);
      length += sizeof(this->ixx);
      length += sizeof(this->ixy);
      length += sizeof(this->ixz);
      length += sizeof(this->iyy);
      length += sizeof(this->iyz);
      length += sizeof(this->izz);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::size_t link_name_pos = link_name.find("\"");
      while(link_name_pos != std::string::npos){
        link_name.replace(link_name_pos, 1,"\\\"");
        link_name_pos = link_name.find("\"", link_name_pos+2);
      }
      string_echo += "\"link_name\":\"";
      string_echo += link_name;
      string_echo += "\",";
      string_echo += "\"com\":";
      string_echo += this->com.echo();
      string_echo += ",";
      std::stringstream ss_gravity_mode; ss_gravity_mode << "\"gravity_mode\":" << gravity_mode <<",";
      string_echo += ss_gravity_mode.str();
      std::stringstream ss_mass; ss_mass << "\"mass\":" << mass <<",";
      string_echo += ss_mass.str();
      std::stringstream ss_ixx; ss_ixx << "\"ixx\":" << ixx <<",";
      string_echo += ss_ixx.str();
      std::stringstream ss_ixy; ss_ixy << "\"ixy\":" << ixy <<",";
      string_echo += ss_ixy.str();
      std::stringstream ss_ixz; ss_ixz << "\"ixz\":" << ixz <<",";
      string_echo += ss_ixz.str();
      std::stringstream ss_iyy; ss_iyy << "\"iyy\":" << iyy <<",";
      string_echo += ss_iyy.str();
      std::stringstream ss_iyz; ss_iyz << "\"iyz\":" << iyz <<",";
      string_echo += ss_iyz.str();
      std::stringstream ss_izz; ss_izz << "\"izz\":" << izz <<"";
      string_echo += ss_izz.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return SETLINKPROPERTIES; }
    virtual std::string getMD5(){ return "03d7e308476601832a9e1ce9d7eab722"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SetLinkPropertiesResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef bool _success_type;
      _success_type success;
      typedef std::string _status_message_type;
      _status_message_type status_message;

    SetLinkPropertiesResponse():
      success(0),
      status_message("")
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      union {
        bool real;
        uint8_t base;
      } u_success;
      u_success.real = this->success;
      *(outbuffer + offset + 0) = (u_success.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->success);
      uint32_t length_status_message = this->status_message.size();
      varToArr(outbuffer + offset, length_status_message);
      offset += 4;
      memcpy(outbuffer + offset, this->status_message.c_str(), length_status_message);
      offset += length_status_message;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      union {
        bool real;
        uint8_t base;
      } u_success;
      u_success.base = 0;
      u_success.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->success = u_success.real;
      offset += sizeof(this->success);
      uint32_t length_status_message;
      arrToVar(length_status_message, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_status_message; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_status_message-1]=0;
      this->status_message = (char *)(inbuffer + offset-1);
      offset += length_status_message;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->success);
      uint32_t length_status_message = this->status_message.size();
      length += 4;
      length += length_status_message;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_success; ss_success << "\"success\":" << success <<",";
      string_echo += ss_success.str();
      std::size_t status_message_pos = status_message.find("\"");
      while(status_message_pos != std::string::npos){
        status_message.replace(status_message_pos, 1,"\\\"");
        status_message_pos = status_message.find("\"", status_message_pos+2);
      }
      string_echo += "\"status_message\":\"";
      string_echo += status_message;
      string_echo += "\"";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return SETLINKPROPERTIES; }
    virtual std::string getMD5(){ return "777dea05607e1bca37e3206f97801d89"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SetLinkProperties {
    public:
    typedef SetLinkPropertiesRequest Request;
    typedef SetLinkPropertiesResponse Response;
  };

}
#endif
