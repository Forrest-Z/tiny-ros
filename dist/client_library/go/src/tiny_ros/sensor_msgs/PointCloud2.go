package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type PointCloud2 struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_height uint32 `json:"height"`
    Go_width uint32 `json:"width"`
    Go_fields []*PointField `json:"fields"`
    Go_is_bigendian bool `json:"is_bigendian"`
    Go_point_step uint32 `json:"point_step"`
    Go_row_step uint32 `json:"row_step"`
    Go_data []uint8 `json:"data"`
    Go_is_dense bool `json:"is_dense"`
}

func NewPointCloud2() (*PointCloud2) {
    newPointCloud2 := new(PointCloud2)
    newPointCloud2.Go_header = std_msgs.NewHeader()
    newPointCloud2.Go_height = 0
    newPointCloud2.Go_width = 0
    newPointCloud2.Go_fields = []*PointField{}
    newPointCloud2.Go_is_bigendian = false
    newPointCloud2.Go_point_step = 0
    newPointCloud2.Go_row_step = 0
    newPointCloud2.Go_data = []uint8{}
    newPointCloud2.Go_is_dense = false
    return newPointCloud2
}

func (self *PointCloud2) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_height = 0
    self.Go_width = 0
    self.Go_fields = []*PointField{}
    self.Go_is_bigendian = false
    self.Go_point_step = 0
    self.Go_row_step = 0
    self.Go_data = []uint8{}
    self.Go_is_dense = false
}

func (self *PointCloud2) Go_serialize(buff []byte) (int) {
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
    length_fields := len(self.Go_fields)
    buff[offset + 0] = byte((length_fields >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_fields >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_fields >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_fields >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_fields; i++ {
        offset += self.Go_fields[i].Go_serialize(buff[offset:])
    }
    if self.Go_is_bigendian {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    buff[offset + 0] = byte((self.Go_point_step >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_point_step >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_point_step >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_point_step >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_row_step >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_row_step >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_row_step >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_row_step >> (8 * 3)) & 0xFF)
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
    if self.Go_is_dense {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *PointCloud2) Go_deserialize(buff []byte) (int) {
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
    length_fields := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_fields |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_fields |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_fields |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_fields = make([]*PointField, length_fields)
    for i := 0; i < length_fields; i++ {
        self.Go_fields[i] = NewPointField()
    }
    for i := 0; i < length_fields; i++ {
        offset += self.Go_fields[i].Go_deserialize(buff[offset:])
    }
    if (buff[offset] & 0xFF) != 0 {
        self.Go_is_bigendian = true
    } else {
        self.Go_is_bigendian = false
    }
    offset += 1
    self.Go_point_step = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_point_step |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_point_step |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_point_step |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_row_step = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_row_step |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_row_step |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_row_step |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
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
    if (buff[offset] & 0xFF) != 0 {
        self.Go_is_dense = true
    } else {
        self.Go_is_dense = false
    }
    offset += 1
    return offset
}

func (self *PointCloud2) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    length_fields := len(self.Go_fields)
    for i := 0; i < length_fields; i++ {
        length += self.Go_fields[i].Go_serializedLength()
    }
    length += 1
    length += 4
    length += 4
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    length += 1
    return length
}

func (self *PointCloud2) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *PointCloud2) Go_getType() (string) { return "sensor_msgs/PointCloud2" }
func (self *PointCloud2) Go_getMD5() (string) { return "6aa926339b282463281af40546b3b3cf" }
func (self *PointCloud2) Go_getID() (uint32) { return 0 }
func (self *PointCloud2) Go_setID(id uint32) { }

