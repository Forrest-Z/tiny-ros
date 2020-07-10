package tinyros_hello

import (
    "encoding/json"
)



type TestRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_input string `json:"input"`
}

func NewTestRequest() (*TestRequest) {
    newTestRequest := new(TestRequest)
    newTestRequest.Go_input = ""
    newTestRequest.__id__ = 0
    return newTestRequest
}

func (self *TestRequest) Go_initialize() {
    self.Go_input = ""
    self.__id__ = 0
}

func (self *TestRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_input := len(self.Go_input)
    buff[offset + 0] = byte((length_input >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_input >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_input >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_input >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_input)], self.Go_input)
    offset += length_input
    return offset
}

func (self *TestRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_input := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_input |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_input |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_input |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_input = string(buff[offset:(offset+length_input)])
    offset += length_input
    return offset
}

func (self *TestRequest) Go_serializedLength() (int) {
    length := 0
    length_input := len(self.Go_input)
    length += 4
    length += length_input
    return length
}

func (self *TestRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TestRequest) Go_getType() (string) { return "tinyros_hello/Test" }
func (self *TestRequest) Go_getMD5() (string) { return "26ee7a44335f1f7b55a5a7490460807d" }
func (self *TestRequest) Go_getID() (uint32) { return self.__id__ }
func (self *TestRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type TestResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_output string `json:"output"`
}

func NewTestResponse() (*TestResponse) {
    newTestResponse := new(TestResponse)
    newTestResponse.Go_output = ""
    newTestResponse.__id__ = 0
    return newTestResponse
}

func (self *TestResponse) Go_initialize() {
    self.Go_output = ""
    self.__id__ = 0
}

func (self *TestResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_output := len(self.Go_output)
    buff[offset + 0] = byte((length_output >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_output >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_output >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_output >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_output)], self.Go_output)
    offset += length_output
    return offset
}

func (self *TestResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_output := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_output |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_output |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_output |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_output = string(buff[offset:(offset+length_output)])
    offset += length_output
    return offset
}

func (self *TestResponse) Go_serializedLength() (int) {
    length := 0
    length_output := len(self.Go_output)
    length += 4
    length += length_output
    return length
}

func (self *TestResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TestResponse) Go_getType() (string) { return "tinyros_hello/Test" }
func (self *TestResponse) Go_getMD5() (string) { return "e2f6296e6ea9df7406f4fac9fb52d44b" }
func (self *TestResponse) Go_getID() (uint32) { return self.__id__ }
func (self *TestResponse) Go_setID(id uint32) { self.__id__ = id }


