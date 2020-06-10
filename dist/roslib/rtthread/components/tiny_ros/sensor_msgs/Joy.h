#ifndef _TINYROS_sensor_msgs_Joy_h
#define _TINYROS_sensor_msgs_Joy_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"

namespace sensor_msgs
{

  class Joy : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      uint32_t axes_length;
      typedef float _axes_type;
      _axes_type st_axes;
      _axes_type * axes;
      uint32_t buttons_length;
      typedef int32_t _buttons_type;
      _buttons_type st_buttons;
      _buttons_type * buttons;

    Joy():
      header(),
      axes_length(0), axes(NULL),
      buttons_length(0), buttons(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->axes_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->axes_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->axes_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->axes_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->axes_length);
      for( uint32_t i = 0; i < axes_length; i++){
      union {
        float real;
        uint32_t base;
      } u_axesi;
      u_axesi.real = this->axes[i];
      *(outbuffer + offset + 0) = (u_axesi.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_axesi.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_axesi.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_axesi.base >> (8 * 3)) & 0xFF;
      offset += sizeof(this->axes[i]);
      }
      *(outbuffer + offset + 0) = (this->buttons_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->buttons_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->buttons_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->buttons_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->buttons_length);
      for( uint32_t i = 0; i < buttons_length; i++){
      union {
        int32_t real;
        uint32_t base;
      } u_buttonsi;
      u_buttonsi.real = this->buttons[i];
      *(outbuffer + offset + 0) = (u_buttonsi.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_buttonsi.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_buttonsi.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_buttonsi.base >> (8 * 3)) & 0xFF;
      offset += sizeof(this->buttons[i]);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      uint32_t axes_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      axes_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      axes_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      axes_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->axes_length);
      if(axes_lengthT > axes_length)
        this->axes = (float*)realloc(this->axes, axes_lengthT * sizeof(float));
      axes_length = axes_lengthT;
      for( uint32_t i = 0; i < axes_length; i++){
      union {
        float real;
        uint32_t base;
      } u_st_axes;
      u_st_axes.base = 0;
      u_st_axes.base |= ((uint32_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_st_axes.base |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_st_axes.base |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_st_axes.base |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      this->st_axes = u_st_axes.real;
      offset += sizeof(this->st_axes);
        memcpy( &(this->axes[i]), &(this->st_axes), sizeof(float));
      }
      uint32_t buttons_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      buttons_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      buttons_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      buttons_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->buttons_length);
      if(buttons_lengthT > buttons_length)
        this->buttons = (int32_t*)realloc(this->buttons, buttons_lengthT * sizeof(int32_t));
      buttons_length = buttons_lengthT;
      for( uint32_t i = 0; i < buttons_length; i++){
      union {
        int32_t real;
        uint32_t base;
      } u_st_buttons;
      u_st_buttons.base = 0;
      u_st_buttons.base |= ((uint32_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_st_buttons.base |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_st_buttons.base |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_st_buttons.base |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      this->st_buttons = u_st_buttons.real;
      offset += sizeof(this->st_buttons);
        memcpy( &(this->buttons[i]), &(this->st_buttons), sizeof(int32_t));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->axes_length);
      for( uint32_t i = 0; i < axes_length; i++){
      length += sizeof(this->axes[i]);
      }
      length += sizeof(this->buttons_length);
      for( uint32_t i = 0; i < buttons_length; i++){
      length += sizeof(this->buttons[i]);
      }
      return length;
    }

    virtual tinyros::string getType(){ return "sensor_msgs/Joy"; }
    virtual tinyros::string getMD5(){ return "da3438323dbe92a4d6e4658e06bf8da1"; }

  };

}
#endif
