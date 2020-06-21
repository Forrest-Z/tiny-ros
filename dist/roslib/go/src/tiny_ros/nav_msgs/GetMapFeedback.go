package nav_msgs

import (
)

type GetMapFeedback struct {
}

func NewGetMapFeedback() (*GetMapFeedback) {
    newGetMapFeedback := new(GetMapFeedback)
    return newGetMapFeedback
}

func (self *GetMapFeedback) Go_serialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *GetMapFeedback) Go_deserialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *GetMapFeedback) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetMapFeedback) Go_echo() (string) { return "" }
func (self *GetMapFeedback) Go_getType() (string) { return "nav_msgs/GetMapFeedback" }
func (self *GetMapFeedback) Go_getMD5() (string) { return "f561626803919fb2f269eb497bfdfea4" }
func (self *GetMapFeedback) Go_getID() (uint32) { return 0 }
func (self *GetMapFeedback) Go_setID(id uint32) { }

