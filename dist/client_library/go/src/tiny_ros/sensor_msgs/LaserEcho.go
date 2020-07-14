package sensor_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type LaserEcho struct {
    Go_echoes []float32 `json:"echoes"`
}

func NewLaserEcho() (*LaserEcho) {
    newLaserEcho := new(LaserEcho)
    newLaserEcho.Go_echoes = []float32{}
    return newLaserEcho
}

func (self *LaserEcho) Go_initialize() {
    self.Go_echoes = []float32{}
}

func (self *LaserEcho) Go_serialize(buff []byte) (int) {
    offset := 0
    length_echoes := len(self.Go_echoes)
    buff[offset + 0] = byte((length_echoes >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_echoes >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_echoes >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_echoes >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_echoes; i++ {
        bits_echoesi := math.Float32bits(self.Go_echoes[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_echoesi)
        offset += 4
    }
    return offset
}

func (self *LaserEcho) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_echoes := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_echoes |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_echoes |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_echoes |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_echoes = make([]float32, length_echoes)
    for i := 0; i < length_echoes; i++ {
        bits_echoesi := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_echoes[i] = math.Float32frombits(bits_echoesi)
        offset += 4
    }
    return offset
}

func (self *LaserEcho) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_echoes := len(self.Go_echoes)
    for i := 0; i < length_echoes; i++ {
        length += 4
    }
    return length
}

func (self *LaserEcho) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LaserEcho) Go_getType() (string) { return "sensor_msgs/LaserEcho" }
func (self *LaserEcho) Go_getMD5() (string) { return "a8537b388573845b3240b44db5bc4905" }
func (self *LaserEcho) Go_getID() (uint32) { return 0 }
func (self *LaserEcho) Go_setID(id uint32) { }

