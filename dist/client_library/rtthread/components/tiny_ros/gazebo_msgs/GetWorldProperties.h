#ifndef _TINYROS_SERVICE_GetWorldProperties_h
#define _TINYROS_SERVICE_GetWorldProperties_h
#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace gazebo_msgs
{

static const char GETWORLDPROPERTIES[] = "gazebo_msgs/GetWorldProperties";

  class GetWorldPropertiesRequest : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:

    GetWorldPropertiesRequest()
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
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      return length;
    }

    virtual tinyros::string getType(){ return GETWORLDPROPERTIES; }
    virtual tinyros::string getMD5(){ return "3aa5de7106eec5dae41ad1c9ae681123"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetWorldPropertiesResponse : public tinyros::Msg
  {
    private:
      typedef uint32_t ___id___type;
      ___id___type __id__;

    public:
      typedef double _sim_time_type;
      _sim_time_type sim_time;
      uint32_t model_names_length;
      typedef tinyros::string _model_names_type;
      _model_names_type st_model_names;
      _model_names_type * model_names;
      typedef bool _rendering_enabled_type;
      _rendering_enabled_type rendering_enabled;
      typedef bool _success_type;
      _success_type success;
      typedef tinyros::string _status_message_type;
      _status_message_type status_message;

    GetWorldPropertiesResponse():
      sim_time(0),
      model_names_length(0), model_names(NULL),
      rendering_enabled(0),
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
      union {
        double real;
        uint64_t base;
      } u_sim_time;
      u_sim_time.real = this->sim_time;
      *(outbuffer + offset + 0) = (u_sim_time.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_sim_time.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_sim_time.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_sim_time.base >> (8 * 3)) & 0xFF;
      *(outbuffer + offset + 4) = (u_sim_time.base >> (8 * 4)) & 0xFF;
      *(outbuffer + offset + 5) = (u_sim_time.base >> (8 * 5)) & 0xFF;
      *(outbuffer + offset + 6) = (u_sim_time.base >> (8 * 6)) & 0xFF;
      *(outbuffer + offset + 7) = (u_sim_time.base >> (8 * 7)) & 0xFF;
      offset += sizeof(this->sim_time);
      *(outbuffer + offset + 0) = (this->model_names_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->model_names_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->model_names_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->model_names_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->model_names_length);
      for( uint32_t i = 0; i < model_names_length; i++) {
        uint32_t length_model_namesi = this->model_names[i].size();
        varToArr(outbuffer + offset, length_model_namesi);
        offset += 4;
        memcpy(outbuffer + offset, this->model_names[i].c_str(), length_model_namesi);
        offset += length_model_namesi;
      }
      union {
        bool real;
        uint8_t base;
      } u_rendering_enabled;
      u_rendering_enabled.real = this->rendering_enabled;
      *(outbuffer + offset + 0) = (u_rendering_enabled.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->rendering_enabled);
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
      union {
        double real;
        uint64_t base;
      } u_sim_time;
      u_sim_time.base = 0;
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 3))) << (8 * 3);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 4))) << (8 * 4);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 5))) << (8 * 5);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 6))) << (8 * 6);
      u_sim_time.base |= ((uint64_t) (*(inbuffer + offset + 7))) << (8 * 7);
      this->sim_time = u_sim_time.real;
      offset += sizeof(this->sim_time);
      uint32_t model_names_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      model_names_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->model_names_length);
      if(model_names_lengthT > model_names_length)
        this->model_names = (tinyros::string*)realloc(this->model_names, model_names_lengthT * sizeof(tinyros::string));
      model_names_length = model_names_lengthT;
      for( uint32_t i = 0; i < model_names_length; i++) {
        uint32_t length_st_model_names;
        arrToVar(length_st_model_names, (inbuffer + offset));
        offset += 4;
        for(unsigned int k= offset; k< offset+length_st_model_names; ++k){
          inbuffer[k-1]=inbuffer[k];
        }
        inbuffer[offset+length_st_model_names-1]=0;
        this->st_model_names = (char *)(inbuffer + offset-1);
        offset += length_st_model_names;
        memcpy( &(this->model_names[i]), &(this->st_model_names), sizeof(tinyros::string));
      }
      union {
        bool real;
        uint8_t base;
      } u_rendering_enabled;
      u_rendering_enabled.base = 0;
      u_rendering_enabled.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->rendering_enabled = u_rendering_enabled.real;
      offset += sizeof(this->rendering_enabled);
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
      length += sizeof(this->sim_time);
      length += sizeof(this->model_names_length);
      for( uint32_t i = 0; i < model_names_length; i++) {
        uint32_t length_model_namesi = this->model_names[i].size();
        length += 4;
        length += length_model_namesi;
      }
      length += sizeof(this->rendering_enabled);
      length += sizeof(this->success);
      uint32_t length_status_message = this->status_message.size();
      length += 4;
      length += length_status_message;
      return length;
    }

    virtual tinyros::string getType(){ return GETWORLDPROPERTIES; }
    virtual tinyros::string getMD5(){ return "fe944c1c210919291ad14bc43b6c10cf"; }
    uint32_t getID() const { return this->__id__; }
    void setID(uint32_t id){ this->__id__ = id; }

  };

  class GetWorldProperties {
    public:
    typedef GetWorldPropertiesRequest Request;
    typedef GetWorldPropertiesResponse Response;
  };

}
#endif
