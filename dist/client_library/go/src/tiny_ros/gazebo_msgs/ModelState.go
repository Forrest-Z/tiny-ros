package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type ModelState struct {
    Go_model_name string `json:"model_name"`
    Go_pose *geometry_msgs.Pose `json:"pose"`
    Go_twist *geometry_msgs.Twist `json:"twist"`
    Go_reference_frame string `json:"reference_frame"`
}

func NewModelState() (*ModelState) {
    newModelState := new(ModelState)
    newModelState.Go_model_name = ""
    newModelState.Go_pose = geometry_msgs.NewPose()
    newModelState.Go_twist = geometry_msgs.NewTwist()
    newModelState.Go_reference_frame = ""
    return newModelState
}

func (self *ModelState) Go_initialize() {
    self.Go_model_name = ""
    self.Go_pose = geometry_msgs.NewPose()
    self.Go_twist = geometry_msgs.NewTwist()
    self.Go_reference_frame = ""
}

func (self *ModelState) Go_serialize(buff []byte) (int) {
    offset := 0
    length_model_name := len(self.Go_model_name)
    buff[offset + 0] = byte((length_model_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_model_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_model_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_model_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_model_name)], self.Go_model_name)
    offset += length_model_name
    offset += self.Go_pose.Go_serialize(buff[offset:])
    offset += self.Go_twist.Go_serialize(buff[offset:])
    length_reference_frame := len(self.Go_reference_frame)
    buff[offset + 0] = byte((length_reference_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_reference_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_reference_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_reference_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_reference_frame)], self.Go_reference_frame)
    offset += length_reference_frame
    return offset
}

func (self *ModelState) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_model_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_model_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_model_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_model_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_model_name = string(buff[offset:(offset+length_model_name)])
    offset += length_model_name
    offset += self.Go_pose.Go_deserialize(buff[offset:])
    offset += self.Go_twist.Go_deserialize(buff[offset:])
    length_reference_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_reference_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_reference_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_reference_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_reference_frame = string(buff[offset:(offset+length_reference_frame)])
    offset += length_reference_frame
    return offset
}

func (self *ModelState) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    length += self.Go_pose.Go_serializedLength()
    length += self.Go_twist.Go_serializedLength()
    length_reference_frame := len(self.Go_reference_frame)
    length += 4
    length += length_reference_frame
    return length
}

func (self *ModelState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ModelState) Go_getType() (string) { return "gazebo_msgs/ModelState" }
func (self *ModelState) Go_getMD5() (string) { return "dee4d802363b4d6bd1ed61e20c2c4635" }
func (self *ModelState) Go_getID() (uint32) { return 0 }
func (self *ModelState) Go_setID(id uint32) { }

