package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type PoseWithCovariance struct {
    Go_pose *Pose `json:"pose"`
    Go_covariance [36]float64 `json:"covariance"`
}

func NewPoseWithCovariance() (*PoseWithCovariance) {
    newPoseWithCovariance := new(PoseWithCovariance)
    newPoseWithCovariance.Go_pose = NewPose()
    newPoseWithCovariance.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    return newPoseWithCovariance
}

func (self *PoseWithCovariance) Go_initialize() {
    self.Go_pose = NewPose()
    self.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (self *PoseWithCovariance) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_pose.Go_serialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := math.Float64bits(self.Go_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *PoseWithCovariance) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_pose.Go_deserialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_covariance[i] = math.Float64frombits(bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *PoseWithCovariance) Go_serializedLength() (int) {
    length := 0
    length += self.Go_pose.Go_serializedLength()
    for i := 0; i < 36; i++ {
        length += 8
    }
    return length
}

func (self *PoseWithCovariance) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PoseWithCovariance) Go_getType() (string) { return "geometry_msgs/PoseWithCovariance" }
func (self *PoseWithCovariance) Go_getMD5() (string) { return "054c6283d50e78f8d9358aaaee5f4c1b" }
func (self *PoseWithCovariance) Go_getID() (uint32) { return 0 }
func (self *PoseWithCovariance) Go_setID(id uint32) { }

