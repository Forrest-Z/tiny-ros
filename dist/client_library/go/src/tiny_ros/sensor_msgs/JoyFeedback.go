package sensor_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)

func Go_TYPE_LED() (uint8) { return  0 }
func Go_TYPE_RUMBLE() (uint8) { return  1 }
func Go_TYPE_BUZZER() (uint8) { return  2 }

type JoyFeedback struct {
    Go_type uint8 `json:"type"`
    Go_id uint8 `json:"id"`
    Go_intensity float32 `json:"intensity"`
}

func NewJoyFeedback() (*JoyFeedback) {
    newJoyFeedback := new(JoyFeedback)
    newJoyFeedback.Go_type = 0
    newJoyFeedback.Go_id = 0
    newJoyFeedback.Go_intensity = 0.0
    return newJoyFeedback
}

func (self *JoyFeedback) Go_initialize() {
    self.Go_type = 0
    self.Go_id = 0
    self.Go_intensity = 0.0
}

func (self *JoyFeedback) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_type >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_id >> (8 * 0)) & 0xFF)
    offset += 1
    bits_intensity := math.Float32bits(self.Go_intensity)
    binary.LittleEndian.PutUint32(buff[offset:], bits_intensity)
    offset += 4
    return offset
}

func (self *JoyFeedback) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_type = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_id = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    bits_intensity := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_intensity = math.Float32frombits(bits_intensity)
    offset += 4
    return offset
}

func (self *JoyFeedback) Go_serializedLength() (int) {
    length := 0
    length += 1
    length += 1
    length += 4
    return length
}

func (self *JoyFeedback) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JoyFeedback) Go_getType() (string) { return "sensor_msgs/JoyFeedback" }
func (self *JoyFeedback) Go_getMD5() (string) { return "206b65e86c8b195f816ccbe40b3568a2" }
func (self *JoyFeedback) Go_getID() (uint32) { return 0 }
func (self *JoyFeedback) Go_setID(id uint32) { }

