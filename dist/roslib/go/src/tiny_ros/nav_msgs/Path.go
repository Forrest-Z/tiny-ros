package nav_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
)


type Path struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_poses []*geometry_msgs.PoseStamped `json:"poses"`
}

func NewPath() (*Path) {
    newPath := new(Path)
    newPath.Go_header = std_msgs.NewHeader()
    newPath.Go_poses = []*geometry_msgs.PoseStamped{}
    return newPath
}

func (self *Path) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_poses = []*geometry_msgs.PoseStamped{}
}

func (self *Path) Go_serialize(buff []byte) (int) {
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

func (self *Path) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_poses := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_poses |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_poses |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_poses |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_poses = make([]*geometry_msgs.PoseStamped, length_poses)
    for i := 0; i < length_poses; i++ {
        self.Go_poses[i] = geometry_msgs.NewPoseStamped()
    }
    for i := 0; i < length_poses; i++ {
        offset += self.Go_poses[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *Path) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_poses := len(self.Go_poses)
    for i := 0; i < length_poses; i++ {
        length += self.Go_poses[i].Go_serializedLength()
    }
    return length
}

func (self *Path) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Path) Go_getType() (string) { return "nav_msgs/Path" }
func (self *Path) Go_getMD5() (string) { return "4a185240c929c496a7e0d6202e3c89af" }
func (self *Path) Go_getID() (uint32) { return 0 }
func (self *Path) Go_setID(id uint32) { }

