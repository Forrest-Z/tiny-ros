package geometry_msgs

import (
    "tiny_ros/std_msgs"
    "geometry_msgs/Inertia"
)

type InertiaStamped struct {
    Go_header std_msgs.Header `json:"header"`
    Go_inertia geometry_msgs.Inertia `json:"inertia"`
}

func NewInertiaStamped() (*InertiaStamped) {
    newInertiaStamped := new(InertiaStamped)
    newInertiaStamped.Go_header = std_msgs.NewHeader()
    newInertiaStamped.Go_inertia = geometry_msgs.NewInertia()
    return newInertiaStamped
}

func (self *InertiaStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_inertia.Go_serialize(buff[offset:])
    return offset
}

func (self *InertiaStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_inertia.Go_deserialize(buff[offset:])
    return offset
}

func (self *InertiaStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_inertia.Go_serializedLength()
    return length
}

func (self *InertiaStamped) Go_echo() (string) { return "" }
func (self *InertiaStamped) Go_getType() (string) { return "geometry_msgs/InertiaStamped" }
func (self *InertiaStamped) Go_getMD5() (string) { return "2b3c9b263c59f65da44508cd041d18a0" }
func (self *InertiaStamped) Go_getID() (uint32) { return 0 }
func (self *InertiaStamped) Go_setID(id uint32) { }

