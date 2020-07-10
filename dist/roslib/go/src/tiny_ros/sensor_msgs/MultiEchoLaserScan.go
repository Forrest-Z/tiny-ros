package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type MultiEchoLaserScan struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_angle_min float32 `json:"angle_min"`
    Go_angle_max float32 `json:"angle_max"`
    Go_angle_increment float32 `json:"angle_increment"`
    Go_time_increment float32 `json:"time_increment"`
    Go_scan_time float32 `json:"scan_time"`
    Go_range_min float32 `json:"range_min"`
    Go_range_max float32 `json:"range_max"`
    Go_ranges []*LaserEcho `json:"ranges"`
    Go_intensities []*LaserEcho `json:"intensities"`
}

func NewMultiEchoLaserScan() (*MultiEchoLaserScan) {
    newMultiEchoLaserScan := new(MultiEchoLaserScan)
    newMultiEchoLaserScan.Go_header = std_msgs.NewHeader()
    newMultiEchoLaserScan.Go_angle_min = 0.0
    newMultiEchoLaserScan.Go_angle_max = 0.0
    newMultiEchoLaserScan.Go_angle_increment = 0.0
    newMultiEchoLaserScan.Go_time_increment = 0.0
    newMultiEchoLaserScan.Go_scan_time = 0.0
    newMultiEchoLaserScan.Go_range_min = 0.0
    newMultiEchoLaserScan.Go_range_max = 0.0
    newMultiEchoLaserScan.Go_ranges = []*LaserEcho{}
    newMultiEchoLaserScan.Go_intensities = []*LaserEcho{}
    return newMultiEchoLaserScan
}

func (self *MultiEchoLaserScan) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_angle_min = 0.0
    self.Go_angle_max = 0.0
    self.Go_angle_increment = 0.0
    self.Go_time_increment = 0.0
    self.Go_scan_time = 0.0
    self.Go_range_min = 0.0
    self.Go_range_max = 0.0
    self.Go_ranges = []*LaserEcho{}
    self.Go_intensities = []*LaserEcho{}
}

func (self *MultiEchoLaserScan) Go_serialize(buff []byte) (int) {
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
        offset += self.Go_ranges[i].Go_serialize(buff[offset:])
    }
    length_intensities := len(self.Go_intensities)
    buff[offset + 0] = byte((length_intensities >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_intensities >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_intensities >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_intensities >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_intensities; i++ {
        offset += self.Go_intensities[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *MultiEchoLaserScan) Go_deserialize(buff []byte) (int) {
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
    self.Go_ranges = make([]*LaserEcho, length_ranges)
    for i := 0; i < length_ranges; i++ {
        self.Go_ranges[i] = NewLaserEcho()
    }
    for i := 0; i < length_ranges; i++ {
        offset += self.Go_ranges[i].Go_deserialize(buff[offset:])
    }
    length_intensities := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_intensities |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_intensities |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_intensities |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_intensities = make([]*LaserEcho, length_intensities)
    for i := 0; i < length_intensities; i++ {
        self.Go_intensities[i] = NewLaserEcho()
    }
    for i := 0; i < length_intensities; i++ {
        offset += self.Go_intensities[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *MultiEchoLaserScan) Go_serializedLength() (int) {
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
        length += self.Go_ranges[i].Go_serializedLength()
    }
    length += 4
    length_intensities := len(self.Go_intensities)
    for i := 0; i < length_intensities; i++ {
        length += self.Go_intensities[i].Go_serializedLength()
    }
    return length
}

func (self *MultiEchoLaserScan) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MultiEchoLaserScan) Go_getType() (string) { return "sensor_msgs/MultiEchoLaserScan" }
func (self *MultiEchoLaserScan) Go_getMD5() (string) { return "92f3933b4fa486e3889b461437899bf5" }
func (self *MultiEchoLaserScan) Go_getID() (uint32) { return 0 }
func (self *MultiEchoLaserScan) Go_setID(id uint32) { }

