package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type AccelStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_accel *Accel `json:"accel"`
}

func NewAccelStamped() (*AccelStamped) {
    newAccelStamped := new(AccelStamped)
    newAccelStamped.Go_header = std_msgs.NewHeader()
    newAccelStamped.Go_accel = NewAccel()
    return newAccelStamped
}

func (self *AccelStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_accel = NewAccel()
}

func (self *AccelStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_accel.Go_serialize(buff[offset:])
    return offset
}

func (self *AccelStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_accel.Go_deserialize(buff[offset:])
    return offset
}

func (self *AccelStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_accel.Go_serializedLength()
    return length
}

func (self *AccelStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *AccelStamped) Go_getType() (string) { return "geometry_msgs/AccelStamped" }
func (self *AccelStamped) Go_getMD5() (string) { return "fa35432963826361a1073b1df905a559" }
func (self *AccelStamped) Go_getID() (uint32) { return 0 }
func (self *AccelStamped) Go_setID(id uint32) { }

