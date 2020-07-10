package std_msgs

import (
    "encoding/json"
)


type Int16 struct {
    Go_data int16 `json:"data"`
}

func NewInt16() (*Int16) {
    newInt16 := new(Int16)
    newInt16.Go_data = 0
    return newInt16
}

func (self *Int16) Go_initialize() {
    self.Go_data = 0
}

func (self *Int16) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data >> (8 * 1)) & 0xFF)
    offset += 2
    return offset
}

func (self *Int16) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = int16(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data |= int16(buff[offset + 1] & 0xFF) << (8 * 1)
    offset += 2
    return offset
}

func (self *Int16) Go_serializedLength() (int) {
    length := 0
    length += 2
    return length
}

func (self *Int16) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Int16) Go_getType() (string) { return "std_msgs/Int16" }
func (self *Int16) Go_getMD5() (string) { return "156735359f7b6d347f113a1090134af3" }
func (self *Int16) Go_getID() (uint32) { return 0 }
func (self *Int16) Go_setID(id uint32) { }

