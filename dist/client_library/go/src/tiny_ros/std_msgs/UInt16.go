package std_msgs

import (
    "encoding/json"
)


type UInt16 struct {
    Go_data uint16 `json:"data"`
}

func NewUInt16() (*UInt16) {
    newUInt16 := new(UInt16)
    newUInt16.Go_data = 0
    return newUInt16
}

func (self *UInt16) Go_initialize() {
    self.Go_data = 0
}

func (self *UInt16) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data >> (8 * 1)) & 0xFF)
    offset += 2
    return offset
}

func (self *UInt16) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = uint16(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data |= uint16(buff[offset + 1] & 0xFF) << (8 * 1)
    offset += 2
    return offset
}

func (self *UInt16) Go_serializedLength() (int) {
    length := 0
    length += 2
    return length
}

func (self *UInt16) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt16) Go_getType() (string) { return "std_msgs/UInt16" }
func (self *UInt16) Go_getMD5() (string) { return "a4caad902eedb84b728d8369c50ffc39" }
func (self *UInt16) Go_getID() (uint32) { return 0 }
func (self *UInt16) Go_setID(id uint32) { }

