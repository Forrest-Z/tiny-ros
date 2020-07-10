package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)

func Go_ULTRASOUND() (uint8) { return 0 }
func Go_INFRARED() (uint8) { return 1 }

type Range struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_radiation_type uint8 `json:"radiation_type"`
    Go_field_of_view float32 `json:"field_of_view"`
    Go_min_range float32 `json:"min_range"`
    Go_max_range float32 `json:"max_range"`
    Go_range float32 `json:"range"`
}

func NewRange() (*Range) {
    newRange := new(Range)
    newRange.Go_header = std_msgs.NewHeader()
    newRange.Go_radiation_type = 0
    newRange.Go_field_of_view = 0.0
    newRange.Go_min_range = 0.0
    newRange.Go_max_range = 0.0
    newRange.Go_range = 0.0
    return newRange
}

func (self *Range) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_radiation_type = 0
    self.Go_field_of_view = 0.0
    self.Go_min_range = 0.0
    self.Go_max_range = 0.0
    self.Go_range = 0.0
}

func (self *Range) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_radiation_type >> (8 * 0)) & 0xFF)
    offset += 1
    bits_field_of_view := math.Float32bits(self.Go_field_of_view)
    binary.LittleEndian.PutUint32(buff[offset:], bits_field_of_view)
    offset += 4
    bits_min_range := math.Float32bits(self.Go_min_range)
    binary.LittleEndian.PutUint32(buff[offset:], bits_min_range)
    offset += 4
    bits_max_range := math.Float32bits(self.Go_max_range)
    binary.LittleEndian.PutUint32(buff[offset:], bits_max_range)
    offset += 4
    bits_range := math.Float32bits(self.Go_range)
    binary.LittleEndian.PutUint32(buff[offset:], bits_range)
    offset += 4
    return offset
}

func (self *Range) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_radiation_type = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    bits_field_of_view := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_field_of_view = math.Float32frombits(bits_field_of_view)
    offset += 4
    bits_min_range := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_min_range = math.Float32frombits(bits_min_range)
    offset += 4
    bits_max_range := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_max_range = math.Float32frombits(bits_max_range)
    offset += 4
    bits_range := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_range = math.Float32frombits(bits_range)
    offset += 4
    return offset
}

func (self *Range) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 1
    length += 4
    length += 4
    length += 4
    length += 4
    return length
}

func (self *Range) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Range) Go_getType() (string) { return "sensor_msgs/Range" }
func (self *Range) Go_getMD5() (string) { return "54d647e3a481f5b87ce39d1b97e84d53" }
func (self *Range) Go_getID() (uint32) { return 0 }
func (self *Range) Go_setID(id uint32) { }

