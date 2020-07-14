package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type WrenchStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_wrench *Wrench `json:"wrench"`
}

func NewWrenchStamped() (*WrenchStamped) {
    newWrenchStamped := new(WrenchStamped)
    newWrenchStamped.Go_header = std_msgs.NewHeader()
    newWrenchStamped.Go_wrench = NewWrench()
    return newWrenchStamped
}

func (self *WrenchStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_wrench = NewWrench()
}

func (self *WrenchStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_wrench.Go_serialize(buff[offset:])
    return offset
}

func (self *WrenchStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_wrench.Go_deserialize(buff[offset:])
    return offset
}

func (self *WrenchStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_wrench.Go_serializedLength()
    return length
}

func (self *WrenchStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *WrenchStamped) Go_getType() (string) { return "geometry_msgs/WrenchStamped" }
func (self *WrenchStamped) Go_getMD5() (string) { return "cf53874aa63609de4155ec8e9cf2c540" }
func (self *WrenchStamped) Go_getID() (uint32) { return 0 }
func (self *WrenchStamped) Go_setID(id uint32) { }

