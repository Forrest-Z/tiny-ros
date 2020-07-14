package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type PointStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_point *Point `json:"point"`
}

func NewPointStamped() (*PointStamped) {
    newPointStamped := new(PointStamped)
    newPointStamped.Go_header = std_msgs.NewHeader()
    newPointStamped.Go_point = NewPoint()
    return newPointStamped
}

func (self *PointStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_point = NewPoint()
}

func (self *PointStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_point.Go_serialize(buff[offset:])
    return offset
}

func (self *PointStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_point.Go_deserialize(buff[offset:])
    return offset
}

func (self *PointStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_point.Go_serializedLength()
    return length
}

func (self *PointStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PointStamped) Go_getType() (string) { return "geometry_msgs/PointStamped" }
func (self *PointStamped) Go_getMD5() (string) { return "d34e83bdbef7bf4b617a6293aab8390e" }
func (self *PointStamped) Go_getID() (uint32) { return 0 }
func (self *PointStamped) Go_setID(id uint32) { }

