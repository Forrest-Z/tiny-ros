package tf

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type tfMessage struct {
    Go_transforms []*geometry_msgs.TransformStamped `json:"transforms"`
}

func NewtfMessage() (*tfMessage) {
    newtfMessage := new(tfMessage)
    newtfMessage.Go_transforms = []*geometry_msgs.TransformStamped{}
    return newtfMessage
}

func (self *tfMessage) Go_initialize() {
    self.Go_transforms = []*geometry_msgs.TransformStamped{}
}

func (self *tfMessage) Go_serialize(buff []byte) (int) {
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

func (self *tfMessage) Go_deserialize(buff []byte) (int) {
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

func (self *tfMessage) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_transforms := len(self.Go_transforms)
    for i := 0; i < length_transforms; i++ {
        length += self.Go_transforms[i].Go_serializedLength()
    }
    return length
}

func (self *tfMessage) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *tfMessage) Go_getType() (string) { return "tf/tfMessage" }
func (self *tfMessage) Go_getMD5() (string) { return "0401e4f6583aa665321e471e02ec671b" }
func (self *tfMessage) Go_getID() (uint32) { return 0 }
func (self *tfMessage) Go_setID(id uint32) { }

