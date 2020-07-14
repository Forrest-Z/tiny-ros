package nav_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
    "encoding/binary"
    "math"
    "tiny_ros/geometry_msgs"
)


type MapMetaData struct {
    Go_map_load_time *rostime.Time `json:"map_load_time"`
    Go_resolution float32 `json:"resolution"`
    Go_width uint32 `json:"width"`
    Go_height uint32 `json:"height"`
    Go_origin *geometry_msgs.Pose `json:"origin"`
}

func NewMapMetaData() (*MapMetaData) {
    newMapMetaData := new(MapMetaData)
    newMapMetaData.Go_map_load_time = rostime.NewTime()
    newMapMetaData.Go_resolution = 0.0
    newMapMetaData.Go_width = 0
    newMapMetaData.Go_height = 0
    newMapMetaData.Go_origin = geometry_msgs.NewPose()
    return newMapMetaData
}

func (self *MapMetaData) Go_initialize() {
    self.Go_map_load_time = rostime.NewTime()
    self.Go_resolution = 0.0
    self.Go_width = 0
    self.Go_height = 0
    self.Go_origin = geometry_msgs.NewPose()
}

func (self *MapMetaData) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_map_load_time.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_map_load_time.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_map_load_time.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_map_load_time.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_map_load_time.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_map_load_time.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_map_load_time.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_map_load_time.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    bits_resolution := math.Float32bits(self.Go_resolution)
    binary.LittleEndian.PutUint32(buff[offset:], bits_resolution)
    offset += 4
    buff[offset + 0] = byte((self.Go_width >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_width >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_width >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_width >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_height >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_height >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_height >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_height >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_origin.Go_serialize(buff[offset:])
    return offset
}

func (self *MapMetaData) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_map_load_time.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_map_load_time.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_map_load_time.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_map_load_time.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_map_load_time.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_map_load_time.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_map_load_time.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_map_load_time.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_resolution := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_resolution = math.Float32frombits(bits_resolution)
    offset += 4
    self.Go_width = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_width |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_width |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_width |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_height = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_height |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_height |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_height |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_origin.Go_deserialize(buff[offset:])
    return offset
}

func (self *MapMetaData) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += self.Go_origin.Go_serializedLength()
    return length
}

func (self *MapMetaData) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MapMetaData) Go_getType() (string) { return "nav_msgs/MapMetaData" }
func (self *MapMetaData) Go_getMD5() (string) { return "328f5a1f2242fff4676d48189bd8b309" }
func (self *MapMetaData) Go_getID() (uint32) { return 0 }
func (self *MapMetaData) Go_setID(id uint32) { }

