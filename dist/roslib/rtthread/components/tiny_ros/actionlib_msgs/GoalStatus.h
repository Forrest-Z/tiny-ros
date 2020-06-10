#ifndef _TINYROS_actionlib_msgs_GoalStatus_h
#define _TINYROS_actionlib_msgs_GoalStatus_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "tiny_ros/ros/msg.h"
#include "tiny_ros/actionlib_msgs/GoalID.h"

namespace actionlib_msgs
{

  class GoalStatus : public tinyros::Msg
  {
    public:
      typedef actionlib_msgs::GoalID _goal_id_type;
      _goal_id_type goal_id;
      typedef uint8_t _status_type;
      _status_type status;
      typedef tinyros::string _text_type;
      _text_type text;
      enum { PENDING =  0    };
      enum { ACTIVE =  1    };
      enum { PREEMPTED =  2    };
      enum { SUCCEEDED =  3    };
      enum { ABORTED =  4    };
      enum { REJECTED =  5    };
      enum { PREEMPTING =  6    };
      enum { RECALLING =  7    };
      enum { RECALLED =  8    };
      enum { LOST =  9    };

    GoalStatus():
      goal_id(),
      status(0),
      text("")
    {
    }

    virtual int serialize(unsigned char *outbuffer) const
    {
      int offset = 0;
      offset += this->goal_id.serialize(outbuffer + offset);
      *(outbuffer + offset + 0) = (this->status >> (8 * 0)) & 0xFF;
      offset += sizeof(this->status);
      uint32_t length_text = this->text.size();
      varToArr(outbuffer + offset, length_text);
      offset += 4;
      memcpy(outbuffer + offset, this->text.c_str(), length_text);
      offset += length_text;
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer)
    {
      int offset = 0;
      offset += this->goal_id.deserialize(inbuffer + offset);
      this->status =  ((uint8_t) (*(inbuffer + offset)));
      offset += sizeof(this->status);
      uint32_t length_text;
      arrToVar(length_text, (inbuffer + offset));
      offset += 4;
      for(unsigned int k= offset; k< offset+length_text; ++k){
          inbuffer[k-1]=inbuffer[k];
      }
      inbuffer[offset+length_text-1]=0;
      this->text = (char *)(inbuffer + offset-1);
      offset += length_text;
      return offset;
    }

    virtual int serializedLength() const
    {
      int length = 0;
      length += this->goal_id.serializedLength();
      length += sizeof(this->status);
      uint32_t length_text = this->text.size();
      length += 4;
      length += length_text;
      return length;
    }

    virtual tinyros::string getType(){ return "actionlib_msgs/GoalStatus"; }
    virtual tinyros::string getMD5(){ return "086be35ea957e692de83fc3477e4ef0b"; }

  };

}
#endif
