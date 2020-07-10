package std_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Float64MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []float64 `json:"data"`
}

func NewFloat64MultiArray() (*Float64MultiArray) {
    newFloat64MultiArray := new(Float64MultiArray)
    newFloat64MultiArray.Go_layout = NewMultiArrayLayout()
    newFloat64MultiArray.Go_data = []float64{}
    return newFloat64MultiArray
}

func (self *Float64MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []float64{}
}

func (self *Float64MultiArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_serialize(buff[offset:])
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_data; i++ {
        bits_datai := math.Float64bits(self.Go_data[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_datai)
        offset += 8
    }
    return offset
}

func (self *Float64MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]float64, length_data)
    for i := 0; i < length_data; i++ {
        bits_datai := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_data[i] = math.Float64frombits(bits_datai)
        offset += 8
    }
    return offset
}

func (self *Float64MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 8
    }
    return length
}

func (self *Float64MultiArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Float64MultiArray) Go_getType() (string) { return "std_msgs/Float64MultiArray" }
func (self *Float64MultiArray) Go_getMD5() (string) { return "e3061da26924f3790a70f9dbf06fc1a5" }
func (self *Float64MultiArray) Go_getID() (uint32) { return 0 }
func (self *Float64MultiArray) Go_setID(id uint32) { }

