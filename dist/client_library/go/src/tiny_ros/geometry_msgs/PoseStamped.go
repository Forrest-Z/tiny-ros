package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type PoseStamped struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_pose *Pose `json:"pose"`
}

func NewPoseStamped() (*PoseStamped) {
    newPoseStamped := new(PoseStamped)
    newPoseStamped.Go_header = std_msgs.NewHeader()
    newPoseStamped.Go_pose = NewPose()
    return newPoseStamped
}

func (self *PoseStamped) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_pose = NewPose()
}

func (self *PoseStamped) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_pose.Go_serialize(buff[offset:])
    return offset
}

func (self *PoseStamped) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_pose.Go_deserialize(buff[offset:])
    return offset
}

func (self *PoseStamped) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_pose.Go_serializedLength()
    return length
}

func (self *PoseStamped) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PoseStamped) Go_getType() (string) { return "geometry_msgs/PoseStamped" }
func (self *PoseStamped) Go_getMD5() (string) { return "c7084e6b27c3d6e62efd9bf6d2f6540f" }
func (self *PoseStamped) Go_getID() (uint32) { return 0 }
func (self *PoseStamped) Go_setID(id uint32) { }

