package map_msgs

import (
    "encoding/json"
    "tiny_ros/sensor_msgs"
)



type GetPointMapRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewGetPointMapRequest() (*GetPointMapRequest) {
    newGetPointMapRequest := new(GetPointMapRequest)
    newGetPointMapRequest.__id__ = 0
    return newGetPointMapRequest
}

func (self *GetPointMapRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *GetPointMapRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *GetPointMapRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *GetPointMapRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetPointMapRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPointMapRequest) Go_getType() (string) { return "map_msgs/GetPointMap" }
func (self *GetPointMapRequest) Go_getMD5() (string) { return "418d4ee8c9d758b7ef1aab3e068c7568" }
func (self *GetPointMapRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetPointMapRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetPointMapResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_map *sensor_msgs.PointCloud2 `json:"map"`
}

func NewGetPointMapResponse() (*GetPointMapResponse) {
    newGetPointMapResponse := new(GetPointMapResponse)
    newGetPointMapResponse.Go_map = sensor_msgs.NewPointCloud2()
    newGetPointMapResponse.__id__ = 0
    return newGetPointMapResponse
}

func (self *GetPointMapResponse) Go_initialize() {
    self.Go_map = sensor_msgs.NewPointCloud2()
    self.__id__ = 0
}

func (self *GetPointMapResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_map.Go_serialize(buff[offset:])
    return offset
}

func (self *GetPointMapResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_map.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetPointMapResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_map.Go_serializedLength()
    return length
}

func (self *GetPointMapResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPointMapResponse) Go_getType() (string) { return "map_msgs/GetPointMap" }
func (self *GetPointMapResponse) Go_getMD5() (string) { return "abc97e6c5b25f536248da280bdf271d7" }
func (self *GetPointMapResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetPointMapResponse) Go_setID(id uint32) { self.__id__ = id }


