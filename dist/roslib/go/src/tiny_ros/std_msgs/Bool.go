package std_msgs

import (
    "encoding/json"
)


type Bool struct {
    Go_data bool `json:"data"`
}

func NewBool() (*Bool) {
    newBool := new(Bool)
    newBool.Go_data = false
    return newBool
}

func (self *Bool) Go_initialize() {
    self.Go_data = false
}

func (self *Bool) Go_serialize(buff []byte) (int) {
    offset := 0
    if self.Go_data {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *Bool) Go_deserialize(buff []byte) (int) {
    offset := 0
    if (buff[offset] & 0xFF) != 0 {
        self.Go_data = true
    } else {
        self.Go_data = false
    }
    offset += 1
    return offset
}

func (self *Bool) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *Bool) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Bool) Go_getType() (string) { return "std_msgs/Bool" }
func (self *Bool) Go_getMD5() (string) { return "cf6f397ea93618cea833f64b7eef203e" }
func (self *Bool) Go_getID() (uint32) { return 0 }
func (self *Bool) Go_setID(id uint32) { }

