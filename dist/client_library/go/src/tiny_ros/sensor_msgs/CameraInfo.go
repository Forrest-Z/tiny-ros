package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type CameraInfo struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_height uint32 `json:"height"`
    Go_width uint32 `json:"width"`
    Go_distortion_model string `json:"distortion_model"`
    Go_D []float64 `json:"D"`
    Go_K [9]float64 `json:"K"`
    Go_R [9]float64 `json:"R"`
    Go_P [12]float64 `json:"P"`
    Go_binning_x uint32 `json:"binning_x"`
    Go_binning_y uint32 `json:"binning_y"`
    Go_roi *RegionOfInterest `json:"roi"`
}

func NewCameraInfo() (*CameraInfo) {
    newCameraInfo := new(CameraInfo)
    newCameraInfo.Go_header = std_msgs.NewHeader()
    newCameraInfo.Go_height = 0
    newCameraInfo.Go_width = 0
    newCameraInfo.Go_distortion_model = ""
    newCameraInfo.Go_D = []float64{}
    newCameraInfo.Go_K = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newCameraInfo.Go_R = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newCameraInfo.Go_P = [12]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    newCameraInfo.Go_binning_x = 0
    newCameraInfo.Go_binning_y = 0
    newCameraInfo.Go_roi = NewRegionOfInterest()
    return newCameraInfo
}

func (self *CameraInfo) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_height = 0
    self.Go_width = 0
    self.Go_distortion_model = ""
    self.Go_D = []float64{}
    self.Go_K = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_R = [9]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_P = [12]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
    self.Go_binning_x = 0
    self.Go_binning_y = 0
    self.Go_roi = NewRegionOfInterest()
}

func (self *CameraInfo) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_height >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_height >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_height >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_height >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_width >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_width >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_width >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_width >> (8 * 3)) & 0xFF)
    offset += 4
    length_distortion_model := len(self.Go_distortion_model)
    buff[offset + 0] = byte((length_distortion_model >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_distortion_model >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_distortion_model >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_distortion_model >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_distortion_model)], self.Go_distortion_model)
    offset += length_distortion_model
    length_D := len(self.Go_D)
    buff[offset + 0] = byte((length_D >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_D >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_D >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_D >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_D; i++ {
        bits_Di := math.Float64bits(self.Go_D[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_Di)
        offset += 8
    }
    for i := 0; i < 9; i++ {
        bits_Ki := math.Float64bits(self.Go_K[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_Ki)
        offset += 8
    }
    for i := 0; i < 9; i++ {
        bits_Ri := math.Float64bits(self.Go_R[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_Ri)
        offset += 8
    }
    for i := 0; i < 12; i++ {
        bits_Pi := math.Float64bits(self.Go_P[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_Pi)
        offset += 8
    }
    buff[offset + 0] = byte((self.Go_binning_x >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_binning_x >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_binning_x >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_binning_x >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_binning_y >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_binning_y >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_binning_y >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_binning_y >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_roi.Go_serialize(buff[offset:])
    return offset
}

func (self *CameraInfo) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_height = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_height |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_height |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_height |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_width = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_width |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_width |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_width |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_distortion_model := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_distortion_model |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_distortion_model |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_distortion_model |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_distortion_model = string(buff[offset:(offset+length_distortion_model)])
    offset += length_distortion_model
    length_D := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_D |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_D |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_D |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_D = make([]float64, length_D)
    for i := 0; i < length_D; i++ {
        bits_Di := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_D[i] = math.Float64frombits(bits_Di)
        offset += 8
    }
    for i := 0; i < 9; i++ {
        bits_Ki := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_K[i] = math.Float64frombits(bits_Ki)
        offset += 8
    }
    for i := 0; i < 9; i++ {
        bits_Ri := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_R[i] = math.Float64frombits(bits_Ri)
        offset += 8
    }
    for i := 0; i < 12; i++ {
        bits_Pi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_P[i] = math.Float64frombits(bits_Pi)
        offset += 8
    }
    self.Go_binning_x = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_binning_x |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_binning_x |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_binning_x |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_binning_y = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_binning_y |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_binning_y |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_binning_y |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_roi.Go_deserialize(buff[offset:])
    return offset
}

func (self *CameraInfo) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length_distortion_model := len(self.Go_distortion_model)
    length += 4
    length += length_distortion_model
    length += 4
    length_D := len(self.Go_D)
    for i := 0; i < length_D; i++ {
        length += 8
    }
    for i := 0; i < 9; i++ {
        length += 8
    }
    for i := 0; i < 9; i++ {
        length += 8
    }
    for i := 0; i < 12; i++ {
        length += 8
    }
    length += 4
    length += 4
    length += self.Go_roi.Go_serializedLength()
    return length
}

func (self *CameraInfo) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *CameraInfo) Go_getType() (string) { return "sensor_msgs/CameraInfo" }
func (self *CameraInfo) Go_getMD5() (string) { return "57d2553deec0a7842f00837f40032798" }
func (self *CameraInfo) Go_getID() (uint32) { return 0 }
func (self *CameraInfo) Go_setID(id uint32) { }

