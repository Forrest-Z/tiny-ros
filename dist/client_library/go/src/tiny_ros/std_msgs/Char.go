package std_msgs

import (
    "encoding/json"
)


type Char struct {
    Go_data byte `json:"data"`
}

func NewChar() (*Char) {
    newChar := new(Char)
    newChar.Go_data = 0
    return newChar
}

func (self *Char) Go_initialize() {
    self.Go_data = 0
}

func (self *Char) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *Char) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data = byte(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    return offset
}

func (self *Char) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *Char) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Char) Go_getType() (string) { return "std_msgs/Char" }
func (self *Char) Go_getMD5() (string) { return "18741a80dcc02fcae20538073fcb0872" }
func (self *Char) Go_getID() (uint32) { return 0 }
func (self *Char) Go_setID(id uint32) { }

