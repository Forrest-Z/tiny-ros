package std_msgs

import (
    "encoding/json"
)


type UInt8MultiArray struct {
    Go_layout *MultiArrayLayout `json:"layout"`
    Go_data []uint8 `json:"data"`
}

func NewUInt8MultiArray() (*UInt8MultiArray) {
    newUInt8MultiArray := new(UInt8MultiArray)
    newUInt8MultiArray.Go_layout = NewMultiArrayLayout()
    newUInt8MultiArray.Go_data = []uint8{}
    return newUInt8MultiArray
}

func (self *UInt8MultiArray) Go_initialize() {
    self.Go_layout = NewMultiArrayLayout()
    self.Go_data = []uint8{}
}

func (self *UInt8MultiArray) Go_serialize(buff []byte) (int) {
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
        offset += 1
    }
    return offset
}

func (self *UInt8MultiArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_layout.Go_deserialize(buff[offset:])
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]uint8, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
        offset += 1
    }
    return offset
}

func (self *UInt8MultiArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_layout.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    return length
}

func (self *UInt8MultiArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *UInt8MultiArray) Go_getType() (string) { return "std_msgs/UInt8MultiArray" }
func (self *UInt8MultiArray) Go_getMD5() (string) { return "ed0dcee956eb67c464fc17fa893003bb" }
func (self *UInt8MultiArray) Go_getID() (uint32) { return 0 }
func (self *UInt8MultiArray) Go_setID(id uint32) { }

