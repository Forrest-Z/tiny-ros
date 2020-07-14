package tf2_msgs

import (
    "encoding/json"
)


type LookupTransformAction struct {
    Go_action_goal *LookupTransformActionGoal `json:"action_goal"`
    Go_action_result *LookupTransformActionResult `json:"action_result"`
    Go_action_feedback *LookupTransformActionFeedback `json:"action_feedback"`
}

func NewLookupTransformAction() (*LookupTransformAction) {
    newLookupTransformAction := new(LookupTransformAction)
    newLookupTransformAction.Go_action_goal = NewLookupTransformActionGoal()
    newLookupTransformAction.Go_action_result = NewLookupTransformActionResult()
    newLookupTransformAction.Go_action_feedback = NewLookupTransformActionFeedback()
    return newLookupTransformAction
}

func (self *LookupTransformAction) Go_initialize() {
    self.Go_action_goal = NewLookupTransformActionGoal()
    self.Go_action_result = NewLookupTransformActionResult()
    self.Go_action_feedback = NewLookupTransformActionFeedback()
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

func (self *LookupTransformAction) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LookupTransformAction) Go_getType() (string) { return "tf2_msgs/LookupTransformAction" }
func (self *LookupTransformAction) Go_getMD5() (string) { return "49655e848adf6c64870a8eb763b94484" }
func (self *LookupTransformAction) Go_getID() (uint32) { return 0 }
func (self *LookupTransformAction) Go_setID(id uint32) { }

