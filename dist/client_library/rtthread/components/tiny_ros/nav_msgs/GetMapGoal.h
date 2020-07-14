#ifndef _TINYROS_nav_msgs_GetMapGoal_h
#define _TINYROS_nav_msgs_GetMapGoal_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace nav_msgs
{

  class GetMapGoal : public tinyros::Msg
  {
    public:

    GetMapGoal()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      return length;
    }

    virtual tinyros::string getType(){ return "nav_msgs/GetMapGoal"; }
    virtual tinyros::string getMD5(){ return "b39e6b705afaad0184bd2c87f4bd870f"; }

  };

}
#endif
