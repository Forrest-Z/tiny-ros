package trajectory_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
    "tiny_ros/tinyros/time"
)


type MultiDOFJointTrajectoryPoint struct {
    Go_transforms []*geometry_msgs.Transform `json:"transforms"`
    Go_velocities []*geometry_msgs.Twist `json:"velocities"`
    Go_accelerations []*geometry_msgs.Twist `json:"accelerations"`
    Go_time_from_start *rostime.Duration `json:"time_from_start"`
}

func NewMultiDOFJointTrajectoryPoint() (*MultiDOFJointTrajectoryPoint) {
    newMultiDOFJointTrajectoryPoint := new(MultiDOFJointTrajectoryPoint)
    newMultiDOFJointTrajectoryPoint.Go_transforms = []*geometry_msgs.Transform{}
    newMultiDOFJointTrajectoryPoint.Go_velocities = []*geometry_msgs.Twist{}
    newMultiDOFJointTrajectoryPoint.Go_accelerations = []*geometry_msgs.Twist{}
    newMultiDOFJointTrajectoryPoint.Go_time_from_start = rostime.NewDuration()
    return newMultiDOFJointTrajectoryPoint
}

func (self *MultiDOFJointTrajectoryPoint) Go_initialize() {
    self.Go_transforms = []*geometry_msgs.Transform{}
    self.Go_velocities = []*geometry_msgs.Twist{}
    self.Go_accelerations = []*geometry_msgs.Twist{}
    self.Go_time_from_start = rostime.NewDuration()
}

func (self *MultiDOFJointTrajectoryPoint) Go_serialize(buff []byte) (int) {
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
    length_velocities := len(self.Go_velocities)
    buff[offset + 0] = byte((length_velocities >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_velocities >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_velocities >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_velocities >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_velocities; i++ {
        offset += self.Go_velocities[i].Go_serialize(buff[offset:])
    }
    length_accelerations := len(self.Go_accelerations)
    buff[offset + 0] = byte((length_accelerations >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_accelerations >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_accelerations >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_accelerations >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_accelerations; i++ {
        offset += self.Go_accelerations[i].Go_serialize(buff[offset:])
    }
    buff[offset + 0] = byte((self.Go_time_from_start.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_from_start.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_from_start.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_from_start.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_time_from_start.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_from_start.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_from_start.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_from_start.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *MultiDOFJointTrajectoryPoint) Go_deserialize(buff []byte) (int) {
    offset := 0
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
    length_velocities := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_velocities |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_velocities |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_velocities |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_velocities = make([]*geometry_msgs.Twist, length_velocities)
    for i := 0; i < length_velocities; i++ {
        self.Go_velocities[i] = geometry_msgs.NewTwist()
    }
    for i := 0; i < length_velocities; i++ {
        offset += self.Go_velocities[i].Go_deserialize(buff[offset:])
    }
    length_accelerations := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_accelerations |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_accelerations |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_accelerations |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_accelerations = make([]*geometry_msgs.Twist, length_accelerations)
    for i := 0; i < length_accelerations; i++ {
        self.Go_accelerations[i] = geometry_msgs.NewTwist()
    }
    for i := 0; i < length_accelerations; i++ {
        offset += self.Go_accelerations[i].Go_deserialize(buff[offset:])
    }
    self.Go_time_from_start.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_time_from_start.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *MultiDOFJointTrajectoryPoint) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_transforms := len(self.Go_transforms)
    for i := 0; i < length_transforms; i++ {
        length += self.Go_transforms[i].Go_serializedLength()
    }
    length += 4
    length_velocities := len(self.Go_velocities)
    for i := 0; i < length_velocities; i++ {
        length += self.Go_velocities[i].Go_serializedLength()
    }
    length += 4
    length_accelerations := len(self.Go_accelerations)
    for i := 0; i < length_accelerations; i++ {
        length += self.Go_accelerations[i].Go_serializedLength()
    }
    length += 4
    length += 4
    return length
}

func (self *MultiDOFJointTrajectoryPoint) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MultiDOFJointTrajectoryPoint) Go_getType() (string) { return "trajectory_msgs/MultiDOFJointTrajectoryPoint" }
func (self *MultiDOFJointTrajectoryPoint) Go_getMD5() (string) { return "f8b4a74b416279b5c5d565029308ff08" }
func (self *MultiDOFJointTrajectoryPoint) Go_getID() (uint32) { return 0 }
func (self *MultiDOFJointTrajectoryPoint) Go_setID(id uint32) { }

