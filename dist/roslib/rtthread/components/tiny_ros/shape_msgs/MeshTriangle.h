#ifndef _TINYROS_shape_msgs_MeshTriangle_h
#define _TINYROS_shape_msgs_MeshTriangle_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace shape_msgs
{

  class MeshTriangle : public tinyros::Msg
  {
    public:
      uint32_t vertex_indices[3];

    MeshTriangle():
      vertex_indices()
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      for( uint32_t i = 0; i < 3; i++) {
        *(outbuffer + offset + 0) = (this->vertex_indices[i] >> (8 * 0)) & 0xFF;
        *(outbuffer + offset + 1) = (this->vertex_indices[i] >> (8 * 1)) & 0xFF;
        *(outbuffer + offset + 2) = (this->vertex_indices[i] >> (8 * 2)) & 0xFF;
        *(outbuffer + offset + 3) = (this->vertex_indices[i] >> (8 * 3)) & 0xFF;
        offset += sizeof(this->vertex_indices[i]);
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      for( uint32_t i = 0; i < 3; i++){
        this->vertex_indices[i] =  ((uint32_t) (*(inbuffer + offset)));
        this->vertex_indices[i] |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
        this->vertex_indices[i] |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
        this->vertex_indices[i] |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
        offset += sizeof(this->vertex_indices[i]);
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      for( uint32_t i = 0; i < 3; i++) {
        length += sizeof(this->vertex_indices[i]);
      }
      return length;
    }

    virtual tinyros::string getType(){ return "shape_msgs/MeshTriangle"; }
    virtual tinyros::string getMD5(){ return "01020cfeb9ad7679dd18bbd7149962ba"; }

  };

}
#endif
