package map_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type ProjectedMapInfo struct {
    Go_frame_id string `json:"frame_id"`
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_width float64 `json:"width"`
    Go_height float64 `json:"height"`
    Go_min_z float64 `json:"min_z"`
    Go_max_z float64 `json:"max_z"`
}

func NewProjectedMapInfo() (*ProjectedMapInfo) {
    newProjectedMapInfo := new(ProjectedMapInfo)
    newProjectedMapInfo.Go_frame_id = ""
    newProjectedMapInfo.Go_x = 0.0
    newProjectedMapInfo.Go_y = 0.0
    newProjectedMapInfo.Go_width = 0.0
    newProjectedMapInfo.Go_height = 0.0
    newProjectedMapInfo.Go_min_z = 0.0
    newProjectedMapInfo.Go_max_z = 0.0
    return newProjectedMapInfo
}

func (self *ProjectedMapInfo) Go_initialize() {
    self.Go_frame_id = ""
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_width = 0.0
    self.Go_height = 0.0
    self.Go_min_z = 0.0
    self.Go_max_z = 0.0
}

func (self *ProjectedMapInfo) Go_serialize(buff []byte) (int) {
    offset := 0
    length_frame_id := len(self.Go_frame_id)
    buff[offset + 0] = byte((length_frame_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_frame_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_frame_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_frame_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_frame_id)], self.Go_frame_id)
    offset += length_frame_id
    bits_x := math.Float64bits(self.Go_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_x)
    offset += 8
    bits_y := math.Float64bits(self.Go_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_y)
    offset += 8
    bits_width := math.Float64bits(self.Go_width)
    binary.LittleEndian.PutUint64(buff[offset:], bits_width)
    offset += 8
    bits_height := math.Float64bits(self.Go_height)
    binary.LittleEndian.PutUint64(buff[offset:], bits_height)
    offset += 8
    bits_min_z := math.Float64bits(self.Go_min_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_min_z)
    offset += 8
    bits_max_z := math.Float64bits(self.Go_max_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_max_z)
    offset += 8
    return offset
}

func (self *ProjectedMapInfo) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_frame_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_frame_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_frame_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_frame_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_frame_id = string(buff[offset:(offset+length_frame_id)])
    offset += length_frame_id
    bits_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_x = math.Float64frombits(bits_x)
    offset += 8
    bits_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_y = math.Float64frombits(bits_y)
    offset += 8
    bits_width := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_width = math.Float64frombits(bits_width)
    offset += 8
    bits_height := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_height = math.Float64frombits(bits_height)
    offset += 8
    bits_min_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_min_z = math.Float64frombits(bits_min_z)
    offset += 8
    bits_max_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_max_z = math.Float64frombits(bits_max_z)
    offset += 8
    return offset
}

func (self *ProjectedMapInfo) Go_serializedLength() (int) {
    length := 0
    length_frame_id := len(self.Go_frame_id)
    length += 4
    length += length_frame_id
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *ProjectedMapInfo) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ProjectedMapInfo) Go_getType() (string) { return "map_msgs/ProjectedMapInfo" }
func (self *ProjectedMapInfo) Go_getMD5() (string) { return "f661365637fb759e63cb5d179a4461e1" }
func (self *ProjectedMapInfo) Go_getID() (uint32) { return 0 }
func (self *ProjectedMapInfo) Go_setID(id uint32) { }

