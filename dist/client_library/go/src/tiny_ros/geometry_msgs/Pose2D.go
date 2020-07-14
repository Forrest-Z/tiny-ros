package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type Pose2D struct {
    Go_x float64 `json:"x"`
    Go_y float64 `json:"y"`
    Go_theta float64 `json:"theta"`
}

func NewPose2D() (*Pose2D) {
    newPose2D := new(Pose2D)
    newPose2D.Go_x = 0.0
    newPose2D.Go_y = 0.0
    newPose2D.Go_theta = 0.0
    return newPose2D
}

func (self *Pose2D) Go_initialize() {
    self.Go_x = 0.0
    self.Go_y = 0.0
    self.Go_theta = 0.0
}

func (self *Pose2D) Go_serialize(buff []byte) (int) {
    offset := 0
    bits_x := math.Float64bits(self.Go_x)
    binary.LittleEndian.PutUint64(buff[offset:], bits_x)
    offset += 8
    bits_y := math.Float64bits(self.Go_y)
    binary.LittleEndian.PutUint64(buff[offset:], bits_y)
    offset += 8
    bits_theta := math.Float64bits(self.Go_theta)
    binary.LittleEndian.PutUint64(buff[offset:], bits_theta)
    offset += 8
    return offset
}

func (self *Pose2D) Go_deserialize(buff []byte) (int) {
    offset := 0
    bits_x := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_x = math.Float64frombits(bits_x)
    offset += 8
    bits_y := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_y = math.Float64frombits(bits_y)
    offset += 8
    bits_theta := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_theta = math.Float64frombits(bits_theta)
    offset += 8
    return offset
}

func (self *Pose2D) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += 8
    return length
}

func (self *Pose2D) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Pose2D) Go_getType() (string) { return "geometry_msgs/Pose2D" }
func (self *Pose2D) Go_getMD5() (string) { return "509f362ff66c4d3df21020fa7c01f8c6" }
func (self *Pose2D) Go_getID() (uint32) { return 0 }
func (self *Pose2D) Go_setID(id uint32) { }

