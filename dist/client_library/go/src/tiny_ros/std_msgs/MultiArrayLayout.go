package std_msgs

import (
    "encoding/json"
)


type MultiArrayLayout struct {
    Go_dim []*MultiArrayDimension `json:"dim"`
    Go_data_offset uint32 `json:"data_offset"`
}

func NewMultiArrayLayout() (*MultiArrayLayout) {
    newMultiArrayLayout := new(MultiArrayLayout)
    newMultiArrayLayout.Go_dim = []*MultiArrayDimension{}
    newMultiArrayLayout.Go_data_offset = 0
    return newMultiArrayLayout
}

func (self *MultiArrayLayout) Go_initialize() {
    self.Go_dim = []*MultiArrayDimension{}
    self.Go_data_offset = 0
}

func (self *MultiArrayLayout) Go_serialize(buff []byte) (int) {
    offset := 0
    length_dim := len(self.Go_dim)
    buff[offset + 0] = byte((length_dim >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_dim >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_dim >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_dim >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_dim; i++ {
        offset += self.Go_dim[i].Go_serialize(buff[offset:])
    }
    buff[offset + 0] = byte((self.Go_data_offset >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data_offset >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data_offset >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data_offset >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *MultiArrayLayout) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_dim := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_dim |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_dim |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_dim |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_dim = make([]*MultiArrayDimension, length_dim)
    for i := 0; i < length_dim; i++ {
        self.Go_dim[i] = NewMultiArrayDimension()
    }
    for i := 0; i < length_dim; i++ {
        offset += self.Go_dim[i].Go_deserialize(buff[offset:])
    }
    self.Go_data_offset = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data_offset |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data_offset |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data_offset |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *MultiArrayLayout) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_dim := len(self.Go_dim)
    for i := 0; i < length_dim; i++ {
        length += self.Go_dim[i].Go_serializedLength()
    }
    length += 4
    return length
}

func (self *MultiArrayLayout) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MultiArrayLayout) Go_getType() (string) { return "std_msgs/MultiArrayLayout" }
func (self *MultiArrayLayout) Go_getMD5() (string) { return "f40f0b5b285a93ca167c98c1012a989a" }
func (self *MultiArrayLayout) Go_getID() (uint32) { return 0 }
func (self *MultiArrayLayout) Go_setID(id uint32) { }

