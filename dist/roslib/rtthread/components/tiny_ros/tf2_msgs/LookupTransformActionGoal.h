#ifndef _TINYROS_tf2_msgs_LookupTransformActionGoal_h
#define _TINYROS_tf2_msgs_LookupTransformActionGoal_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/actionlib_msgs/GoalID.h"
#include "tiny_ros/tf2_msgs/LookupTransformGoal.h"

namespace tf2_msgs
{

  class LookupTransformActionGoal : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef actionlib_msgs::GoalID _goal_id_type;
      _goal_id_type goal_id;
      typedef tf2_msgs::LookupTransformGoal _goal_type;
      _goal_type goal;

    LookupTransformActionGoal():
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

    virtual tinyros::string getType(){ return "tf2_msgs/LookupTransformActionGoal"; }
    virtual tinyros::string getMD5(){ return "b8a7d4ffa64f063b4df7b1dd3fc2bf79"; }

  };

}
#endif
