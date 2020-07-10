package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type Image struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_height uint32 `json:"height"`
    Go_width uint32 `json:"width"`
    Go_encoding string `json:"encoding"`
    Go_is_bigendian uint8 `json:"is_bigendian"`
    Go_step uint32 `json:"step"`
    Go_data []uint8 `json:"data"`
}

func NewImage() (*Image) {
    newImage := new(Image)
    newImage.Go_header = std_msgs.NewHeader()
    newImage.Go_height = 0
    newImage.Go_width = 0
    newImage.Go_encoding = ""
    newImage.Go_is_bigendian = 0
    newImage.Go_step = 0
    newImage.Go_data = []uint8{}
    return newImage
}

func (self *Image) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_height = 0
    self.Go_width = 0
    self.Go_encoding = ""
    self.Go_is_bigendian = 0
    self.Go_step = 0
    self.Go_data = []uint8{}
}

func (self *Image) Go_serialize(buff []byte) (int) {
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
    length_encoding := len(self.Go_encoding)
    buff[offset + 0] = byte((length_encoding >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_encoding >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_encoding >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_encoding >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_encoding)], self.Go_encoding)
    offset += length_encoding
    buff[offset + 0] = byte((self.Go_is_bigendian >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_step >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_step >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_step >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_step >> (8 * 3)) & 0xFF)
    offset += 4
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_data; i++ {
        buff[offset + 0] = byte((self.Go_data[i] >> (8 * 0)) & 0xFF)
        offset += 1
    }
    return offset
}

func (self *Image) Go_deserialize(buff []byte) (int) {
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
    length_encoding := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_encoding |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_encoding |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_encoding |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_encoding = string(buff[offset:(offset+length_encoding)])
    offset += length_encoding
    self.Go_is_bigendian = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_step = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_step |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_step |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_step |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]uint8, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
        offset += 1
    }
    return offset
}

func (self *Image) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length_encoding := len(self.Go_encoding)
    length += 4
    length += length_encoding
    length += 1
    length += 4
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    return length
}

func (self *Image) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Image) Go_getType() (string) { return "sensor_msgs/Image" }
func (self *Image) Go_getMD5() (string) { return "886f928dc81bf7f1496a8b452057c5b2" }
func (self *Image) Go_getID() (uint32) { return 0 }
func (self *Image) Go_setID(id uint32) { }

