package std_msgs

import (
    "encoding/json"
)


type Int8 struct {
    Go_data int8 `json:"data"`
}

func NewInt8() (*Int8) {
    newInt8 := new(Int8)
    newInt8.Go_data = 0
    return newInt8
}

func (self *Int8) Go_initialize() {
    self.Go_data = 0
}

func (self *Int8) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte(uint8(self.Go_data >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *Int8) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = int8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    return offset
}

func (self *Int8) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *Int8) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Int8) Go_getType() (string) { return "std_msgs/Int8" }
func (self *Int8) Go_getMD5() (string) { return "36c967b84bec6cd7c260bffc7f4dfbe0" }
func (self *Int8) Go_getID() (uint32) { return 0 }
func (self *Int8) Go_setID(id uint32) { }

