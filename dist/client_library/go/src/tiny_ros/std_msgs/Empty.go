package std_msgs

import (
    "encoding/json"
)


type Empty struct {
}

func NewEmpty() (*Empty) {
    newEmpty := new(Empty)
    return newEmpty
}

func (self *Empty) Go_initialize() {
}

func (self *Empty) Go_serialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *Empty) Go_deserialize(buff []byte) (int) {
    offset := 0
    return offset
}

func (self *Empty) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *Empty) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Empty) Go_getType() (string) { return "std_msgs/Empty" }
func (self *Empty) Go_getMD5() (string) { return "140bdcb7bbc50b4a936e90ff2350c8d3" }
func (self *Empty) Go_getID() (uint32) { return 0 }
func (self *Empty) Go_setID(id uint32) { }

