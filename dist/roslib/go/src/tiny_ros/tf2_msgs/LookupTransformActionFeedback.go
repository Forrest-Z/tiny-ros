package tf2_msgs

import (
    "tiny_ros/std_msgs"
    "actionlib_msgs/GoalStatus"
    "tf2_msgs/LookupTransformFeedback"
)

type LookupTransformActionFeedback struct {
    Go_header std_msgs.Header `json:"header"`
    Go_status actionlib_msgs.GoalStatus `json:"status"`
    Go_feedback tf2_msgs.LookupTransformFeedback `json:"feedback"`
}

func NewLookupTransformActionFeedback() (*LookupTransformActionFeedback) {
    newLookupTransformActionFeedback := new(LookupTransformActionFeedback)
    newLookupTransformActionFeedback.Go_header = std_msgs.NewHeader()
    newLookupTransformActionFeedback.Go_status = actionlib_msgs.NewGoalStatus()
    newLookupTransformActionFeedback.Go_feedback = tf2_msgs.NewLookupTransformFeedback()
    return newLookupTransformActionFeedback
}

func (self *LookupTransformActionFeedback) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_status.Go_serialize(buff[offset:])
    offset += self.Go_feedback.Go_serialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionFeedback) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_status.Go_deserialize(buff[offset:])
    offset += self.Go_feedback.Go_deserialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionFeedback) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_status.Go_serializedLength()
    length += self.Go_feedback.Go_serializedLength()
    return length
}

func (self *LookupTransformActionFeedback) Go_echo() (string) { return "" }
func (self *LookupTransformActionFeedback) Go_getType() (string) { return "tf2_msgs/LookupTransformActionFeedback" }
func (self *LookupTransformActionFeedback) Go_getMD5() (string) { return "f432c3fa33a3ac9586ace248b606de3a" }
func (self *LookupTransformActionFeedback) Go_getID() (uint32) { return 0 }
func (self *LookupTransformActionFeedback) Go_setID(id uint32) { }

