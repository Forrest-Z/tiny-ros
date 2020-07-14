package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type JointState struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_name []string `json:"name"`
    Go_position []float64 `json:"position"`
    Go_velocity []float64 `json:"velocity"`
    Go_effort []float64 `json:"effort"`
}

func NewJointState() (*JointState) {
    newJointState := new(JointState)
    newJointState.Go_header = std_msgs.NewHeader()
    newJointState.Go_name = []string{}
    newJointState.Go_position = []float64{}
    newJointState.Go_velocity = []float64{}
    newJointState.Go_effort = []float64{}
    return newJointState
}

func (self *JointState) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_name = []string{}
    self.Go_position = []float64{}
    self.Go_velocity = []float64{}
    self.Go_effort = []float64{}
}

func (self *JointState) Go_serialize(buff []byte) (int) {
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
    length_position := len(self.Go_position)
    buff[offset + 0] = byte((length_position >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_position >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_position >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_position >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_position; i++ {
        bits_positioni := math.Float64bits(self.Go_position[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_positioni)
        offset += 8
    }
    length_velocity := len(self.Go_velocity)
    buff[offset + 0] = byte((length_velocity >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_velocity >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_velocity >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_velocity >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_velocity; i++ {
        bits_velocityi := math.Float64bits(self.Go_velocity[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_velocityi)
        offset += 8
    }
    length_effort := len(self.Go_effort)
    buff[offset + 0] = byte((length_effort >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_effort >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_effort >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_effort >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_effort; i++ {
        bits_efforti := math.Float64bits(self.Go_effort[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_efforti)
        offset += 8
    }
    return offset
}

func (self *JointState) Go_deserialize(buff []byte) (int) {
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
    length_position := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_position |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_position |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_position |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_position = make([]float64, length_position)
    for i := 0; i < length_position; i++ {
        bits_positioni := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_position[i] = math.Float64frombits(bits_positioni)
        offset += 8
    }
    length_velocity := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_velocity |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_velocity |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_velocity |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_velocity = make([]float64, length_velocity)
    for i := 0; i < length_velocity; i++ {
        bits_velocityi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_velocity[i] = math.Float64frombits(bits_velocityi)
        offset += 8
    }
    length_effort := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_effort |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_effort |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_effort |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_effort = make([]float64, length_effort)
    for i := 0; i < length_effort; i++ {
        bits_efforti := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_effort[i] = math.Float64frombits(bits_efforti)
        offset += 8
    }
    return offset
}

func (self *JointState) Go_serializedLength() (int) {
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
    length_position := len(self.Go_position)
    for i := 0; i < length_position; i++ {
        length += 8
    }
    length += 4
    length_velocity := len(self.Go_velocity)
    for i := 0; i < length_velocity; i++ {
        length += 8
    }
    length += 4
    length_effort := len(self.Go_effort)
    for i := 0; i < length_effort; i++ {
        length += 8
    }
    return length
}

func (self *JointState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JointState) Go_getType() (string) { return "sensor_msgs/JointState" }
func (self *JointState) Go_getMD5() (string) { return "6df7130a6d6a4c2f2037ce4a6e061fb9" }
func (self *JointState) Go_getID() (uint32) { return 0 }
func (self *JointState) Go_setID(id uint32) { }

