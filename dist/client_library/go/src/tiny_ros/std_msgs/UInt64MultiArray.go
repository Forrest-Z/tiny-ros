package std_msgs

import (
    "encoding/json"
)


type UInt64MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []uint64 `json:"data"`
}

func NewUInt64MultiArray() (*UInt64MultiArray) {
    newUInt64MultiArray := new(UInt64MultiArray)
    newUInt64MultiArray.Go_layout = NewMultiArrayLayout()
    newUInt64MultiArray.Go_data = []uint64{}
    return newUInt64MultiArray
}

func (self *UInt64MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []uint64{}
}

func (self *UInt64MultiArray) Go_serialize(buff []byte) (int) {
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

func (self *UInt64MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]uint64, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = uint64(buff[offset + 0] & 0xFF) << (8 * 0)
        self.Go_data[i] |= uint64(buff[offset + 1] & 0xFF) << (8 * 1)
        self.Go_data[i] |= uint64(buff[offset + 2] & 0xFF) << (8 * 2)
        self.Go_data[i] |= uint64(buff[offset + 3] & 0xFF) << (8 * 3)
        self.Go_data[i] |= uint64(buff[offset + 4] & 0xFF) << (8 * 4)
        self.Go_data[i] |= uint64(buff[offset + 5] & 0xFF) << (8 * 5)
        self.Go_data[i] |= uint64(buff[offset + 6] & 0xFF) << (8 * 6)
        self.Go_data[i] |= uint64(buff[offset + 7] & 0xFF) << (8 * 7)
        offset += 8
    }
    return offset
}

func (self *UInt64MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 8
    }
    return length
}

func (self *UInt64MultiArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt64MultiArray) Go_getType() (string) { return "std_msgs/UInt64MultiArray" }
func (self *UInt64MultiArray) Go_getMD5() (string) { return "a0fe08f13f702ecfc59c58150f66678a" }
func (self *UInt64MultiArray) Go_getID() (uint32) { return 0 }
func (self *UInt64MultiArray) Go_setID(id uint32) { }

