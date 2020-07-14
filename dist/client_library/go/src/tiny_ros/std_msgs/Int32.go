package std_msgs

import (
    "encoding/json"
)


type Int32 struct {
    Go_data int32 `json:"data"`
}

func NewInt32() (*Int32) {
    newInt32 := new(Int32)
    newInt32.Go_data = 0
    return newInt32
}

func (self *Int32) Go_initialize() {
    self.Go_data = 0
}

func (self *Int32) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *Int32) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *Int32) Go_serializedLength() (int) {
    length := 0
    length += 4
    return length
}

func (self *Int32) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Int32) Go_getType() (string) { return "std_msgs/Int32" }
func (self *Int32) Go_getMD5() (string) { return "8e99256d77b8f00635db2852db07fc9f" }
func (self *Int32) Go_getID() (uint32) { return 0 }
func (self *Int32) Go_setID(id uint32) { }

