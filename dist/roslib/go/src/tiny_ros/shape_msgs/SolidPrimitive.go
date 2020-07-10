package shape_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)

func Go_BOX() (uint8) { return 1 }
func Go_SPHERE() (uint8) { return 2 }
func Go_CYLINDER() (uint8) { return 3 }
func Go_CONE() (uint8) { return 4 }
func Go_BOX_X() (uint8) { return 0 }
func Go_BOX_Y() (uint8) { return 1 }
func Go_BOX_Z() (uint8) { return 2 }
func Go_SPHERE_RADIUS() (uint8) { return 0 }
func Go_CYLINDER_HEIGHT() (uint8) { return 0 }
func Go_CYLINDER_RADIUS() (uint8) { return 1 }
func Go_CONE_HEIGHT() (uint8) { return 0 }
func Go_CONE_RADIUS() (uint8) { return 1 }

type SolidPrimitive struct {
    Go_type uint8 `json:"type"`
    Go_dimensions []float64 `json:"dimensions"`
}

func NewSolidPrimitive() (*SolidPrimitive) {
    newSolidPrimitive := new(SolidPrimitive)
    newSolidPrimitive.Go_type = 0
    newSolidPrimitive.Go_dimensions = []float64{}
    return newSolidPrimitive
}

func (self *SolidPrimitive) Go_initialize() {
    self.Go_type = 0
    self.Go_dimensions = []float64{}
}

func (self *SolidPrimitive) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_type >> (8 * 0)) & 0xFF)
    offset += 1
    length_dimensions := len(self.Go_dimensions)
    buff[offset + 0] = byte((length_dimensions >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_dimensions >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_dimensions >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_dimensions >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_dimensions; i++ {
        bits_dimensionsi := math.Float64bits(self.Go_dimensions[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_dimensionsi)
        offset += 8
    }
    return offset
}

func (self *SolidPrimitive) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_type = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_dimensions := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_dimensions |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_dimensions |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_dimensions |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_dimensions = make([]float64, length_dimensions)
    for i := 0; i < length_dimensions; i++ {
        bits_dimensionsi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_dimensions[i] = math.Float64frombits(bits_dimensionsi)
        offset += 8
    }
    return offset
}

func (self *SolidPrimitive) Go_serializedLength() (int) {
    length := 0
    length += 1
    length += 4
    length_dimensions := len(self.Go_dimensions)
    for i := 0; i < length_dimensions; i++ {
        length += 8
    }
    return length
}

func (self *SolidPrimitive) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SolidPrimitive) Go_getType() (string) { return "shape_msgs/SolidPrimitive" }
func (self *SolidPrimitive) Go_getMD5() (string) { return "f40805922065416be24909fd8fd751b5" }
func (self *SolidPrimitive) Go_getID() (uint32) { return 0 }
func (self *SolidPrimitive) Go_setID(id uint32) { }

