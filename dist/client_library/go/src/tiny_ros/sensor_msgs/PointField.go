package sensor_msgs

import (
    "encoding/json"
)

func Go_INT8() (uint8) { return  1 }
func Go_UINT8() (uint8) { return  2 }
func Go_INT16() (uint8) { return  3 }
func Go_UINT16() (uint8) { return  4 }
func Go_INT32() (uint8) { return  5 }
func Go_UINT32() (uint8) { return  6 }
func Go_FLOAT32() (uint8) { return  7 }
func Go_FLOAT64() (uint8) { return  8 }

type PointField struct {
    Go_name string `json:"name"`
    Go_offset uint32 `json:"offset"`
    Go_datatype uint8 `json:"datatype"`
    Go_count uint32 `json:"count"`
}

func NewPointField() (*PointField) {
    newPointField := new(PointField)
    newPointField.Go_name = ""
    newPointField.Go_offset = 0
    newPointField.Go_datatype = 0
    newPointField.Go_count = 0
    return newPointField
}

func (self *PointField) Go_initialize() {
    self.Go_name = ""
    self.Go_offset = 0
    self.Go_datatype = 0
    self.Go_count = 0
}

func (self *PointField) Go_serialize(buff []byte) (int) {
    offset := 0
    length_name := len(self.Go_name)
    buff[offset + 0] = byte((length_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_name)], self.Go_name)
    offset += length_name
    buff[offset + 0] = byte((self.Go_offset >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_offset >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_offset >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_offset >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_datatype >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_count >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_count >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_count >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_count >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *PointField) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_name = string(buff[offset:(offset+length_name)])
    offset += length_name
    self.Go_offset = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_offset |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_offset |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_offset |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_datatype = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_count = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_count |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_count |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_count |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *PointField) Go_serializedLength() (int) {
    length := 0
    length_name := len(self.Go_name)
    length += 4
    length += length_name
    length += 4
    length += 1
    length += 4
    return length
}

func (self *PointField) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PointField) Go_getType() (string) { return "sensor_msgs/PointField" }
func (self *PointField) Go_getMD5() (string) { return "039974f05fdf0470d9dc865fd01dcc3e" }
func (self *PointField) Go_getID() (uint32) { return 0 }
func (self *PointField) Go_setID(id uint32) { }

