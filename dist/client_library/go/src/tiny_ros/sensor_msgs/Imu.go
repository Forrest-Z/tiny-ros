package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
    "encoding/binary"
    "math"
)


type Imu struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_orientation *geometry_msgs.Quaternion `json:"orientation"`
    Go_orientation_covariance [9]float64 `json:"orientation_covariance"`
    Go_angular_velocity *geometry_msgs.Vector3 `json:"angular_velocity"`
    Go_angular_velocity_covariance [9]float64 `json:"angular_velocity_covariance"`
    Go_linear_acceleration *geometry_msgs.Vector3 `json:"linear_acceleration"`
    Go_linear_acceleration_covariance [9]float64 `json:"linear_acceleration_covariance"`
}

func NewImu() (*Imu) {
    newImu := new(Imu)
    newImu.Go_header = std_msgs.NewHeader()
    newImu.Go_orientation = geometry_msgs.NewQuaternion()
    newImu.Go_orientation_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newImu.Go_angular_velocity = geometry_msgs.NewVector3()
    newImu.Go_angular_velocity_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newImu.Go_linear_acceleration = geometry_msgs.NewVector3()
    newImu.Go_linear_acceleration_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    return newImu
}

func (self *Imu) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_orientation = geometry_msgs.NewQuaternion()
    self.Go_orientation_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_angular_velocity = geometry_msgs.NewVector3()
    self.Go_angular_velocity_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_linear_acceleration = geometry_msgs.NewVector3()
    self.Go_linear_acceleration_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (self *Imu) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_orientation.Go_serialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_orientation_covariancei := math.Float64bits(self.Go_orientation_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_orientation_covariancei)
        offset += 8
    }
    offset += self.Go_angular_velocity.Go_serialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_angular_velocity_covariancei := math.Float64bits(self.Go_angular_velocity_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_angular_velocity_covariancei)
        offset += 8
    }
    offset += self.Go_linear_acceleration.Go_serialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_linear_acceleration_covariancei := math.Float64bits(self.Go_linear_acceleration_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_linear_acceleration_covariancei)
        offset += 8
    }
    return offset
}

func (self *Imu) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_orientation.Go_deserialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_orientation_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_orientation_covariance[i] = math.Float64frombits(bits_orientation_covariancei)
        offset += 8
    }
    offset += self.Go_angular_velocity.Go_deserialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_angular_velocity_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_angular_velocity_covariance[i] = math.Float64frombits(bits_angular_velocity_covariancei)
        offset += 8
    }
    offset += self.Go_linear_acceleration.Go_deserialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_linear_acceleration_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_linear_acceleration_covariance[i] = math.Float64frombits(bits_linear_acceleration_covariancei)
        offset += 8
    }
    return offset
}

func (self *Imu) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_orientation.Go_serializedLength()
    for i := 0; i < 9; i++ {
        length += 8
    }
    length += self.Go_angular_velocity.Go_serializedLength()
    for i := 0; i < 9; i++ {
        length += 8
    }
    length += self.Go_linear_acceleration.Go_serializedLength()
    for i := 0; i < 9; i++ {
        length += 8
    }
    return length
}

func (self *Imu) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Imu) Go_getType() (string) { return "sensor_msgs/Imu" }
func (self *Imu) Go_getMD5() (string) { return "a42c1ab94665a5807834c0ea19a6d16a" }
func (self *Imu) Go_getID() (uint32) { return 0 }
func (self *Imu) Go_setID(id uint32) { }

