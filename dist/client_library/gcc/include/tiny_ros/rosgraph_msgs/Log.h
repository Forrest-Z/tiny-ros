#ifndef _TINYROS_rosgraph_msgs_Log_h
#define _TINYROS_rosgraph_msgs_Log_h

#include <stdint.h>
#include <string>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/std_msgs/Header.h"

namespace rosgraph_msgs
{

  class Log : public tinyros::Msg
  {
    public:
      typedef std_msgs::Header _header_type;
      _header_type header;
      typedef int8_t _level_type;
      _level_type level;
      typedef std::string _name_type;
      _name_type name;
      typedef std::string _msg_type;
      _msg_type msg;
      typedef std::string _file_type;
      _file_type file;
      typedef std::string _function_type;
      _function_type function;
      typedef uint32_t _line_type;
      _line_type line;
      uint32_t topics_length;
      typedef std::string _topics_type;
      _topics_type st_topics;
      _topics_type * topics;
      enum { DEBUG = 1  };
      enum { INFO = 2   };
      enum { WARN = 4   };
      enum { ERROR = 8  };
      enum { FATAL = 16  };

    Log():
      header(),
      level(0),
      name(""),
      msg(""),
      file(""),
      function(""),
      line(0),
      topics_length(0), topics(NULL)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->header.serialize(outbuffer + offset);
      union {
        int8_t real;
        uint8_t base;
      } u_level;
      u_level.real = this->level;
      *(outbuffer + offset + 0) = (u_level.base >> (8 * 0)) & 0xFF;
      offset += sizeof(this->level);
      uint32_t length_name = this->name.size();
      varToArr(outbuffer + offset, length_name);
      offset += 4;
      memcpy(outbuffer + offset, this->name.c_str(), length_name);
      offset += length_name;
      uint32_t length_msg = this->msg.size();
      varToArr(outbuffer + offset, length_msg);
      offset += 4;
      memcpy(outbuffer + offset, this->msg.c_str(), length_msg);
      offset += length_msg;
      uint32_t length_file = this->file.size();
      varToArr(outbuffer + offset, length_file);
      offset += 4;
      memcpy(outbuffer + offset, this->file.c_str(), length_file);
      offset += length_file;
      uint32_t length_function = this->function.size();
      varToArr(outbuffer + offset, length_function);
      offset += 4;
      memcpy(outbuffer + offset, this->function.c_str(), length_function);
      offset += length_function;
      *(outbuffer + offset + 0) = (this->line >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->line >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->line >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->line >> (8 * 3)) & 0xFF;
      offset += sizeof(this->line);
      *(outbuffer + offset + 0) = (this->topics_length >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (this->topics_length >> (8 * 1)) & 0xFF;
      *(outbuffer + offset + 2) = (this->topics_length >> (8 * 2)) & 0xFF;
      *(outbuffer + offset + 3) = (this->topics_length >> (8 * 3)) & 0xFF;
      offset += sizeof(this->topics_length);
      for( uint32_t i = 0; i < topics_length; i++) {
        uint32_t length_topicsi = this->topics[i].size();
        varToArr(outbuffer + offset, length_topicsi);
        offset += 4;
        memcpy(outbuffer + offset, this->topics[i].c_str(), length_topicsi);
        offset += length_topicsi;
      }
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->header.deserialize(inbuffer + offset);
      union {
        int8_t real;
        uint8_t base;
      } u_level;
      u_level.base = 0;
      u_level.base |= ((uint8_t) (*(inbuffer + offset + 0))) << (8 * 0);
      this->level = u_level.real;
      offset += sizeof(this->level);
      uint32_t length_name;
      arrToVar(length_name, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_name; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_name-1]=0;
      this->name = (char *)(inbuffer + offset-1);
      offset += length_name;
      uint32_t length_msg;
      arrToVar(length_msg, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_msg; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_msg-1]=0;
      this->msg = (char *)(inbuffer + offset-1);
      offset += length_msg;
      uint32_t length_file;
      arrToVar(length_file, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_file; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_file-1]=0;
      this->file = (char *)(inbuffer + offset-1);
      offset += length_file;
      uint32_t length_function;
      arrToVar(length_function, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_function; ++k){
        inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_function-1]=0;
      this->function = (char *)(inbuffer + offset-1);
      offset += length_function;
      this->line =  ((uint32_t) (*(inbuffer + offset)));
      this->line |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->line |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);
      this->line |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);
      offset += sizeof(this->line);
      uint32_t topics_lengthT = ((uint32_t) (*(inbuffer + offset))); 
      topics_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); 
      topics_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); 
      topics_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); 
      offset += sizeof(this->topics_length);
      if(topics_lengthT > topics_length)
        this->topics = (std::string*)realloc(this->topics, topics_lengthT * sizeof(std::string));
      topics_length = topics_lengthT;
      for( uint32_t i = 0; i < topics_length; i++) {
        uint32_t length_st_topics;
        arrToVar(length_st_topics, (inbuffer + offset));
        offset += 4;
        for(unsigned int k= offset; k< offset+length_st_topics; ++k){
          inbuffer[k-1]=inbuffer[k];
        }
        inbuffer[offset+length_st_topics-1]=0;
        this->st_topics = (char *)(inbuffer + offset-1);
        offset += length_st_topics;
        memcpy( &(this->topics[i]), &(this->st_topics), sizeof(std::string));
      }
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->header.serializedLength();
      length += sizeof(this->level);
      uint32_t length_name = this->name.size();
      length += 4;
      length += length_name;
      uint32_t length_msg = this->msg.size();
      length += 4;
      length += length_msg;
      uint32_t length_file = this->file.size();
      length += 4;
      length += length_file;
      uint32_t length_function = this->function.size();
      length += 4;
      length += length_function;
      length += sizeof(this->line);
      length += sizeof(this->topics_length);
      for( uint32_t i = 0; i < topics_length; i++) {
        uint32_t length_topicsi = this->topics[i].size();
        length += 4;
        length += length_topicsi;
      }
      return length;
    }

    virtual std::string echo()
    {
      std::string string_echo = "{";
      string_echo += "\"header\":";
      string_echo += this->header.echo();
      string_echo += ",";
      std::stringstream ss_level; ss_level << "\"level\":" << (int16_t)level <<",";
      string_echo += ss_level.str();
      std::size_t name_pos = name.find("\"");
      while(name_pos != std::string::npos){
        name.replace(name_pos, 1,"\\\"");
        name_pos = name.find("\"", name_pos+2);
      }
      string_echo += "\"name\":\"";
      string_echo += name;
      string_echo += "\",";
      std::size_t msg_pos = msg.find("\"");
      while(msg_pos != std::string::npos){
        msg.replace(msg_pos, 1,"\\\"");
        msg_pos = msg.find("\"", msg_pos+2);
      }
      string_echo += "\"msg\":\"";
      string_echo += msg;
      string_echo += "\",";
      std::size_t file_pos = file.find("\"");
      while(file_pos != std::string::npos){
        file.replace(file_pos, 1,"\\\"");
        file_pos = file.find("\"", file_pos+2);
      }
      string_echo += "\"file\":\"";
      string_echo += file;
      string_echo += "\",";
      std::size_t function_pos = function.find("\"");
      while(function_pos != std::string::npos){
        function.replace(function_pos, 1,"\\\"");
        function_pos = function.find("\"", function_pos+2);
      }
      string_echo += "\"function\":\"";
      string_echo += function;
      string_echo += "\",";
      std::stringstream ss_line; ss_line << "\"line\":" << line <<",";
      string_echo += ss_line.str();
      string_echo += "topics:[";
      for( uint32_t i = 0; i < topics_length; i++) {
        if( i == (topics_length - 1)) {
          std::stringstream ss_topicsi; ss_topicsi << "\"";
          string_echo += ss_topicsi.str();
          std::size_t topicsi_pos = topics[i].find("\"");
          while(topicsi_pos != std::string::npos){
            topics[i].replace(topicsi_pos, 1,"\\\"");
            topicsi_pos = topics[i].find("\"", topicsi_pos+2);
          }
          string_echo += topics[i];
          string_echo += "\"";
        } else {
          std::stringstream ss_topicsi; ss_topicsi << "\"";
          string_echo += ss_topicsi.str();
          std::size_t topicsi_pos = topics[i].find("\"");
          while(topicsi_pos != std::string::npos){
            topics[i].replace(topicsi_pos, 1,"\\\"");
            topicsi_pos = topics[i].find("\"", topicsi_pos+2);
          }
          string_echo += topics[i];
          string_echo += "\",";
        }
      }
      string_echo += "]";
      string_echo += "}";
      return string_echo;
    }

    virtual std::string getType(){ return "rosgraph_msgs/Log"; }
    virtual std::string getMD5(){ return "2de9daf47e984009074d74dbdd492d49"; }

  };

}
#endif
