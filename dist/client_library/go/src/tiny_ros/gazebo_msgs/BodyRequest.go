package gazebo_msgs

import (
    "encoding/json"
)



type BodyRequestRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_body_name string `json:"body_name"`
}

func NewBodyRequestRequest() (*BodyRequestRequest) {
    newBodyRequestRequest := new(BodyRequestRequest)
    newBodyRequestRequest.Go_body_name = ""
    newBodyRequestRequest.__id__ = 0
    return newBodyRequestRequest
}

func (self *BodyRequestRequest) Go_initialize() {
    self.Go_body_name = ""
    self.__id__ = 0
}

func (self *BodyRequestRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_body_name := len(self.Go_body_name)
    buff[offset + 0] = byte((length_body_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_body_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_body_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_body_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_body_name)], self.Go_body_name)
    offset += length_body_name
    return offset
}

func (self *BodyRequestRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_body_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_body_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_body_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_body_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_body_name = string(buff[offset:(offset+length_body_name)])
    offset += length_body_name
    return offset
}

func (self *BodyRequestRequest) Go_serializedLength() (int) {
    length := 0
    length_body_name := len(self.Go_body_name)
    length += 4
    length += length_body_name
    return length
}

func (self *BodyRequestRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *BodyRequestRequest) Go_getType() (string) { return "gazebo_msgs/BodyRequest" }
func (self *BodyRequestRequest) Go_getMD5() (string) { return "d1c66fbceb0ee1b51e3b09ac030dedec" }
func (self *BodyRequestRequest) Go_getID() (uint32) { return self.__id__ }
func (self *BodyRequestRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type BodyRequestResponse struct {
    __id__ uint32 `json:"__id__"`
}

func NewBodyRequestResponse() (*BodyRequestResponse) {
    newBodyRequestResponse := new(BodyRequestResponse)
    newBodyRequestResponse.__id__ = 0
    return newBodyRequestResponse
}

func (self *BodyRequestResponse) Go_initialize() {
    self.__id__ = 0
}

func (self *BodyRequestResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *BodyRequestResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *BodyRequestResponse) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *BodyRequestResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *BodyRequestResponse) Go_getType() (string) { return "gazebo_msgs/BodyRequest" }
func (self *BodyRequestResponse) Go_getMD5() (string) { return "e0caf2eb9951542b962f95924c6eeb34" }
func (self *BodyRequestResponse) Go_getID() (uint32) { return self.__id__ }
func (self *BodyRequestResponse) Go_setID(id uint32) { self.__id__ = id }


