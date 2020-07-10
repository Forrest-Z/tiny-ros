package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type Illuminance struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_illuminance float64 `json:"illuminance"`
    Go_variance float64 `json:"variance"`
}

func NewIlluminance() (*Illuminance) {
    newIlluminance := new(Illuminance)
    newIlluminance.Go_header = std_msgs.NewHeader()
    newIlluminance.Go_illuminance = 0.0
    newIlluminance.Go_variance = 0.0
    return newIlluminance
}

func (self *Illuminance) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_illuminance = 0.0
    self.Go_variance = 0.0
}

func (self *Illuminance) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_illuminance := math.Float64bits(self.Go_illuminance)
    binary.LittleEndian.PutUint64(buff[offset:], bits_illuminance)
    offset += 8
    bits_variance := math.Float64bits(self.Go_variance)
    binary.LittleEndian.PutUint64(buff[offset:], bits_variance)
    offset += 8
    return offset
}

func (self *Illuminance) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_illuminance := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_illuminance = math.Float64frombits(bits_illuminance)
    offset += 8
    bits_variance := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_variance = math.Float64frombits(bits_variance)
    offset += 8
    return offset
}

func (self *Illuminance) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 8
    length += 8
    return length
}

func (self *Illuminance) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Illuminance) Go_getType() (string) { return "sensor_msgs/Illuminance" }
func (self *Illuminance) Go_getMD5() (string) { return "08ab9e71fcfbed30d5e337886c3f07f2" }
func (self *Illuminance) Go_getID() (uint32) { return 0 }
func (self *Illuminance) Go_setID(id uint32) { }

