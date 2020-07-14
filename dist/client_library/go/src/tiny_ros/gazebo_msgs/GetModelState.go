package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)



type GetModelStateRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
    Go_relative_entity_name string `json:"relative_entity_name"`
}

func NewGetModelStateRequest() (*GetModelStateRequest) {
    newGetModelStateRequest := new(GetModelStateRequest)
    newGetModelStateRequest.Go_model_name = ""
    newGetModelStateRequest.Go_relative_entity_name = ""
    newGetModelStateRequest.__id__ = 0
    return newGetModelStateRequest
}

func (self *GetModelStateRequest) Go_initialize() {
    self.Go_model_name = ""
    self.Go_relative_entity_name = ""
    self.__id__ = 0
}

func (self *GetModelStateRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_model_name := len(self.Go_model_name)
    buff[offset + 0] = byte((length_model_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_model_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_model_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_model_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_model_name)], self.Go_model_name)
    offset += length_model_name
    length_relative_entity_name := len(self.Go_relative_entity_name)
    buff[offset + 0] = byte((length_relative_entity_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_relative_entity_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_relative_entity_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_relative_entity_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_relative_entity_name)], self.Go_relative_entity_name)
    offset += length_relative_entity_name
    return offset
}

func (self *GetModelStateRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_model_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_model_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_model_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_model_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_model_name = string(buff[offset:(offset+length_model_name)])
    offset += length_model_name
    length_relative_entity_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_relative_entity_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_relative_entity_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_relative_entity_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_relative_entity_name = string(buff[offset:(offset+length_relative_entity_name)])
    offset += length_relative_entity_name
    return offset
}

func (self *GetModelStateRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    length_relative_entity_name := len(self.Go_relative_entity_name)
    length += 4
    length += length_relative_entity_name
    return length
}

func (self *GetModelStateRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetModelStateRequest) Go_getType() (string) { return "gazebo_msgs/GetModelState" }
func (self *GetModelStateRequest) Go_getMD5() (string) { return "92a8c6443abf59a40e396c81c0e29d40" }
func (self *GetModelStateRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetModelStateRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetModelStateResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_pose *geometry_msgs.Pose `json:"pose"`
    Go_twist *geometry_msgs.Twist `json:"twist"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetModelStateResponse() (*GetModelStateResponse) {
    newGetModelStateResponse := new(GetModelStateResponse)
    newGetModelStateResponse.Go_pose = geometry_msgs.NewPose()
    newGetModelStateResponse.Go_twist = geometry_msgs.NewTwist()
    newGetModelStateResponse.Go_success = false
    newGetModelStateResponse.Go_status_message = ""
    newGetModelStateResponse.__id__ = 0
    return newGetModelStateResponse
}

func (self *GetModelStateResponse) Go_initialize() {
    self.Go_pose = geometry_msgs.NewPose()
    self.Go_twist = geometry_msgs.NewTwist()
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetModelStateResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_pose.Go_serialize(buff[offset:])
    offset += self.Go_twist.Go_serialize(buff[offset:])
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

func (self *GetModelStateResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_pose.Go_deserialize(buff[offset:])
    offset += self.Go_twist.Go_deserialize(buff[offset:])
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

func (self *GetModelStateResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_pose.Go_serializedLength()
    length += self.Go_twist.Go_serializedLength()
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetModelStateResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetModelStateResponse) Go_getType() (string) { return "gazebo_msgs/GetModelState" }
func (self *GetModelStateResponse) Go_getMD5() (string) { return "3fd873975bc823929b01f7473704b974" }
func (self *GetModelStateResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetModelStateResponse) Go_setID(id uint32) { self.__id__ = id }


