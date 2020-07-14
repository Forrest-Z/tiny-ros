#ifndef _TINYROS_SERVICE_GetPlan_h
#define _TINYROS_SERVICE_GetPlan_h
#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/PoseStamped.h"
#include "tiny_ros/nav_msgs/Path.h"

namespace nav_msgs
{

static const char GETPLAN[] = "nav_msgs/GetPlan";

  class GetPlanRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef geometry_msgs::PoseStamped _start_type;
      _start_type start;
      typedef geometry_msgs::PoseStamped _goal_type;
      _goal_type goal;
      typedef float _tolerance_type;
      _tolerance_type tolerance;

    GetPlanRequest():
      start(),
      goal(),
      tolerance(0)
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
      offset += this->start.serialize(outbuffer + offset);
      offset += this->goal.serialize(outbuffer + offset);
      union {
        float real;
        uint32_t base;
      } u_tolerance;
      u_tolerance.real = this->tolerance;
      *(outbuffer + offset + 0) = (u_tolerance.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_tolerance.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_tolerance.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_tolerance.base >> (8 * 3)) & 0xFF;
      offset += sizeof(this->tolerance);
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
      offset += this->start.deserialize(inbuffer + offset);
      offset += this->goal.deserialize(inbuffer + offset);
      union {
        float real;
        uint32_t base;
      } u_tolerance;
      u_tolerance.base = 0;
      u_tolerance.base |= ((uint32_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_tolerance.base |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_tolerance.base |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_tolerance.base |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      this->tolerance = u_tolerance.real;
      offset += sizeof(this->tolerance);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->start.serializedLength();
      length += this->goal.serializedLength();
      length += sizeof(this->tolerance);
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"start\":";
      string_echo += this->start.echo();
      string_echo += ",";
      string_echo += "\"goal\":";
      string_echo += this->goal.echo();
      string_echo += ",";
      std::stringstream ss_tolerance; ss_tolerance << "\"tolerance\":" << tolerance <<"";
      string_echo += ss_tolerance.str();
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return GETPLAN; }
    virtual std::string getMD5(){ return "557d5ea947f7761284cf7abef1cd7227"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetPlanResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef nav_msgs::Path _plan_type;
      _plan_type plan;

    GetPlanResponse():
      plan()
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
      offset += this->plan.serialize(outbuffer + offset);
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
      offset += this->plan.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->plan.serializedLength();
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"plan\":";
      string_echo += this->plan.echo();
      string_echo += "";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return GETPLAN; }
    virtual std::string getMD5(){ return "67c62b8c931eabfe35c88aed4b8f1258"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetPlan {
    public:
    typedef GetPlanRequest Request;
    typedef GetPlanResponse Response;
  };

}
#endif
