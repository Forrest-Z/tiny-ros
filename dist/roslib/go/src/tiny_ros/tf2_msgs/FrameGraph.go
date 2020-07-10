package tf2_msgs

import (
    "encoding/json"
)



type FrameGraphRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewFrameGraphRequest() (*FrameGraphRequest) {
    newFrameGraphRequest := new(FrameGraphRequest)
    newFrameGraphRequest.__id__ = 0
    return newFrameGraphRequest
}

func (self *FrameGraphRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *FrameGraphRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *FrameGraphRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *FrameGraphRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *FrameGraphRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *FrameGraphRequest) Go_getType() (string) { return "tf2_msgs/FrameGraph" }
func (self *FrameGraphRequest) Go_getMD5() (string) { return "aa023d3c31410363e0583979223258c8" }
func (self *FrameGraphRequest) Go_getID() (uint32) { return self.__id__ }
func (self *FrameGraphRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type FrameGraphResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_frame_yaml string `json:"frame_yaml"`
}

func NewFrameGraphResponse() (*FrameGraphResponse) {
    newFrameGraphResponse := new(FrameGraphResponse)
    newFrameGraphResponse.Go_frame_yaml = ""
    newFrameGraphResponse.__id__ = 0
    return newFrameGraphResponse
}

func (self *FrameGraphResponse) Go_initialize() {
    self.Go_frame_yaml = ""
    self.__id__ = 0
}

func (self *FrameGraphResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_frame_yaml := len(self.Go_frame_yaml)
    buff[offset + 0] = byte((length_frame_yaml >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_frame_yaml >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_frame_yaml >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_frame_yaml >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_frame_yaml)], self.Go_frame_yaml)
    offset += length_frame_yaml
    return offset
}

func (self *FrameGraphResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_frame_yaml := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_frame_yaml |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_frame_yaml |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_frame_yaml |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_frame_yaml = string(buff[offset:(offset+length_frame_yaml)])
    offset += length_frame_yaml
    return offset
}

func (self *FrameGraphResponse) Go_serializedLength() (int) {
    length := 0
    length_frame_yaml := len(self.Go_frame_yaml)
    length += 4
    length += length_frame_yaml
    return length
}

func (self *FrameGraphResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *FrameGraphResponse) Go_getType() (string) { return "tf2_msgs/FrameGraph" }
func (self *FrameGraphResponse) Go_getMD5() (string) { return "97e361486f8cc8fb1a460cf17555126b" }
func (self *FrameGraphResponse) Go_getID() (uint32) { return self.__id__ }
func (self *FrameGraphResponse) Go_setID(id uint32) { self.__id__ = id }


