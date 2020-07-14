package rosgraph_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type Clock struct {
    Go_clock *rostime.Time `json:"clock"`
}

func NewClock() (*Clock) {
    newClock := new(Clock)
    newClock.Go_clock = rostime.NewTime()
    return newClock
}

func (self *Clock) Go_initialize() {
    self.Go_clock = rostime.NewTime()
}

func (self *Clock) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_clock.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_clock.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_clock.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_clock.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_clock.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_clock.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_clock.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_clock.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *Clock) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_clock.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_clock.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_clock.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_clock.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_clock.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_clock.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_clock.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_clock.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *Clock) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    return length
}

func (self *Clock) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Clock) Go_getType() (string) { return "rosgraph_msgs/Clock" }
func (self *Clock) Go_getMD5() (string) { return "d3bedbe03b904b8181e3fef4bbe0a73e" }
func (self *Clock) Go_getID() (uint32) { return 0 }
func (self *Clock) Go_setID(id uint32) { }

