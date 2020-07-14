package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/geometry_msgs"
    "math"
)



type GetPhysicsPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewGetPhysicsPropertiesRequest() (*GetPhysicsPropertiesRequest) {
    newGetPhysicsPropertiesRequest := new(GetPhysicsPropertiesRequest)
    newGetPhysicsPropertiesRequest.__id__ = 0
    return newGetPhysicsPropertiesRequest
}

func (self *GetPhysicsPropertiesRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *GetPhysicsPropertiesRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *GetPhysicsPropertiesRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *GetPhysicsPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetPhysicsPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPhysicsPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/GetPhysicsProperties" }
func (self *GetPhysicsPropertiesRequest) Go_getMD5() (string) { return "cba4b83a6824644ef787d2ac8bb7aa60" }
func (self *GetPhysicsPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetPhysicsPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetPhysicsPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_time_step float64 `json:"time_step"`
    Go_pause bool `json:"pause"`
    Go_max_update_rate float64 `json:"max_update_rate"`
    Go_gravity *geometry_msgs.Vector3 `json:"gravity"`
    Go_ode_config *ODEPhysics `json:"ode_config"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetPhysicsPropertiesResponse() (*GetPhysicsPropertiesResponse) {
    newGetPhysicsPropertiesResponse := new(GetPhysicsPropertiesResponse)
    newGetPhysicsPropertiesResponse.Go_time_step = 0.0
    newGetPhysicsPropertiesResponse.Go_pause = false
    newGetPhysicsPropertiesResponse.Go_max_update_rate = 0.0
    newGetPhysicsPropertiesResponse.Go_gravity = geometry_msgs.NewVector3()
    newGetPhysicsPropertiesResponse.Go_ode_config = NewODEPhysics()
    newGetPhysicsPropertiesResponse.Go_success = false
    newGetPhysicsPropertiesResponse.Go_status_message = ""
    newGetPhysicsPropertiesResponse.__id__ = 0
    return newGetPhysicsPropertiesResponse
}

func (self *GetPhysicsPropertiesResponse) Go_initialize() {
    self.Go_time_step = 0.0
    self.Go_pause = false
    self.Go_max_update_rate = 0.0
    self.Go_gravity = geometry_msgs.NewVector3()
    self.Go_ode_config = NewODEPhysics()
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetPhysicsPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    bits_time_step := math.Float64bits(self.Go_time_step)
    binary.LittleEndian.PutUint64(buff[offset:], bits_time_step)
    offset += 8
    if self.Go_pause {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    bits_max_update_rate := math.Float64bits(self.Go_max_update_rate)
    binary.LittleEndian.PutUint64(buff[offset:], bits_max_update_rate)
    offset += 8
    offset += self.Go_gravity.Go_serialize(buff[offset:])
    offset += self.Go_ode_config.Go_serialize(buff[offset:])
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

func (self *GetPhysicsPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_time_step := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_time_step = math.Float64frombits(bits_time_step)
    offset += 8
    if (buff[offset] & 0xFF) != 0 {
        self.Go_pause = true
    } else {
        self.Go_pause = false
    }
    offset += 1
    bits_max_update_rate := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_max_update_rate = math.Float64frombits(bits_max_update_rate)
    offset += 8
    offset += self.Go_gravity.Go_deserialize(buff[offset:])
    offset += self.Go_ode_config.Go_deserialize(buff[offset:])
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

func (self *GetPhysicsPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 1
    length += 8
    length += self.Go_gravity.Go_serializedLength()
    length += self.Go_ode_config.Go_serializedLength()
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetPhysicsPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPhysicsPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/GetPhysicsProperties" }
func (self *GetPhysicsPropertiesResponse) Go_getMD5() (string) { return "8f17f751935ff418006bf6b24982cb08" }
func (self *GetPhysicsPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetPhysicsPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


