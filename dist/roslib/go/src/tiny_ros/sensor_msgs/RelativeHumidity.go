package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type RelativeHumidity struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_relative_humidity float64 `json:"relative_humidity"`
    Go_variance float64 `json:"variance"`
}

func NewRelativeHumidity() (*RelativeHumidity) {
    newRelativeHumidity := new(RelativeHumidity)
    newRelativeHumidity.Go_header = std_msgs.NewHeader()
    newRelativeHumidity.Go_relative_humidity = 0.0
    newRelativeHumidity.Go_variance = 0.0
    return newRelativeHumidity
}

func (self *RelativeHumidity) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_relative_humidity = 0.0
    self.Go_variance = 0.0
}

func (self *RelativeHumidity) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_relative_humidity := math.Float64bits(self.Go_relative_humidity)
    binary.LittleEndian.PutUint64(buff[offset:], bits_relative_humidity)
    offset += 8
    bits_variance := math.Float64bits(self.Go_variance)
    binary.LittleEndian.PutUint64(buff[offset:], bits_variance)
    offset += 8
    return offset
}

func (self *RelativeHumidity) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_relative_humidity := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_relative_humidity = math.Float64frombits(bits_relative_humidity)
    offset += 8
    bits_variance := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_variance = math.Float64frombits(bits_variance)
    offset += 8
    return offset
}

func (self *RelativeHumidity) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 8
    length += 8
    return length
}

func (self *RelativeHumidity) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *RelativeHumidity) Go_getType() (string) { return "sensor_msgs/RelativeHumidity" }
func (self *RelativeHumidity) Go_getMD5() (string) { return "d9a3a4b2c3c0c55eede767d38b460110" }
func (self *RelativeHumidity) Go_getID() (uint32) { return 0 }
func (self *RelativeHumidity) Go_setID(id uint32) { }

