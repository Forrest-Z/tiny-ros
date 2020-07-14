package tf2_msgs

import (
    "encoding/json"
)


type LookupTransformFeedback struct {
}

func NewLookupTransformFeedback() (*LookupTransformFeedback) {
    newLookupTransformFeedback := new(LookupTransformFeedback)
    return newLookupTransformFeedback
}

func (self *LookupTransformFeedback) Go_initialize() {
}

func (self *LookupTransformFeedback) Go_serialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *LookupTransformFeedback) Go_deserialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *LookupTransformFeedback) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *LookupTransformFeedback) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LookupTransformFeedback) Go_getType() (string) { return "tf2_msgs/LookupTransformFeedback" }
func (self *LookupTransformFeedback) Go_getMD5() (string) { return "e6217f8a8e77aa218a8d6f594d08ba08" }
func (self *LookupTransformFeedback) Go_getID() (uint32) { return 0 }
func (self *LookupTransformFeedback) Go_setID(id uint32) { }

