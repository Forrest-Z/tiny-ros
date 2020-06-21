package tf2_msgs

import (
    "geometry_msgs/TransformStamped"
    "tf2_msgs/TF2Error"
)

type LookupTransformResult struct {
    Go_transform geometry_msgs.TransformStamped `json:"transform"`
    Go_error tf2_msgs.TF2Error `json:"error"`
}

func NewLookupTransformResult() (*LookupTransformResult) {
    newLookupTransformResult := new(LookupTransformResult)
    newLookupTransformResult.Go_transform = geometry_msgs.NewTransformStamped()
    newLookupTransformResult.Go_error = tf2_msgs.NewTF2Error()
    return newLookupTransformResult
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

func (self *LookupTransformResult) Go_echo() (string) { return "" }
func (self *LookupTransformResult) Go_getType() (string) { return "tf2_msgs/LookupTransformResult" }
func (self *LookupTransformResult) Go_getMD5() (string) { return "7be4fc6719f512bc94491db1ccda6aee" }
func (self *LookupTransformResult) Go_getID() (uint32) { return 0 }
func (self *LookupTransformResult) Go_setID(id uint32) { }

