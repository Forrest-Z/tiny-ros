package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)



type GetWorldPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewGetWorldPropertiesRequest() (*GetWorldPropertiesRequest) {
    newGetWorldPropertiesRequest := new(GetWorldPropertiesRequest)
    newGetWorldPropertiesRequest.__id__ = 0
    return newGetWorldPropertiesRequest
}

func (self *GetWorldPropertiesRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *GetWorldPropertiesRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *GetWorldPropertiesRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *GetWorldPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *GetWorldPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetWorldPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/GetWorldProperties" }
func (self *GetWorldPropertiesRequest) Go_getMD5() (string) { return "3aa5de7106eec5dae41ad1c9ae681123" }
func (self *GetWorldPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetWorldPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetWorldPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_sim_time float64 `json:"sim_time"`
    Go_model_names []string `json:"model_names"`
    Go_rendering_enabled bool `json:"rendering_enabled"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetWorldPropertiesResponse() (*GetWorldPropertiesResponse) {
    newGetWorldPropertiesResponse := new(GetWorldPropertiesResponse)
    newGetWorldPropertiesResponse.Go_sim_time = 0.0
    newGetWorldPropertiesResponse.Go_model_names = []string{}
    newGetWorldPropertiesResponse.Go_rendering_enabled = false
    newGetWorldPropertiesResponse.Go_success = false
    newGetWorldPropertiesResponse.Go_status_message = ""
    newGetWorldPropertiesResponse.__id__ = 0
    return newGetWorldPropertiesResponse
}

func (self *GetWorldPropertiesResponse) Go_initialize() {
    self.Go_sim_time = 0.0
    self.Go_model_names = []string{}
    self.Go_rendering_enabled = false
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetWorldPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    bits_sim_time := math.Float64bits(self.Go_sim_time)
    binary.LittleEndian.PutUint64(buff[offset:], bits_sim_time)
    offset += 8
    length_model_names := len(self.Go_model_names)
    buff[offset + 0] = byte((length_model_names >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_model_names >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_model_names >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_model_names >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_model_names; i++ {
        length_model_namesi := len(self.Go_model_names[i])
        buff[offset + 0] = byte((length_model_namesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_model_namesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_model_namesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_model_namesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_model_namesi)], self.Go_model_names[i])
        offset += length_model_namesi
    }
    if self.Go_rendering_enabled {
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

func (self *GetWorldPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    bits_sim_time := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_sim_time = math.Float64frombits(bits_sim_time)
    offset += 8
    length_model_names := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_model_names |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_model_names |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_model_names |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_model_names = make([]string, length_model_names)
    for i := 0; i < length_model_names; i++ {
        length_model_namesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_model_namesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_model_namesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_model_namesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_model_names[i] = string(buff[offset:(offset+length_model_namesi)])
        offset += length_model_namesi
    }
    if (buff[offset] & 0xFF) != 0 {
        self.Go_rendering_enabled = true
    } else {
        self.Go_rendering_enabled = false
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

func (self *GetWorldPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 8
    length += 4
    length_model_names := len(self.Go_model_names)
    for i := 0; i < length_model_names; i++ {
        length_model_namesi := len(self.Go_model_names[i])
        length += 4
        length += length_model_namesi
    }
    length += 1
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetWorldPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetWorldPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/GetWorldProperties" }
func (self *GetWorldPropertiesResponse) Go_getMD5() (string) { return "fe944c1c210919291ad14bc43b6c10cf" }
func (self *GetWorldPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetWorldPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


