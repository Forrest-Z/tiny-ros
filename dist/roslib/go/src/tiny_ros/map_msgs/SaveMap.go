package map_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)



type SaveMapRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_filename *std_msgs.String `json:"filename"`
}

func NewSaveMapRequest() (*SaveMapRequest) {
    newSaveMapRequest := new(SaveMapRequest)
    newSaveMapRequest.Go_filename = std_msgs.NewString()
    newSaveMapRequest.__id__ = 0
    return newSaveMapRequest
}

func (self *SaveMapRequest) Go_initialize() {
    self.Go_filename = std_msgs.NewString()
    self.__id__ = 0
}

func (self *SaveMapRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_filename.Go_serialize(buff[offset:])
    return offset
}

func (self *SaveMapRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_filename.Go_deserialize(buff[offset:])
    return offset
}

func (self *SaveMapRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_filename.Go_serializedLength()
    return length
}

func (self *SaveMapRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SaveMapRequest) Go_getType() (string) { return "map_msgs/SaveMap" }
func (self *SaveMapRequest) Go_getMD5() (string) { return "6643d8ede81a23998690e6a3ff657316" }
func (self *SaveMapRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SaveMapRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SaveMapResponse struct {
    __id__ uint32 `json:"__id__"`
}

func NewSaveMapResponse() (*SaveMapResponse) {
    newSaveMapResponse := new(SaveMapResponse)
    newSaveMapResponse.__id__ = 0
    return newSaveMapResponse
}

func (self *SaveMapResponse) Go_initialize() {
    self.__id__ = 0
}

func (self *SaveMapResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *SaveMapResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *SaveMapResponse) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *SaveMapResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SaveMapResponse) Go_getType() (string) { return "map_msgs/SaveMap" }
func (self *SaveMapResponse) Go_getMD5() (string) { return "9cd07446fa1bd59b4758dadf19f196e9" }
func (self *SaveMapResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SaveMapResponse) Go_setID(id uint32) { self.__id__ = id }


