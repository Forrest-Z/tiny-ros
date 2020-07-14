package nav_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/actionlib_msgs"
)


type GetMapActionFeedback struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_status *actionlib_msgs.GoalStatus `json:"status"`
    Go_feedback *GetMapFeedback `json:"feedback"`
}

func NewGetMapActionFeedback() (*GetMapActionFeedback) {
    newGetMapActionFeedback := new(GetMapActionFeedback)
    newGetMapActionFeedback.Go_header = std_msgs.NewHeader()
    newGetMapActionFeedback.Go_status = actionlib_msgs.NewGoalStatus()
    newGetMapActionFeedback.Go_feedback = NewGetMapFeedback()
    return newGetMapActionFeedback
}

func (self *GetMapActionFeedback) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_status = actionlib_msgs.NewGoalStatus()
    self.Go_feedback = NewGetMapFeedback()
}

func (self *GetMapActionFeedback) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_status.Go_serialize(buff[offset:])
    offset += self.Go_feedback.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapActionFeedback) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_status.Go_deserialize(buff[offset:])
    offset += self.Go_feedback.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapActionFeedback) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_status.Go_serializedLength()
    length += self.Go_feedback.Go_serializedLength()
    return length
}

func (self *GetMapActionFeedback) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapActionFeedback) Go_getType() (string) { return "nav_msgs/GetMapActionFeedback" }
func (self *GetMapActionFeedback) Go_getMD5() (string) { return "9ebb88ff2cf2120160bf2197071a69b6" }
func (self *GetMapActionFeedback) Go_getID() (uint32) { return 0 }
func (self *GetMapActionFeedback) Go_setID(id uint32) { }

