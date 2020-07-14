package tf2_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/actionlib_msgs"
)


type LookupTransformActionGoal struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_goal_id *actionlib_msgs.GoalID `json:"goal_id"`
    Go_goal *LookupTransformGoal `json:"goal"`
}

func NewLookupTransformActionGoal() (*LookupTransformActionGoal) {
    newLookupTransformActionGoal := new(LookupTransformActionGoal)
    newLookupTransformActionGoal.Go_header = std_msgs.NewHeader()
    newLookupTransformActionGoal.Go_goal_id = actionlib_msgs.NewGoalID()
    newLookupTransformActionGoal.Go_goal = NewLookupTransformGoal()
    return newLookupTransformActionGoal
}

func (self *LookupTransformActionGoal) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_goal_id = actionlib_msgs.NewGoalID()
    self.Go_goal = NewLookupTransformGoal()
}

func (self *LookupTransformActionGoal) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_goal_id.Go_serialize(buff[offset:])
    offset += self.Go_goal.Go_serialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionGoal) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_goal_id.Go_deserialize(buff[offset:])
    offset += self.Go_goal.Go_deserialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionGoal) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_goal_id.Go_serializedLength()
    length += self.Go_goal.Go_serializedLength()
    return length
}

func (self *LookupTransformActionGoal) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LookupTransformActionGoal) Go_getType() (string) { return "tf2_msgs/LookupTransformActionGoal" }
func (self *LookupTransformActionGoal) Go_getMD5() (string) { return "b8a7d4ffa64f063b4df7b1dd3fc2bf79" }
func (self *LookupTransformActionGoal) Go_getID() (uint32) { return 0 }
func (self *LookupTransformActionGoal) Go_setID(id uint32) { }

