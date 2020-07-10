package std_msgs

import (
    "encoding/json"
)


type Int64MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []int64 `json:"data"`
}

func NewInt64MultiArray() (*Int64MultiArray) {
    newInt64MultiArray := new(Int64MultiArray)
    newInt64MultiArray.Go_layout = NewMultiArrayLayout()
    newInt64MultiArray.Go_data = []int64{}
    return newInt64MultiArray
}

func (self *Int64MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []int64{}
}

func (self *Int64MultiArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_serialize(buff[offset:])
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_data; i++ {
        buff[offset + 0] = byte((self.Go_data[i] >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((self.Go_data[i] >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((self.Go_data[i] >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((self.Go_data[i] >> (8 * 3)) & 0xFF)
        buff[offset + 4] = byte((self.Go_data[i] >> (8 * 4)) & 0xFF)
        buff[offset + 5] = byte((self.Go_data[i] >> (8 * 5)) & 0xFF)
        buff[offset + 6] = byte((self.Go_data[i] >> (8 * 6)) & 0xFF)
        buff[offset + 7] = byte((self.Go_data[i] >> (8 * 7)) & 0xFF)
        offset += 8
    }
    return offset
}

func (self *Int64MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]int64, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = int64(buff[offset + 0] & 0xFF) << (8 * 0)
        self.Go_data[i] |= int64(buff[offset + 1] & 0xFF) << (8 * 1)
        self.Go_data[i] |= int64(buff[offset + 2] & 0xFF) << (8 * 2)
        self.Go_data[i] |= int64(buff[offset + 3] & 0xFF) << (8 * 3)
        self.Go_data[i] |= int64(buff[offset + 4] & 0xFF) << (8 * 4)
        self.Go_data[i] |= int64(buff[offset + 5] & 0xFF) << (8 * 5)
        self.Go_data[i] |= int64(buff[offset + 6] & 0xFF) << (8 * 6)
        self.Go_data[i] |= int64(buff[offset + 7] & 0xFF) << (8 * 7)
        offset += 8
    }
    return offset
}

func (self *Int64MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 8
    }
    return length
}

func (self *Int64MultiArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Int64MultiArray) Go_getType() (string) { return "std_msgs/Int64MultiArray" }
func (self *Int64MultiArray) Go_getMD5() (string) { return "c06d166bcec6ac0d57a5122b314c9f5c" }
func (self *Int64MultiArray) Go_getID() (uint32) { return 0 }
func (self *Int64MultiArray) Go_setID(id uint32) { }

