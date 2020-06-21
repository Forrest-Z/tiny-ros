package nav_msgs

import (
    "tiny_ros/std_msgs"
    "nav_msgs/MapMetaData"
)

type OccupancyGrid struct {
    Go_header std_msgs.Header `json:"header"`
    Go_info nav_msgs.MapMetaData `json:"info"`
    Go_data []int8 `json:"data"`
}

func NewOccupancyGrid() (*OccupancyGrid) {
    newOccupancyGrid := new(OccupancyGrid)
    newOccupancyGrid.Go_header = std_msgs.NewHeader()
    newOccupancyGrid.Go_info = nav_msgs.NewMapMetaData()
    newOccupancyGrid.Go_data = []int8{}
    return newOccupancyGrid
}

func (self *OccupancyGrid) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_info.Go_serialize(buff[offset:])
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

func (self *OccupancyGrid) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_info.Go_deserialize(buff[offset:])
    length_data := int((buff[offset + 0] & 0xFF) << (8 * 0))
    length_data |= int((buff[offset + 1] & 0xFF) << (8 * 1))
    length_data |= int((buff[offset + 2] & 0xFF) << (8 * 2))
    length_data |= int((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    self.Go_data = make([]int8, length_data, length_data)
    for i := 0; i < length_data; i++ {
        self.Go_data[i] = int8((buff[offset + 0] & 0xFF) << (8 * 0))
        offset += 1
    }
    return offset
}

func (self *OccupancyGrid) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_info.Go_serializedLength()
    length += 4
    length_data := len(self.Go_data)
    for i := 0; i < length_data; i++ {
        length += 1
    }
    return length
}

func (self *OccupancyGrid) Go_echo() (string) { return "" }
func (self *OccupancyGrid) Go_getType() (string) { return "nav_msgs/OccupancyGrid" }
func (self *OccupancyGrid) Go_getMD5() (string) { return "e489a26457224a97799696f3642f16a0" }
func (self *OccupancyGrid) Go_getID() (uint32) { return 0 }
func (self *OccupancyGrid) Go_setID(id uint32) { }

