package std_srvs

import (
    "encoding/json"
)



type EmptyRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewEmptyRequest() (*EmptyRequest) {
    newEmptyRequest := new(EmptyRequest)
    newEmptyRequest.__id__ = 0
    return newEmptyRequest
}

func (self *EmptyRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *EmptyRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *EmptyRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *EmptyRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *EmptyRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *EmptyRequest) Go_getType() (string) { return "std_srvs/Empty" }
func (self *EmptyRequest) Go_getMD5() (string) { return "6a9da448a5a2256d30e815f50a75bc57" }
func (self *EmptyRequest) Go_getID() (uint32) { return self.__id__ }
func (self *EmptyRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type EmptyResponse struct {
    __id__ uint32 `json:"__id__"`
}

func NewEmptyResponse() (*EmptyResponse) {
    newEmptyResponse := new(EmptyResponse)
    newEmptyResponse.__id__ = 0
    return newEmptyResponse
}

func (self *EmptyResponse) Go_initialize() {
    self.__id__ = 0
}

func (self *EmptyResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *EmptyResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *EmptyResponse) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *EmptyResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *EmptyResponse) Go_getType() (string) { return "std_srvs/Empty" }
func (self *EmptyResponse) Go_getMD5() (string) { return "1b9fedd0c70a6b7846f790471f388d15" }
func (self *EmptyResponse) Go_getID() (uint32) { return self.__id__ }
func (self *EmptyResponse) Go_setID(id uint32) { self.__id__ = id }


