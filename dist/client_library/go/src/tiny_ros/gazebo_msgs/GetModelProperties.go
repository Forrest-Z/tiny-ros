package gazebo_msgs

import (
    "encoding/json"
)



type GetModelPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
}

func NewGetModelPropertiesRequest() (*GetModelPropertiesRequest) {
    newGetModelPropertiesRequest := new(GetModelPropertiesRequest)
    newGetModelPropertiesRequest.Go_model_name = ""
    newGetModelPropertiesRequest.__id__ = 0
    return newGetModelPropertiesRequest
}

func (self *GetModelPropertiesRequest) Go_initialize() {
    self.Go_model_name = ""
    self.__id__ = 0
}

func (self *GetModelPropertiesRequest) Go_serialize(buff []byte) (int) {
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
    return offset
}

func (self *GetModelPropertiesRequest) Go_deserialize(buff []byte) (int) {
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
    return offset
}

func (self *GetModelPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    return length
}

func (self *GetModelPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetModelPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/GetModelProperties" }
func (self *GetModelPropertiesRequest) Go_getMD5() (string) { return "fe0194bf75c917c89b820b09c12fe6c1" }
func (self *GetModelPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetModelPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetModelPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_parent_model_name string `json:"parent_model_name"`
    Go_canonical_body_name string `json:"canonical_body_name"`
    Go_body_names []string `json:"body_names"`
    Go_geom_names []string `json:"geom_names"`
    Go_joint_names []string `json:"joint_names"`
    Go_child_model_names []string `json:"child_model_names"`
    Go_is_static bool `json:"is_static"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetModelPropertiesResponse() (*GetModelPropertiesResponse) {
    newGetModelPropertiesResponse := new(GetModelPropertiesResponse)
    newGetModelPropertiesResponse.Go_parent_model_name = ""
    newGetModelPropertiesResponse.Go_canonical_body_name = ""
    newGetModelPropertiesResponse.Go_body_names = []string{}
    newGetModelPropertiesResponse.Go_geom_names = []string{}
    newGetModelPropertiesResponse.Go_joint_names = []string{}
    newGetModelPropertiesResponse.Go_child_model_names = []string{}
    newGetModelPropertiesResponse.Go_is_static = false
    newGetModelPropertiesResponse.Go_success = false
    newGetModelPropertiesResponse.Go_status_message = ""
    newGetModelPropertiesResponse.__id__ = 0
    return newGetModelPropertiesResponse
}

func (self *GetModelPropertiesResponse) Go_initialize() {
    self.Go_parent_model_name = ""
    self.Go_canonical_body_name = ""
    self.Go_body_names = []string{}
    self.Go_geom_names = []string{}
    self.Go_joint_names = []string{}
    self.Go_child_model_names = []string{}
    self.Go_is_static = false
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetModelPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_parent_model_name := len(self.Go_parent_model_name)
    buff[offset + 0] = byte((length_parent_model_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_parent_model_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_parent_model_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_parent_model_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_parent_model_name)], self.Go_parent_model_name)
    offset += length_parent_model_name
    length_canonical_body_name := len(self.Go_canonical_body_name)
    buff[offset + 0] = byte((length_canonical_body_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_canonical_body_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_canonical_body_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_canonical_body_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_canonical_body_name)], self.Go_canonical_body_name)
    offset += length_canonical_body_name
    length_body_names := len(self.Go_body_names)
    buff[offset + 0] = byte((length_body_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_body_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_body_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_body_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_body_names; i++ {
        length_body_namesi := len(self.Go_body_names[i])
        buff[offset + 0] = byte((length_body_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_body_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_body_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_body_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_body_namesi)], self.Go_body_names[i])
        offset += length_body_namesi
    }
    length_geom_names := len(self.Go_geom_names)
    buff[offset + 0] = byte((length_geom_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_geom_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_geom_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_geom_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_geom_names; i++ {
        length_geom_namesi := len(self.Go_geom_names[i])
        buff[offset + 0] = byte((length_geom_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_geom_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_geom_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_geom_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_geom_namesi)], self.Go_geom_names[i])
        offset += length_geom_namesi
    }
    length_joint_names := len(self.Go_joint_names)
    buff[offset + 0] = byte((length_joint_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_joint_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_joint_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_joint_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := len(self.Go_joint_names[i])
        buff[offset + 0] = byte((length_joint_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_joint_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_joint_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_joint_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_joint_namesi)], self.Go_joint_names[i])
        offset += length_joint_namesi
    }
    length_child_model_names := len(self.Go_child_model_names)
    buff[offset + 0] = byte((length_child_model_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_child_model_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_child_model_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_child_model_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_child_model_names; i++ {
        length_child_model_namesi := len(self.Go_child_model_names[i])
        buff[offset + 0] = byte((length_child_model_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_child_model_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_child_model_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_child_model_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_child_model_namesi)], self.Go_child_model_names[i])
        offset += length_child_model_namesi
    }
    if self.Go_is_static {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
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

func (self *GetModelPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_parent_model_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_parent_model_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_parent_model_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_parent_model_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_parent_model_name = string(buff[offset:(offset+length_parent_model_name)])
    offset += length_parent_model_name
    length_canonical_body_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_canonical_body_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_canonical_body_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_canonical_body_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_canonical_body_name = string(buff[offset:(offset+length_canonical_body_name)])
    offset += length_canonical_body_name
    length_body_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_body_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_body_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_body_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_body_names = make([]string, length_body_names)
    for i := 0; i < length_body_names; i++ {
        length_body_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_body_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_body_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_body_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_body_names[i] = string(buff[offset:(offset+length_body_namesi)])
        offset += length_body_namesi
    }
    length_geom_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_geom_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_geom_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_geom_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_geom_names = make([]string, length_geom_names)
    for i := 0; i < length_geom_names; i++ {
        length_geom_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_geom_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_geom_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_geom_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_geom_names[i] = string(buff[offset:(offset+length_geom_namesi)])
        offset += length_geom_namesi
    }
    length_joint_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_joint_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_joint_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_joint_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_joint_names = make([]string, length_joint_names)
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_joint_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_joint_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_joint_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_joint_names[i] = string(buff[offset:(offset+length_joint_namesi)])
        offset += length_joint_namesi
    }
    length_child_model_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_child_model_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_child_model_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_child_model_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_child_model_names = make([]string, length_child_model_names)
    for i := 0; i < length_child_model_names; i++ {
        length_child_model_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_child_model_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_child_model_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_child_model_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_child_model_names[i] = string(buff[offset:(offset+length_child_model_namesi)])
        offset += length_child_model_namesi
    }
    if (buff[offset] & 0xFF) != 0 {
        self.Go_is_static = true
    } else {
        self.Go_is_static = false
    }
    offset += 1
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

func (self *GetModelPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length_parent_model_name := len(self.Go_parent_model_name)
    length += 4
    length += length_parent_model_name
    length_canonical_body_name := len(self.Go_canonical_body_name)
    length += 4
    length += length_canonical_body_name
    length += 4
    length_body_names := len(self.Go_body_names)
    for i := 0; i < length_body_names; i++ {
        length_body_namesi := len(self.Go_body_names[i])
        length += 4
        length += length_body_namesi
    }
    length += 4
    length_geom_names := len(self.Go_geom_names)
    for i := 0; i < length_geom_names; i++ {
        length_geom_namesi := len(self.Go_geom_names[i])
        length += 4
        length += length_geom_namesi
    }
    length += 4
    length_joint_names := len(self.Go_joint_names)
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := len(self.Go_joint_names[i])
        length += 4
        length += length_joint_namesi
    }
    length += 4
    length_child_model_names := len(self.Go_child_model_names)
    for i := 0; i < length_child_model_names; i++ {
        length_child_model_namesi := len(self.Go_child_model_names[i])
        length += 4
        length += length_child_model_namesi
    }
    length += 1
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetModelPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetModelPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/GetModelProperties" }
func (self *GetModelPropertiesResponse) Go_getMD5() (string) { return "d8f16b08abaf0220a551cf9025748602" }
func (self *GetModelPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetModelPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


