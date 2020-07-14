package map_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/sensor_msgs"
)

func Go_ADD() (uint32) { return 0 }
func Go_DELETE() (uint32) { return 1 }

type PointCloud2Update struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_type uint32 `json:"type"`
    Go_points *sensor_msgs.PointCloud2 `json:"points"`
}

func NewPointCloud2Update() (*PointCloud2Update) {
    newPointCloud2Update := new(PointCloud2Update)
    newPointCloud2Update.Go_header = std_msgs.NewHeader()
    newPointCloud2Update.Go_type = 0
    newPointCloud2Update.Go_points = sensor_msgs.NewPointCloud2()
    return newPointCloud2Update
}

func (self *PointCloud2Update) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_type = 0
    self.Go_points = sensor_msgs.NewPointCloud2()
}

func (self *PointCloud2Update) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_type >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_type >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_type >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_type >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_points.Go_serialize(buff[offset:])
    return offset
}

func (self *PointCloud2Update) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_type = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_type |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_type |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_type |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_points.Go_deserialize(buff[offset:])
    return offset
}

func (self *PointCloud2Update) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += self.Go_points.Go_serializedLength()
    return length
}

func (self *PointCloud2Update) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PointCloud2Update) Go_getType() (string) { return "map_msgs/PointCloud2Update" }
func (self *PointCloud2Update) Go_getMD5() (string) { return "e79dfbefd7336861352e1bc7148491c4" }
func (self *PointCloud2Update) Go_getID() (uint32) { return 0 }
func (self *PointCloud2Update) Go_setID(id uint32) { }

