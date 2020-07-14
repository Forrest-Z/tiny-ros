package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type TwistStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_twist *Twist `json:"twist"`
}

func NewTwistStamped() (*TwistStamped) {
    newTwistStamped := new(TwistStamped)
    newTwistStamped.Go_header = std_msgs.NewHeader()
    newTwistStamped.Go_twist = NewTwist()
    return newTwistStamped
}

func (self *TwistStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_twist = NewTwist()
}

func (self *TwistStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_twist.Go_serialize(buff[offset:])
    return offset
}

func (self *TwistStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_twist.Go_deserialize(buff[offset:])
    return offset
}

func (self *TwistStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_twist.Go_serializedLength()
    return length
}

func (self *TwistStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TwistStamped) Go_getType() (string) { return "geometry_msgs/TwistStamped" }
func (self *TwistStamped) Go_getMD5() (string) { return "2e3e0a57a69306091cb5c65e92d048e1" }
func (self *TwistStamped) Go_getID() (uint32) { return 0 }
func (self *TwistStamped) Go_setID(id uint32) { }

