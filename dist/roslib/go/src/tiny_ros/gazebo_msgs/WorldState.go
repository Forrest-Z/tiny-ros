package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
)


type WorldState struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_name []string `json:"name"`
    Go_pose []*geometry_msgs.Pose `json:"pose"`
    Go_twist []*geometry_msgs.Twist `json:"twist"`
    Go_wrench []*geometry_msgs.Wrench `json:"wrench"`
}

func NewWorldState() (*WorldState) {
    newWorldState := new(WorldState)
    newWorldState.Go_header = std_msgs.NewHeader()
    newWorldState.Go_name = []string{}
    newWorldState.Go_pose = []*geometry_msgs.Pose{}
    newWorldState.Go_twist = []*geometry_msgs.Twist{}
    newWorldState.Go_wrench = []*geometry_msgs.Wrench{}
    return newWorldState
}

func (self *WorldState) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_name = []string{}
    self.Go_pose = []*geometry_msgs.Pose{}
    self.Go_twist = []*geometry_msgs.Twist{}
    self.Go_wrench = []*geometry_msgs.Wrench{}
}

func (self *WorldState) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
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

func (self *WorldState) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
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

func (self *WorldState) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
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
    length += 4
    length_wrench := len(self.Go_wrench)
    for i := 0; i < length_wrench; i++ {
        length += self.Go_wrench[i].Go_serializedLength()
    }
    return length
}

func (self *WorldState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *WorldState) Go_getType() (string) { return "gazebo_msgs/WorldState" }
func (self *WorldState) Go_getMD5() (string) { return "0e1997127271c4d021f99645c28f1c09" }
func (self *WorldState) Go_getID() (uint32) { return 0 }
func (self *WorldState) Go_setID(id uint32) { }

