package geometry_msgs

import (
    "encoding/json"
)


type Polygon struct {
    Go_points []*Point32 `json:"points"`
}

func NewPolygon() (*Polygon) {
    newPolygon := new(Polygon)
    newPolygon.Go_points = []*Point32{}
    return newPolygon
}

func (self *Polygon) Go_initialize() {
    self.Go_points = []*Point32{}
}

func (self *Polygon) Go_serialize(buff []byte) (int) {
    offset := 0
    length_points := len(self.Go_points)
    buff[offset + 0] = byte((length_points >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_points >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_points >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_points >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *Polygon) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_points := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_points |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_points |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_points |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_points = make([]*Point32, length_points)
    for i := 0; i < length_points; i++ {
        self.Go_points[i] = NewPoint32()
    }
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *Polygon) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_points := len(self.Go_points)
    for i := 0; i < length_points; i++ {
        length += self.Go_points[i].Go_serializedLength()
    }
    return length
}

func (self *Polygon) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Polygon) Go_getType() (string) { return "geometry_msgs/Polygon" }
func (self *Polygon) Go_getMD5() (string) { return "f94a78a947b7879954bd14397db4bc9d" }
func (self *Polygon) Go_getID() (uint32) { return 0 }
func (self *Polygon) Go_setID(id uint32) { }

