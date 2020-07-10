package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type CompressedImage struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_format string `json:"format"`
    Go_data []uint8 `json:"data"`
}

func NewCompressedImage() (*CompressedImage) {
    newCompressedImage := new(CompressedImage)
    newCompressedImage.Go_header = std_msgs.NewHeader()
    newCompressedImage.Go_format = ""
    newCompressedImage.Go_data = []uint8{}
    return newCompressedImage
}

func (self *CompressedImage) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_format = ""
    self.Go_data = []uint8{}
}

func (self *CompressedImage) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_format := len(self.Go_format)
    buff[offset + 0] = byte((length_format >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_format >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_format >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_format >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_format)], self.Go_format)
    offset += length_format
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

func (self *CompressedImage) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_format := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_format |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_format |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_format |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_format = string(buff[offset:(offset+length_format)])
    offset += length_format
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

func (self *CompressedImage) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length_format := len(self.Go_format)
    length += 4
    length += length_format
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    return length
}

func (self *CompressedImage) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *CompressedImage) Go_getType() (string) { return "sensor_msgs/CompressedImage" }
func (self *CompressedImage) Go_getMD5() (string) { return "eed57d856457441995644e6294152301" }
func (self *CompressedImage) Go_getID() (uint32) { return 0 }
func (self *CompressedImage) Go_setID(id uint32) { }

