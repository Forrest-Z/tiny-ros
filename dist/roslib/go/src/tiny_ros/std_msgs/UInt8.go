package std_msgs

import (
    "encoding/json"
)


type UInt8 struct {
    Go_data uint8 `json:"data"`
}

func NewUInt8() (*UInt8) {
    newUInt8 := new(UInt8)
    newUInt8.Go_data = 0
    return newUInt8
}

func (self *UInt8) Go_initialize() {
    self.Go_data = 0
}

func (self *UInt8) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *UInt8) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    return offset
}

func (self *UInt8) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *UInt8) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt8) Go_getType() (string) { return "std_msgs/UInt8" }
func (self *UInt8) Go_getMD5() (string) { return "6f90555707d539e508484b884b2acc65" }
func (self *UInt8) Go_getID() (uint32) { return 0 }
func (self *UInt8) Go_setID(id uint32) { }

