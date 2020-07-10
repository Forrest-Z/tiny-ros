package std_msgs

import (
    "encoding/json"
)


type Int32MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []int32 `json:"data"`
}

func NewInt32MultiArray() (*Int32MultiArray) {
    newInt32MultiArray := new(Int32MultiArray)
    newInt32MultiArray.Go_layout = NewMultiArrayLayout()
    newInt32MultiArray.Go_data = []int32{}
    return newInt32MultiArray
}

func (self *Int32MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []int32{}
}

func (self *Int32MultiArray) Go_serialize(buff []byte) (int) {
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
        offset += 4
    }
    return offset
}

func (self *Int32MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]int32, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = int32(buff[offset + 0] & 0xFF) << (8 * 0)
        self.Go_data[i] |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
        self.Go_data[i] |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
        self.Go_data[i] |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
    }
    return offset
}

func (self *Int32MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 4
    }
    return length
}

func (self *Int32MultiArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Int32MultiArray) Go_getType() (string) { return "std_msgs/Int32MultiArray" }
func (self *Int32MultiArray) Go_getMD5() (string) { return "d09b4e71ff2be4c9c957c4e4d889034d" }
func (self *Int32MultiArray) Go_getID() (uint32) { return 0 }
func (self *Int32MultiArray) Go_setID(id uint32) { }

