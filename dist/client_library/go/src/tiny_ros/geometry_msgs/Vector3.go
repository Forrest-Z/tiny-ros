package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Vector3 struct {
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_z float64 `json:"z"`
}

func NewVector3() (*Vector3) {
    newVector3 := new(Vector3)
    newVector3.Go_x = 0.0
    newVector3.Go_y = 0.0
    newVector3.Go_z = 0.0
    return newVector3
}

func (self *Vector3) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_z = 0.0
}

func (self *Vector3) Go_serialize(buff []byte) (int) {
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

func (self *Vector3) Go_deserialize(buff []byte) (int) {
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

func (self *Vector3) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    return length
}

func (self *Vector3) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Vector3) Go_getType() (string) { return "geometry_msgs/Vector3" }
func (self *Vector3) Go_getMD5() (string) { return "b5c8c5b484ec7d5e36a4d9de9124c561" }
func (self *Vector3) Go_getID() (uint32) { return 0 }
func (self *Vector3) Go_setID(id uint32) { }

