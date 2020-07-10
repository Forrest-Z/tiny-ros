package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type LaserScan struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_angle_min float32 `json:"angle_min"`
    Go_angle_max float32 `json:"angle_max"`
    Go_angle_increment float32 `json:"angle_increment"`
    Go_time_increment float32 `json:"time_increment"`
    Go_scan_time float32 `json:"scan_time"`
    Go_range_min float32 `json:"range_min"`
    Go_range_max float32 `json:"range_max"`
    Go_ranges []float32 `json:"ranges"`
    Go_intensities []float32 `json:"intensities"`
}

func NewLaserScan() (*LaserScan) {
    newLaserScan := new(LaserScan)
    newLaserScan.Go_header = std_msgs.NewHeader()
    newLaserScan.Go_angle_min = 0.0
    newLaserScan.Go_angle_max = 0.0
    newLaserScan.Go_angle_increment = 0.0
    newLaserScan.Go_time_increment = 0.0
    newLaserScan.Go_scan_time = 0.0
    newLaserScan.Go_range_min = 0.0
    newLaserScan.Go_range_max = 0.0
    newLaserScan.Go_ranges = []float32{}
    newLaserScan.Go_intensities = []float32{}
    return newLaserScan
}

func (self *LaserScan) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_angle_min = 0.0
    self.Go_angle_max = 0.0
    self.Go_angle_increment = 0.0
    self.Go_time_increment = 0.0
    self.Go_scan_time = 0.0
    self.Go_range_min = 0.0
    self.Go_range_max = 0.0
    self.Go_ranges = []float32{}
    self.Go_intensities = []float32{}
}

func (self *LaserScan) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_angle_min := math.Float32bits(self.Go_angle_min)
    binary.LittleEndian.PutUint32(buff[offset:], bits_angle_min)
    offset += 4
    bits_angle_max := math.Float32bits(self.Go_angle_max)
    binary.LittleEndian.PutUint32(buff[offset:], bits_angle_max)
    offset += 4
    bits_angle_increment := math.Float32bits(self.Go_angle_increment)
    binary.LittleEndian.PutUint32(buff[offset:], bits_angle_increment)
    offset += 4
    bits_time_increment := math.Float32bits(self.Go_time_increment)
    binary.LittleEndian.PutUint32(buff[offset:], bits_time_increment)
    offset += 4
    bits_scan_time := math.Float32bits(self.Go_scan_time)
    binary.LittleEndian.PutUint32(buff[offset:], bits_scan_time)
    offset += 4
    bits_range_min := math.Float32bits(self.Go_range_min)
    binary.LittleEndian.PutUint32(buff[offset:], bits_range_min)
    offset += 4
    bits_range_max := math.Float32bits(self.Go_range_max)
    binary.LittleEndian.PutUint32(buff[offset:], bits_range_max)
    offset += 4
    length_ranges := len(self.Go_ranges)
    buff[offset + 0] = byte((length_ranges >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_ranges >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_ranges >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_ranges >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_ranges; i++ {
        bits_rangesi := math.Float32bits(self.Go_ranges[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_rangesi)
        offset += 4
    }
    length_intensities := len(self.Go_intensities)
    buff[offset + 0] = byte((length_intensities >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_intensities >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_intensities >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_intensities >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_intensities; i++ {
        bits_intensitiesi := math.Float32bits(self.Go_intensities[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_intensitiesi)
        offset += 4
    }
    return offset
}

func (self *LaserScan) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_angle_min := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_angle_min = math.Float32frombits(bits_angle_min)
    offset += 4
    bits_angle_max := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_angle_max = math.Float32frombits(bits_angle_max)
    offset += 4
    bits_angle_increment := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_angle_increment = math.Float32frombits(bits_angle_increment)
    offset += 4
    bits_time_increment := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_time_increment = math.Float32frombits(bits_time_increment)
    offset += 4
    bits_scan_time := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_scan_time = math.Float32frombits(bits_scan_time)
    offset += 4
    bits_range_min := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_range_min = math.Float32frombits(bits_range_min)
    offset += 4
    bits_range_max := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_range_max = math.Float32frombits(bits_range_max)
    offset += 4
    length_ranges := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_ranges |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_ranges |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_ranges |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_ranges = make([]float32, length_ranges)
    for i := 0; i < length_ranges; i++ {
        bits_rangesi := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_ranges[i] = math.Float32frombits(bits_rangesi)
        offset += 4
    }
    length_intensities := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_intensities |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_intensities |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_intensities |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_intensities = make([]float32, length_intensities)
    for i := 0; i < length_intensities; i++ {
        bits_intensitiesi := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_intensities[i] = math.Float32frombits(bits_intensitiesi)
        offset += 4
    }
    return offset
}

func (self *LaserScan) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length_ranges := len(self.Go_ranges)
    for i := 0; i < length_ranges; i++ {
        length += 4
    }
    length += 4
    length_intensities := len(self.Go_intensities)
    for i := 0; i < length_intensities; i++ {
        length += 4
    }
    return length
}

func (self *LaserScan) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LaserScan) Go_getType() (string) { return "sensor_msgs/LaserScan" }
func (self *LaserScan) Go_getMD5() (string) { return "9387943977c16b7fa134689acd87f1a2" }
func (self *LaserScan) Go_getID() (uint32) { return 0 }
func (self *LaserScan) Go_setID(id uint32) { }

