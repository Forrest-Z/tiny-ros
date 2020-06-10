#ifndef _TINYROS_nav_msgs_GetMapFeedback_h
#define _TINYROS_nav_msgs_GetMapFeedback_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace nav_msgs
{

  class GetMapFeedback : public tinyros::Msg
  {
    public:

    GetMapFeedback()
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

    virtual tinyros::string getType(){ return "nav_msgs/GetMapFeedback"; }
    virtual tinyros::string getMD5(){ return "f561626803919fb2f269eb497bfdfea4"; }

  };

}
#endif
