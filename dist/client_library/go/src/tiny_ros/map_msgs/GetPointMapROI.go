package map_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/sensor_msgs"
    "math"
)



type GetPointMapROIRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_z float64 `json:"z"`
    Go_r float64 `json:"r"`
    Go_l_x float64 `json:"l_x"`
    Go_l_y float64 `json:"l_y"`
    Go_l_z float64 `json:"l_z"`
}

func NewGetPointMapROIRequest() (*GetPointMapROIRequest) {
    newGetPointMapROIRequest := new(GetPointMapROIRequest)
    newGetPointMapROIRequest.Go_x = 0.0
    newGetPointMapROIRequest.Go_y = 0.0
    newGetPointMapROIRequest.Go_z = 0.0
    newGetPointMapROIRequest.Go_r = 0.0
    newGetPointMapROIRequest.Go_l_x = 0.0
    newGetPointMapROIRequest.Go_l_y = 0.0
    newGetPointMapROIRequest.Go_l_z = 0.0
    newGetPointMapROIRequest.__id__ = 0
    return newGetPointMapROIRequest
}

func (self *GetPointMapROIRequest) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_z = 0.0
    self.Go_r = 0.0
    self.Go_l_x = 0.0
    self.Go_l_y = 0.0
    self.Go_l_z = 0.0
    self.__id__ = 0
}

func (self *GetPointMapROIRequest) Go_serialize(buff []byte) (int) {
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
    bits_z := math.Float64bits(self.Go_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_z)
    offset += 8
    bits_r := math.Float64bits(self.Go_r)
    binary.LittleEndian.PutUint64(buff[offset:], bits_r)
    offset += 8
    bits_l_x := math.Float64bits(self.Go_l_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_l_x)
    offset += 8
    bits_l_y := math.Float64bits(self.Go_l_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_l_y)
    offset += 8
    bits_l_z := math.Float64bits(self.Go_l_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_l_z)
    offset += 8
    return offset
}

func (self *GetPointMapROIRequest) Go_deserialize(buff []byte) (int) {
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
    bits_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_z = math.Float64frombits(bits_z)
    offset += 8
    bits_r := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_r = math.Float64frombits(bits_r)
    offset += 8
    bits_l_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_l_x = math.Float64frombits(bits_l_x)
    offset += 8
    bits_l_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_l_y = math.Float64frombits(bits_l_y)
    offset += 8
    bits_l_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_l_z = math.Float64frombits(bits_l_z)
    offset += 8
    return offset
}

func (self *GetPointMapROIRequest) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *GetPointMapROIRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPointMapROIRequest) Go_getType() (string) { return "map_msgs/GetPointMapROI" }
func (self *GetPointMapROIRequest) Go_getMD5() (string) { return "c338f5616930e00a72c38486f77e9810" }
func (self *GetPointMapROIRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetPointMapROIRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetPointMapROIResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_sub_map *sensor_msgs.PointCloud2 `json:"sub_map"`
}

func NewGetPointMapROIResponse() (*GetPointMapROIResponse) {
    newGetPointMapROIResponse := new(GetPointMapROIResponse)
    newGetPointMapROIResponse.Go_sub_map = sensor_msgs.NewPointCloud2()
    newGetPointMapROIResponse.__id__ = 0
    return newGetPointMapROIResponse
}

func (self *GetPointMapROIResponse) Go_initialize() {
    self.Go_sub_map = sensor_msgs.NewPointCloud2()
    self.__id__ = 0
}

func (self *GetPointMapROIResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_sub_map.Go_serialize(buff[offset:])
    return offset
}

func (self *GetPointMapROIResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_sub_map.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetPointMapROIResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_sub_map.Go_serializedLength()
    return length
}

func (self *GetPointMapROIResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPointMapROIResponse) Go_getType() (string) { return "map_msgs/GetPointMapROI" }
func (self *GetPointMapROIResponse) Go_getMD5() (string) { return "de10fb68f23987142a0ffbdb59b6e079" }
func (self *GetPointMapROIResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetPointMapROIResponse) Go_setID(id uint32) { self.__id__ = id }


