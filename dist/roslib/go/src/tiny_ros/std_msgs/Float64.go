package std_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Float64 struct {
    Go_data float64 `json:"data"`
}

func NewFloat64() (*Float64) {
    newFloat64 := new(Float64)
    newFloat64.Go_data = 0.0
    return newFloat64
}

func (self *Float64) Go_initialize() {
    self.Go_data = 0.0
}

func (self *Float64) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_data := math.Float64bits(self.Go_data)
    binary.LittleEndian.PutUint64(buff[offset:], bits_data)
    offset += 8
    return offset
}

func (self *Float64) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_data := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_data = math.Float64frombits(bits_data)
    offset += 8
    return offset
}

func (self *Float64) Go_serializedLength() (int) {
    length := 0
    length += 8
    return length
}

func (self *Float64) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Float64) Go_getType() (string) { return "std_msgs/Float64" }
func (self *Float64) Go_getMD5() (string) { return "3439fe880debfd63cf4e61e62e285345" }
func (self *Float64) Go_getID() (uint32) { return 0 }
func (self *Float64) Go_setID(id uint32) { }

