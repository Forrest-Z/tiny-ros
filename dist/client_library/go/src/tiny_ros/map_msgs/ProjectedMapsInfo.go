package map_msgs

import (
    "encoding/json"
)



type ProjectedMapsInfoRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_projected_maps_info []*ProjectedMapInfo `json:"projected_maps_info"`
}

func NewProjectedMapsInfoRequest() (*ProjectedMapsInfoRequest) {
    newProjectedMapsInfoRequest := new(ProjectedMapsInfoRequest)
    newProjectedMapsInfoRequest.Go_projected_maps_info = []*ProjectedMapInfo{}
    newProjectedMapsInfoRequest.__id__ = 0
    return newProjectedMapsInfoRequest
}

func (self *ProjectedMapsInfoRequest) Go_initialize() {
    self.Go_projected_maps_info = []*ProjectedMapInfo{}
    self.__id__ = 0
}

func (self *ProjectedMapsInfoRequest) Go_serialize(buff []byte) (int) {
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

func (self *ProjectedMapsInfoRequest) Go_deserialize(buff []byte) (int) {
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

func (self *ProjectedMapsInfoRequest) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_projected_maps_info := len(self.Go_projected_maps_info)
    for i := 0; i < length_projected_maps_info; i++ {
        length += self.Go_projected_maps_info[i].Go_serializedLength()
    }
    return length
}

func (self *ProjectedMapsInfoRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ProjectedMapsInfoRequest) Go_getType() (string) { return "map_msgs/ProjectedMapsInfo" }
func (self *ProjectedMapsInfoRequest) Go_getMD5() (string) { return "59778fc7286f314a408be52b4611a8b4" }
func (self *ProjectedMapsInfoRequest) Go_getID() (uint32) { return self.__id__ }
func (self *ProjectedMapsInfoRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type ProjectedMapsInfoResponse struct {
    __id__ uint32 `json:"__id__"`
}

func NewProjectedMapsInfoResponse() (*ProjectedMapsInfoResponse) {
    newProjectedMapsInfoResponse := new(ProjectedMapsInfoResponse)
    newProjectedMapsInfoResponse.__id__ = 0
    return newProjectedMapsInfoResponse
}

func (self *ProjectedMapsInfoResponse) Go_initialize() {
    self.__id__ = 0
}

func (self *ProjectedMapsInfoResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *ProjectedMapsInfoResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *ProjectedMapsInfoResponse) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *ProjectedMapsInfoResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ProjectedMapsInfoResponse) Go_getType() (string) { return "map_msgs/ProjectedMapsInfo" }
func (self *ProjectedMapsInfoResponse) Go_getMD5() (string) { return "223a7c48f052c5181dd525823dcc67fc" }
func (self *ProjectedMapsInfoResponse) Go_getID() (uint32) { return self.__id__ }
func (self *ProjectedMapsInfoResponse) Go_setID(id uint32) { self.__id__ = id }


