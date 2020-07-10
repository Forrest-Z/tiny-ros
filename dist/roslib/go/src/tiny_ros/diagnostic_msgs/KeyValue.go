package diagnostic_msgs

import (
    "encoding/json"
)


type KeyValue struct {
    Go_key string `json:"key"`
    Go_value string `json:"value"`
}

func NewKeyValue() (*KeyValue) {
    newKeyValue := new(KeyValue)
    newKeyValue.Go_key = ""
    newKeyValue.Go_value = ""
    return newKeyValue
}

func (self *KeyValue) Go_initialize() {
    self.Go_key = ""
    self.Go_value = ""
}

func (self *KeyValue) Go_serialize(buff []byte) (int) {
    offset := 0
    length_key := len(self.Go_key)
    buff[offset + 0] = byte((length_key >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_key >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_key >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_key >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_key)], self.Go_key)
    offset += length_key
    length_value := len(self.Go_value)
    buff[offset + 0] = byte((length_value >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_value >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_value >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_value >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_value)], self.Go_value)
    offset += length_value
    return offset
}

func (self *KeyValue) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_key := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_key |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_key |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_key |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_key = string(buff[offset:(offset+length_key)])
    offset += length_key
    length_value := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_value |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_value |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_value |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_value = string(buff[offset:(offset+length_value)])
    offset += length_value
    return offset
}

func (self *KeyValue) Go_serializedLength() (int) {
    length := 0
    length_key := len(self.Go_key)
    length += 4
    length += length_key
    length_value := len(self.Go_value)
    length += 4
    length += length_value
    return length
}

func (self *KeyValue) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *KeyValue) Go_getType() (string) { return "diagnostic_msgs/KeyValue" }
func (self *KeyValue) Go_getMD5() (string) { return "1baa904b80c685c77d1a42a872ca1d07" }
func (self *KeyValue) Go_getID() (uint32) { return 0 }
func (self *KeyValue) Go_setID(id uint32) { }

