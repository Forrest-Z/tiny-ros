package shape_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Plane struct {
    Go_coef [4]float64 `json:"coef"`
}

func NewPlane() (*Plane) {
    newPlane := new(Plane)
    newPlane.Go_coef = [4]float64{0.0, 0.0, 0.0, 0.0}
    return newPlane
}

func (self *Plane) Go_initialize() {
    self.Go_coef = [4]float64{0.0, 0.0, 0.0, 0.0}
}

func (self *Plane) Go_serialize(buff []byte) (int) {
    offset := 0
    for i := 0; i < 4; i++ {
        bits_coefi := math.Float64bits(self.Go_coef[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_coefi)
        offset += 8
    }
    return offset
}

func (self *Plane) Go_deserialize(buff []byte) (int) {
    offset := 0
    for i := 0; i < 4; i++ {
        bits_coefi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_coef[i] = math.Float64frombits(bits_coefi)
        offset += 8
    }
    return offset
}

func (self *Plane) Go_serializedLength() (int) {
    length := 0
    for i := 0; i < 4; i++ {
        length += 8
    }
    return length
}

func (self *Plane) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Plane) Go_getType() (string) { return "shape_msgs/Plane" }
func (self *Plane) Go_getMD5() (string) { return "770421286b7c90effe8aac9f1c37eac0" }
func (self *Plane) Go_getID() (uint32) { return 0 }
func (self *Plane) Go_setID(id uint32) { }

