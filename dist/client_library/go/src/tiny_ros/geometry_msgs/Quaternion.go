package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Quaternion struct {
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_z float64 `json:"z"`
    Go_w float64 `json:"w"`
}

func NewQuaternion() (*Quaternion) {
    newQuaternion := new(Quaternion)
    newQuaternion.Go_x = 0.0
    newQuaternion.Go_y = 0.0
    newQuaternion.Go_z = 0.0
    newQuaternion.Go_w = 0.0
    return newQuaternion
}

func (self *Quaternion) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_z = 0.0
    self.Go_w = 0.0
}

func (self *Quaternion) Go_serialize(buff []byte) (int) {
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
    bits_w := math.Float64bits(self.Go_w)
    binary.LittleEndian.PutUint64(buff[offset:], bits_w)
    offset += 8
    return offset
}

func (self *Quaternion) Go_deserialize(buff []byte) (int) {
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
    bits_w := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_w = math.Float64frombits(bits_w)
    offset += 8
    return offset
}

func (self *Quaternion) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *Quaternion) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Quaternion) Go_getType() (string) { return "geometry_msgs/Quaternion" }
func (self *Quaternion) Go_getMD5() (string) { return "175c1571887d10ebed42ba6c042ddd88" }
func (self *Quaternion) Go_getID() (uint32) { return 0 }
func (self *Quaternion) Go_setID(id uint32) { }

