package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Point32 struct {
    Go_x float32 `json:"x"`
    Go_y float32 `json:"y"`
    Go_z float32 `json:"z"`
}

func NewPoint32() (*Point32) {
    newPoint32 := new(Point32)
    newPoint32.Go_x = 0.0
    newPoint32.Go_y = 0.0
    newPoint32.Go_z = 0.0
    return newPoint32
}

func (self *Point32) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_z = 0.0
}

func (self *Point32) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_x := math.Float32bits(self.Go_x)
    binary.LittleEndian.PutUint32(buff[offset:], bits_x)
    offset += 4
    bits_y := math.Float32bits(self.Go_y)
    binary.LittleEndian.PutUint32(buff[offset:], bits_y)
    offset += 4
    bits_z := math.Float32bits(self.Go_z)
    binary.LittleEndian.PutUint32(buff[offset:], bits_z)
    offset += 4
    return offset
}

func (self *Point32) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_x := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_x = math.Float32frombits(bits_x)
    offset += 4
    bits_y := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_y = math.Float32frombits(bits_y)
    offset += 4
    bits_z := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_z = math.Float32frombits(bits_z)
    offset += 4
    return offset
}

func (self *Point32) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    return length
}

func (self *Point32) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Point32) Go_getType() (string) { return "geometry_msgs/Point32" }
func (self *Point32) Go_getMD5() (string) { return "b17f2230f465fce816e3773d7d59a841" }
func (self *Point32) Go_getID() (uint32) { return 0 }
func (self *Point32) Go_setID(id uint32) { }

