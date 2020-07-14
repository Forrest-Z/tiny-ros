package map_msgs

import (
    "encoding/json"
)



type SetMapProjectionsRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewSetMapProjectionsRequest() (*SetMapProjectionsRequest) {
    newSetMapProjectionsRequest := new(SetMapProjectionsRequest)
    newSetMapProjectionsRequest.__id__ = 0
    return newSetMapProjectionsRequest
}

func (self *SetMapProjectionsRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *SetMapProjectionsRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *SetMapProjectionsRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *SetMapProjectionsRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *SetMapProjectionsRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetMapProjectionsRequest) Go_getType() (string) { return "map_msgs/SetMapProjections" }
func (self *SetMapProjectionsRequest) Go_getMD5() (string) { return "26dbba584adf9695d68b8667830be463" }
func (self *SetMapProjectionsRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetMapProjectionsRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetMapProjectionsResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_projected_maps_info []*ProjectedMapInfo `json:"projected_maps_info"`
}

func NewSetMapProjectionsResponse() (*SetMapProjectionsResponse) {
    newSetMapProjectionsResponse := new(SetMapProjectionsResponse)
    newSetMapProjectionsResponse.Go_projected_maps_info = []*ProjectedMapInfo{}
    newSetMapProjectionsResponse.__id__ = 0
    return newSetMapProjectionsResponse
}

func (self *SetMapProjectionsResponse) Go_initialize() {
    self.Go_projected_maps_info = []*ProjectedMapInfo{}
    self.__id__ = 0
}

func (self *SetMapProjectionsResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_projected_maps_info := len(self.Go_projected_maps_info)
    buff[offset + 0] = byte((length_projected_maps_info >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_projected_maps_info >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_projected_maps_info >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_projected_maps_info >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_projected_maps_info; i++ {
        offset += self.Go_projected_maps_info[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *SetMapProjectionsResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_projected_maps_info := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_projected_maps_info |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_projected_maps_info |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_projected_maps_info |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_projected_maps_info = make([]*ProjectedMapInfo, length_projected_maps_info)
    for i := 0; i < length_projected_maps_info; i++ {
        self.Go_projected_maps_info[i] = NewProjectedMapInfo()
    }
    for i := 0; i < length_projected_maps_info; i++ {
        offset += self.Go_projected_maps_info[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *SetMapProjectionsResponse) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_projected_maps_info := len(self.Go_projected_maps_info)
    for i := 0; i < length_projected_maps_info; i++ {
        length += self.Go_projected_maps_info[i].Go_serializedLength()
    }
    return length
}

func (self *SetMapProjectionsResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetMapProjectionsResponse) Go_getType() (string) { return "map_msgs/SetMapProjections" }
func (self *SetMapProjectionsResponse) Go_getMD5() (string) { return "47b7931263dc316e5b0f0264428853e9" }
func (self *SetMapProjectionsResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetMapProjectionsResponse) Go_setID(id uint32) { self.__id__ = id }


