package std_msgs

import (
    "encoding/json"
)


type UInt32 struct {
    Go_data uint32 `json:"data"`
}

func NewUInt32() (*UInt32) {
    newUInt32 := new(UInt32)
    newUInt32.Go_data = 0
    return newUInt32
}

func (self *UInt32) Go_initialize() {
    self.Go_data = 0
}

func (self *UInt32) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *UInt32) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *UInt32) Go_serializedLength() (int) {
    length := 0
    length += 4
    return length
}

func (self *UInt32) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt32) Go_getType() (string) { return "std_msgs/UInt32" }
func (self *UInt32) Go_getMD5() (string) { return "d4e8dc9c9e9d5076e594a3e342c2d4e3" }
func (self *UInt32) Go_getID() (uint32) { return 0 }
func (self *UInt32) Go_setID(id uint32) { }

