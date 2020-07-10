package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
    "encoding/binary"
    "math"
)


type ContactState struct {
    Go_info string `json:"info"`
    Go_collision1_name string `json:"collision1_name"`
    Go_collision2_name string `json:"collision2_name"`
    Go_wrenches []*geometry_msgs.Wrench `json:"wrenches"`
    Go_total_wrench *geometry_msgs.Wrench `json:"total_wrench"`
    Go_contact_positions []*geometry_msgs.Vector3 `json:"contact_positions"`
    Go_contact_normals []*geometry_msgs.Vector3 `json:"contact_normals"`
    Go_depths []float64 `json:"depths"`
}

func NewContactState() (*ContactState) {
    newContactState := new(ContactState)
    newContactState.Go_info = ""
    newContactState.Go_collision1_name = ""
    newContactState.Go_collision2_name = ""
    newContactState.Go_wrenches = []*geometry_msgs.Wrench{}
    newContactState.Go_total_wrench = geometry_msgs.NewWrench()
    newContactState.Go_contact_positions = []*geometry_msgs.Vector3{}
    newContactState.Go_contact_normals = []*geometry_msgs.Vector3{}
    newContactState.Go_depths = []float64{}
    return newContactState
}

func (self *ContactState) Go_initialize() {
    self.Go_info = ""
    self.Go_collision1_name = ""
    self.Go_collision2_name = ""
    self.Go_wrenches = []*geometry_msgs.Wrench{}
    self.Go_total_wrench = geometry_msgs.NewWrench()
    self.Go_contact_positions = []*geometry_msgs.Vector3{}
    self.Go_contact_normals = []*geometry_msgs.Vector3{}
    self.Go_depths = []float64{}
}

func (self *ContactState) Go_serialize(buff []byte) (int) {
    offset := 0
    length_info := len(self.Go_info)
    buff[offset + 0] = byte((length_info >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_info >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_info >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_info >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_info)], self.Go_info)
    offset += length_info
    length_collision1_name := len(self.Go_collision1_name)
    buff[offset + 0] = byte((length_collision1_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_collision1_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_collision1_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_collision1_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_collision1_name)], self.Go_collision1_name)
    offset += length_collision1_name
    length_collision2_name := len(self.Go_collision2_name)
    buff[offset + 0] = byte((length_collision2_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_collision2_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_collision2_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_collision2_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_collision2_name)], self.Go_collision2_name)
    offset += length_collision2_name
    length_wrenches := len(self.Go_wrenches)
    buff[offset + 0] = byte((length_wrenches >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_wrenches >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_wrenches >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_wrenches >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_wrenches; i++ {
        offset += self.Go_wrenches[i].Go_serialize(buff[offset:])
    }
    offset += self.Go_total_wrench.Go_serialize(buff[offset:])
    length_contact_positions := len(self.Go_contact_positions)
    buff[offset + 0] = byte((length_contact_positions >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_contact_positions >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_contact_positions >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_contact_positions >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_contact_positions; i++ {
        offset += self.Go_contact_positions[i].Go_serialize(buff[offset:])
    }
    length_contact_normals := len(self.Go_contact_normals)
    buff[offset + 0] = byte((length_contact_normals >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_contact_normals >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_contact_normals >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_contact_normals >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_contact_normals; i++ {
        offset += self.Go_contact_normals[i].Go_serialize(buff[offset:])
    }
    length_depths := len(self.Go_depths)
    buff[offset + 0] = byte((length_depths >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_depths >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_depths >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_depths >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_depths; i++ {
        bits_depthsi := math.Float64bits(self.Go_depths[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_depthsi)
        offset += 8
    }
    return offset
}

func (self *ContactState) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_info := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_info |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_info |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_info |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_info = string(buff[offset:(offset+length_info)])
    offset += length_info
    length_collision1_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_collision1_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_collision1_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_collision1_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_collision1_name = string(buff[offset:(offset+length_collision1_name)])
    offset += length_collision1_name
    length_collision2_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_collision2_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_collision2_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_collision2_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_collision2_name = string(buff[offset:(offset+length_collision2_name)])
    offset += length_collision2_name
    length_wrenches := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_wrenches |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_wrenches |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_wrenches |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_wrenches = make([]*geometry_msgs.Wrench, length_wrenches)
    for i := 0; i < length_wrenches; i++ {
        self.Go_wrenches[i] = geometry_msgs.NewWrench()
    }
    for i := 0; i < length_wrenches; i++ {
        offset += self.Go_wrenches[i].Go_deserialize(buff[offset:])
    }
    offset += self.Go_total_wrench.Go_deserialize(buff[offset:])
    length_contact_positions := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_contact_positions |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_contact_positions |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_contact_positions |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_contact_positions = make([]*geometry_msgs.Vector3, length_contact_positions)
    for i := 0; i < length_contact_positions; i++ {
        self.Go_contact_positions[i] = geometry_msgs.NewVector3()
    }
    for i := 0; i < length_contact_positions; i++ {
        offset += self.Go_contact_positions[i].Go_deserialize(buff[offset:])
    }
    length_contact_normals := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_contact_normals |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_contact_normals |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_contact_normals |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_contact_normals = make([]*geometry_msgs.Vector3, length_contact_normals)
    for i := 0; i < length_contact_normals; i++ {
        self.Go_contact_normals[i] = geometry_msgs.NewVector3()
    }
    for i := 0; i < length_contact_normals; i++ {
        offset += self.Go_contact_normals[i].Go_deserialize(buff[offset:])
    }
    length_depths := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_depths |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_depths |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_depths |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_depths = make([]float64, length_depths)
    for i := 0; i < length_depths; i++ {
        bits_depthsi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_depths[i] = math.Float64frombits(bits_depthsi)
        offset += 8
    }
    return offset
}

func (self *ContactState) Go_serializedLength() (int) {
    length := 0
    length_info := len(self.Go_info)
    length += 4
    length += length_info
    length_collision1_name := len(self.Go_collision1_name)
    length += 4
    length += length_collision1_name
    length_collision2_name := len(self.Go_collision2_name)
    length += 4
    length += length_collision2_name
    length += 4
    length_wrenches := len(self.Go_wrenches)
    for i := 0; i < length_wrenches; i++ {
        length += self.Go_wrenches[i].Go_serializedLength()
    }
    length += self.Go_total_wrench.Go_serializedLength()
    length += 4
    length_contact_positions := len(self.Go_contact_positions)
    for i := 0; i < length_contact_positions; i++ {
        length += self.Go_contact_positions[i].Go_serializedLength()
    }
    length += 4
    length_contact_normals := len(self.Go_contact_normals)
    for i := 0; i < length_contact_normals; i++ {
        length += self.Go_contact_normals[i].Go_serializedLength()
    }
    length += 4
    length_depths := len(self.Go_depths)
    for i := 0; i < length_depths; i++ {
        length += 8
    }
    return length
}

func (self *ContactState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ContactState) Go_getType() (string) { return "gazebo_msgs/ContactState" }
func (self *ContactState) Go_getMD5() (string) { return "d82d0f0cae88aebf6b2cc86caea33a2b" }
func (self *ContactState) Go_getID() (uint32) { return 0 }
func (self *ContactState) Go_setID(id uint32) { }

