package std_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type ColorRGBA struct {
    Go_r float32 `json:"r"`
    Go_g float32 `json:"g"`
    Go_b float32 `json:"b"`
    Go_a float32 `json:"a"`
}

func NewColorRGBA() (*ColorRGBA) {
    newColorRGBA := new(ColorRGBA)
    newColorRGBA.Go_r = 0.0
    newColorRGBA.Go_g = 0.0
    newColorRGBA.Go_b = 0.0
    newColorRGBA.Go_a = 0.0
    return newColorRGBA
}

func (self *ColorRGBA) Go_initialize() {
    self.Go_r = 0.0
    self.Go_g = 0.0
    self.Go_b = 0.0
    self.Go_a = 0.0
}

func (self *ColorRGBA) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_r := math.Float32bits(self.Go_r)
    binary.LittleEndian.PutUint32(buff[offset:], bits_r)
    offset += 4
    bits_g := math.Float32bits(self.Go_g)
    binary.LittleEndian.PutUint32(buff[offset:], bits_g)
    offset += 4
    bits_b := math.Float32bits(self.Go_b)
    binary.LittleEndian.PutUint32(buff[offset:], bits_b)
    offset += 4
    bits_a := math.Float32bits(self.Go_a)
    binary.LittleEndian.PutUint32(buff[offset:], bits_a)
    offset += 4
    return offset
}

func (self *ColorRGBA) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_r := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_r = math.Float32frombits(bits_r)
    offset += 4
    bits_g := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_g = math.Float32frombits(bits_g)
    offset += 4
    bits_b := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_b = math.Float32frombits(bits_b)
    offset += 4
    bits_a := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_a = math.Float32frombits(bits_a)
    offset += 4
    return offset
}

func (self *ColorRGBA) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    length += 4
    return length
}

func (self *ColorRGBA) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ColorRGBA) Go_getType() (string) { return "std_msgs/ColorRGBA" }
func (self *ColorRGBA) Go_getMD5() (string) { return "3c740aa311a337bfa4133c69405e0aed" }
func (self *ColorRGBA) Go_getID() (uint32) { return 0 }
func (self *ColorRGBA) Go_setID(id uint32) { }

