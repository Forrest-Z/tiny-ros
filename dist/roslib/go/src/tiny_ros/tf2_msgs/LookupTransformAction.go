package tf2_msgs

import (
    "tf2_msgs/LookupTransformActionGoal"
    "tf2_msgs/LookupTransformActionResult"
    "tf2_msgs/LookupTransformActionFeedback"
)

type LookupTransformAction struct {
    Go_action_goal tf2_msgs.LookupTransformActionGoal `json:"action_goal"`
    Go_action_result tf2_msgs.LookupTransformActionResult `json:"action_result"`
    Go_action_feedback tf2_msgs.LookupTransformActionFeedback `json:"action_feedback"`
}

func NewLookupTransformAction() (*LookupTransformAction) {
    newLookupTransformAction := new(LookupTransformAction)
    newLookupTransformAction.Go_action_goal = tf2_msgs.NewLookupTransformActionGoal()
    newLookupTransformAction.Go_action_result = tf2_msgs.NewLookupTransformActionResult()
    newLookupTransformAction.Go_action_feedback = tf2_msgs.NewLookupTransformActionFeedback()
    return newLookupTransformAction
}

func (self *LookupTransformAction) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_action_goal.Go_serialize(buff[offset:])
    offset += self.Go_action_result.Go_serialize(buff[offset:])
    offset += self.Go_action_feedback.Go_serialize(buff[offset:])
    return offset
}

func (self *LookupTransformAction) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_action_goal.Go_deserialize(buff[offset:])
    offset += self.Go_action_result.Go_deserialize(buff[offset:])
    offset += self.Go_action_feedback.Go_deserialize(buff[offset:])
    return offset
}

func (self *LookupTransformAction) Go_serializedLength() (int) {
    length := 0
    length += self.Go_action_goal.Go_serializedLength()
    length += self.Go_action_result.Go_serializedLength()
    length += self.Go_action_feedback.Go_serializedLength()
    return length
}

func (self *LookupTransformAction) Go_echo() (string) { return "" }
func (self *LookupTransformAction) Go_getType() (string) { return "tf2_msgs/LookupTransformAction" }
func (self *LookupTransformAction) Go_getMD5() (string) { return "49655e848adf6c64870a8eb763b94484" }
func (self *LookupTransformAction) Go_getID() (uint32) { return 0 }
func (self *LookupTransformAction) Go_setID(id uint32) { }

