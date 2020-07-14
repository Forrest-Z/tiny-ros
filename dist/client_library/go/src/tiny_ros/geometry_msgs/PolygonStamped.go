package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type PolygonStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_polygon *Polygon `json:"polygon"`
}

func NewPolygonStamped() (*PolygonStamped) {
    newPolygonStamped := new(PolygonStamped)
    newPolygonStamped.Go_header = std_msgs.NewHeader()
    newPolygonStamped.Go_polygon = NewPolygon()
    return newPolygonStamped
}

func (self *PolygonStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_polygon = NewPolygon()
}

func (self *PolygonStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_polygon.Go_serialize(buff[offset:])
    return offset
}

func (self *PolygonStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_polygon.Go_deserialize(buff[offset:])
    return offset
}

func (self *PolygonStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_polygon.Go_serializedLength()
    return length
}

func (self *PolygonStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PolygonStamped) Go_getType() (string) { return "geometry_msgs/PolygonStamped" }
func (self *PolygonStamped) Go_getMD5() (string) { return "33bdf94066425e572879b25c9a51ed50" }
func (self *PolygonStamped) Go_getID() (uint32) { return 0 }
func (self *PolygonStamped) Go_setID(id uint32) { }

