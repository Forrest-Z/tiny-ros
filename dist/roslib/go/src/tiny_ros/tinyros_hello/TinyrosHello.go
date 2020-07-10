package tinyros_hello

import (
    "encoding/json"
)


type TinyrosHello struct {
    Go_hello string `json:"hello"`
}

func NewTinyrosHello() (*TinyrosHello) {
    newTinyrosHello := new(TinyrosHello)
    newTinyrosHello.Go_hello = ""
    return newTinyrosHello
}

func (self *TinyrosHello) Go_initialize() {
    self.Go_hello = ""
}

func (self *TinyrosHello) Go_serialize(buff []byte) (int) {
    offset := 0
    length_hello := len(self.Go_hello)
    buff[offset + 0] = byte((length_hello >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_hello >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_hello >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_hello >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_hello)], self.Go_hello)
    offset += length_hello
    return offset
}

func (self *TinyrosHello) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_hello := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_hello |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_hello |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_hello |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_hello = string(buff[offset:(offset+length_hello)])
    offset += length_hello
    return offset
}

func (self *TinyrosHello) Go_serializedLength() (int) {
    length := 0
    length_hello := len(self.Go_hello)
    length += 4
    length += length_hello
    return length
}

func (self *TinyrosHello) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TinyrosHello) Go_getType() (string) { return "tinyros_hello/TinyrosHello" }
func (self *TinyrosHello) Go_getMD5() (string) { return "0c68e66a217802ad0c9f648b7a69d090" }
func (self *TinyrosHello) Go_getID() (uint32) { return 0 }
func (self *TinyrosHello) Go_setID(id uint32) { }

