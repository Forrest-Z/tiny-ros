package tf2_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type TFMessage struct {
    Go_transforms []*geometry_msgs.TransformStamped `json:"transforms"`
}

func NewTFMessage() (*TFMessage) {
    newTFMessage := new(TFMessage)
    newTFMessage.Go_transforms = []*geometry_msgs.TransformStamped{}
    return newTFMessage
}

func (self *TFMessage) Go_initialize() {
    self.Go_transforms = []*geometry_msgs.TransformStamped{}
}

func (self *TFMessage) Go_serialize(buff []byte) (int) {
    offset := 0
    length_transforms := len(self.Go_transforms)
    buff[offset + 0] = byte((length_transforms >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_transforms >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_transforms >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_transforms >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_transforms; i++ {
        offset += self.Go_transforms[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *TFMessage) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_transforms := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_transforms |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_transforms |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_transforms |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_transforms = make([]*geometry_msgs.TransformStamped, length_transforms)
    for i := 0; i < length_transforms; i++ {
        self.Go_transforms[i] = geometry_msgs.NewTransformStamped()
    }
    for i := 0; i < length_transforms; i++ {
        offset += self.Go_transforms[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *TFMessage) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_transforms := len(self.Go_transforms)
    for i := 0; i < length_transforms; i++ {
        length += self.Go_transforms[i].Go_serializedLength()
    }
    return length
}

func (self *TFMessage) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TFMessage) Go_getType() (string) { return "tf2_msgs/TFMessage" }
func (self *TFMessage) Go_getMD5() (string) { return "cb93cfe6a141f8d8af7cc34997ec99fe" }
func (self *TFMessage) Go_getID() (uint32) { return 0 }
func (self *TFMessage) Go_setID(id uint32) { }

