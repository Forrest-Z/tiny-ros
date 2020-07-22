#ifndef _TINYROS_tinyros_msgs_TopicInfo_h
#define _TINYROS_tinyros_msgs_TopicInfo_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"

namespace tinyros_msgs
{

  class TopicInfo : public tinyros::Msg
  {
    public:
      typedef uint32_t _topic_id_type;
      _topic_id_type topic_id;
      typedef std::string _topic_name_type;
      _topic_name_type topic_name;
      typedef std::string _message_type_type;
      _message_type_type message_type;
      typedef std::string _md5sum_type;
      _md5sum_type md5sum;
      typedef int32_t _buffer_size_type;
      _buffer_size_type buffer_size;
      typedef bool _negotiated_type;
      _negotiated_type negotiated;
      typedef std::string _node_type;
      _node_type node;
      enum { ID_PUBLISHER = 0 };
      enum { ID_SUBSCRIBER = 1 };
      enum { ID_SERVICE_SERVER = 2 };
      enum { ID_SERVICE_CLIENT = 4 };
      enum { ID_ROSTOPIC_REQUEST = 6 };
      enum { ID_ROSSERVICE_REQUEST = 7 };
      enum { ID_LOG = 8 };
      enum { ID_TIME = 9 };
      enum { ID_NEGOTIATED = 10 };
      enum { ID_SESSION_ID = 11 };

    TopicInfo():
      topic_id(0),
      topic_name(""),
      message_type(""),
      md5sum(""),
      buffer_size(0),
      negotiated(0),
      node("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      *(outbuffer + offset + 0) = (this->topic_id >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->topic_id >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->topic_id >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->topic_id >> (8 * 3)) & 0xFF;
      offset += sizeof(this->topic_id);
      uint32_t length_topic_name = this->topic_name.size();
      varToArr(outbuffer + offset, length_topic_name);
      offset += 4;
      memcpy(outbuffer + offset, this->topic_name.c_str(), length_topic_name);
      offset += length_topic_name;
      uint32_t length_message_type = this->message_type.size();
      varToArr(outbuffer + offset, length_message_type);
      offset += 4;
      memcpy(outbuffer + offset, this->message_type.c_str(), length_message_type);
      offset += length_message_type;
      uint32_t length_md5sum = this->md5sum.size();
      varToArr(outbuffer + offset, length_md5sum);
      offset += 4;
      memcpy(outbuffer + offset, this->md5sum.c_str(), length_md5sum);
      offset += length_md5sum;
      union {
        int32_t real;
        uint32_t base;
      } u_buffer_size;
      u_buffer_size.real = this->buffer_size;
      *(outbuffer + offset + 0) = (u_buffer_size.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_buffer_size.base >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (u_buffer_size.base >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (u_buffer_size.base >> (8 * 3)) & 0xFF;
      offset += sizeof(this->buffer_size);
      union {
        bool real;
        uint8_t base;
      } u_negotiated;
      u_negotiated.real = this->negotiated;
      *(outbuffer + offset + 0) = (u_negotiated.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->negotiated);
      uint32_t length_node = this->node.size();
      varToArr(outbuffer + offset, length_node);
      offset += 4;
      memcpy(outbuffer + offset, this->node.c_str(), length_node);
      offset += length_node;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      this->topic_id =  ((uint32_t) (*(inbuffer + offset)));
      this->topic_id |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->topic_id |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->topic_id |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->topic_id);
      uint32_t length_topic_name;
      arrToVar(length_topic_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_topic_name; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_topic_name-1]=0;
      this->topic_name = (char *)(inbuffer + offset-1);
      offset += length_topic_name;
      uint32_t length_message_type;
      arrToVar(length_message_type, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_message_type; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_message_type-1]=0;
      this->message_type = (char *)(inbuffer + offset-1);
      offset += length_message_type;
      uint32_t length_md5sum;
      arrToVar(length_md5sum, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_md5sum; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_md5sum-1]=0;
      this->md5sum = (char *)(inbuffer + offset-1);
      offset += length_md5sum;
      union {
        int32_t real;
        uint32_t base;
      } u_buffer_size;
      u_buffer_size.base = 0;
      u_buffer_size.base |= ((uint32_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_buffer_size.base |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      u_buffer_size.base |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      u_buffer_size.base |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      this->buffer_size = u_buffer_size.real;
      offset += sizeof(this->buffer_size);
      union {
        bool real;
        uint8_t base;
      } u_negotiated;
      u_negotiated.base = 0;
      u_negotiated.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->negotiated = u_negotiated.real;
      offset += sizeof(this->negotiated);
      uint32_t length_node;
      arrToVar(length_node, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_node; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_node-1]=0;
      this->node = (char *)(inbuffer + offset-1);
      offset += length_node;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += sizeof(this->topic_id);
      uint32_t length_topic_name = this->topic_name.size();
      length += 4;
      length += length_topic_name;
      uint32_t length_message_type = this->message_type.size();
      length += 4;
      length += length_message_type;
      uint32_t length_md5sum = this->md5sum.size();
      length += 4;
      length += length_md5sum;
      length += sizeof(this->buffer_size);
      length += sizeof(this->negotiated);
      uint32_t length_node = this->node.size();
      length += 4;
      length += length_node;
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      std::stringstream ss_topic_id; ss_topic_id << "\"topic_id\":" << topic_id <<",";
      string_echo += ss_topic_id.str();
      std::size_t topic_name_pos = topic_name.find("\"");
      while(topic_name_pos != std::string::npos){
        topic_name.replace(topic_name_pos, 1,"\\\"");
        topic_name_pos = topic_name.find("\"", topic_name_pos+2);
      }
      string_echo += "\"topic_name\":\"";
      string_echo += topic_name;
      string_echo += "\",";
      std::size_t message_type_pos = message_type.find("\"");
      while(message_type_pos != std::string::npos){
        message_type.replace(message_type_pos, 1,"\\\"");
        message_type_pos = message_type.find("\"", message_type_pos+2);
      }
      string_echo += "\"message_type\":\"";
      string_echo += message_type;
      string_echo += "\",";
      std::size_t md5sum_pos = md5sum.find("\"");
      while(md5sum_pos != std::string::npos){
        md5sum.replace(md5sum_pos, 1,"\\\"");
        md5sum_pos = md5sum.find("\"", md5sum_pos+2);
      }
      string_echo += "\"md5sum\":\"";
      string_echo += md5sum;
      string_echo += "\",";
      std::stringstream ss_buffer_size; ss_buffer_size << "\"buffer_size\":" << buffer_size <<",";
      string_echo += ss_buffer_size.str();
      std::stringstream ss_negotiated; ss_negotiated << "\"negotiated\":" << negotiated <<",";
      string_echo += ss_negotiated.str();
      std::size_t node_pos = node.find("\"");
      while(node_pos != std::string::npos){
        node.replace(node_pos, 1,"\\\"");
        node_pos = node.find("\"", node_pos+2);
      }
      string_echo += "\"node\":\"";
      string_echo += node;
      string_echo += "\"";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "tinyros_msgs/TopicInfo"; }
    virtual std::string getMD5(){ return "76d40676946fcde66f228def7575451a"; }

  };

}
#endif
