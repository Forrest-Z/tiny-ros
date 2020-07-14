package nav_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
    "tiny_ros/geometry_msgs"
)


type GridCells struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_cell_width float32 `json:"cell_width"`
    Go_cell_height float32 `json:"cell_height"`
    Go_cells []*geometry_msgs.Point `json:"cells"`
}

func NewGridCells() (*GridCells) {
    newGridCells := new(GridCells)
    newGridCells.Go_header = std_msgs.NewHeader()
    newGridCells.Go_cell_width = 0.0
    newGridCells.Go_cell_height = 0.0
    newGridCells.Go_cells = []*geometry_msgs.Point{}
    return newGridCells
}

func (self *GridCells) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_cell_width = 0.0
    self.Go_cell_height = 0.0
    self.Go_cells = []*geometry_msgs.Point{}
}

func (self *GridCells) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_cell_width := math.Float32bits(self.Go_cell_width)
    binary.LittleEndian.PutUint32(buff[offset:], bits_cell_width)
    offset += 4
    bits_cell_height := math.Float32bits(self.Go_cell_height)
    binary.LittleEndian.PutUint32(buff[offset:], bits_cell_height)
    offset += 4
    length_cells := len(self.Go_cells)
    buff[offset + 0] = byte((length_cells >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_cells >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_cells >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_cells >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_cells; i++ {
        offset += self.Go_cells[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *GridCells) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_cell_width := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_cell_width = math.Float32frombits(bits_cell_width)
    offset += 4
    bits_cell_height := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_cell_height = math.Float32frombits(bits_cell_height)
    offset += 4
    length_cells := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_cells |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_cells |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_cells |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_cells = make([]*geometry_msgs.Point, length_cells)
    for i := 0; i < length_cells; i++ {
        self.Go_cells[i] = geometry_msgs.NewPoint()
    }
    for i := 0; i < length_cells; i++ {
        offset += self.Go_cells[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *GridCells) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    length_cells := len(self.Go_cells)
    for i := 0; i < length_cells; i++ {
        length += self.Go_cells[i].Go_serializedLength()
    }
    return length
}

func (self *GridCells) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GridCells) Go_getType() (string) { return "nav_msgs/GridCells" }
func (self *GridCells) Go_getMD5() (string) { return "13ce9063aaf922c39d3a2207d3926427" }
func (self *GridCells) Go_getID() (uint32) { return 0 }
func (self *GridCells) Go_setID(id uint32) { }

