package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type Vector3Stamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_vector *Vector3 `json:"vector"`
}

func NewVector3Stamped() (*Vector3Stamped) {
    newVector3Stamped := new(Vector3Stamped)
    newVector3Stamped.Go_header = std_msgs.NewHeader()
    newVector3Stamped.Go_vector = NewVector3()
    return newVector3Stamped
}

func (self *Vector3Stamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_vector = NewVector3()
}

func (self *Vector3Stamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_vector.Go_serialize(buff[offset:])
    return offset
}

func (self *Vector3Stamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_vector.Go_deserialize(buff[offset:])
    return offset
}

func (self *Vector3Stamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_vector.Go_serializedLength()
    return length
}

func (self *Vector3Stamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Vector3Stamped) Go_getType() (string) { return "geometry_msgs/Vector3Stamped" }
func (self *Vector3Stamped) Go_getMD5() (string) { return "4b85025eb6f70f6b1e0cefbb75f69ac2" }
func (self *Vector3Stamped) Go_getID() (uint32) { return 0 }
func (self *Vector3Stamped) Go_setID(id uint32) { }

