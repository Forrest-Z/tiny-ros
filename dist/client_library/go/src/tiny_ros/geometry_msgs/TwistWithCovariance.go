package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type TwistWithCovariance struct {
    Go_twist *Twist `json:"twist"`
    Go_covariance [36]float64 `json:"covariance"`
}

func NewTwistWithCovariance() (*TwistWithCovariance) {
    newTwistWithCovariance := new(TwistWithCovariance)
    newTwistWithCovariance.Go_twist = NewTwist()
    newTwistWithCovariance.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    return newTwistWithCovariance
}

func (self *TwistWithCovariance) Go_initialize() {
    self.Go_twist = NewTwist()
    self.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (self *TwistWithCovariance) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_twist.Go_serialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := math.Float64bits(self.Go_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *TwistWithCovariance) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_twist.Go_deserialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_covariance[i] = math.Float64frombits(bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *TwistWithCovariance) Go_serializedLength() (int) {
    length := 0
    length += self.Go_twist.Go_serializedLength()
    for i := 0; i < 36; i++ {
        length += 8
    }
    return length
}

func (self *TwistWithCovariance) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TwistWithCovariance) Go_getType() (string) { return "geometry_msgs/TwistWithCovariance" }
func (self *TwistWithCovariance) Go_getMD5() (string) { return "0421bae691707888d99987e0bbcf4c55" }
func (self *TwistWithCovariance) Go_getID() (uint32) { return 0 }
func (self *TwistWithCovariance) Go_setID(id uint32) { }

