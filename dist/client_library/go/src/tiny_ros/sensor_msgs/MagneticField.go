package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/geometry_msgs"
    "encoding/binary"
    "math"
)


type MagneticField struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_magnetic_field *geometry_msgs.Vector3 `json:"magnetic_field"`
    Go_magnetic_field_covariance [9]float64 `json:"magnetic_field_covariance"`
}

func NewMagneticField() (*MagneticField) {
    newMagneticField := new(MagneticField)
    newMagneticField.Go_header = std_msgs.NewHeader()
    newMagneticField.Go_magnetic_field = geometry_msgs.NewVector3()
    newMagneticField.Go_magnetic_field_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    return newMagneticField
}

func (self *MagneticField) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_magnetic_field = geometry_msgs.NewVector3()
    self.Go_magnetic_field_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (self *MagneticField) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_magnetic_field.Go_serialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_magnetic_field_covariancei := math.Float64bits(self.Go_magnetic_field_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_magnetic_field_covariancei)
        offset += 8
    }
    return offset
}

func (self *MagneticField) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_magnetic_field.Go_deserialize(buff[offset:])
    for i := 0; i < 9; i++ {
        bits_magnetic_field_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_magnetic_field_covariance[i] = math.Float64frombits(bits_magnetic_field_covariancei)
        offset += 8
    }
    return offset
}

func (self *MagneticField) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_magnetic_field.Go_serializedLength()
    for i := 0; i < 9; i++ {
        length += 8
    }
    return length
}

func (self *MagneticField) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MagneticField) Go_getType() (string) { return "sensor_msgs/MagneticField" }
func (self *MagneticField) Go_getMD5() (string) { return "f8e051d776de1349146122759c66db92" }
func (self *MagneticField) Go_getID() (uint32) { return 0 }
func (self *MagneticField) Go_setID(id uint32) { }

