package nav_msgs

import (
    "encoding/json"
)



type GetMapRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewGetMapRequest() (*GetMapRequest) {
    newGetMapRequest := new(GetMapRequest)
    newGetMapRequest.__id__ = 0
    return newGetMapRequest
}

func (self *GetMapRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *GetMapRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *GetMapRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *GetMapRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetMapRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapRequest) Go_getType() (string) { return "nav_msgs/GetMap" }
func (self *GetMapRequest) Go_getMD5() (string) { return "8ea512c109a0b3a9eca8de407dd02d2a" }
func (self *GetMapRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetMapRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetMapResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_map *OccupancyGrid `json:"map"`
}

func NewGetMapResponse() (*GetMapResponse) {
    newGetMapResponse := new(GetMapResponse)
    newGetMapResponse.Go_map = NewOccupancyGrid()
    newGetMapResponse.__id__ = 0
    return newGetMapResponse
}

func (self *GetMapResponse) Go_initialize() {
    self.Go_map = NewOccupancyGrid()
    self.__id__ = 0
}

func (self *GetMapResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_map.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_map.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_map.Go_serializedLength()
    return length
}

func (self *GetMapResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapResponse) Go_getType() (string) { return "nav_msgs/GetMap" }
func (self *GetMapResponse) Go_getMD5() (string) { return "5bf895a6cdaff312c69ca1cef20fed64" }
func (self *GetMapResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetMapResponse) Go_setID(id uint32) { self.__id__ = id }


