package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type Temperature struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_temperature float64 `json:"temperature"`
    Go_variance float64 `json:"variance"`
}

func NewTemperature() (*Temperature) {
    newTemperature := new(Temperature)
    newTemperature.Go_header = std_msgs.NewHeader()
    newTemperature.Go_temperature = 0.0
    newTemperature.Go_variance = 0.0
    return newTemperature
}

func (self *Temperature) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_temperature = 0.0
    self.Go_variance = 0.0
}

func (self *Temperature) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_temperature := math.Float64bits(self.Go_temperature)
    binary.LittleEndian.PutUint64(buff[offset:], bits_temperature)
    offset += 8
    bits_variance := math.Float64bits(self.Go_variance)
    binary.LittleEndian.PutUint64(buff[offset:], bits_variance)
    offset += 8
    return offset
}

func (self *Temperature) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_temperature := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_temperature = math.Float64frombits(bits_temperature)
    offset += 8
    bits_variance := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_variance = math.Float64frombits(bits_variance)
    offset += 8
    return offset
}

func (self *Temperature) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 8
    length += 8
    return length
}

func (self *Temperature) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Temperature) Go_getType() (string) { return "sensor_msgs/Temperature" }
func (self *Temperature) Go_getMD5() (string) { return "898a0e5950c8e4073c0c3cf2d7e7bf26" }
func (self *Temperature) Go_getID() (uint32) { return 0 }
func (self *Temperature) Go_setID(id uint32) { }

