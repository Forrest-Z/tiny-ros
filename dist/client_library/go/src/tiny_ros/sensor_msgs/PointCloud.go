package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
)


type PointCloud struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_points []*geometry_msgs.Point32 `json:"points"`
    Go_channels []*ChannelFloat32 `json:"channels"`
}

func NewPointCloud() (*PointCloud) {
    newPointCloud := new(PointCloud)
    newPointCloud.Go_header = std_msgs.NewHeader()
    newPointCloud.Go_points = []*geometry_msgs.Point32{}
    newPointCloud.Go_channels = []*ChannelFloat32{}
    return newPointCloud
}

func (self *PointCloud) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_points = []*geometry_msgs.Point32{}
    self.Go_channels = []*ChannelFloat32{}
}

func (self *PointCloud) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_points := len(self.Go_points)
    buff[offset + 0] = byte((length_points >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_points >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_points >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_points >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_serialize(buff[offset:])
    }
    length_channels := len(self.Go_channels)
    buff[offset + 0] = byte((length_channels >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_channels >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_channels >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_channels >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_channels; i++ {
        offset += self.Go_channels[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *PointCloud) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_points := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_points |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_points |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_points |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_points = make([]*geometry_msgs.Point32, length_points)
    for i := 0; i < length_points; i++ {
        self.Go_points[i] = geometry_msgs.NewPoint32()
    }
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_deserialize(buff[offset:])
    }
    length_channels := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_channels |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_channels |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_channels |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_channels = make([]*ChannelFloat32, length_channels)
    for i := 0; i < length_channels; i++ {
        self.Go_channels[i] = NewChannelFloat32()
    }
    for i := 0; i < length_channels; i++ {
        offset += self.Go_channels[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *PointCloud) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_points := len(self.Go_points)
    for i := 0; i < length_points; i++ {
        length += self.Go_points[i].Go_serializedLength()
    }
    length += 4
    length_channels := len(self.Go_channels)
    for i := 0; i < length_channels; i++ {
        length += self.Go_channels[i].Go_serializedLength()
    }
    return length
}

func (self *PointCloud) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PointCloud) Go_getType() (string) { return "sensor_msgs/PointCloud" }
func (self *PointCloud) Go_getMD5() (string) { return "b01249148cae0106a561ab36cd1e48a8" }
func (self *PointCloud) Go_getID() (uint32) { return 0 }
func (self *PointCloud) Go_setID(id uint32) { }

