#ifndef _TINYROS_trajectory_msgs_MultiDOFJointTrajectory_h
#define _TINYROS_trajectory_msgs_MultiDOFJointTrajectory_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/trajectory_msgs/MultiDOFJointTrajectoryPoint.h"

namespace trajectory_msgs
{

  class MultiDOFJointTrajectory : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t joint_names_length;
      typedef std::string _joint_names_type;
      _joint_names_type st_joint_names;
      _joint_names_type * joint_names;
      uint32_t points_length;
      typedef trajectory_msgs::MultiDOFJointTrajectoryPoint _points_type;
      _points_type st_points;
      _points_type * points;

    MultiDOFJointTrajectory():
      header(),
      joint_names_length(0), joint_names(NULL),
      points_length(0), points(NULL)
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
      for( uint32_t i = 0; i < joint_names_length; i++){
      uint32_t length_joint_namesi = this->joint_names[i].size();
      varToArr(outbuffer + offset, length_joint_namesi);
      offset += 4;
      memcpy(outbuffer + offset, this->joint_names[i].c_str(), length_joint_namesi);
      offset += length_joint_namesi;
      }
      *(outbuffer + offset + 0) = (this->points_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->points_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->points_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->points_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->points_length);
      for( uint32_t i = 0; i < points_length; i++){
      offset += this->points[i].serialize(outbuffer + offset);
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
        this->joint_names = (std::string*)realloc(this->joint_names, joint_names_lengthT * sizeof(std::string));
      joint_names_length = joint_names_lengthT;
      for( uint32_t i = 0; i < joint_names_length; i++){
      uint32_t length_st_joint_names;
      arrToVar(length_st_joint_names, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_joint_names; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_joint_names-1]=0;
      this->st_joint_names = (char *)(inbuffer + offset-1);
      offset += length_st_joint_names;
        memcpy( &(this->joint_names[i]), &(this->st_joint_names), sizeof(std::string));
      }
      uint32_t points_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      points_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      points_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      points_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->points_length);
      if(points_lengthT > points_length)
        this->points = (trajectory_msgs::MultiDOFJointTrajectoryPoint*)realloc(this->points, points_lengthT * sizeof(trajectory_msgs::MultiDOFJointTrajectoryPoint));
      points_length = points_lengthT;
      for( uint32_t i = 0; i < points_length; i++){
      offset += this->st_points.deserialize(inbuffer + offset);
        memcpy( &(this->points[i]), &(this->st_points), sizeof(trajectory_msgs::MultiDOFJointTrajectoryPoint));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->joint_names_length);
      for( uint32_t i = 0; i < joint_names_length; i++){
      uint32_t length_joint_namesi = this->joint_names[i].size();
      length += 4;
      length += length_joint_namesi;
      }
      length += sizeof(this->points_length);
      for( uint32_t i = 0; i < points_length; i++){
      length += this->points[i].serializedLength();
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\": {";
      string_echo += this->header.echo();
      string_echo += "}, ";
      string_echo += "joint_names: [";
      for( uint32_t i = 0; i < joint_names_length; i++){
      if( i == (joint_names_length - 1)) {
      string_echo += "\"joint_names[i]\": \"";
      string_echo += joint_names[i];
      string_echo += "\"";
      } else {
      string_echo += "\"joint_names[i]\": \"";
      string_echo += joint_names[i];
      string_echo += "\", ";
      }
      }
      string_echo += "], ";
      string_echo += "points: [";
      for( uint32_t i = 0; i < points_length; i++){
      if( i == (points_length - 1)) {
      std::stringstream ss_pointsi; ss_pointsi << "{\"points" << i <<"\": {";
      string_echo += ss_pointsi.str();
      string_echo += this->points[i].echo();
      string_echo += "}}";
      } else {
      std::stringstream ss_pointsi; ss_pointsi << "{\"points" << i <<"\": {";
      string_echo += ss_pointsi.str();
      string_echo += this->points[i].echo();
      string_echo += "}}, ";
      }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "trajectory_msgs/MultiDOFJointTrajectory"; }
    virtual std::string getMD5(){ return "19f97d4cdcae9d0ef55126c1495f6c91"; }

  };

}
#endif
