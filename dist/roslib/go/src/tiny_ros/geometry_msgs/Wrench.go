package geometry_msgs

import (
    "encoding/json"
)


type Wrench struct {
    Go_force *Vector3 `json:"force"`
    Go_torque *Vector3 `json:"torque"`
}

func NewWrench() (*Wrench) {
    newWrench := new(Wrench)
    newWrench.Go_force = NewVector3()
    newWrench.Go_torque = NewVector3()
    return newWrench
}

func (self *Wrench) Go_initialize() {
    self.Go_force = NewVector3()
    self.Go_torque = NewVector3()
}

func (self *Wrench) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_force.Go_serialize(buff[offset:])
    offset += self.Go_torque.Go_serialize(buff[offset:])
    return offset
}

func (self *Wrench) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_force.Go_deserialize(buff[offset:])
    offset += self.Go_torque.Go_deserialize(buff[offset:])
    return offset
}

func (self *Wrench) Go_serializedLength() (int) {
    length := 0
    length += self.Go_force.Go_serializedLength()
    length += self.Go_torque.Go_serializedLength()
    return length
}

func (self *Wrench) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Wrench) Go_getType() (string) { return "geometry_msgs/Wrench" }
func (self *Wrench) Go_getMD5() (string) { return "02d01d4a8dc253c7b42d4c9866201aee" }
func (self *Wrench) Go_getID() (uint32) { return 0 }
func (self *Wrench) Go_setID(id uint32) { }

