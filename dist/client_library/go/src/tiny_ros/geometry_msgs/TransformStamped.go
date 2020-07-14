package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type TransformStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_child_frame_id string `json:"child_frame_id"`
    Go_transform *Transform `json:"transform"`
}

func NewTransformStamped() (*TransformStamped) {
    newTransformStamped := new(TransformStamped)
    newTransformStamped.Go_header = std_msgs.NewHeader()
    newTransformStamped.Go_child_frame_id = ""
    newTransformStamped.Go_transform = NewTransform()
    return newTransformStamped
}

func (self *TransformStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_child_frame_id = ""
    self.Go_transform = NewTransform()
}

func (self *TransformStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_child_frame_id := len(self.Go_child_frame_id)
    buff[offset + 0] = byte((length_child_frame_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_child_frame_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_child_frame_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_child_frame_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_child_frame_id)], self.Go_child_frame_id)
    offset += length_child_frame_id
    offset += self.Go_transform.Go_serialize(buff[offset:])
    return offset
}

func (self *TransformStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_child_frame_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_child_frame_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_child_frame_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_child_frame_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_child_frame_id = string(buff[offset:(offset+length_child_frame_id)])
    offset += length_child_frame_id
    offset += self.Go_transform.Go_deserialize(buff[offset:])
    return offset
}

func (self *TransformStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length_child_frame_id := len(self.Go_child_frame_id)
    length += 4
    length += length_child_frame_id
    length += self.Go_transform.Go_serializedLength()
    return length
}

func (self *TransformStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TransformStamped) Go_getType() (string) { return "geometry_msgs/TransformStamped" }
func (self *TransformStamped) Go_getMD5() (string) { return "e46d447d8e8afc726d6013a3ae4146dd" }
func (self *TransformStamped) Go_getID() (uint32) { return 0 }
func (self *TransformStamped) Go_setID(id uint32) { }

