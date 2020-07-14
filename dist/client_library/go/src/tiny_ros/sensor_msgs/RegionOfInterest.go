package sensor_msgs

import (
    "encoding/json"
)


type RegionOfInterest struct {
    Go_x_offset uint32 `json:"x_offset"`
    Go_y_offset uint32 `json:"y_offset"`
    Go_height uint32 `json:"height"`
    Go_width uint32 `json:"width"`
    Go_do_rectify bool `json:"do_rectify"`
}

func NewRegionOfInterest() (*RegionOfInterest) {
    newRegionOfInterest := new(RegionOfInterest)
    newRegionOfInterest.Go_x_offset = 0
    newRegionOfInterest.Go_y_offset = 0
    newRegionOfInterest.Go_height = 0
    newRegionOfInterest.Go_width = 0
    newRegionOfInterest.Go_do_rectify = false
    return newRegionOfInterest
}

func (self *RegionOfInterest) Go_initialize() {
    self.Go_x_offset = 0
    self.Go_y_offset = 0
    self.Go_height = 0
    self.Go_width = 0
    self.Go_do_rectify = false
}

func (self *RegionOfInterest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_x_offset >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_x_offset >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_x_offset >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_x_offset >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_y_offset >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_y_offset >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_y_offset >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_y_offset >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_height >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_height >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_height >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_height >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_width >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_width >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_width >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_width >> (8 * 3)) & 0xFF)
    offset += 4
    if self.Go_do_rectify {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *RegionOfInterest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_x_offset = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_x_offset |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_x_offset |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_x_offset |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_y_offset = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_y_offset |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_y_offset |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_y_offset |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_height = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_height |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_height |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_height |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_width = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_width |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_width |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_width |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    if (buff[offset] & 0xFF) != 0 {
        self.Go_do_rectify = true
    } else {
        self.Go_do_rectify = false
    }
    offset += 1
    return offset
}

func (self *RegionOfInterest) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    length += 4
    length += 1
    return length
}

func (self *RegionOfInterest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *RegionOfInterest) Go_getType() (string) { return "sensor_msgs/RegionOfInterest" }
func (self *RegionOfInterest) Go_getMD5() (string) { return "8370dc286f915405c906299aef5bb442" }
func (self *RegionOfInterest) Go_getID() (uint32) { return 0 }
func (self *RegionOfInterest) Go_setID(id uint32) { }

