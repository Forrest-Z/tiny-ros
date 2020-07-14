package std_msgs

import (
    "encoding/json"
)


type UInt64 struct {
    Go_data uint64 `json:"data"`
}

func NewUInt64() (*UInt64) {
    newUInt64 := new(UInt64)
    newUInt64.Go_data = 0
    return newUInt64
}

func (self *UInt64) Go_initialize() {
    self.Go_data = 0
}

func (self *UInt64) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data >> (8 * 3)) & 0xFF)
    buff[offset + 4] = byte((self.Go_data >> (8 * 4)) & 0xFF)
    buff[offset + 5] = byte((self.Go_data >> (8 * 5)) & 0xFF)
    buff[offset + 6] = byte((self.Go_data >> (8 * 6)) & 0xFF)
    buff[offset + 7] = byte((self.Go_data >> (8 * 7)) & 0xFF)
    offset += 8
    return offset
}

func (self *UInt64) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = uint64(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data |= uint64(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data |= uint64(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data |= uint64(buff[offset + 3] & 0xFF) << (8 * 3)
    self.Go_data |= uint64(buff[offset + 4] & 0xFF) << (8 * 4)
    self.Go_data |= uint64(buff[offset + 5] & 0xFF) << (8 * 5)
    self.Go_data |= uint64(buff[offset + 6] & 0xFF) << (8 * 6)
    self.Go_data |= uint64(buff[offset + 7] & 0xFF) << (8 * 7)
    offset += 8
    return offset
}

func (self *UInt64) Go_serializedLength() (int) {
    length := 0
    length += 8
    return length
}

func (self *UInt64) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt64) Go_getType() (string) { return "std_msgs/UInt64" }
func (self *UInt64) Go_getMD5() (string) { return "e815672a5006369d73f91f25ee5609ac" }
func (self *UInt64) Go_getID() (uint32) { return 0 }
func (self *UInt64) Go_setID(id uint32) { }

