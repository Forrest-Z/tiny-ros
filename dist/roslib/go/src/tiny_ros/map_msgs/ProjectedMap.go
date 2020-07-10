package map_msgs

import (
    "encoding/json"
    "tiny_ros/nav_msgs"
    "encoding/binary"
    "math"
)


type ProjectedMap struct {
    Go_map *nav_msgs.OccupancyGrid `json:"map"`
    Go_min_z float64 `json:"min_z"`
    Go_max_z float64 `json:"max_z"`
}

func NewProjectedMap() (*ProjectedMap) {
    newProjectedMap := new(ProjectedMap)
    newProjectedMap.Go_map = nav_msgs.NewOccupancyGrid()
    newProjectedMap.Go_min_z = 0.0
    newProjectedMap.Go_max_z = 0.0
    return newProjectedMap
}

func (self *ProjectedMap) Go_initialize() {
    self.Go_map = nav_msgs.NewOccupancyGrid()
    self.Go_min_z = 0.0
    self.Go_max_z = 0.0
}

func (self *ProjectedMap) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_map.Go_serialize(buff[offset:])
    bits_min_z := math.Float64bits(self.Go_min_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_min_z)
    offset += 8
    bits_max_z := math.Float64bits(self.Go_max_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_max_z)
    offset += 8
    return offset
}

func (self *ProjectedMap) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_map.Go_deserialize(buff[offset:])
    bits_min_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_min_z = math.Float64frombits(bits_min_z)
    offset += 8
    bits_max_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_max_z = math.Float64frombits(bits_max_z)
    offset += 8
    return offset
}

func (self *ProjectedMap) Go_serializedLength() (int) {
    length := 0
    length += self.Go_map.Go_serializedLength()
    length += 8
    length += 8
    return length
}

func (self *ProjectedMap) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ProjectedMap) Go_getType() (string) { return "map_msgs/ProjectedMap" }
func (self *ProjectedMap) Go_getMD5() (string) { return "cbd5598c259cc16f5aa07335587a7367" }
func (self *ProjectedMap) Go_getID() (uint32) { return 0 }
func (self *ProjectedMap) Go_setID(id uint32) { }

