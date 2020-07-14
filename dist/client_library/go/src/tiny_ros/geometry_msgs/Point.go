package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Point struct {
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_z float64 `json:"z"`
}

func NewPoint() (*Point) {
    newPoint := new(Point)
    newPoint.Go_x = 0.0
    newPoint.Go_y = 0.0
    newPoint.Go_z = 0.0
    return newPoint
}

func (self *Point) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_z = 0.0
}

func (self *Point) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_x := math.Float64bits(self.Go_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_x)
    offset += 8
    bits_y := math.Float64bits(self.Go_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_y)
    offset += 8
    bits_z := math.Float64bits(self.Go_z)
    binary.LittleEndian.PutUint64(buff[offset:], bits_z)
    offset += 8
    return offset
}

func (self *Point) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_x = math.Float64frombits(bits_x)
    offset += 8
    bits_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_y = math.Float64frombits(bits_y)
    offset += 8
    bits_z := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_z = math.Float64frombits(bits_z)
    offset += 8
    return offset
}

func (self *Point) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    return length
}

func (self *Point) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Point) Go_getType() (string) { return "geometry_msgs/Point" }
func (self *Point) Go_getMD5() (string) { return "f75eead1a8b17241f0c81a1de081b731" }
func (self *Point) Go_getID() (uint32) { return 0 }
func (self *Point) Go_setID(id uint32) { }

