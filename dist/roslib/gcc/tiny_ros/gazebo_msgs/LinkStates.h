#ifndef _TINYROS_gazebo_msgs_LinkStates_h
#define _TINYROS_gazebo_msgs_LinkStates_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/geometry_msgs/Pose.h"
#include "tiny_ros/geometry_msgs/Twist.h"

namespace gazebo_msgs
{

  class LinkStates : public tinyros::Msg
  {
    public:
      uint32_t name_length;
      typedef std::string _name_type;
      _name_type st_name;
      _name_type * name;
      uint32_t pose_length;
      typedef geometry_msgs::Pose _pose_type;
      _pose_type st_pose;
      _pose_type * pose;
      uint32_t twist_length;
      typedef geometry_msgs::Twist _twist_type;
      _twist_type st_twist;
      _twist_type * twist;

    LinkStates():
      name_length(0), name(NULL),
      pose_length(0), pose(NULL),
      twist_length(0), twist(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->name_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->name_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->name_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->name_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->name_length);
      for( uint32_t i = 0; i < name_length; i++){
      uint32_t length_namei = this->name[i].size();
      varToArr(outbuffer + offset, length_namei);
      offset += 4;
      memcpy(outbuffer + offset, this->name[i].c_str(), length_namei);
      offset += length_namei;
      }
      *(outbuffer + offset + 0) = (this->pose_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->pose_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->pose_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->pose_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->pose_length);
      for( uint32_t i = 0; i < pose_length; i++){
      offset += this->pose[i].serialize(outbuffer + offset);
      }
      *(outbuffer + offset + 0) = (this->twist_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->twist_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->twist_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->twist_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->twist_length);
      for( uint32_t i = 0; i < twist_length; i++){
      offset += this->twist[i].serialize(outbuffer + offset);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      uint32_t name_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      name_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      name_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      name_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->name_length);
      if(name_lengthT > name_length)
        this->name = (std::string*)realloc(this->name, name_lengthT * sizeof(std::string));
      name_length = name_lengthT;
      for( uint32_t i = 0; i < name_length; i++){
      uint32_t length_st_name;
      arrToVar(length_st_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_name; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_name-1]=0;
      this->st_name = (char *)(inbuffer + offset-1);
      offset += length_st_name;
        memcpy( &(this->name[i]), &(this->st_name), sizeof(std::string));
      }
      uint32_t pose_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      pose_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      pose_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      pose_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->pose_length);
      if(pose_lengthT > pose_length)
        this->pose = (geometry_msgs::Pose*)realloc(this->pose, pose_lengthT * sizeof(geometry_msgs::Pose));
      pose_length = pose_lengthT;
      for( uint32_t i = 0; i < pose_length; i++){
      offset += this->st_pose.deserialize(inbuffer + offset);
        memcpy( &(this->pose[i]), &(this->st_pose), sizeof(geometry_msgs::Pose));
      }
      uint32_t twist_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      twist_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->twist_length);
      if(twist_lengthT > twist_length)
        this->twist = (geometry_msgs::Twist*)realloc(this->twist, twist_lengthT * sizeof(geometry_msgs::Twist));
      twist_length = twist_lengthT;
      for( uint32_t i = 0; i < twist_length; i++){
      offset += this->st_twist.deserialize(inbuffer + offset);
        memcpy( &(this->twist[i]), &(this->st_twist), sizeof(geometry_msgs::Twist));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->name_length);
      for( uint32_t i = 0; i < name_length; i++){
      uint32_t length_namei = this->name[i].size();
      length += 4;
      length += length_namei;
      }
      length += sizeof(this->pose_length);
      for( uint32_t i = 0; i < pose_length; i++){
      length += this->pose[i].serializedLength();
      }
      length += sizeof(this->twist_length);
      for( uint32_t i = 0; i < twist_length; i++){
      length += this->twist[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "name: [";
      for( uint32_t i = 0; i < name_length; i++){
      if( i == (name_length - 1)) {
      string_echo += "\"name[i]\": \"";
      string_echo += name[i];
      string_echo += "\"";
      } else {
      string_echo += "\"name[i]\": \"";
      string_echo += name[i];
      string_echo += "\", ";
      }
      }
      string_echo += "], ";
      string_echo += "pose: [";
      for( uint32_t i = 0; i < pose_length; i++){
      if( i == (pose_length - 1)) {
      std::stringstream ss_posei; ss_posei << "{\"pose" << i <<"\": {";
      string_echo += ss_posei.str();
      string_echo += this->pose[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_posei; ss_posei << "{\"pose" << i <<"\": {";
      string_echo += ss_posei.str();
      string_echo += this->pose[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "], ";
      string_echo += "twist: [";
      for( uint32_t i = 0; i < twist_length; i++){
      if( i == (twist_length - 1)) {
      std::stringstream ss_twisti; ss_twisti << "{\"twist" << i <<"\": {";
      string_echo += ss_twisti.str();
      string_echo += this->twist[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_twisti; ss_twisti << "{\"twist" << i <<"\": {";
      string_echo += ss_twisti.str();
      string_echo += this->twist[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "gazebo_msgs/LinkStates"; }
    virtual std::string getMD5(){ return "a6f8cc7b3dee31015716313fe2d419eb"; }

  };

}
#endif