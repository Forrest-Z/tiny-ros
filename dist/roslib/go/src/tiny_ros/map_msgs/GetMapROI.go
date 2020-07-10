package map_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/nav_msgs"
    "math"
)



type GetMapROIRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_l_x float64 `json:"l_x"`
    Go_l_y float64 `json:"l_y"`
}

func NewGetMapROIRequest() (*GetMapROIRequest) {
    newGetMapROIRequest := new(GetMapROIRequest)
    newGetMapROIRequest.Go_x = 0.0
    newGetMapROIRequest.Go_y = 0.0
    newGetMapROIRequest.Go_l_x = 0.0
    newGetMapROIRequest.Go_l_y = 0.0
    newGetMapROIRequest.__id__ = 0
    return newGetMapROIRequest
}

func (self *GetMapROIRequest) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_l_x = 0.0
    self.Go_l_y = 0.0
    self.__id__ = 0
}

func (self *GetMapROIRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    bits_x := math.Float64bits(self.Go_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_x)
    offset += 8
    bits_y := math.Float64bits(self.Go_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_y)
    offset += 8
    bits_l_x := math.Float64bits(self.Go_l_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_l_x)
    offset += 8
    bits_l_y := math.Float64bits(self.Go_l_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_l_y)
    offset += 8
    return offset
}

func (self *GetMapROIRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_x = math.Float64frombits(bits_x)
    offset += 8
    bits_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_y = math.Float64frombits(bits_y)
    offset += 8
    bits_l_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_l_x = math.Float64frombits(bits_l_x)
    offset += 8
    bits_l_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_l_y = math.Float64frombits(bits_l_y)
    offset += 8
    return offset
}

func (self *GetMapROIRequest) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *GetMapROIRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapROIRequest) Go_getType() (string) { return "map_msgs/GetMapROI" }
func (self *GetMapROIRequest) Go_getMD5() (string) { return "f74ea8c8dc9b857aae7ea10033520d28" }
func (self *GetMapROIRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetMapROIRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetMapROIResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_sub_map *nav_msgs.OccupancyGrid `json:"sub_map"`
}

func NewGetMapROIResponse() (*GetMapROIResponse) {
    newGetMapROIResponse := new(GetMapROIResponse)
    newGetMapROIResponse.Go_sub_map = nav_msgs.NewOccupancyGrid()
    newGetMapROIResponse.__id__ = 0
    return newGetMapROIResponse
}

func (self *GetMapROIResponse) Go_initialize() {
    self.Go_sub_map = nav_msgs.NewOccupancyGrid()
    self.__id__ = 0
}

func (self *GetMapROIResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_sub_map.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapROIResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_sub_map.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapROIResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_sub_map.Go_serializedLength()
    return length
}

func (self *GetMapROIResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapROIResponse) Go_getType() (string) { return "map_msgs/GetMapROI" }
func (self *GetMapROIResponse) Go_getMD5() (string) { return "a178ec520c3d0d99d9d85c70ed4b535a" }
func (self *GetMapROIResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetMapROIResponse) Go_setID(id uint32) { self.__id__ = id }


