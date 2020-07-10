package geometry_msgs

import (
    "encoding/json"
)


type Twist struct {
    Go_linear *Vector3 `json:"linear"`
    Go_angular *Vector3 `json:"angular"`
}

func NewTwist() (*Twist) {
    newTwist := new(Twist)
    newTwist.Go_linear = NewVector3()
    newTwist.Go_angular = NewVector3()
    return newTwist
}

func (self *Twist) Go_initialize() {
    self.Go_linear = NewVector3()
    self.Go_angular = NewVector3()
}

func (self *Twist) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_linear.Go_serialize(buff[offset:])
    offset += self.Go_angular.Go_serialize(buff[offset:])
    return offset
}

func (self *Twist) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_linear.Go_deserialize(buff[offset:])
    offset += self.Go_angular.Go_deserialize(buff[offset:])
    return offset
}

func (self *Twist) Go_serializedLength() (int) {
    length := 0
    length += self.Go_linear.Go_serializedLength()
    length += self.Go_angular.Go_serializedLength()
    return length
}

func (self *Twist) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Twist) Go_getType() (string) { return "geometry_msgs/Twist" }
func (self *Twist) Go_getMD5() (string) { return "29e7e4839b73f684ad08b19dc12c9c70" }
func (self *Twist) Go_getID() (uint32) { return 0 }
func (self *Twist) Go_setID(id uint32) { }

