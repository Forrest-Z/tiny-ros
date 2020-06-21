package std_msgs

import (
)

type Bool struct {
    Go_data bool `json:"data"`
}

func NewBool() (*Bool) {
    newBool := new(Bool)
    newBool.Go_data = false
    return newBool
}

func (self *Bool) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *Bool) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = bool((buff[offset + 0] & 0xFF) << (8 * 0))
    offset += 1
    return offset
}

func (self *Bool) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *Bool) Go_echo() (string) { return "" }
func (self *Bool) Go_getType() (string) { return "std_msgs/Bool" }
func (self *Bool) Go_getMD5() (string) { return "cf6f397ea93618cea833f64b7eef203e" }
func (self *Bool) Go_getID() (uint32) { return 0 }
func (self *Bool) Go_setID(id uint32) { }

