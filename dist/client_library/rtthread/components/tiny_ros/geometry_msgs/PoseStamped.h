#ifndef _TINYROS_geometry_msgs_PoseStamped_h
#define _TINYROS_geometry_msgs_PoseStamped_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/geometry_msgs/Pose.h"

namespace geometry_msgs
{

  class PoseStamped : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef geometry_msgs::Pose _pose_type;
      _pose_type pose;

    PoseStamped():
      header(),
      pose()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->pose.serialize(outbuffer + offset);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->pose.deserialize(inbuffer + offset);
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->pose.serializedLength();
      return length;
    }

    virtual tinyros::string getType(){ return "geometry_msgs/PoseStamped"; }
    virtual tinyros::string getMD5(){ return "c7084e6b27c3d6e62efd9bf6d2f6540f"; }

  };

}
#endif
