package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)


type ODEPhysics struct {
    Go_auto_disable_bodies bool `json:"auto_disable_bodies"`
    Go_sor_pgs_precon_iters uint32 `json:"sor_pgs_precon_iters"`
    Go_sor_pgs_iters uint32 `json:"sor_pgs_iters"`
    Go_sor_pgs_w float64 `json:"sor_pgs_w"`
    Go_sor_pgs_rms_error_tol float64 `json:"sor_pgs_rms_error_tol"`
    Go_contact_surface_layer float64 `json:"contact_surface_layer"`
    Go_contact_max_correcting_vel float64 `json:"contact_max_correcting_vel"`
    Go_cfm float64 `json:"cfm"`
    Go_erp float64 `json:"erp"`
    Go_max_contacts uint32 `json:"max_contacts"`
}

func NewODEPhysics() (*ODEPhysics) {
    newODEPhysics := new(ODEPhysics)
    newODEPhysics.Go_auto_disable_bodies = false
    newODEPhysics.Go_sor_pgs_precon_iters = 0
    newODEPhysics.Go_sor_pgs_iters = 0
    newODEPhysics.Go_sor_pgs_w = 0.0
    newODEPhysics.Go_sor_pgs_rms_error_tol = 0.0
    newODEPhysics.Go_contact_surface_layer = 0.0
    newODEPhysics.Go_contact_max_correcting_vel = 0.0
    newODEPhysics.Go_cfm = 0.0
    newODEPhysics.Go_erp = 0.0
    newODEPhysics.Go_max_contacts = 0
    return newODEPhysics
}

func (self *ODEPhysics) Go_initialize() {
    self.Go_auto_disable_bodies = false
    self.Go_sor_pgs_precon_iters = 0
    self.Go_sor_pgs_iters = 0
    self.Go_sor_pgs_w = 0.0
    self.Go_sor_pgs_rms_error_tol = 0.0
    self.Go_contact_surface_layer = 0.0
    self.Go_contact_max_correcting_vel = 0.0
    self.Go_cfm = 0.0
    self.Go_erp = 0.0
    self.Go_max_contacts = 0
}

func (self *ODEPhysics) Go_serialize(buff []byte) (int) {
    offset := 0
    if self.Go_auto_disable_bodies {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    buff[offset + 0] = byte((self.Go_sor_pgs_precon_iters >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_sor_pgs_precon_iters >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_sor_pgs_precon_iters >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_sor_pgs_precon_iters >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_sor_pgs_iters >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_sor_pgs_iters >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_sor_pgs_iters >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_sor_pgs_iters >> (8 * 3)) & 0xFF)
    offset += 4
    bits_sor_pgs_w := math.Float64bits(self.Go_sor_pgs_w)
    binary.LittleEndian.PutUint64(buff[offset:], bits_sor_pgs_w)
    offset += 8
    bits_sor_pgs_rms_error_tol := math.Float64bits(self.Go_sor_pgs_rms_error_tol)
    binary.LittleEndian.PutUint64(buff[offset:], bits_sor_pgs_rms_error_tol)
    offset += 8
    bits_contact_surface_layer := math.Float64bits(self.Go_contact_surface_layer)
    binary.LittleEndian.PutUint64(buff[offset:], bits_contact_surface_layer)
    offset += 8
    bits_contact_max_correcting_vel := math.Float64bits(self.Go_contact_max_correcting_vel)
    binary.LittleEndian.PutUint64(buff[offset:], bits_contact_max_correcting_vel)
    offset += 8
    bits_cfm := math.Float64bits(self.Go_cfm)
    binary.LittleEndian.PutUint64(buff[offset:], bits_cfm)
    offset += 8
    bits_erp := math.Float64bits(self.Go_erp)
    binary.LittleEndian.PutUint64(buff[offset:], bits_erp)
    offset += 8
    buff[offset + 0] = byte((self.Go_max_contacts >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_max_contacts >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_max_contacts >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_max_contacts >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *ODEPhysics) Go_deserialize(buff []byte) (int) {
    offset := 0
    if (buff[offset] & 0xFF) != 0 {
        self.Go_auto_disable_bodies = true
    } else {
        self.Go_auto_disable_bodies = false
    }
    offset += 1
    self.Go_sor_pgs_precon_iters = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_sor_pgs_precon_iters |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_sor_pgs_precon_iters |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_sor_pgs_precon_iters |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_sor_pgs_iters = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_sor_pgs_iters |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_sor_pgs_iters |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_sor_pgs_iters |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_sor_pgs_w := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_sor_pgs_w = math.Float64frombits(bits_sor_pgs_w)
    offset += 8
    bits_sor_pgs_rms_error_tol := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_sor_pgs_rms_error_tol = math.Float64frombits(bits_sor_pgs_rms_error_tol)
    offset += 8
    bits_contact_surface_layer := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_contact_surface_layer = math.Float64frombits(bits_contact_surface_layer)
    offset += 8
    bits_contact_max_correcting_vel := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_contact_max_correcting_vel = math.Float64frombits(bits_contact_max_correcting_vel)
    offset += 8
    bits_cfm := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_cfm = math.Float64frombits(bits_cfm)
    offset += 8
    bits_erp := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_erp = math.Float64frombits(bits_erp)
    offset += 8
    self.Go_max_contacts = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_max_contacts |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_max_contacts |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_max_contacts |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *ODEPhysics) Go_serializedLength() (int) {
    length := 0
    length += 1
    length += 4
    length += 4
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 4
    return length
}

func (self *ODEPhysics) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ODEPhysics) Go_getType() (string) { return "gazebo_msgs/ODEPhysics" }
func (self *ODEPhysics) Go_getMD5() (string) { return "67a077e58362b50f63dc189c25d01418" }
func (self *ODEPhysics) Go_getID() (uint32) { return 0 }
func (self *ODEPhysics) Go_setID(id uint32) { }

