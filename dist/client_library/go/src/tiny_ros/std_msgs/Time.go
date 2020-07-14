package std_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type Time struct {
    Go_data *rostime.Time `json:"data"`
}

func NewTime() (*Time) {
    newTime := new(Time)
    newTime.Go_data = rostime.NewTime()
    return newTime
}

func (self *Time) Go_initialize() {
    self.Go_data = rostime.NewTime()
}

func (self *Time) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_data.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_data.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *Time) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_data.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *Time) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    return length
}

func (self *Time) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Time) Go_getType() (string) { return "std_msgs/Time" }
func (self *Time) Go_getMD5() (string) { return "64602ed67393e1e61260ab68d6fa2045" }
func (self *Time) Go_getID() (uint32) { return 0 }
func (self *Time) Go_setID(id uint32) { }

