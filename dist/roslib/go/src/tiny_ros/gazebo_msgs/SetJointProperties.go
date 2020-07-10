package gazebo_msgs

import (
    "encoding/json"
)



type SetJointPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_joint_name string `json:"joint_name"`
    Go_ode_joint_config *ODEJointProperties `json:"ode_joint_config"`
}

func NewSetJointPropertiesRequest() (*SetJointPropertiesRequest) {
    newSetJointPropertiesRequest := new(SetJointPropertiesRequest)
    newSetJointPropertiesRequest.Go_joint_name = ""
    newSetJointPropertiesRequest.Go_ode_joint_config = NewODEJointProperties()
    newSetJointPropertiesRequest.__id__ = 0
    return newSetJointPropertiesRequest
}

func (self *SetJointPropertiesRequest) Go_initialize() {
    self.Go_joint_name = ""
    self.Go_ode_joint_config = NewODEJointProperties()
    self.__id__ = 0
}

func (self *SetJointPropertiesRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_joint_name := len(self.Go_joint_name)
    buff[offset + 0] = byte((length_joint_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_joint_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_joint_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_joint_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_joint_name)], self.Go_joint_name)
    offset += length_joint_name
    offset += self.Go_ode_joint_config.Go_serialize(buff[offset:])
    return offset
}

func (self *SetJointPropertiesRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_joint_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_joint_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_joint_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_joint_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_joint_name = string(buff[offset:(offset+length_joint_name)])
    offset += length_joint_name
    offset += self.Go_ode_joint_config.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetJointPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length_joint_name := len(self.Go_joint_name)
    length += 4
    length += length_joint_name
    length += self.Go_ode_joint_config.Go_serializedLength()
    return length
}

func (self *SetJointPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetJointPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/SetJointProperties" }
func (self *SetJointPropertiesRequest) Go_getMD5() (string) { return "e9063603bda4e7bdd2b5530ad7705661" }
func (self *SetJointPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetJointPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetJointPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetJointPropertiesResponse() (*SetJointPropertiesResponse) {
    newSetJointPropertiesResponse := new(SetJointPropertiesResponse)
    newSetJointPropertiesResponse.Go_success = false
    newSetJointPropertiesResponse.Go_status_message = ""
    newSetJointPropertiesResponse.__id__ = 0
    return newSetJointPropertiesResponse
}

func (self *SetJointPropertiesResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetJointPropertiesResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetJointPropertiesResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetJointPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetJointPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetJointPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/SetJointProperties" }
func (self *SetJointPropertiesResponse) Go_getMD5() (string) { return "7e8650b70fd2dfc24b249dddf8f45cee" }
func (self *SetJointPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetJointPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


