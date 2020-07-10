package geometry_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type AccelWithCovariance struct {
    Go_accel *Accel `json:"accel"`
    Go_covariance [36]float64 `json:"covariance"`
}

func NewAccelWithCovariance() (*AccelWithCovariance) {
    newAccelWithCovariance := new(AccelWithCovariance)
    newAccelWithCovariance.Go_accel = NewAccel()
    newAccelWithCovariance.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    return newAccelWithCovariance
}

func (self *AccelWithCovariance) Go_initialize() {
    self.Go_accel = NewAccel()
    self.Go_covariance = [36]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (self *AccelWithCovariance) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_accel.Go_serialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := math.Float64bits(self.Go_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *AccelWithCovariance) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_accel.Go_deserialize(buff[offset:])
    for i := 0; i < 36; i++ {
        bits_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_covariance[i] = math.Float64frombits(bits_covariancei)
        offset += 8
    }
    return offset
}

func (self *AccelWithCovariance) Go_serializedLength() (int) {
    length := 0
    length += self.Go_accel.Go_serializedLength()
    for i := 0; i < 36; i++ {
        length += 8
    }
    return length
}

func (self *AccelWithCovariance) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *AccelWithCovariance) Go_getType() (string) { return "geometry_msgs/AccelWithCovariance" }
func (self *AccelWithCovariance) Go_getMD5() (string) { return "6c9c3b4380e0391a48b0a1be79b38ac6" }
func (self *AccelWithCovariance) Go_getID() (uint32) { return 0 }
func (self *AccelWithCovariance) Go_setID(id uint32) { }

