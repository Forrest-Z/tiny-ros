package tf

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

func (self *FrameGraphRequest) Go_getType() (string) { return "tf/FrameGraph" }
func (self *FrameGraphRequest) Go_getMD5() (string) { return "d66179e20d21cee31291f40363976e1d" }
func (self *FrameGraphRequest) Go_getID() (uint32) { return self.__id__ }
func (self *FrameGraphRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type FrameGraphResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_dot_graph string `json:"dot_graph"`
}

func NewFrameGraphResponse() (*FrameGraphResponse) {
    newFrameGraphResponse := new(FrameGraphResponse)
    newFrameGraphResponse.Go_dot_graph = ""
    newFrameGraphResponse.__id__ = 0
    return newFrameGraphResponse
}

func (self *FrameGraphResponse) Go_initialize() {
    self.Go_dot_graph = ""
    self.__id__ = 0
}

func (self *FrameGraphResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_dot_graph := len(self.Go_dot_graph)
    buff[offset + 0] = byte((length_dot_graph >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_dot_graph >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_dot_graph >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_dot_graph >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_dot_graph)], self.Go_dot_graph)
    offset += length_dot_graph
    return offset
}

func (self *FrameGraphResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_dot_graph := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_dot_graph |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_dot_graph |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_dot_graph |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_dot_graph = string(buff[offset:(offset+length_dot_graph)])
    offset += length_dot_graph
    return offset
}

func (self *FrameGraphResponse) Go_serializedLength() (int) {
    length := 0
    length_dot_graph := len(self.Go_dot_graph)
    length += 4
    length += length_dot_graph
    return length
}

func (self *FrameGraphResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *FrameGraphResponse) Go_getType() (string) { return "tf/FrameGraph" }
func (self *FrameGraphResponse) Go_getMD5() (string) { return "8528671f80a5c0815f9700a446efbc36" }
func (self *FrameGraphResponse) Go_getID() (uint32) { return self.__id__ }
func (self *FrameGraphResponse) Go_setID(id uint32) { self.__id__ = id }


