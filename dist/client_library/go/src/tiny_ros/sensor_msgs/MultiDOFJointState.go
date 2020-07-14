package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
)


type MultiDOFJointState struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_joint_names []string `json:"joint_names"`
    Go_transforms []*geometry_msgs.Transform `json:"transforms"`
    Go_twist []*geometry_msgs.Twist `json:"twist"`
    Go_wrench []*geometry_msgs.Wrench `json:"wrench"`
}

func NewMultiDOFJointState() (*MultiDOFJointState) {
    newMultiDOFJointState := new(MultiDOFJointState)
    newMultiDOFJointState.Go_header = std_msgs.NewHeader()
    newMultiDOFJointState.Go_joint_names = []string{}
    newMultiDOFJointState.Go_transforms = []*geometry_msgs.Transform{}
    newMultiDOFJointState.Go_twist = []*geometry_msgs.Twist{}
    newMultiDOFJointState.Go_wrench = []*geometry_msgs.Wrench{}
    return newMultiDOFJointState
}

func (self *MultiDOFJointState) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_joint_names = []string{}
    self.Go_transforms = []*geometry_msgs.Transform{}
    self.Go_twist = []*geometry_msgs.Twist{}
    self.Go_wrench = []*geometry_msgs.Wrench{}
}

func (self *MultiDOFJointState) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_joint_names := len(self.Go_joint_names)
    buff[offset + 0] = byte((length_joint_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_joint_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_joint_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_joint_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := len(self.Go_joint_names[i])
        buff[offset + 0] = byte((length_joint_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_joint_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_joint_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_joint_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_joint_namesi)], self.Go_joint_names[i])
        offset += length_joint_namesi
    }
    length_transforms := len(self.Go_transforms)
    buff[offset + 0] = byte((length_transforms >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_transforms >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_transforms >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_transforms >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_transforms; i++ {
        offset += self.Go_transforms[i].Go_serialize(buff[offset:])
    }
    length_twist := len(self.Go_twist)
    buff[offset + 0] = byte((length_twist >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_twist >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_twist >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_twist >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_twist; i++ {
        offset += self.Go_twist[i].Go_serialize(buff[offset:])
    }
    length_wrench := len(self.Go_wrench)
    buff[offset + 0] = byte((length_wrench >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_wrench >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_wrench >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_wrench >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_wrench; i++ {
        offset += self.Go_wrench[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *MultiDOFJointState) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_joint_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_joint_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_joint_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_joint_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_joint_names = make([]string, length_joint_names)
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_joint_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_joint_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_joint_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_joint_names[i] = string(buff[offset:(offset+length_joint_namesi)])
        offset += length_joint_namesi
    }
    length_transforms := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_transforms |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_transforms |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_transforms |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_transforms = make([]*geometry_msgs.Transform, length_transforms)
    for i := 0; i < length_transforms; i++ {
        self.Go_transforms[i] = geometry_msgs.NewTransform()
    }
    for i := 0; i < length_transforms; i++ {
        offset += self.Go_transforms[i].Go_deserialize(buff[offset:])
    }
    length_twist := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_twist |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_twist |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_twist |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_twist = make([]*geometry_msgs.Twist, length_twist)
    for i := 0; i < length_twist; i++ {
        self.Go_twist[i] = geometry_msgs.NewTwist()
    }
    for i := 0; i < length_twist; i++ {
        offset += self.Go_twist[i].Go_deserialize(buff[offset:])
    }
    length_wrench := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_wrench |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_wrench |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_wrench |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_wrench = make([]*geometry_msgs.Wrench, length_wrench)
    for i := 0; i < length_wrench; i++ {
        self.Go_wrench[i] = geometry_msgs.NewWrench()
    }
    for i := 0; i < length_wrench; i++ {
        offset += self.Go_wrench[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *MultiDOFJointState) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_joint_names := len(self.Go_joint_names)
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := len(self.Go_joint_names[i])
        length += 4
        length += length_joint_namesi
    }
    length += 4
    length_transforms := len(self.Go_transforms)
    for i := 0; i < length_transforms; i++ {
        length += self.Go_transforms[i].Go_serializedLength()
    }
    length += 4
    length_twist := len(self.Go_twist)
    for i := 0; i < length_twist; i++ {
        length += self.Go_twist[i].Go_serializedLength()
    }
    length += 4
    length_wrench := len(self.Go_wrench)
    for i := 0; i < length_wrench; i++ {
        length += self.Go_wrench[i].Go_serializedLength()
    }
    return length
}

func (self *MultiDOFJointState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MultiDOFJointState) Go_getType() (string) { return "sensor_msgs/MultiDOFJointState" }
func (self *MultiDOFJointState) Go_getMD5() (string) { return "c1b0232d8e5071b24db76eb143f286eb" }
func (self *MultiDOFJointState) Go_getID() (uint32) { return 0 }
func (self *MultiDOFJointState) Go_setID(id uint32) { }

