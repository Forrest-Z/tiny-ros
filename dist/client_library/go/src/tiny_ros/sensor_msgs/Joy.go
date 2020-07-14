package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type Joy struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_axes []float32 `json:"axes"`
    Go_buttons []int32 `json:"buttons"`
}

func NewJoy() (*Joy) {
    newJoy := new(Joy)
    newJoy.Go_header = std_msgs.NewHeader()
    newJoy.Go_axes = []float32{}
    newJoy.Go_buttons = []int32{}
    return newJoy
}

func (self *Joy) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_axes = []float32{}
    self.Go_buttons = []int32{}
}

func (self *Joy) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_axes := len(self.Go_axes)
    buff[offset + 0] = byte((length_axes >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_axes >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_axes >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_axes >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_axes; i++ {
        bits_axesi := math.Float32bits(self.Go_axes[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_axesi)
        offset += 4
    }
    length_buttons := len(self.Go_buttons)
    buff[offset + 0] = byte((length_buttons >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_buttons >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_buttons >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_buttons >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_buttons; i++ {
        buff[offset + 0] = byte((self.Go_buttons[i] >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((self.Go_buttons[i] >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((self.Go_buttons[i] >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((self.Go_buttons[i] >> (8 * 3)) & 0xFF)
        offset += 4
    }
    return offset
}

func (self *Joy) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_axes := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_axes |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_axes |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_axes |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_axes = make([]float32, length_axes)
    for i := 0; i < length_axes; i++ {
        bits_axesi := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_axes[i] = math.Float32frombits(bits_axesi)
        offset += 4
    }
    length_buttons := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_buttons |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_buttons |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_buttons |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_buttons = make([]int32, length_buttons)
    for i := 0; i < length_buttons; i++ {
        self.Go_buttons[i] = int32(buff[offset + 0] & 0xFF) << (8 * 0)
        self.Go_buttons[i] |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
        self.Go_buttons[i] |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
        self.Go_buttons[i] |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
    }
    return offset
}

func (self *Joy) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_axes := len(self.Go_axes)
    for i := 0; i < length_axes; i++ {
        length += 4
    }
    length += 4
    length_buttons := len(self.Go_buttons)
    for i := 0; i < length_buttons; i++ {
        length += 4
    }
    return length
}

func (self *Joy) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Joy) Go_getType() (string) { return "sensor_msgs/Joy" }
func (self *Joy) Go_getMD5() (string) { return "da3438323dbe92a4d6e4658e06bf8da1" }
func (self *Joy) Go_getID() (uint32) { return 0 }
func (self *Joy) Go_setID(id uint32) { }

