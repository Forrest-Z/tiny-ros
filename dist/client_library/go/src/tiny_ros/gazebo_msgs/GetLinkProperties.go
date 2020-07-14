package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/geometry_msgs"
    "math"
)



type GetLinkPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_link_name string `json:"link_name"`
}

func NewGetLinkPropertiesRequest() (*GetLinkPropertiesRequest) {
    newGetLinkPropertiesRequest := new(GetLinkPropertiesRequest)
    newGetLinkPropertiesRequest.Go_link_name = ""
    newGetLinkPropertiesRequest.__id__ = 0
    return newGetLinkPropertiesRequest
}

func (self *GetLinkPropertiesRequest) Go_initialize() {
    self.Go_link_name = ""
    self.__id__ = 0
}

func (self *GetLinkPropertiesRequest) Go_serialize(buff []byte) (int) {
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
    return offset
}

func (self *GetLinkPropertiesRequest) Go_deserialize(buff []byte) (int) {
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
    return offset
}

func (self *GetLinkPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length_link_name := len(self.Go_link_name)
    length += 4
    length += length_link_name
    return length
}

func (self *GetLinkPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetLinkPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/GetLinkProperties" }
func (self *GetLinkPropertiesRequest) Go_getMD5() (string) { return "30b187ce76c283d51c190e8f7d59f0ff" }
func (self *GetLinkPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetLinkPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetLinkPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_com *geometry_msgs.Pose `json:"com"`
    Go_gravity_mode bool `json:"gravity_mode"`
    Go_mass float64 `json:"mass"`
    Go_ixx float64 `json:"ixx"`
    Go_ixy float64 `json:"ixy"`
    Go_ixz float64 `json:"ixz"`
    Go_iyy float64 `json:"iyy"`
    Go_iyz float64 `json:"iyz"`
    Go_izz float64 `json:"izz"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetLinkPropertiesResponse() (*GetLinkPropertiesResponse) {
    newGetLinkPropertiesResponse := new(GetLinkPropertiesResponse)
    newGetLinkPropertiesResponse.Go_com = geometry_msgs.NewPose()
    newGetLinkPropertiesResponse.Go_gravity_mode = false
    newGetLinkPropertiesResponse.Go_mass = 0.0
    newGetLinkPropertiesResponse.Go_ixx = 0.0
    newGetLinkPropertiesResponse.Go_ixy = 0.0
    newGetLinkPropertiesResponse.Go_ixz = 0.0
    newGetLinkPropertiesResponse.Go_iyy = 0.0
    newGetLinkPropertiesResponse.Go_iyz = 0.0
    newGetLinkPropertiesResponse.Go_izz = 0.0
    newGetLinkPropertiesResponse.Go_success = false
    newGetLinkPropertiesResponse.Go_status_message = ""
    newGetLinkPropertiesResponse.__id__ = 0
    return newGetLinkPropertiesResponse
}

func (self *GetLinkPropertiesResponse) Go_initialize() {
    self.Go_com = geometry_msgs.NewPose()
    self.Go_gravity_mode = false
    self.Go_mass = 0.0
    self.Go_ixx = 0.0
    self.Go_ixy = 0.0
    self.Go_ixz = 0.0
    self.Go_iyy = 0.0
    self.Go_iyz = 0.0
    self.Go_izz = 0.0
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetLinkPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
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

func (self *GetLinkPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
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

func (self *GetLinkPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_com.Go_serializedLength()
    length += 1
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetLinkPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetLinkPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/GetLinkProperties" }
func (self *GetLinkPropertiesResponse) Go_getMD5() (string) { return "d45a9e2f72bfb95f519b2d0cbaac4512" }
func (self *GetLinkPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetLinkPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


