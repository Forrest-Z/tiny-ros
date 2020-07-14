package actionlib_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type GoalStatusArray struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_status_list []*GoalStatus `json:"status_list"`
}

func NewGoalStatusArray() (*GoalStatusArray) {
    newGoalStatusArray := new(GoalStatusArray)
    newGoalStatusArray.Go_header = std_msgs.NewHeader()
    newGoalStatusArray.Go_status_list = []*GoalStatus{}
    return newGoalStatusArray
}

func (self *GoalStatusArray) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_status_list = []*GoalStatus{}
}

func (self *GoalStatusArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_status_list := len(self.Go_status_list)
    buff[offset + 0] = byte((length_status_list >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_status_list >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_status_list >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_status_list >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_status_list; i++ {
        offset += self.Go_status_list[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *GoalStatusArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_status_list := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_status_list |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_status_list |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_status_list |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_status_list = make([]*GoalStatus, length_status_list)
    for i := 0; i < length_status_list; i++ {
        self.Go_status_list[i] = NewGoalStatus()
    }
    for i := 0; i < length_status_list; i++ {
        offset += self.Go_status_list[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *GoalStatusArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_status_list := len(self.Go_status_list)
    for i := 0; i < length_status_list; i++ {
        length += self.Go_status_list[i].Go_serializedLength()
    }
    return length
}

func (self *GoalStatusArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GoalStatusArray) Go_getType() (string) { return "actionlib_msgs/GoalStatusArray" }
func (self *GoalStatusArray) Go_getMD5() (string) { return "53f6501f7c14f5f3963638de4bbe3a71" }
func (self *GoalStatusArray) Go_getID() (uint32) { return 0 }
func (self *GoalStatusArray) Go_setID(id uint32) { }

