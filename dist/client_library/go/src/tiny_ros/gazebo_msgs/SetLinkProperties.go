package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/geometry_msgs"
    "math"
)



type SetLinkPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_link_name string `json:"link_name"`
    Go_com *geometry_msgs.Pose `json:"com"`
    Go_gravity_mode bool `json:"gravity_mode"`
    Go_mass float64 `json:"mass"`
    Go_ixx float64 `json:"ixx"`
    Go_ixy float64 `json:"ixy"`
    Go_ixz float64 `json:"ixz"`
    Go_iyy float64 `json:"iyy"`
    Go_iyz float64 `json:"iyz"`
    Go_izz float64 `json:"izz"`
}

func NewSetLinkPropertiesRequest() (*SetLinkPropertiesRequest) {
    newSetLinkPropertiesRequest := new(SetLinkPropertiesRequest)
    newSetLinkPropertiesRequest.Go_link_name = ""
    newSetLinkPropertiesRequest.Go_com = geometry_msgs.NewPose()
    newSetLinkPropertiesRequest.Go_gravity_mode = false
    newSetLinkPropertiesRequest.Go_mass = 0.0
    newSetLinkPropertiesRequest.Go_ixx = 0.0
    newSetLinkPropertiesRequest.Go_ixy = 0.0
    newSetLinkPropertiesRequest.Go_ixz = 0.0
    newSetLinkPropertiesRequest.Go_iyy = 0.0
    newSetLinkPropertiesRequest.Go_iyz = 0.0
    newSetLinkPropertiesRequest.Go_izz = 0.0
    newSetLinkPropertiesRequest.__id__ = 0
    return newSetLinkPropertiesRequest
}

func (self *SetLinkPropertiesRequest) Go_initialize() {
    self.Go_link_name = ""
    self.Go_com = geometry_msgs.NewPose()
    self.Go_gravity_mode = false
    self.Go_mass = 0.0
    self.Go_ixx = 0.0
    self.Go_ixy = 0.0
    self.Go_ixz = 0.0
    self.Go_iyy = 0.0
    self.Go_iyz = 0.0
    self.Go_izz = 0.0
    self.__id__ = 0
}

func (self *SetLinkPropertiesRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_link_name := len(self.Go_link_name)
    buff[offset + 0] = byte((length_link_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_link_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_link_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_link_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_link_name)], self.Go_link_name)
    offset += length_link_name
    offset += self.Go_com.Go_serialize(buff[offset:])
    if self.Go_gravity_mode {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    bits_mass := math.Float64bits(self.Go_mass)
    binary.LittleEndian.PutUint64(buff[offset:], bits_mass)
    offset += 8
    bits_ixx := math.Float64bits(self.Go_ixx)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixx)
    offset += 8
    bits_ixy := math.Float64bits(self.Go_ixy)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixy)
    offset += 8
    bits_ixz := math.Float64bits(self.Go_ixz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_ixz)
    offset += 8
    bits_iyy := math.Float64bits(self.Go_iyy)
    binary.LittleEndian.PutUint64(buff[offset:], bits_iyy)
    offset += 8
    bits_iyz := math.Float64bits(self.Go_iyz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_iyz)
    offset += 8
    bits_izz := math.Float64bits(self.Go_izz)
    binary.LittleEndian.PutUint64(buff[offset:], bits_izz)
    offset += 8
    return offset
}

func (self *SetLinkPropertiesRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_link_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_link_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_link_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_link_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_link_name = string(buff[offset:(offset+length_link_name)])
    offset += length_link_name
    offset += self.Go_com.Go_deserialize(buff[offset:])
    if (buff[offset] & 0xFF) != 0 {
        self.Go_gravity_mode = true
    } else {
        self.Go_gravity_mode = false
    }
    offset += 1
    bits_mass := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_mass = math.Float64frombits(bits_mass)
    offset += 8
    bits_ixx := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixx = math.Float64frombits(bits_ixx)
    offset += 8
    bits_ixy := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixy = math.Float64frombits(bits_ixy)
    offset += 8
    bits_ixz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_ixz = math.Float64frombits(bits_ixz)
    offset += 8
    bits_iyy := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_iyy = math.Float64frombits(bits_iyy)
    offset += 8
    bits_iyz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_iyz = math.Float64frombits(bits_iyz)
    offset += 8
    bits_izz := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_izz = math.Float64frombits(bits_izz)
    offset += 8
    return offset
}

func (self *SetLinkPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length_link_name := len(self.Go_link_name)
    length += 4
    length += length_link_name
    length += self.Go_com.Go_serializedLength()
    length += 1
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length
}

func (self *SetLinkPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetLinkPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/SetLinkProperties" }
func (self *SetLinkPropertiesRequest) Go_getMD5() (string) { return "03d7e308476601832a9e1ce9d7eab722" }
func (self *SetLinkPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetLinkPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetLinkPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetLinkPropertiesResponse() (*SetLinkPropertiesResponse) {
    newSetLinkPropertiesResponse := new(SetLinkPropertiesResponse)
    newSetLinkPropertiesResponse.Go_success = false
    newSetLinkPropertiesResponse.Go_status_message = ""
    newSetLinkPropertiesResponse.__id__ = 0
    return newSetLinkPropertiesResponse
}

func (self *SetLinkPropertiesResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetLinkPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    if self.Go_success {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    length_status_message := len(self.Go_status_message)
    buff[offset + 0] = byte((length_status_message >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_status_message >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_status_message >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_status_message >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_status_message)], self.Go_status_message)
    offset += length_status_message
    return offset
}

func (self *SetLinkPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    if (buff[offset] & 0xFF) != 0 {
        self.Go_success = true
    } else {
        self.Go_success = false
    }
    offset += 1
    length_status_message := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_status_message |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_status_message |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_status_message |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_status_message = string(buff[offset:(offset+length_status_message)])
    offset += length_status_message
    return offset
}

func (self *SetLinkPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetLinkPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetLinkPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/SetLinkProperties" }
func (self *SetLinkPropertiesResponse) Go_getMD5() (string) { return "777dea05607e1bca37e3206f97801d89" }
func (self *SetLinkPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetLinkPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


