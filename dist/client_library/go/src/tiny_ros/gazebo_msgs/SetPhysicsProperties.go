package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/geometry_msgs"
    "math"
)



type SetPhysicsPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_time_step float64 `json:"time_step"`
    Go_max_update_rate float64 `json:"max_update_rate"`
    Go_gravity *geometry_msgs.Vector3 `json:"gravity"`
    Go_ode_config *ODEPhysics `json:"ode_config"`
}

func NewSetPhysicsPropertiesRequest() (*SetPhysicsPropertiesRequest) {
    newSetPhysicsPropertiesRequest := new(SetPhysicsPropertiesRequest)
    newSetPhysicsPropertiesRequest.Go_time_step = 0.0
    newSetPhysicsPropertiesRequest.Go_max_update_rate = 0.0
    newSetPhysicsPropertiesRequest.Go_gravity = geometry_msgs.NewVector3()
    newSetPhysicsPropertiesRequest.Go_ode_config = NewODEPhysics()
    newSetPhysicsPropertiesRequest.__id__ = 0
    return newSetPhysicsPropertiesRequest
}

func (self *SetPhysicsPropertiesRequest) Go_initialize() {
    self.Go_time_step = 0.0
    self.Go_max_update_rate = 0.0
    self.Go_gravity = geometry_msgs.NewVector3()
    self.Go_ode_config = NewODEPhysics()
    self.__id__ = 0
}

func (self *SetPhysicsPropertiesRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    bits_time_step := math.Float64bits(self.Go_time_step)
    binary.LittleEndian.PutUint64(buff[offset:], bits_time_step)
    offset += 8
    bits_max_update_rate := math.Float64bits(self.Go_max_update_rate)
    binary.LittleEndian.PutUint64(buff[offset:], bits_max_update_rate)
    offset += 8
    offset += self.Go_gravity.Go_serialize(buff[offset:])
    offset += self.Go_ode_config.Go_serialize(buff[offset:])
    return offset
}

func (self *SetPhysicsPropertiesRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_time_step := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_time_step = math.Float64frombits(bits_time_step)
    offset += 8
    bits_max_update_rate := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_max_update_rate = math.Float64frombits(bits_max_update_rate)
    offset += 8
    offset += self.Go_gravity.Go_deserialize(buff[offset:])
    offset += self.Go_ode_config.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetPhysicsPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 8
    length += self.Go_gravity.Go_serializedLength()
    length += self.Go_ode_config.Go_serializedLength()
    return length
}

func (self *SetPhysicsPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetPhysicsPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/SetPhysicsProperties" }
func (self *SetPhysicsPropertiesRequest) Go_getMD5() (string) { return "373e5371b10119be0a008429a9660679" }
func (self *SetPhysicsPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetPhysicsPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetPhysicsPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetPhysicsPropertiesResponse() (*SetPhysicsPropertiesResponse) {
    newSetPhysicsPropertiesResponse := new(SetPhysicsPropertiesResponse)
    newSetPhysicsPropertiesResponse.Go_success = false
    newSetPhysicsPropertiesResponse.Go_status_message = ""
    newSetPhysicsPropertiesResponse.__id__ = 0
    return newSetPhysicsPropertiesResponse
}

func (self *SetPhysicsPropertiesResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetPhysicsPropertiesResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetPhysicsPropertiesResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetPhysicsPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetPhysicsPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetPhysicsPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/SetPhysicsProperties" }
func (self *SetPhysicsPropertiesResponse) Go_getMD5() (string) { return "5b1d14bf828ba119319cc03e2bb3473a" }
func (self *SetPhysicsPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetPhysicsPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


