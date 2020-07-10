package shape_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)


type Mesh struct {
    Go_triangles []*MeshTriangle `json:"triangles"`
    Go_vertices []*geometry_msgs.Point `json:"vertices"`
}

func NewMesh() (*Mesh) {
    newMesh := new(Mesh)
    newMesh.Go_triangles = []*MeshTriangle{}
    newMesh.Go_vertices = []*geometry_msgs.Point{}
    return newMesh
}

func (self *Mesh) Go_initialize() {
    self.Go_triangles = []*MeshTriangle{}
    self.Go_vertices = []*geometry_msgs.Point{}
}

func (self *Mesh) Go_serialize(buff []byte) (int) {
    offset := 0
    length_triangles := len(self.Go_triangles)
    buff[offset + 0] = byte((length_triangles >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_triangles >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_triangles >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_triangles >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_triangles; i++ {
        offset += self.Go_triangles[i].Go_serialize(buff[offset:])
    }
    length_vertices := len(self.Go_vertices)
    buff[offset + 0] = byte((length_vertices >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_vertices >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_vertices >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_vertices >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_vertices; i++ {
        offset += self.Go_vertices[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *Mesh) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_triangles := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_triangles |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_triangles |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_triangles |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_triangles = make([]*MeshTriangle, length_triangles)
    for i := 0; i < length_triangles; i++ {
        self.Go_triangles[i] = NewMeshTriangle()
    }
    for i := 0; i < length_triangles; i++ {
        offset += self.Go_triangles[i].Go_deserialize(buff[offset:])
    }
    length_vertices := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_vertices |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_vertices |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_vertices |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_vertices = make([]*geometry_msgs.Point, length_vertices)
    for i := 0; i < length_vertices; i++ {
        self.Go_vertices[i] = geometry_msgs.NewPoint()
    }
    for i := 0; i < length_vertices; i++ {
        offset += self.Go_vertices[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *Mesh) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_triangles := len(self.Go_triangles)
    for i := 0; i < length_triangles; i++ {
        length += self.Go_triangles[i].Go_serializedLength()
    }
    length += 4
    length_vertices := len(self.Go_vertices)
    for i := 0; i < length_vertices; i++ {
        length += self.Go_vertices[i].Go_serializedLength()
    }
    return length
}

func (self *Mesh) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Mesh) Go_getType() (string) { return "shape_msgs/Mesh" }
func (self *Mesh) Go_getMD5() (string) { return "1579670b316f622b47d6700cd4f7e18d" }
func (self *Mesh) Go_getID() (uint32) { return 0 }
func (self *Mesh) Go_setID(id uint32) { }

