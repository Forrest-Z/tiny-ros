package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)

func Go_COVARIANCE_TYPE_UNKNOWN() (uint8) { return  0 }
func Go_COVARIANCE_TYPE_APPROXIMATED() (uint8) { return  1 }
func Go_COVARIANCE_TYPE_DIAGONAL_KNOWN() (uint8) { return  2 }
func Go_COVARIANCE_TYPE_KNOWN() (uint8) { return  3 }

type NavSatFix struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_status *NavSatStatus `json:"status"`
    Go_latitude float64 `json:"latitude"`
    Go_longitude float64 `json:"longitude"`
    Go_altitude float64 `json:"altitude"`
    Go_position_covariance [9]float64 `json:"position_covariance"`
    Go_position_covariance_type uint8 `json:"position_covariance_type"`
}

func NewNavSatFix() (*NavSatFix) {
    newNavSatFix := new(NavSatFix)
    newNavSatFix.Go_header = std_msgs.NewHeader()
    newNavSatFix.Go_status = NewNavSatStatus()
    newNavSatFix.Go_latitude = 0.0
    newNavSatFix.Go_longitude = 0.0
    newNavSatFix.Go_altitude = 0.0
    newNavSatFix.Go_position_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newNavSatFix.Go_position_covariance_type = 0
    return newNavSatFix
}

func (self *NavSatFix) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_status = NewNavSatStatus()
    self.Go_latitude = 0.0
    self.Go_longitude = 0.0
    self.Go_altitude = 0.0
    self.Go_position_covariance = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_position_covariance_type = 0
}

func (self *NavSatFix) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_status.Go_serialize(buff[offset:])
    bits_latitude := math.Float64bits(self.Go_latitude)
    binary.LittleEndian.PutUint64(buff[offset:], bits_latitude)
    offset += 8
    bits_longitude := math.Float64bits(self.Go_longitude)
    binary.LittleEndian.PutUint64(buff[offset:], bits_longitude)
    offset += 8
    bits_altitude := math.Float64bits(self.Go_altitude)
    binary.LittleEndian.PutUint64(buff[offset:], bits_altitude)
    offset += 8
    for i := 0; i < 9; i++ {
        bits_position_covariancei := math.Float64bits(self.Go_position_covariance[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_position_covariancei)
        offset += 8
    }
    buff[offset + 0] = byte((self.Go_position_covariance_type >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *NavSatFix) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_status.Go_deserialize(buff[offset:])
    bits_latitude := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_latitude = math.Float64frombits(bits_latitude)
    offset += 8
    bits_longitude := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_longitude = math.Float64frombits(bits_longitude)
    offset += 8
    bits_altitude := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_altitude = math.Float64frombits(bits_altitude)
    offset += 8
    for i := 0; i < 9; i++ {
        bits_position_covariancei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_position_covariance[i] = math.Float64frombits(bits_position_covariancei)
        offset += 8
    }
    self.Go_position_covariance_type = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    return offset
}

func (self *NavSatFix) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_status.Go_serializedLength()
    length += 8
    length += 8
    length += 8
    for i := 0; i < 9; i++ {
        length += 8
    }
    length += 1
    return length
}

func (self *NavSatFix) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *NavSatFix) Go_getType() (string) { return "sensor_msgs/NavSatFix" }
func (self *NavSatFix) Go_getMD5() (string) { return "adc27ca9d8634ed087021b82f3c43576" }
func (self *NavSatFix) Go_getID() (uint32) { return 0 }
func (self *NavSatFix) Go_setID(id uint32) { }

