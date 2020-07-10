package nav_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
)


type Odometry struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_child_frame_id string `json:"child_frame_id"`
    Go_pose *geometry_msgs.PoseWithCovariance `json:"pose"`
    Go_twist *geometry_msgs.TwistWithCovariance `json:"twist"`
}

func NewOdometry() (*Odometry) {
    newOdometry := new(Odometry)
    newOdometry.Go_header = std_msgs.NewHeader()
    newOdometry.Go_child_frame_id = ""
    newOdometry.Go_pose = geometry_msgs.NewPoseWithCovariance()
    newOdometry.Go_twist = geometry_msgs.NewTwistWithCovariance()
    return newOdometry
}

func (self *Odometry) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_child_frame_id = ""
    self.Go_pose = geometry_msgs.NewPoseWithCovariance()
    self.Go_twist = geometry_msgs.NewTwistWithCovariance()
}

func (self *Odometry) Go_serialize(buff []byte) (int) {
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
    offset += self.Go_pose.Go_serialize(buff[offset:])
    offset += self.Go_twist.Go_serialize(buff[offset:])
    return offset
}

func (self *Odometry) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_child_frame_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_child_frame_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_child_frame_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_child_frame_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_child_frame_id = string(buff[offset:(offset+length_child_frame_id)])
    offset += length_child_frame_id
    offset += self.Go_pose.Go_deserialize(buff[offset:])
    offset += self.Go_twist.Go_deserialize(buff[offset:])
    return offset
}

func (self *Odometry) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length_child_frame_id := len(self.Go_child_frame_id)
    length += 4
    length += length_child_frame_id
    length += self.Go_pose.Go_serializedLength()
    length += self.Go_twist.Go_serializedLength()
    return length
}

func (self *Odometry) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Odometry) Go_getType() (string) { return "nav_msgs/Odometry" }
func (self *Odometry) Go_getMD5() (string) { return "8fbd8c2e0caeb7be9b30b66a3e735193" }
func (self *Odometry) Go_getID() (uint32) { return 0 }
func (self *Odometry) Go_setID(id uint32) { }

