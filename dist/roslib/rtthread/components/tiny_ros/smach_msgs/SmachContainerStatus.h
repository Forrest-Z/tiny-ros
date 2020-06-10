#ifndef _TINYROS_smach_msgs_SmachContainerStatus_h
#define _TINYROS_smach_msgs_SmachContainerStatus_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"

namespace smach_msgs
{

  class SmachContainerStatus : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef tinyros::string _path_type;
      _path_type path;
      uint32_t initial_states_length;
      typedef tinyros::string _initial_states_type;
      _initial_states_type st_initial_states;
      _initial_states_type * initial_states;
      uint32_t active_states_length;
      typedef tinyros::string _active_states_type;
      _active_states_type st_active_states;
      _active_states_type * active_states;
      typedef tinyros::string _local_data_type;
      _local_data_type local_data;
      typedef tinyros::string _info_type;
      _info_type info;

    SmachContainerStatus():
      header(),
      path(""),
      initial_states_length(0), initial_states(NULL),
      active_states_length(0), active_states(NULL),
      local_data(""),
      info("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      uint32_t length_path = this->path.size();
      varToArr(outbuffer + offset, length_path);
      offset += 4;
      memcpy(outbuffer + offset, this->path.c_str(), length_path);
      offset += length_path;
      *(outbuffer + offset + 0) = (this->initial_states_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->initial_states_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->initial_states_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->initial_states_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->initial_states_length);
      for( uint32_t i = 0; i < initial_states_length; i++){
      uint32_t length_initial_statesi = this->initial_states[i].size();
      varToArr(outbuffer + offset, length_initial_statesi);
      offset += 4;
      memcpy(outbuffer + offset, this->initial_states[i].c_str(), length_initial_statesi);
      offset += length_initial_statesi;
      }
      *(outbuffer + offset + 0) = (this->active_states_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->active_states_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->active_states_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->active_states_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->active_states_length);
      for( uint32_t i = 0; i < active_states_length; i++){
      uint32_t length_active_statesi = this->active_states[i].size();
      varToArr(outbuffer + offset, length_active_statesi);
      offset += 4;
      memcpy(outbuffer + offset, this->active_states[i].c_str(), length_active_statesi);
      offset += length_active_statesi;
      }
      uint32_t length_local_data = this->local_data.size();
      varToArr(outbuffer + offset, length_local_data);
      offset += 4;
      memcpy(outbuffer + offset, this->local_data.c_str(), length_local_data);
      offset += length_local_data;
      uint32_t length_info = this->info.size();
      varToArr(outbuffer + offset, length_info);
      offset += 4;
      memcpy(outbuffer + offset, this->info.c_str(), length_info);
      offset += length_info;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t length_path;
      arrToVar(length_path, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_path; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_path-1]=0;
      this->path = (char *)(inbuffer + offset-1);
      offset += length_path;
      uint32_t initial_states_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      initial_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      initial_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      initial_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->initial_states_length);
      if(initial_states_lengthT > initial_states_length)
        this->initial_states = (tinyros::string*)realloc(this->initial_states, initial_states_lengthT * sizeof(tinyros::string));
      initial_states_length = initial_states_lengthT;
      for( uint32_t i = 0; i < initial_states_length; i++){
      uint32_t length_st_initial_states;
      arrToVar(length_st_initial_states, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_initial_states; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_initial_states-1]=0;
      this->st_initial_states = (char *)(inbuffer + offset-1);
      offset += length_st_initial_states;
        memcpy( &(this->initial_states[i]), &(this->st_initial_states), sizeof(tinyros::string));
      }
      uint32_t active_states_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      active_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      active_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      active_states_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->active_states_length);
      if(active_states_lengthT > active_states_length)
        this->active_states = (tinyros::string*)realloc(this->active_states, active_states_lengthT * sizeof(tinyros::string));
      active_states_length = active_states_lengthT;
      for( uint32_t i = 0; i < active_states_length; i++){
      uint32_t length_st_active_states;
      arrToVar(length_st_active_states, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_active_states; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_active_states-1]=0;
      this->st_active_states = (char *)(inbuffer + offset-1);
      offset += length_st_active_states;
        memcpy( &(this->active_states[i]), &(this->st_active_states), sizeof(tinyros::string));
      }
      uint32_t length_local_data;
      arrToVar(length_local_data, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_local_data; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_local_data-1]=0;
      this->local_data = (char *)(inbuffer + offset-1);
      offset += length_local_data;
      uint32_t length_info;
      arrToVar(length_info, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_info; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_info-1]=0;
      this->info = (char *)(inbuffer + offset-1);
      offset += length_info;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      uint32_t length_path = this->path.size();
      length += 4;
      length += length_path;
      length += sizeof(this->initial_states_length);
      for( uint32_t i = 0; i < initial_states_length; i++){
      uint32_t length_initial_statesi = this->initial_states[i].size();
      length += 4;
      length += length_initial_statesi;
      }
      length += sizeof(this->active_states_length);
      for( uint32_t i = 0; i < active_states_length; i++){
      uint32_t length_active_statesi = this->active_states[i].size();
      length += 4;
      length += length_active_statesi;
      }
      uint32_t length_local_data = this->local_data.size();
      length += 4;
      length += length_local_data;
      uint32_t length_info = this->info.size();
      length += 4;
      length += length_info;
      return length;
    }

    virtual tinyros::string getType(){ return "smach_msgs/SmachContainerStatus"; }
    virtual tinyros::string getMD5(){ return "3e74cf72da18311be27e7a5970ea6415"; }

  };

}
#endif
