package geometry_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type PoseArray struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_poses []*Pose `json:"poses"`
}

func NewPoseArray() (*PoseArray) {
    newPoseArray := new(PoseArray)
    newPoseArray.Go_header = std_msgs.NewHeader()
    newPoseArray.Go_poses = []*Pose{}
    return newPoseArray
}

func (self *PoseArray) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_poses = []*Pose{}
}

func (self *PoseArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_poses := len(self.Go_poses)
    buff[offset + 0] = byte((length_poses >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_poses >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_poses >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_poses >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_poses; i++ {
        offset += self.Go_poses[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *PoseArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_poses := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_poses |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_poses |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_poses |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_poses = make([]*Pose, length_poses)
    for i := 0; i < length_poses; i++ {
        self.Go_poses[i] = NewPose()
    }
    for i := 0; i < length_poses; i++ {
        offset += self.Go_poses[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *PoseArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_poses := len(self.Go_poses)
    for i := 0; i < length_poses; i++ {
        length += self.Go_poses[i].Go_serializedLength()
    }
    return length
}

func (self *PoseArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PoseArray) Go_getType() (string) { return "geometry_msgs/PoseArray" }
func (self *PoseArray) Go_getMD5() (string) { return "184f43246f3bc9cb5d0613694e6641a6" }
func (self *PoseArray) Go_getID() (uint32) { return 0 }
func (self *PoseArray) Go_setID(id uint32) { }

