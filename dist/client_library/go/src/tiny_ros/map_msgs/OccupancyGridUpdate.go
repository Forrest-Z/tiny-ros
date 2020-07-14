package map_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type OccupancyGridUpdate struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_x int32 `json:"x"`
    Go_y int32 `json:"y"`
    Go_width uint32 `json:"width"`
    Go_height uint32 `json:"height"`
    Go_data []int8 `json:"data"`
}

func NewOccupancyGridUpdate() (*OccupancyGridUpdate) {
    newOccupancyGridUpdate := new(OccupancyGridUpdate)
    newOccupancyGridUpdate.Go_header = std_msgs.NewHeader()
    newOccupancyGridUpdate.Go_x = 0
    newOccupancyGridUpdate.Go_y = 0
    newOccupancyGridUpdate.Go_width = 0
    newOccupancyGridUpdate.Go_height = 0
    newOccupancyGridUpdate.Go_data = []int8{}
    return newOccupancyGridUpdate
}

func (self *OccupancyGridUpdate) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_x = 0
    self.Go_y = 0
    self.Go_width = 0
    self.Go_height = 0
    self.Go_data = []int8{}
}

func (self *OccupancyGridUpdate) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_x >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_x >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_x >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_x >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_y >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_y >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_y >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_y >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_width >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_width >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_width >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_width >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_height >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_height >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_height >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_height >> (8 * 3)) & 0xFF)
    offset += 4
    length_data := len(self.Go_data)
    buff[offset + 0] = byte((length_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_data >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_data; i++ {
        buff[offset + 0] = byte(uint8(self.Go_data[i] >> (8 * 0)) & 0xFF)
        offset += 1
    }
    return offset
}

func (self *OccupancyGridUpdate) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_x = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_x |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_x |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_x |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_y = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_y |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_y |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_y |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_width = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_width |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_width |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_width |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_height = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_height |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_height |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_height |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data = make([]int8, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = int8(buff[offset + 0] & 0xFF) << (8 * 0)
        offset += 1
    }
    return offset
}

func (self *OccupancyGridUpdate) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    return length
}

func (self *OccupancyGridUpdate) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *OccupancyGridUpdate) Go_getType() (string) { return "map_msgs/OccupancyGridUpdate" }
func (self *OccupancyGridUpdate) Go_getMD5() (string) { return "159b2d7856932f2e2cad9b082ed99ec2" }
func (self *OccupancyGridUpdate) Go_getID() (uint32) { return 0 }
func (self *OccupancyGridUpdate) Go_setID(id uint32) { }

