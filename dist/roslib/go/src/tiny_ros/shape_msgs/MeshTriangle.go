package shape_msgs

import (
    "encoding/json"
)


type MeshTriangle struct {
    Go_vertex_indices [3]uint32 `json:"vertex_indices"`
}

func NewMeshTriangle() (*MeshTriangle) {
    newMeshTriangle := new(MeshTriangle)
    newMeshTriangle.Go_vertex_indices = [3]uint32{0, 0, 0}
    return newMeshTriangle
}

func (self *MeshTriangle) Go_initialize() {
    self.Go_vertex_indices = [3]uint32{0, 0, 0}
}

func (self *MeshTriangle) Go_serialize(buff []byte) (int) {
    offset := 0
    for i := 0; i < 3; i++ {
        buff[offset + 0] = byte((self.Go_vertex_indices[i] >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((self.Go_vertex_indices[i] >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((self.Go_vertex_indices[i] >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((self.Go_vertex_indices[i] >> (8 * 3)) & 0xFF)
        offset += 4
    }
    return offset
}

func (self *MeshTriangle) Go_deserialize(buff []byte) (int) {
    offset := 0
    for i := 0; i < 3; i++ {
        self.Go_vertex_indices[i] = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
        self.Go_vertex_indices[i] |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
        self.Go_vertex_indices[i] |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
        self.Go_vertex_indices[i] |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
    }
    return offset
}

func (self *MeshTriangle) Go_serializedLength() (int) {
    length := 0
    for i := 0; i < 3; i++ {
        length += 4
    }
    return length
}

func (self *MeshTriangle) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *MeshTriangle) Go_getType() (string) { return "shape_msgs/MeshTriangle" }
func (self *MeshTriangle) Go_getMD5() (string) { return "01020cfeb9ad7679dd18bbd7149962ba" }
func (self *MeshTriangle) Go_getID() (uint32) { return 0 }
func (self *MeshTriangle) Go_setID(id uint32) { }

