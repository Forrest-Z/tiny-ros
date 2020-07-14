#ifndef _TINYROS_nav_msgs_OccupancyGrid_h
#define _TINYROS_nav_msgs_OccupancyGrid_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"
#include "tiny_ros/nav_msgs/MapMetaData.h"

namespace nav_msgs
{

  class OccupancyGrid : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef nav_msgs::MapMetaData _info_type;
      _info_type info;
      uint32_t data_length;
      typedef int8_t _data_type;
      _data_type st_data;
      _data_type * data;

    OccupancyGrid():
      header(),
      info(),
      data_length(0), data(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      offset += this->info.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->data_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->data_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->data_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->data_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->data_length);
      for( uint32_t i = 0; i < data_length; i++) {
        union {
          int8_t real;
          uint8_t base;
        } u_datai;
        u_datai.real = this->data[i];
        *(outbuffer + offset + 0) = (u_datai.base >> (8 * 0)) & 0xFF;
        offset += sizeof(this->data[i]);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      offset += this->info.deserialize(inbuffer + offset);
      uint32_t data_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      data_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      data_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      data_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->data_length);
      if(data_lengthT > data_length)
        this->data = (int8_t*)realloc(this->data, data_lengthT * sizeof(int8_t));
      data_length = data_lengthT;
      for( uint32_t i = 0; i < data_length; i++) {
        union {
          int8_t real;
          uint8_t base;
        } u_st_data;
        u_st_data.base = 0;
        u_st_data.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
        this->st_data = u_st_data.real;
        offset += sizeof(this->st_data);
        memcpy( &(this->data[i]), &(this->st_data), sizeof(int8_t));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += this->info.serializedLength();
      length += sizeof(this->data_length);
      for( uint32_t i = 0; i < data_length; i++) {
        length += sizeof(this->data[i]);
      }
      return length;
    }

    virtual tinyros::string getType(){ return "nav_msgs/OccupancyGrid"; }
    virtual tinyros::string getMD5(){ return "e489a26457224a97799696f3642f16a0"; }

  };

}
#endif
