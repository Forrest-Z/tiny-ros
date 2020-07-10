package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type QuaternionStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_quaternion *Quaternion `json:"quaternion"`
}

func NewQuaternionStamped() (*QuaternionStamped) {
    newQuaternionStamped := new(QuaternionStamped)
    newQuaternionStamped.Go_header = std_msgs.NewHeader()
    newQuaternionStamped.Go_quaternion = NewQuaternion()
    return newQuaternionStamped
}

func (self *QuaternionStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_quaternion = NewQuaternion()
}

func (self *QuaternionStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_quaternion.Go_serialize(buff[offset:])
    return offset
}

func (self *QuaternionStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_quaternion.Go_deserialize(buff[offset:])
    return offset
}

func (self *QuaternionStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_quaternion.Go_serializedLength()
    return length
}

func (self *QuaternionStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *QuaternionStamped) Go_getType() (string) { return "geometry_msgs/QuaternionStamped" }
func (self *QuaternionStamped) Go_getMD5() (string) { return "69e39922feb9ec6eaf93755f93fce2cf" }
func (self *QuaternionStamped) Go_getID() (uint32) { return 0 }
func (self *QuaternionStamped) Go_setID(id uint32) { }

