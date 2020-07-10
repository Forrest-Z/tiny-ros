package nav_msgs

import (
    "encoding/json"
)


type GetMapGoal struct {
}

func NewGetMapGoal() (*GetMapGoal) {
    newGetMapGoal := new(GetMapGoal)
    return newGetMapGoal
}

func (self *GetMapGoal) Go_initialize() {
}

func (self *GetMapGoal) Go_serialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *GetMapGoal) Go_deserialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *GetMapGoal) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetMapGoal) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapGoal) Go_getType() (string) { return "nav_msgs/GetMapGoal" }
func (self *GetMapGoal) Go_getMD5() (string) { return "b39e6b705afaad0184bd2c87f4bd870f" }
func (self *GetMapGoal) Go_getID() (uint32) { return 0 }
func (self *GetMapGoal) Go_setID(id uint32) { }

