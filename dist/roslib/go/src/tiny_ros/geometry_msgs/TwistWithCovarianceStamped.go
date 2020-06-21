package geometry_msgs

import (
    "tiny_ros/std_msgs"
    "geometry_msgs/TwistWithCovariance"
)

type TwistWithCovarianceStamped struct {
    Go_header std_msgs.Header `json:"header"`
    Go_twist geometry_msgs.TwistWithCovariance `json:"twist"`
}

func NewTwistWithCovarianceStamped() (*TwistWithCovarianceStamped) {
    newTwistWithCovarianceStamped := new(TwistWithCovarianceStamped)
    newTwistWithCovarianceStamped.Go_header = std_msgs.NewHeader()
    newTwistWithCovarianceStamped.Go_twist = geometry_msgs.NewTwistWithCovariance()
    return newTwistWithCovarianceStamped
}

func (self *TwistWithCovarianceStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_twist.Go_serialize(buff[offset:])
    return offset
}

func (self *TwistWithCovarianceStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_twist.Go_deserialize(buff[offset:])
    return offset
}

func (self *TwistWithCovarianceStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_twist.Go_serializedLength()
    return length
}

func (self *TwistWithCovarianceStamped) Go_echo() (string) { return "" }
func (self *TwistWithCovarianceStamped) Go_getType() (string) { return "geometry_msgs/TwistWithCovarianceStamped" }
func (self *TwistWithCovarianceStamped) Go_getMD5() (string) { return "2cbcab62cac39de1d1d01785b99ba778" }
func (self *TwistWithCovarianceStamped) Go_getID() (uint32) { return 0 }
func (self *TwistWithCovarianceStamped) Go_setID(id uint32) { }

