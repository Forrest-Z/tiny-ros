package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type ODEJointProperties struct {
    Go_damping []float64 `json:"damping"`
    Go_hiStop []float64 `json:"hiStop"`
    Go_loStop []float64 `json:"loStop"`
    Go_erp []float64 `json:"erp"`
    Go_cfm []float64 `json:"cfm"`
    Go_stop_erp []float64 `json:"stop_erp"`
    Go_stop_cfm []float64 `json:"stop_cfm"`
    Go_fudge_factor []float64 `json:"fudge_factor"`
    Go_fmax []float64 `json:"fmax"`
    Go_vel []float64 `json:"vel"`
}

func NewODEJointProperties() (*ODEJointProperties) {
    newODEJointProperties := new(ODEJointProperties)
    newODEJointProperties.Go_damping = []float64{}
    newODEJointProperties.Go_hiStop = []float64{}
    newODEJointProperties.Go_loStop = []float64{}
    newODEJointProperties.Go_erp = []float64{}
    newODEJointProperties.Go_cfm = []float64{}
    newODEJointProperties.Go_stop_erp = []float64{}
    newODEJointProperties.Go_stop_cfm = []float64{}
    newODEJointProperties.Go_fudge_factor = []float64{}
    newODEJointProperties.Go_fmax = []float64{}
    newODEJointProperties.Go_vel = []float64{}
    return newODEJointProperties
}

func (self *ODEJointProperties) Go_initialize() {
    self.Go_damping = []float64{}
    self.Go_hiStop = []float64{}
    self.Go_loStop = []float64{}
    self.Go_erp = []float64{}
    self.Go_cfm = []float64{}
    self.Go_stop_erp = []float64{}
    self.Go_stop_cfm = []float64{}
    self.Go_fudge_factor = []float64{}
    self.Go_fmax = []float64{}
    self.Go_vel = []float64{}
}

func (self *ODEJointProperties) Go_serialize(buff []byte) (int) {
    offset := 0
    length_damping := len(self.Go_damping)
    buff[offset + 0] = byte((length_damping >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_damping >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_damping >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_damping >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_damping; i++ {
        bits_dampingi := math.Float64bits(self.Go_damping[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_dampingi)
        offset += 8
    }
    length_hiStop := len(self.Go_hiStop)
    buff[offset + 0] = byte((length_hiStop >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_hiStop >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_hiStop >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_hiStop >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_hiStop; i++ {
        bits_hiStopi := math.Float64bits(self.Go_hiStop[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_hiStopi)
        offset += 8
    }
    length_loStop := len(self.Go_loStop)
    buff[offset + 0] = byte((length_loStop >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_loStop >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_loStop >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_loStop >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_loStop; i++ {
        bits_loStopi := math.Float64bits(self.Go_loStop[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_loStopi)
        offset += 8
    }
    length_erp := len(self.Go_erp)
    buff[offset + 0] = byte((length_erp >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_erp >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_erp >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_erp >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_erp; i++ {
        bits_erpi := math.Float64bits(self.Go_erp[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_erpi)
        offset += 8
    }
    length_cfm := len(self.Go_cfm)
    buff[offset + 0] = byte((length_cfm >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_cfm >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_cfm >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_cfm >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_cfm; i++ {
        bits_cfmi := math.Float64bits(self.Go_cfm[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_cfmi)
        offset += 8
    }
    length_stop_erp := len(self.Go_stop_erp)
    buff[offset + 0] = byte((length_stop_erp >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_stop_erp >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_stop_erp >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_stop_erp >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_stop_erp; i++ {
        bits_stop_erpi := math.Float64bits(self.Go_stop_erp[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_stop_erpi)
        offset += 8
    }
    length_stop_cfm := len(self.Go_stop_cfm)
    buff[offset + 0] = byte((length_stop_cfm >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_stop_cfm >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_stop_cfm >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_stop_cfm >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_stop_cfm; i++ {
        bits_stop_cfmi := math.Float64bits(self.Go_stop_cfm[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_stop_cfmi)
        offset += 8
    }
    length_fudge_factor := len(self.Go_fudge_factor)
    buff[offset + 0] = byte((length_fudge_factor >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_fudge_factor >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_fudge_factor >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_fudge_factor >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_fudge_factor; i++ {
        bits_fudge_factori := math.Float64bits(self.Go_fudge_factor[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_fudge_factori)
        offset += 8
    }
    length_fmax := len(self.Go_fmax)
    buff[offset + 0] = byte((length_fmax >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_fmax >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_fmax >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_fmax >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_fmax; i++ {
        bits_fmaxi := math.Float64bits(self.Go_fmax[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_fmaxi)
        offset += 8
    }
    length_vel := len(self.Go_vel)
    buff[offset + 0] = byte((length_vel >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_vel >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_vel >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_vel >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_vel; i++ {
        bits_veli := math.Float64bits(self.Go_vel[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_veli)
        offset += 8
    }
    return offset
}

func (self *ODEJointProperties) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_damping := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_damping |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_damping |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_damping |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_damping = make([]float64, length_damping)
    for i := 0; i < length_damping; i++ {
        bits_dampingi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_damping[i] = math.Float64frombits(bits_dampingi)
        offset += 8
    }
    length_hiStop := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_hiStop |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_hiStop |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_hiStop |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_hiStop = make([]float64, length_hiStop)
    for i := 0; i < length_hiStop; i++ {
        bits_hiStopi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_hiStop[i] = math.Float64frombits(bits_hiStopi)
        offset += 8
    }
    length_loStop := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_loStop |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_loStop |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_loStop |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_loStop = make([]float64, length_loStop)
    for i := 0; i < length_loStop; i++ {
        bits_loStopi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_loStop[i] = math.Float64frombits(bits_loStopi)
        offset += 8
    }
    length_erp := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_erp |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_erp |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_erp |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_erp = make([]float64, length_erp)
    for i := 0; i < length_erp; i++ {
        bits_erpi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_erp[i] = math.Float64frombits(bits_erpi)
        offset += 8
    }
    length_cfm := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_cfm |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_cfm |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_cfm |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_cfm = make([]float64, length_cfm)
    for i := 0; i < length_cfm; i++ {
        bits_cfmi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_cfm[i] = math.Float64frombits(bits_cfmi)
        offset += 8
    }
    length_stop_erp := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_stop_erp |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_stop_erp |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_stop_erp |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stop_erp = make([]float64, length_stop_erp)
    for i := 0; i < length_stop_erp; i++ {
        bits_stop_erpi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_stop_erp[i] = math.Float64frombits(bits_stop_erpi)
        offset += 8
    }
    length_stop_cfm := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_stop_cfm |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_stop_cfm |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_stop_cfm |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stop_cfm = make([]float64, length_stop_cfm)
    for i := 0; i < length_stop_cfm; i++ {
        bits_stop_cfmi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_stop_cfm[i] = math.Float64frombits(bits_stop_cfmi)
        offset += 8
    }
    length_fudge_factor := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_fudge_factor |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_fudge_factor |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_fudge_factor |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_fudge_factor = make([]float64, length_fudge_factor)
    for i := 0; i < length_fudge_factor; i++ {
        bits_fudge_factori := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_fudge_factor[i] = math.Float64frombits(bits_fudge_factori)
        offset += 8
    }
    length_fmax := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_fmax |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_fmax |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_fmax |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_fmax = make([]float64, length_fmax)
    for i := 0; i < length_fmax; i++ {
        bits_fmaxi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_fmax[i] = math.Float64frombits(bits_fmaxi)
        offset += 8
    }
    length_vel := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_vel |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_vel |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_vel |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_vel = make([]float64, length_vel)
    for i := 0; i < length_vel; i++ {
        bits_veli := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_vel[i] = math.Float64frombits(bits_veli)
        offset += 8
    }
    return offset
}

func (self *ODEJointProperties) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_damping := len(self.Go_damping)
    for i := 0; i < length_damping; i++ {
        length += 8
    }
    length += 4
    length_hiStop := len(self.Go_hiStop)
    for i := 0; i < length_hiStop; i++ {
        length += 8
    }
    length += 4
    length_loStop := len(self.Go_loStop)
    for i := 0; i < length_loStop; i++ {
        length += 8
    }
    length += 4
    length_erp := len(self.Go_erp)
    for i := 0; i < length_erp; i++ {
        length += 8
    }
    length += 4
    length_cfm := len(self.Go_cfm)
    for i := 0; i < length_cfm; i++ {
        length += 8
    }
    length += 4
    length_stop_erp := len(self.Go_stop_erp)
    for i := 0; i < length_stop_erp; i++ {
        length += 8
    }
    length += 4
    length_stop_cfm := len(self.Go_stop_cfm)
    for i := 0; i < length_stop_cfm; i++ {
        length += 8
    }
    length += 4
    length_fudge_factor := len(self.Go_fudge_factor)
    for i := 0; i < length_fudge_factor; i++ {
        length += 8
    }
    length += 4
    length_fmax := len(self.Go_fmax)
    for i := 0; i < length_fmax; i++ {
        length += 8
    }
    length += 4
    length_vel := len(self.Go_vel)
    for i := 0; i < length_vel; i++ {
        length += 8
    }
    return length
}

func (self *ODEJointProperties) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ODEJointProperties) Go_getType() (string) { return "gazebo_msgs/ODEJointProperties" }
func (self *ODEJointProperties) Go_getMD5() (string) { return "a9e264dbf3eff8e202d2bebecf081639" }
func (self *ODEJointProperties) Go_getID() (uint32) { return 0 }
func (self *ODEJointProperties) Go_setID(id uint32) { }

