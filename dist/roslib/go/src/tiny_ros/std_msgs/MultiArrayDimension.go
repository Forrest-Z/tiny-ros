package std_msgs

import (
    "encoding/json"
)


type MultiArrayDimension struct {
    Go_label string `json:"label"`
    Go_size uint32 `json:"size"`
    Go_stride uint32 `json:"stride"`
}

func NewMultiArrayDimension() (*MultiArrayDimension) {
    newMultiArrayDimension := new(MultiArrayDimension)
    newMultiArrayDimension.Go_label = ""
    newMultiArrayDimension.Go_size = 0
    newMultiArrayDimension.Go_stride = 0
    return newMultiArrayDimension
}

func (self *MultiArrayDimension) Go_initialize() {
    self.Go_label = ""
    self.Go_size = 0
    self.Go_stride = 0
}

func (self *MultiArrayDimension) Go_serialize(buff []byte) (int) {
    offset := 0
    length_label := len(self.Go_label)
    buff[offset + 0] = byte((length_label >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_label >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_label >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_label >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_label)], self.Go_label)
    offset += length_label
    buff[offset + 0] = byte((self.Go_size >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_size >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_size >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_size >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stride >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stride >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stride >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stride >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *MultiArrayDimension) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_label := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_label |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_label |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_label |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_label = string(buff[offset:(offset+length_label)])
    offset += length_label
    self.Go_size = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_size |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_size |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_size |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stride = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stride |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stride |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stride |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *MultiArrayDimension) Go_serializedLength() (int) {
    length := 0
    length_label := len(self.Go_label)
    length += 4
    length += length_label
    length += 4
    length += 4
    return length
}

func (self *MultiArrayDimension) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MultiArrayDimension) Go_getType() (string) { return "std_msgs/MultiArrayDimension" }
func (self *MultiArrayDimension) Go_getMD5() (string) { return "c2aacf83d49c7aa4a8504bd32158e990" }
func (self *MultiArrayDimension) Go_getID() (uint32) { return 0 }
func (self *MultiArrayDimension) Go_setID(id uint32) { }

