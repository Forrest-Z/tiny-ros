package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type AccelWithCovarianceStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_accel *AccelWithCovariance `json:"accel"`
}

func NewAccelWithCovarianceStamped() (*AccelWithCovarianceStamped) {
    newAccelWithCovarianceStamped := new(AccelWithCovarianceStamped)
    newAccelWithCovarianceStamped.Go_header = std_msgs.NewHeader()
    newAccelWithCovarianceStamped.Go_accel = NewAccelWithCovariance()
    return newAccelWithCovarianceStamped
}

func (self *AccelWithCovarianceStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_accel = NewAccelWithCovariance()
}

func (self *AccelWithCovarianceStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_accel.Go_serialize(buff[offset:])
    return offset
}

func (self *AccelWithCovarianceStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_accel.Go_deserialize(buff[offset:])
    return offset
}

func (self *AccelWithCovarianceStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_accel.Go_serializedLength()
    return length
}

func (self *AccelWithCovarianceStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *AccelWithCovarianceStamped) Go_getType() (string) { return "geometry_msgs/AccelWithCovarianceStamped" }
func (self *AccelWithCovarianceStamped) Go_getMD5() (string) { return "efd9e7d0b5ca262cc8b05aa8e97c984f" }
func (self *AccelWithCovarianceStamped) Go_getID() (uint32) { return 0 }
func (self *AccelWithCovarianceStamped) Go_setID(id uint32) { }

