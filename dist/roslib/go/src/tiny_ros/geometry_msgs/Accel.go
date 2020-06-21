package geometry_msgs

import (
    "geometry_msgs/Vector3"
)

type Accel struct {
    Go_linear geometry_msgs.Vector3 `json:"linear"`
    Go_angular geometry_msgs.Vector3 `json:"angular"`
}

func NewAccel() (*Accel) {
    newAccel := new(Accel)
    newAccel.Go_linear = geometry_msgs.NewVector3()
    newAccel.Go_angular = geometry_msgs.NewVector3()
    return newAccel
}

func (self *Accel) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_linear.Go_serialize(buff[offset:])
    offset += self.Go_angular.Go_serialize(buff[offset:])
    return offset
}

func (self *Accel) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_linear.Go_deserialize(buff[offset:])
    offset += self.Go_angular.Go_deserialize(buff[offset:])
    return offset
}

func (self *Accel) Go_serializedLength() (int) {
    length := 0
    length += self.Go_linear.Go_serializedLength()
    length += self.Go_angular.Go_serializedLength()
    return length
}

func (self *Accel) Go_echo() (string) { return "" }
func (self *Accel) Go_getType() (string) { return "geometry_msgs/Accel" }
func (self *Accel) Go_getMD5() (string) { return "580cbad5f3bd2e9f0ca71e14b7ab1b0f" }
func (self *Accel) Go_getID() (uint32) { return 0 }
func (self *Accel) Go_setID(id uint32) { }

