package tf2_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type LookupTransformResult struct {
    Go_transform *geometry_msgs.TransformStamped `json:"transform"`
    Go_error *TF2Error `json:"error"`
}

func NewLookupTransformResult() (*LookupTransformResult) {
    newLookupTransformResult := new(LookupTransformResult)
    newLookupTransformResult.Go_transform = geometry_msgs.NewTransformStamped()
    newLookupTransformResult.Go_error = NewTF2Error()
    return newLookupTransformResult
}

func (self *LookupTransformResult) Go_initialize() {
    self.Go_transform = geometry_msgs.NewTransformStamped()
    self.Go_error = NewTF2Error()
}

func (self *LookupTransformResult) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_transform.Go_serialize(buff[offset:])
    offset += self.Go_error.Go_serialize(buff[offset:])
    return offset
}

func (self *LookupTransformResult) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_transform.Go_deserialize(buff[offset:])
    offset += self.Go_error.Go_deserialize(buff[offset:])
    return offset
}

func (self *LookupTransformResult) Go_serializedLength() (int) {
    length := 0
    length += self.Go_transform.Go_serializedLength()
    length += self.Go_error.Go_serializedLength()
    return length
}

func (self *LookupTransformResult) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LookupTransformResult) Go_getType() (string) { return "tf2_msgs/LookupTransformResult" }
func (self *LookupTransformResult) Go_getMD5() (string) { return "7be4fc6719f512bc94491db1ccda6aee" }
func (self *LookupTransformResult) Go_getID() (uint32) { return 0 }
func (self *LookupTransformResult) Go_setID(id uint32) { }

