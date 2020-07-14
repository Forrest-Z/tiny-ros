package stereo_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/sensor_msgs"
    "encoding/binary"
    "math"
)


type DisparityImage struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_image *sensor_msgs.Image `json:"image"`
    Go_f float32 `json:"f"`
    Go_T float32 `json:"T"`
    Go_valid_window *sensor_msgs.RegionOfInterest `json:"valid_window"`
    Go_min_disparity float32 `json:"min_disparity"`
    Go_max_disparity float32 `json:"max_disparity"`
    Go_delta_d float32 `json:"delta_d"`
}

func NewDisparityImage() (*DisparityImage) {
    newDisparityImage := new(DisparityImage)
    newDisparityImage.Go_header = std_msgs.NewHeader()
    newDisparityImage.Go_image = sensor_msgs.NewImage()
    newDisparityImage.Go_f = 0.0
    newDisparityImage.Go_T = 0.0
    newDisparityImage.Go_valid_window = sensor_msgs.NewRegionOfInterest()
    newDisparityImage.Go_min_disparity = 0.0
    newDisparityImage.Go_max_disparity = 0.0
    newDisparityImage.Go_delta_d = 0.0
    return newDisparityImage
}

func (self *DisparityImage) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_image = sensor_msgs.NewImage()
    self.Go_f = 0.0
    self.Go_T = 0.0
    self.Go_valid_window = sensor_msgs.NewRegionOfInterest()
    self.Go_min_disparity = 0.0
    self.Go_max_disparity = 0.0
    self.Go_delta_d = 0.0
}

func (self *DisparityImage) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_image.Go_serialize(buff[offset:])
    bits_f := math.Float32bits(self.Go_f)
    binary.LittleEndian.PutUint32(buff[offset:], bits_f)
    offset += 4
    bits_T := math.Float32bits(self.Go_T)
    binary.LittleEndian.PutUint32(buff[offset:], bits_T)
    offset += 4
    offset += self.Go_valid_window.Go_serialize(buff[offset:])
    bits_min_disparity := math.Float32bits(self.Go_min_disparity)
    binary.LittleEndian.PutUint32(buff[offset:], bits_min_disparity)
    offset += 4
    bits_max_disparity := math.Float32bits(self.Go_max_disparity)
    binary.LittleEndian.PutUint32(buff[offset:], bits_max_disparity)
    offset += 4
    bits_delta_d := math.Float32bits(self.Go_delta_d)
    binary.LittleEndian.PutUint32(buff[offset:], bits_delta_d)
    offset += 4
    return offset
}

func (self *DisparityImage) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_image.Go_deserialize(buff[offset:])
    bits_f := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_f = math.Float32frombits(bits_f)
    offset += 4
    bits_T := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_T = math.Float32frombits(bits_T)
    offset += 4
    offset += self.Go_valid_window.Go_deserialize(buff[offset:])
    bits_min_disparity := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_min_disparity = math.Float32frombits(bits_min_disparity)
    offset += 4
    bits_max_disparity := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_max_disparity = math.Float32frombits(bits_max_disparity)
    offset += 4
    bits_delta_d := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_delta_d = math.Float32frombits(bits_delta_d)
    offset += 4
    return offset
}

func (self *DisparityImage) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_image.Go_serializedLength()
    length += 4
    length += 4
    length += self.Go_valid_window.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    return length
}

func (self *DisparityImage) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *DisparityImage) Go_getType() (string) { return "stereo_msgs/DisparityImage" }
func (self *DisparityImage) Go_getMD5() (string) { return "03545cef8df8d20bea21fdbbf9482b4b" }
func (self *DisparityImage) Go_getID() (uint32) { return 0 }
func (self *DisparityImage) Go_setID(id uint32) { }

