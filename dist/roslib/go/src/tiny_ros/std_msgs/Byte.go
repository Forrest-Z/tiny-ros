package std_msgs

import (
    "encoding/json"
)


type Byte struct {
    Go_data byte `json:"data"`
}

func NewByte() (*Byte) {
    newByte := new(Byte)
    newByte.Go_data = 0
    return newByte
}

func (self *Byte) Go_initialize() {
    self.Go_data = 0
}

func (self *Byte) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *Byte) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = byte(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    return offset
}

func (self *Byte) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *Byte) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Byte) Go_getType() (string) { return "std_msgs/Byte" }
func (self *Byte) Go_getMD5() (string) { return "8c5affe485c5af9bd37408a1a905a49c" }
func (self *Byte) Go_getID() (uint32) { return 0 }
func (self *Byte) Go_setID(id uint32) { }

