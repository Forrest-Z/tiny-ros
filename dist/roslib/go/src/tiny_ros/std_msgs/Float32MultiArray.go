package std_msgs

import (
    "encoding/binary"
    "math"
)


type Float32MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []float32 `json:"data"`
}

func NewFloat32MultiArray() (*Float32MultiArray) {
    newFloat32MultiArray := new(Float32MultiArray)
    newFloat32MultiArray.Go_layout = NewMultiArrayLayout()
    newFloat32MultiArray.Go_data = []float32{}
    return newFloat32MultiArray
}

func (self *Float32MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []float32{}
}

func (self *Float32MultiArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_serialize(buff[offset:])
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_data; i++ {
        bits_datai := math.Float32bits(self.Go_data[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_datai)
        offset += 4
    }
    return offset
}

func (self *Float32MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]float32, length_data, length_data)
    for i := 0; i < length_data; i++ {
        bits_datai := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_data[i] = math.Float32frombits(bits_datai)
        offset += 4
    }
    return offset
}

func (self *Float32MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 4
    }
    return length
}

func (self *Float32MultiArray) Go_echo() (string) { return "" }
func (self *Float32MultiArray) Go_getType() (string) { return "std_msgs/Float32MultiArray" }
func (self *Float32MultiArray) Go_getMD5() (string) { return "224f9a21761656b5f5da2b311973577f" }
func (self *Float32MultiArray) Go_getID() (uint32) { return 0 }
func (self *Float32MultiArray) Go_setID(id uint32) { }

