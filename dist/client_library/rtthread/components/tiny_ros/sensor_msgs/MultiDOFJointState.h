#ifndef _TINYROS_sensor_msgs_MultiDOFJointState_h
#define _TINYROS_sensor_msgs_MultiDOFJointState_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/geometry_msgs/Transform.h"
#include "tiny_ros/geometry_msgs/Twist.h"
#include "tiny_ros/geometry_msgs/Wrench.h"

namespace sensor_msgs
{

  class MultiDOFJointState : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t joint_names_length;
      typedef tinyros::string _joint_names_type;
      _joint_names_type st_joint_names;
      _joint_names_type * joint_names;
      uint32_t transforms_length;
      typedef geometry_msgs::Transform _transforms_type;
      _transforms_type st_transforms;
      _transforms_type * transforms;
      uint32_t twist_length;
      typedef geometry_msgs::Twist _twist_type;
      _twist_type st_twist;
      _twist_type * twist;
      uint32_t wrench_length;
      typedef geometry_msgs::Wrench _wrench_type;
      _wrench_type st_wrench;
      _wrench_type * wrench;

    MultiDOFJointState():
      header(),
      joint_names_length(0), joint_names(NULL),
      transforms_length(0), transforms(NULL),
      twist_length(0), twist(NULL),
      wrench_length(0), wrench(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->joint_names_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->joint_names_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->joint_names_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->joint_names_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->joint_names_length);
      for( uint32_t i = 0; i < joint_names_length; i++) {
        uint32_t length_joint_namesi = this->joint_names[i].size();
        varToArr(outbuffer + offset, length_joint_namesi);
        offset += 4;
        memcpy(outbuffer + offset, this->joint_names[i].c_str(), length_joint_namesi);
        offset += length_joint_namesi;
      }
      *(outbuffer + offset + 0) = (this->transforms_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->transforms_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->transforms_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->transforms_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->transforms_length);
      for( uint32_t i = 0; i < transforms_length; i++) {
        offset += this->transforms[i].serialize(outbuffer + offset);
      }
      *(outbuffer + offset + 0) = (this->twist_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->twist_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->twist_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->twist_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->twist_length);
      for( uint32_t i = 0; i < twist_length; i++) {
        offset += this->twist[i].serialize(outbuffer + offset);
      }
      *(outbuffer + offset + 0) = (this->wrench_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->wrench_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->wrench_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->wrench_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->wrench_length);
      for( uint32_t i = 0; i < wrench_length; i++) {
        offset += this->wrench[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t joint_names_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      joint_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      joint_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      joint_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->joint_names_length);
      if(joint_names_lengthT > joint_names_length)
        this->joint_names = (tinyros::string*)realloc(this->joint_names, joint_names_lengthT * sizeof(tinyros::string));
      joint_names_length = joint_names_lengthT;
      for( uint32_t i = 0; i < joint_names_length; i++) {
        uint32_t length_st_joint_names;
        arrToVar(length_st_joint_names, (inbuffer + offset));
        offset += 4;
        for(unsigned int k= offset; k< offset+length_st_joint_names; ++k){
          inbuffer[k-1]=inbuffer[k];
        }
        inbuffer[offset+length_st_joint_names-1]=0;
        this->st_joint_names = (char *)(inbuffer + offset-1);
        offset += length_st_joint_names;
        memcpy( &(this->joint_names[i]), &(this->st_joint_names), sizeof(tinyros::string));
      }
      uint32_t transforms_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      transforms_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->transforms_length);
      if(transforms_lengthT > transforms_length)
        this->transforms = (geometry_msgs::Transform*)realloc(this->transforms, transforms_lengthT * sizeof(geometry_msgs::Transform));
      transforms_length = transforms_lengthT;
      for( uint32_t i = 0; i < transforms_length; i++) {
        offset += this->st_transforms.deserialize(inbuffer + offset);
        memcpy( &(this->transforms[i]), &(this->st_transforms), sizeof(geometry_msgs::Transform));
      }
      uint32_t twist_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->twist_length);
      if(twist_lengthT > twist_length)
        this->twist = (geometry_msgs::Twist*)realloc(this->twist, twist_lengthT * sizeof(geometry_msgs::Twist));
      twist_length = twist_lengthT;
      for( uint32_t i = 0; i < twist_length; i++) {
        offset += this->st_twist.deserialize(inbuffer + offset);
        memcpy( &(this->twist[i]), &(this->st_twist), sizeof(geometry_msgs::Twist));
      }
      uint32_t wrench_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      wrench_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      wrench_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      wrench_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->wrench_length);
      if(wrench_lengthT > wrench_length)
        this->wrench = (geometry_msgs::Wrench*)realloc(this->wrench, wrench_lengthT * sizeof(geometry_msgs::Wrench));
      wrench_length = wrench_lengthT;
      for( uint32_t i = 0; i < wrench_length; i++) {
        offset += this->st_wrench.deserialize(inbuffer + offset);
        memcpy( &(this->wrench[i]), &(this->st_wrench), sizeof(geometry_msgs::Wrench));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->joint_names_length);
      for( uint32_t i = 0; i < joint_names_length; i++) {
        uint32_t length_joint_namesi = this->joint_names[i].size();
        length += 4;
        length += length_joint_namesi;
      }
      length += sizeof(this->transforms_length);
      for( uint32_t i = 0; i < transforms_length; i++) {
        length += this->transforms[i].serializedLength();
      }
      length += sizeof(this->twist_length);
      for( uint32_t i = 0; i < twist_length; i++) {
        length += this->twist[i].serializedLength();
      }
      length += sizeof(this->wrench_length);
      for( uint32_t i = 0; i < wrench_length; i++) {
        length += this->wrench[i].serializedLength();
      }
      return length;
    }

    virtual tinyros::string getType(){ return "sensor_msgs/MultiDOFJointState"; }
    virtual tinyros::string getMD5(){ return "c1b0232d8e5071b24db76eb143f286eb"; }

  };

}
#endif
