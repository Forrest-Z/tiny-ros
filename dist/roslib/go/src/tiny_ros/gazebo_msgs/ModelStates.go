package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type ModelStates struct {
    Go_name []string `json:"name"`
    Go_pose []*geometry_msgs.Pose `json:"pose"`
    Go_twist []*geometry_msgs.Twist `json:"twist"`
}

func NewModelStates() (*ModelStates) {
    newModelStates := new(ModelStates)
    newModelStates.Go_name = []string{}
    newModelStates.Go_pose = []*geometry_msgs.Pose{}
    newModelStates.Go_twist = []*geometry_msgs.Twist{}
    return newModelStates
}

func (self *ModelStates) Go_initialize() {
    self.Go_name = []string{}
    self.Go_pose = []*geometry_msgs.Pose{}
    self.Go_twist = []*geometry_msgs.Twist{}
}

func (self *ModelStates) Go_serialize(buff []byte) (int) {
    offset := 0
    length_name := len(self.Go_name)
    buff[offset + 0] = byte((length_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_name >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_name; i++ {
        length_namei := len(self.Go_name[i])
        buff[offset + 0] = byte((length_namei >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_namei >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_namei >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_namei >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_namei)], self.Go_name[i])
        offset += length_namei
    }
    length_pose := len(self.Go_pose)
    buff[offset + 0] = byte((length_pose >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_pose >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_pose >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_pose >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_pose; i++ {
        offset += self.Go_pose[i].Go_serialize(buff[offset:])
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
    return offset
}

func (self *ModelStates) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_name = make([]string, length_name)
    for i := 0; i < length_name; i++ {
        length_namei := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_namei |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_namei |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_namei |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_name[i] = string(buff[offset:(offset+length_namei)])
        offset += length_namei
    }
    length_pose := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_pose |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_pose |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_pose |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_pose = make([]*geometry_msgs.Pose, length_pose)
    for i := 0; i < length_pose; i++ {
        self.Go_pose[i] = geometry_msgs.NewPose()
    }
    for i := 0; i < length_pose; i++ {
        offset += self.Go_pose[i].Go_deserialize(buff[offset:])
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
    return offset
}

func (self *ModelStates) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_name := len(self.Go_name)
    for i := 0; i < length_name; i++ {
        length_namei := len(self.Go_name[i])
        length += 4
        length += length_namei
    }
    length += 4
    length_pose := len(self.Go_pose)
    for i := 0; i < length_pose; i++ {
        length += self.Go_pose[i].Go_serializedLength()
    }
    length += 4
    length_twist := len(self.Go_twist)
    for i := 0; i < length_twist; i++ {
        length += self.Go_twist[i].Go_serializedLength()
    }
    return length
}

func (self *ModelStates) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ModelStates) Go_getType() (string) { return "gazebo_msgs/ModelStates" }
func (self *ModelStates) Go_getMD5() (string) { return "05074231128e3d50825a5f46d9217fda" }
func (self *ModelStates) Go_getID() (uint32) { return 0 }
func (self *ModelStates) Go_setID(id uint32) { }

