package trajectory_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type JointTrajectory struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_joint_names []string `json:"joint_names"`
    Go_points []*JointTrajectoryPoint `json:"points"`
}

func NewJointTrajectory() (*JointTrajectory) {
    newJointTrajectory := new(JointTrajectory)
    newJointTrajectory.Go_header = std_msgs.NewHeader()
    newJointTrajectory.Go_joint_names = []string{}
    newJointTrajectory.Go_points = []*JointTrajectoryPoint{}
    return newJointTrajectory
}

func (self *JointTrajectory) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_joint_names = []string{}
    self.Go_points = []*JointTrajectoryPoint{}
}

func (self *JointTrajectory) Go_serialize(buff []byte) (int) {
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
    length_points := len(self.Go_points)
    buff[offset + 0] = byte((length_points >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_points >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_points >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_points >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *JointTrajectory) Go_deserialize(buff []byte) (int) {
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
    length_points := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_points |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_points |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_points |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_points = make([]*JointTrajectoryPoint, length_points)
    for i := 0; i < length_points; i++ {
        self.Go_points[i] = NewJointTrajectoryPoint()
    }
    for i := 0; i < length_points; i++ {
        offset += self.Go_points[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *JointTrajectory) Go_serializedLength() (int) {
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
    length_points := len(self.Go_points)
    for i := 0; i < length_points; i++ {
        length += self.Go_points[i].Go_serializedLength()
    }
    return length
}

func (self *JointTrajectory) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JointTrajectory) Go_getType() (string) { return "trajectory_msgs/JointTrajectory" }
func (self *JointTrajectory) Go_getMD5() (string) { return "33e07cd7f5a6df021dad1271b3770d66" }
func (self *JointTrajectory) Go_getID() (uint32) { return 0 }
func (self *JointTrajectory) Go_setID(id uint32) { }

