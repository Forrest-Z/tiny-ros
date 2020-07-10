package std_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Float32 struct {
    Go_data float32 `json:"data"`
}

func NewFloat32() (*Float32) {
    newFloat32 := new(Float32)
    newFloat32.Go_data = 0.0
    return newFloat32
}

func (self *Float32) Go_initialize() {
    self.Go_data = 0.0
}

func (self *Float32) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_data := math.Float32bits(self.Go_data)
    binary.LittleEndian.PutUint32(buff[offset:], bits_data)
    offset += 4
    return offset
}

func (self *Float32) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_data := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_data = math.Float32frombits(bits_data)
    offset += 4
    return offset
}

func (self *Float32) Go_serializedLength() (int) {
    length := 0
    length += 4
    return length
}

func (self *Float32) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Float32) Go_getType() (string) { return "std_msgs/Float32" }
func (self *Float32) Go_getMD5() (string) { return "2aff5d2343e8e80ceea1362fc770035c" }
func (self *Float32) Go_getID() (uint32) { return 0 }
func (self *Float32) Go_setID(id uint32) { }

