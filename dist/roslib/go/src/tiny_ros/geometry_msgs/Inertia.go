package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Inertia struct {
    Go_m float64 `json:"m"`
    Go_com *Vector3 `json:"com"`
    Go_ixx float64 `json:"ixx"`
    Go_ixy float64 `json:"ixy"`
    Go_ixz float64 `json:"ixz"`
    Go_iyy float64 `json:"iyy"`
    Go_iyz float64 `json:"iyz"`
    Go_izz float64 `json:"izz"`
}

func NewInertia() (*Inertia) {
    newInertia := new(Inertia)
    newInertia.Go_m = 0.0
    newInertia.Go_com = NewVector3()
    newInertia.Go_ixx = 0.0
    newInertia.Go_ixy = 0.0
    newInertia.Go_ixz = 0.0
    newInertia.Go_iyy = 0.0
    newInertia.Go_iyz = 0.0
    newInertia.Go_izz = 0.0
    return newInertia
}

func (self *Inertia) Go_initialize() {
    self.Go_m = 0.0
    self.Go_com = NewVector3()
    self.Go_ixx = 0.0
    self.Go_ixy = 0.0
    self.Go_ixz = 0.0
    self.Go_iyy = 0.0
    self.Go_iyz = 0.0
    self.Go_izz = 0.0
}

func (self *Inertia) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_m := math.Float64bits(self.Go_m)
    binary.LittleEndian.PutUint64(buff[offset:], bits_m)
    offset += 8
    offset += self.Go_com.Go_serialize(buff[offset:])
    bits_ixx := math.Float64bits(self.Go_ixx)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixx)
    offset += 8
    bits_ixy := math.Float64bits(self.Go_ixy)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixy)
    offset += 8
    bits_ixz := math.Float64bits(self.Go_ixz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixz)
    offset += 8
    bits_iyy := math.Float64bits(self.Go_iyy)
    binary.LittleEndian.PutUint64(buff[offset:], bits_iyy)
    offset += 8
    bits_iyz := math.Float64bits(self.Go_iyz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_iyz)
    offset += 8
    bits_izz := math.Float64bits(self.Go_izz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_izz)
    offset += 8
    return offset
}

func (self *Inertia) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_m := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_m = math.Float64frombits(bits_m)
    offset += 8
    offset += self.Go_com.Go_deserialize(buff[offset:])
    bits_ixx := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixx = math.Float64frombits(bits_ixx)
    offset += 8
    bits_ixy := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixy = math.Float64frombits(bits_ixy)
    offset += 8
    bits_ixz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixz = math.Float64frombits(bits_ixz)
    offset += 8
    bits_iyy := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_iyy = math.Float64frombits(bits_iyy)
    offset += 8
    bits_iyz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_iyz = math.Float64frombits(bits_iyz)
    offset += 8
    bits_izz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_izz = math.Float64frombits(bits_izz)
    offset += 8
    return offset
}

func (self *Inertia) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += self.Go_com.Go_serializedLength()
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *Inertia) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Inertia) Go_getType() (string) { return "geometry_msgs/Inertia" }
func (self *Inertia) Go_getMD5() (string) { return "9116c935782bc29999dad1927624dff0" }
func (self *Inertia) Go_getID() (uint32) { return 0 }
func (self *Inertia) Go_setID(id uint32) { }

