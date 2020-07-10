package geometry_msgs

import (
    "encoding/json"
)


type Transform struct {
    Go_translation *Vector3 `json:"translation"`
    Go_rotation *Quaternion `json:"rotation"`
}

func NewTransform() (*Transform) {
    newTransform := new(Transform)
    newTransform.Go_translation = NewVector3()
    newTransform.Go_rotation = NewQuaternion()
    return newTransform
}

func (self *Transform) Go_initialize() {
    self.Go_translation = NewVector3()
    self.Go_rotation = NewQuaternion()
}

func (self *Transform) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_translation.Go_serialize(buff[offset:])
    offset += self.Go_rotation.Go_serialize(buff[offset:])
    return offset
}

func (self *Transform) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_translation.Go_deserialize(buff[offset:])
    offset += self.Go_rotation.Go_deserialize(buff[offset:])
    return offset
}

func (self *Transform) Go_serializedLength() (int) {
    length := 0
    length += self.Go_translation.Go_serializedLength()
    length += self.Go_rotation.Go_serializedLength()
    return length
}

func (self *Transform) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Transform) Go_getType() (string) { return "geometry_msgs/Transform" }
func (self *Transform) Go_getMD5() (string) { return "2526ee1b1cc2e723e386c3c1b048ba72" }
func (self *Transform) Go_getID() (uint32) { return 0 }
func (self *Transform) Go_setID(id uint32) { }

