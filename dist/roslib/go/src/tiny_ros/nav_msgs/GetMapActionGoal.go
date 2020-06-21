package nav_msgs

import (
    "tiny_ros/std_msgs"
    "actionlib_msgs/GoalID"
    "nav_msgs/GetMapGoal"
)

type GetMapActionGoal struct {
    Go_header std_msgs.Header `json:"header"`
    Go_goal_id actionlib_msgs.GoalID `json:"goal_id"`
    Go_goal nav_msgs.GetMapGoal `json:"goal"`
}

func NewGetMapActionGoal() (*GetMapActionGoal) {
    newGetMapActionGoal := new(GetMapActionGoal)
    newGetMapActionGoal.Go_header = std_msgs.NewHeader()
    newGetMapActionGoal.Go_goal_id = actionlib_msgs.NewGoalID()
    newGetMapActionGoal.Go_goal = nav_msgs.NewGetMapGoal()
    return newGetMapActionGoal
}

func (self *GetMapActionGoal) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_goal_id.Go_serialize(buff[offset:])
    offset += self.Go_goal.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapActionGoal) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_goal_id.Go_deserialize(buff[offset:])
    offset += self.Go_goal.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapActionGoal) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_goal_id.Go_serializedLength()
    length += self.Go_goal.Go_serializedLength()
    return length
}

func (self *GetMapActionGoal) Go_echo() (string) { return "" }
func (self *GetMapActionGoal) Go_getType() (string) { return "nav_msgs/GetMapActionGoal" }
func (self *GetMapActionGoal) Go_getMD5() (string) { return "8aea83336b4ee626241742bb14b14d90" }
func (self *GetMapActionGoal) Go_getID() (uint32) { return 0 }
func (self *GetMapActionGoal) Go_setID(id uint32) { }

