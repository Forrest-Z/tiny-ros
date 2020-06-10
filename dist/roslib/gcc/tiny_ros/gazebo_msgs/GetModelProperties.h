#ifndef _TINYROS_SERVICE_GetModelProperties_h
#define _TINYROS_SERVICE_GetModelProperties_h
#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace gazebo_msgs
{

static const char GETMODELPROPERTIES[] = "gazebo_msgs/GetModelProperties";

  class GetModelPropertiesRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef std::string _model_name_type;
      _model_name_type model_name;

    GetModelPropertiesRequest():
      model_name("")
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      uint32_t length_model_name = this->model_name.size();
      varToArr(outbuffer + offset, length_model_name);
      offset += 4;
      memcpy(outbuffer + offset, this->model_name.c_str(), length_model_name);
      offset += length_model_name;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      uint32_t length_model_name;
      arrToVar(length_model_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_model_name; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_model_name-1]=0;
      this->model_name = (char *)(inbuffer + offset-1);
      offset += length_model_name;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_model_name = this->model_name.size();
      length += 4;
      length += length_model_name;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"model_name\": \"";
      string_echo += model_name;
      string_echo += "\"";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return GETMODELPROPERTIES; }
    virtual std::string getMD5(){ return "fe0194bf75c917c89b820b09c12fe6c1"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetModelPropertiesResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef std::string _parent_model_name_type;
      _parent_model_name_type parent_model_name;
      typedef std::string _canonical_body_name_type;
      _canonical_body_name_type canonical_body_name;
      uint32_t body_names_length;
      typedef std::string _body_names_type;
      _body_names_type st_body_names;
      _body_names_type * body_names;
      uint32_t geom_names_length;
      typedef std::string _geom_names_type;
      _geom_names_type st_geom_names;
      _geom_names_type * geom_names;
      uint32_t joint_names_length;
      typedef std::string _joint_names_type;
      _joint_names_type st_joint_names;
      _joint_names_type * joint_names;
      uint32_t child_model_names_length;
      typedef std::string _child_model_names_type;
      _child_model_names_type st_child_model_names;
      _child_model_names_type * child_model_names;
      typedef bool _is_static_type;
      _is_static_type is_static;
      typedef bool _success_type;
      _success_type success;
      typedef std::string _status_message_type;
      _status_message_type status_message;

    GetModelPropertiesResponse():
      parent_model_name(""),
      canonical_body_name(""),
      body_names_length(0), body_names(NULL),
      geom_names_length(0), geom_names(NULL),
      joint_names_length(0), joint_names(NULL),
      child_model_names_length(0), child_model_names(NULL),
      is_static(0),
      success(0),
      status_message("")
    {
      this->__id__ = 0;
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;
      offset += sizeof(this->__id__);
      uint32_t length_parent_model_name = this->parent_model_name.size();
      varToArr(outbuffer + offset, length_parent_model_name);
      offset += 4;
      memcpy(outbuffer + offset, this->parent_model_name.c_str(), length_parent_model_name);
      offset += length_parent_model_name;
      uint32_t length_canonical_body_name = this->canonical_body_name.size();
      varToArr(outbuffer + offset, length_canonical_body_name);
      offset += 4;
      memcpy(outbuffer + offset, this->canonical_body_name.c_str(), length_canonical_body_name);
      offset += length_canonical_body_name;
      *(outbuffer + offset + 0) = (this->body_names_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->body_names_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->body_names_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->body_names_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->body_names_length);
      for( uint32_t i = 0; i < body_names_length; i++){
      uint32_t length_body_namesi = this->body_names[i].size();
      varToArr(outbuffer + offset, length_body_namesi);
      offset += 4;
      memcpy(outbuffer + offset, this->body_names[i].c_str(), length_body_namesi);
      offset += length_body_namesi;
      }
      *(outbuffer + offset + 0) = (this->geom_names_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->geom_names_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->geom_names_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->geom_names_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->geom_names_length);
      for( uint32_t i = 0; i < geom_names_length; i++){
      uint32_t length_geom_namesi = this->geom_names[i].size();
      varToArr(outbuffer + offset, length_geom_namesi);
      offset += 4;
      memcpy(outbuffer + offset, this->geom_names[i].c_str(), length_geom_namesi);
      offset += length_geom_namesi;
      }
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
      *(outbuffer + offset + 0) = (this->child_model_names_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->child_model_names_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->child_model_names_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->child_model_names_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->child_model_names_length);
      for( uint32_t i = 0; i < child_model_names_length; i++){
      uint32_t length_child_model_namesi = this->child_model_names[i].size();
      varToArr(outbuffer + offset, length_child_model_namesi);
      offset += 4;
      memcpy(outbuffer + offset, this->child_model_names[i].c_str(), length_child_model_namesi);
      offset += length_child_model_namesi;
      }
      union {
        bool real;
        uint8_t base;
      } u_is_static;
      u_is_static.real = this->is_static;
      *(outbuffer + offset + 0) = (u_is_static.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->is_static);
      union {
        bool real;
        uint8_t base;
      } u_success;
      u_success.real = this->success;
      *(outbuffer + offset + 0) = (u_success.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->success);
      uint32_t length_status_message = this->status_message.size();
      varToArr(outbuffer + offset, length_status_message);
      offset += 4;
      memcpy(outbuffer + offset, this->status_message.c_str(), length_status_message);
      offset += length_status_message;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->__id__);
      uint32_t length_parent_model_name;
      arrToVar(length_parent_model_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_parent_model_name; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_parent_model_name-1]=0;
      this->parent_model_name = (char *)(inbuffer + offset-1);
      offset += length_parent_model_name;
      uint32_t length_canonical_body_name;
      arrToVar(length_canonical_body_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_canonical_body_name; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_canonical_body_name-1]=0;
      this->canonical_body_name = (char *)(inbuffer + offset-1);
      offset += length_canonical_body_name;
      uint32_t body_names_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      body_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      body_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      body_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->body_names_length);
      if(body_names_lengthT > body_names_length)
        this->body_names = (std::string*)realloc(this->body_names, body_names_lengthT * sizeof(std::string));
      body_names_length = body_names_lengthT;
      for( uint32_t i = 0; i < body_names_length; i++){
      uint32_t length_st_body_names;
      arrToVar(length_st_body_names, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_body_names; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_body_names-1]=0;
      this->st_body_names = (char *)(inbuffer + offset-1);
      offset += length_st_body_names;
        memcpy( &(this->body_names[i]), &(this->st_body_names), sizeof(std::string));
      }
      uint32_t geom_names_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      geom_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      geom_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      geom_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->geom_names_length);
      if(geom_names_lengthT > geom_names_length)
        this->geom_names = (std::string*)realloc(this->geom_names, geom_names_lengthT * sizeof(std::string));
      geom_names_length = geom_names_lengthT;
      for( uint32_t i = 0; i < geom_names_length; i++){
      uint32_t length_st_geom_names;
      arrToVar(length_st_geom_names, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_geom_names; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_geom_names-1]=0;
      this->st_geom_names = (char *)(inbuffer + offset-1);
      offset += length_st_geom_names;
        memcpy( &(this->geom_names[i]), &(this->st_geom_names), sizeof(std::string));
      }
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
      uint32_t child_model_names_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      child_model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      child_model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      child_model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->child_model_names_length);
      if(child_model_names_lengthT > child_model_names_length)
        this->child_model_names = (std::string*)realloc(this->child_model_names, child_model_names_lengthT * sizeof(std::string));
      child_model_names_length = child_model_names_lengthT;
      for( uint32_t i = 0; i < child_model_names_length; i++){
      uint32_t length_st_child_model_names;
      arrToVar(length_st_child_model_names, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_st_child_model_names; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_st_child_model_names-1]=0;
      this->st_child_model_names = (char *)(inbuffer + offset-1);
      offset += length_st_child_model_names;
        memcpy( &(this->child_model_names[i]), &(this->st_child_model_names), sizeof(std::string));
      }
      union {
        bool real;
        uint8_t base;
      } u_is_static;
      u_is_static.base = 0;
      u_is_static.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->is_static = u_is_static.real;
      offset += sizeof(this->is_static);
      union {
        bool real;
        uint8_t base;
      } u_success;
      u_success.base = 0;
      u_success.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->success = u_success.real;
      offset += sizeof(this->success);
      uint32_t length_status_message;
      arrToVar(length_status_message, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_status_message; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_status_message-1]=0;
      this->status_message = (char *)(inbuffer + offset-1);
      offset += length_status_message;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      uint32_t length_parent_model_name = this->parent_model_name.size();
      length += 4;
      length += length_parent_model_name;
      uint32_t length_canonical_body_name = this->canonical_body_name.size();
      length += 4;
      length += length_canonical_body_name;
      length += sizeof(this->body_names_length);
      for( uint32_t i = 0; i < body_names_length; i++){
      uint32_t length_body_namesi = this->body_names[i].size();
      length += 4;
      length += length_body_namesi;
      }
      length += sizeof(this->geom_names_length);
      for( uint32_t i = 0; i < geom_names_length; i++){
      uint32_t length_geom_namesi = this->geom_names[i].size();
      length += 4;
      length += length_geom_namesi;
      }
      length += sizeof(this->joint_names_length);
      for( uint32_t i = 0; i < joint_names_length; i++){
      uint32_t length_joint_namesi = this->joint_names[i].size();
      length += 4;
      length += length_joint_namesi;
      }
      length += sizeof(this->child_model_names_length);
      for( uint32_t i = 0; i < child_model_names_length; i++){
      uint32_t length_child_model_namesi = this->child_model_names[i].size();
      length += 4;
      length += length_child_model_namesi;
      }
      length += sizeof(this->is_static);
      length += sizeof(this->success);
      uint32_t length_status_message = this->status_message.size();
      length += 4;
      length += length_status_message;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"parent_model_name\": \"";
      string_echo += parent_model_name;
      string_echo += "\", ";
      string_echo += "\"canonical_body_name\": \"";
      string_echo += canonical_body_name;
      string_echo += "\", ";
      string_echo += "body_names: [";
      for( uint32_t i = 0; i < body_names_length; i++){
      if( i == (body_names_length - 1)) {
      string_echo += "\"body_names[i]\": \"";
      string_echo += body_names[i];
      string_echo += "\"";
      } else {
      string_echo += "\"body_names[i]\": \"";
      string_echo += body_names[i];
      string_echo += "\", ";
      }
      }
      string_echo += "], ";
      string_echo += "geom_names: [";
      for( uint32_t i = 0; i < geom_names_length; i++){
      if( i == (geom_names_length - 1)) {
      string_echo += "\"geom_names[i]\": \"";
      string_echo += geom_names[i];
      string_echo += "\"";
      } else {
      string_echo += "\"geom_names[i]\": \"";
      string_echo += geom_names[i];
      string_echo += "\", ";
      }
      }
      string_echo += "], ";
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
      string_echo += "child_model_names: [";
      for( uint32_t i = 0; i < child_model_names_length; i++){
      if( i == (child_model_names_length - 1)) {
      string_echo += "\"child_model_names[i]\": \"";
      string_echo += child_model_names[i];
      string_echo += "\"";
      } else {
      string_echo += "\"child_model_names[i]\": \"";
      string_echo += child_model_names[i];
      string_echo += "\", ";
      }
      }
      string_echo += "], ";
      std::stringstream ss_is_static; ss_is_static << "\"is_static\": " << is_static <<", ";
      string_echo += ss_is_static.str();
      std::stringstream ss_success; ss_success << "\"success\": " << success <<", ";
      string_echo += ss_success.str();
      string_echo += "\"status_message\": \"";
      string_echo += status_message;
      string_echo += "\"";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return GETMODELPROPERTIES; }
    virtual std::string getMD5(){ return "d8f16b08abaf0220a551cf9025748602"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetModelProperties {
    public:
    typedef GetModelPropertiesRequest Request;
    typedef GetModelPropertiesResponse Response;
  };

}
#endif
