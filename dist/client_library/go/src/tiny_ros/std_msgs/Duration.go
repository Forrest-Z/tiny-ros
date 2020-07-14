package std_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type Duration struct {
    Go_data *rostime.Duration `json:"data"`
}

func NewDuration() (*Duration) {
    newDuration := new(Duration)
    newDuration.Go_data = rostime.NewDuration()
    return newDuration
}

func (self *Duration) Go_initialize() {
    self.Go_data = rostime.NewDuration()
}

func (self *Duration) Go_serialize(buff []byte) (int) {
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

func (self *Duration) Go_deserialize(buff []byte) (int) {
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

func (self *Duration) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    return length
}

func (self *Duration) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Duration) Go_getType() (string) { return "std_msgs/Duration" }
func (self *Duration) Go_getMD5() (string) { return "5dc598a1f3dfbcf0b7343c41b15d6ab2" }
func (self *Duration) Go_getID() (uint32) { return 0 }
func (self *Duration) Go_setID(id uint32) { }

