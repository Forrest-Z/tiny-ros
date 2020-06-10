#ifndef _TINYROS_SERVICE_SetPhysicsProperties_h
#define _TINYROS_SERVICE_SetPhysicsProperties_h
#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Vector3.h"
#include "tiny_ros/gazebo_msgs/ODEPhysics.h"

namespace gazebo_msgs
{

static const char SETPHYSICSPROPERTIES[] = "gazebo_msgs/SetPhysicsProperties";

  class SetPhysicsPropertiesRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef double _time_step_type;
      _time_step_type time_step;
      typedef double _max_update_rate_type;
      _max_update_rate_type max_update_rate;
      typedef geometry_msgs::Vector3 _gravity_type;
      _gravity_type gravity;
      typedef gazebo_msgs::ODEPhysics _ode_config_type;
      _ode_config_type ode_config;

    SetPhysicsPropertiesRequest():
      time_step(0),
      max_update_rate(0),
      gravity(),
      ode_config()
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
        double real;
        uint64_t base;
      } u_time_step;
      u_time_step.real = this->time_step;
      *(outbuffer + offset + 0) = (u_time_step.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_time_step.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_time_step.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_time_step.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_time_step.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_time_step.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_time_step.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_time_step.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->time_step);
      union {
        double real;
        uint64_t base;
      } u_max_update_rate;
      u_max_update_rate.real = this->max_update_rate;
      *(outbuffer + offset + 0) = (u_max_update_rate.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_max_update_rate.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_max_update_rate.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_max_update_rate.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_max_update_rate.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_max_update_rate.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_max_update_rate.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_max_update_rate.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->max_update_rate);
      offset += this->gravity.serialize(outbuffer + offset);
      offset += this->ode_config.serialize(outbuffer + offset);
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
        double real;
        uint64_t base;
      } u_time_step;
      u_time_step.base = 0;
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_time_step.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->time_step = u_time_step.real;
      offset += sizeof(this->time_step);
      union {
        double real;
        uint64_t base;
      } u_max_update_rate;
      u_max_update_rate.base = 0;
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_max_update_rate.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->max_update_rate = u_max_update_rate.real;
      offset += sizeof(this->max_update_rate);
      offset += this->gravity.deserialize(inbuffer + offset);
      offset += this->ode_config.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->time_step);
      length += sizeof(this->max_update_rate);
      length += this->gravity.serializedLength();
      length += this->ode_config.serializedLength();
      return length;
    }

    virtual tinyros::string getType(){ return SETPHYSICSPROPERTIES; }
    virtual tinyros::string getMD5(){ return "373e5371b10119be0a008429a9660679"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SetPhysicsPropertiesResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef bool _success_type;
      _success_type success;
      typedef tinyros::string _status_message_type;
      _status_message_type status_message;

    SetPhysicsPropertiesResponse():
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

    virtual tinyros::string getType(){ return SETPHYSICSPROPERTIES; }
    virtual tinyros::string getMD5(){ return "5b1d14bf828ba119319cc03e2bb3473a"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class SetPhysicsProperties {
    public:
    typedef SetPhysicsPropertiesRequest Request;
    typedef SetPhysicsPropertiesResponse Response;
  };

}
#endif
