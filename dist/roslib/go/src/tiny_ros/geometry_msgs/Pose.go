package geometry_msgs

import (
    "geometry_msgs/Point"
    "geometry_msgs/Quaternion"
)

type Pose struct {
    Go_position geometry_msgs.Point `json:"position"`
    Go_orientation geometry_msgs.Quaternion `json:"orientation"`
}

func NewPose() (*Pose) {
    newPose := new(Pose)
    newPose.Go_position = geometry_msgs.NewPoint()
    newPose.Go_orientation = geometry_msgs.NewQuaternion()
    return newPose
}

func (self *Pose) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_position.Go_serialize(buff[offset:])
    offset += self.Go_orientation.Go_serialize(buff[offset:])
    return offset
}

func (self *Pose) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_position.Go_deserialize(buff[offset:])
    offset += self.Go_orientation.Go_deserialize(buff[offset:])
    return offset
}

func (self *Pose) Go_serializedLength() (int) {
    length := 0
    length += self.Go_position.Go_serializedLength()
    length += self.Go_orientation.Go_serializedLength()
    return length
}

func (self *Pose) Go_echo() (string) { return "" }
func (self *Pose) Go_getType() (string) { return "geometry_msgs/Pose" }
func (self *Pose) Go_getMD5() (string) { return "0b42fb88be8cac0efa6e446e13befcae" }
func (self *Pose) Go_getID() (uint32) { return 0 }
func (self *Pose) Go_setID(id uint32) { }

