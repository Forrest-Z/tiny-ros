#ifndef _TINYROS_nav_msgs_GetMapActionGoal_h
#define _TINYROS_nav_msgs_GetMapActionGoal_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/actionlib_msgs/GoalID.h"
#include "tiny_ros/nav_msgs/GetMapGoal.h"

namespace nav_msgs
{

  class GetMapActionGoal : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef actionlib_msgs::GoalID _goal_id_type;
      _goal_id_type goal_id;
      typedef nav_msgs::GetMapGoal _goal_type;
      _goal_type goal;

    GetMapActionGoal():
      header(),
      goal_id(),
      goal()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->goal_id.serialize(outbuffer + offset);
      offset += this->goal.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->goal_id.deserialize(inbuffer + offset);
      offset += this->goal.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->goal_id.serializedLength();
      length += this->goal.serializedLength();
      return length;
    }

    virtual tinyros::string getType(){ return "nav_msgs/GetMapActionGoal"; }
    virtual tinyros::string getMD5(){ return "8aea83336b4ee626241742bb14b14d90"; }

  };

}
#endif
