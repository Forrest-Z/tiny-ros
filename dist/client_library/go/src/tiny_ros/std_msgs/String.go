package std_msgs

import (
    "encoding/json"
)


type String struct {
    Go_data string `json:"data"`
}

func NewString() (*String) {
    newString := new(String)
    newString.Go_data = ""
    return newString
}

func (self *String) Go_initialize() {
    self.Go_data = ""
}

func (self *String) Go_serialize(buff []byte) (int) {
    offset := 0
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_data)], self.Go_data)
    offset += length_data
    return offset
}

func (self *String) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = string(buff[offset:(offset+length_data)])
    offset += length_data
    return offset
}

func (self *String) Go_serializedLength() (int) {
    length := 0
    length_data := len(self.Go_data)
    length += 4
    length += length_data
    return length
}

func (self *String) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *String) Go_getType() (string) { return "std_msgs/String" }
func (self *String) Go_getMD5() (string) { return "44b822292b4a9ed05e241aa225458f97" }
func (self *String) Go_getID() (uint32) { return 0 }
func (self *String) Go_setID(id uint32) { }

