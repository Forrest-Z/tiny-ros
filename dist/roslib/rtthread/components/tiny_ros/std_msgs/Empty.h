#ifndef _TINYROS_std_msgs_Empty_h
#define _TINYROS_std_msgs_Empty_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace std_msgs
{

  class Empty : public tinyros::Msg
  {
    public:

    Empty()
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

    virtual tinyros::string getType(){ return "std_msgs/Empty"; }
    virtual tinyros::string getMD5(){ return "140bdcb7bbc50b4a936e90ff2350c8d3"; }

  };

}
#endif
