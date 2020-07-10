package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)



type SetModelConfigurationRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
    Go_urdf_param_name string `json:"urdf_param_name"`
    Go_joint_names []string `json:"joint_names"`
    Go_joint_positions []float64 `json:"joint_positions"`
}

func NewSetModelConfigurationRequest() (*SetModelConfigurationRequest) {
    newSetModelConfigurationRequest := new(SetModelConfigurationRequest)
    newSetModelConfigurationRequest.Go_model_name = ""
    newSetModelConfigurationRequest.Go_urdf_param_name = ""
    newSetModelConfigurationRequest.Go_joint_names = []string{}
    newSetModelConfigurationRequest.Go_joint_positions = []float64{}
    newSetModelConfigurationRequest.__id__ = 0
    return newSetModelConfigurationRequest
}

func (self *SetModelConfigurationRequest) Go_initialize() {
    self.Go_model_name = ""
    self.Go_urdf_param_name = ""
    self.Go_joint_names = []string{}
    self.Go_joint_positions = []float64{}
    self.__id__ = 0
}

func (self *SetModelConfigurationRequest) Go_serialize(buff []byte) (int) {
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
    length_urdf_param_name := len(self.Go_urdf_param_name)
    buff[offset + 0] = byte((length_urdf_param_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_urdf_param_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_urdf_param_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_urdf_param_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_urdf_param_name)], self.Go_urdf_param_name)
    offset += length_urdf_param_name
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
    length_joint_positions := len(self.Go_joint_positions)
    buff[offset + 0] = byte((length_joint_positions >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_joint_positions >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_joint_positions >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_joint_positions >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_joint_positions; i++ {
        bits_joint_positionsi := math.Float64bits(self.Go_joint_positions[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_joint_positionsi)
        offset += 8
    }
    return offset
}

func (self *SetModelConfigurationRequest) Go_deserialize(buff []byte) (int) {
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
    length_urdf_param_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_urdf_param_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_urdf_param_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_urdf_param_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_urdf_param_name = string(buff[offset:(offset+length_urdf_param_name)])
    offset += length_urdf_param_name
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
    length_joint_positions := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_joint_positions |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_joint_positions |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_joint_positions |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_joint_positions = make([]float64, length_joint_positions)
    for i := 0; i < length_joint_positions; i++ {
        bits_joint_positionsi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_joint_positions[i] = math.Float64frombits(bits_joint_positionsi)
        offset += 8
    }
    return offset
}

func (self *SetModelConfigurationRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    length_urdf_param_name := len(self.Go_urdf_param_name)
    length += 4
    length += length_urdf_param_name
    length += 4
    length_joint_names := len(self.Go_joint_names)
    for i := 0; i < length_joint_names; i++ {
        length_joint_namesi := len(self.Go_joint_names[i])
        length += 4
        length += length_joint_namesi
    }
    length += 4
    length_joint_positions := len(self.Go_joint_positions)
    for i := 0; i < length_joint_positions; i++ {
        length += 8
    }
    return length
}

func (self *SetModelConfigurationRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetModelConfigurationRequest) Go_getType() (string) { return "gazebo_msgs/SetModelConfiguration" }
func (self *SetModelConfigurationRequest) Go_getMD5() (string) { return "74db6184ae83468b540d4c02d244ada7" }
func (self *SetModelConfigurationRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetModelConfigurationRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetModelConfigurationResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetModelConfigurationResponse() (*SetModelConfigurationResponse) {
    newSetModelConfigurationResponse := new(SetModelConfigurationResponse)
    newSetModelConfigurationResponse.Go_success = false
    newSetModelConfigurationResponse.Go_status_message = ""
    newSetModelConfigurationResponse.__id__ = 0
    return newSetModelConfigurationResponse
}

func (self *SetModelConfigurationResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetModelConfigurationResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetModelConfigurationResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetModelConfigurationResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetModelConfigurationResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetModelConfigurationResponse) Go_getType() (string) { return "gazebo_msgs/SetModelConfiguration" }
func (self *SetModelConfigurationResponse) Go_getMD5() (string) { return "6f12aefa315c8b37040d5d47471e39ee" }
func (self *SetModelConfigurationResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetModelConfigurationResponse) Go_setID(id uint32) { self.__id__ = id }


