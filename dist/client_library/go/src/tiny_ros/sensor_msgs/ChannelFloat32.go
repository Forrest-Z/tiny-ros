package sensor_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type ChannelFloat32 struct {
    Go_name string `json:"name"`
    Go_values []float32 `json:"values"`
}

func NewChannelFloat32() (*ChannelFloat32) {
    newChannelFloat32 := new(ChannelFloat32)
    newChannelFloat32.Go_name = ""
    newChannelFloat32.Go_values = []float32{}
    return newChannelFloat32
}

func (self *ChannelFloat32) Go_initialize() {
    self.Go_name = ""
    self.Go_values = []float32{}
}

func (self *ChannelFloat32) Go_serialize(buff []byte) (int) {
    offset := 0
    length_name := len(self.Go_name)
    buff[offset + 0] = byte((length_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_name)], self.Go_name)
    offset += length_name
    length_values := len(self.Go_values)
    buff[offset + 0] = byte((length_values >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_values >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_values >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_values >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_values; i++ {
        bits_valuesi := math.Float32bits(self.Go_values[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_valuesi)
        offset += 4
    }
    return offset
}

func (self *ChannelFloat32) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_name = string(buff[offset:(offset+length_name)])
    offset += length_name
    length_values := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_values |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_values |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_values |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_values = make([]float32, length_values)
    for i := 0; i < length_values; i++ {
        bits_valuesi := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_values[i] = math.Float32frombits(bits_valuesi)
        offset += 4
    }
    return offset
}

func (self *ChannelFloat32) Go_serializedLength() (int) {
    length := 0
    length_name := len(self.Go_name)
    length += 4
    length += length_name
    length += 4
    length_values := len(self.Go_values)
    for i := 0; i < length_values; i++ {
        length += 4
    }
    return length
}

func (self *ChannelFloat32) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ChannelFloat32) Go_getType() (string) { return "sensor_msgs/ChannelFloat32" }
func (self *ChannelFloat32) Go_getMD5() (string) { return "c4cf01c81334c609dca1afd3a227daff" }
func (self *ChannelFloat32) Go_getID() (uint32) { return 0 }
func (self *ChannelFloat32) Go_setID(id uint32) { }

