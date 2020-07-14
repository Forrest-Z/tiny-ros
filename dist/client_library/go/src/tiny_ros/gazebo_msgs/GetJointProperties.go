package gazebo_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
)



type GetJointPropertiesRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_joint_name string `json:"joint_name"`
}

func NewGetJointPropertiesRequest() (*GetJointPropertiesRequest) {
    newGetJointPropertiesRequest := new(GetJointPropertiesRequest)
    newGetJointPropertiesRequest.Go_joint_name = ""
    newGetJointPropertiesRequest.__id__ = 0
    return newGetJointPropertiesRequest
}

func (self *GetJointPropertiesRequest) Go_initialize() {
    self.Go_joint_name = ""
    self.__id__ = 0
}

func (self *GetJointPropertiesRequest) Go_serialize(buff []byte) (int) {
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
    return offset
}

func (self *GetJointPropertiesRequest) Go_deserialize(buff []byte) (int) {
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
    return offset
}

func (self *GetJointPropertiesRequest) Go_serializedLength() (int) {
    length := 0
    length_joint_name := len(self.Go_joint_name)
    length += 4
    length += length_joint_name
    return length
}

func (self *GetJointPropertiesRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetJointPropertiesRequest) Go_getType() (string) { return "gazebo_msgs/GetJointProperties" }
func (self *GetJointPropertiesRequest) Go_getMD5() (string) { return "b07a3ce5fb5aba1cfc56577c9cb3ecc6" }
func (self *GetJointPropertiesRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetJointPropertiesRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////

func Go_REVOLUTE() (uint8) { return  0                 }
func Go_CONTINUOUS() (uint8) { return  1                 }
func Go_PRISMATIC() (uint8) { return  2                 }
func Go_FIXED() (uint8) { return  3                 }
func Go_BALL() (uint8) { return  4                 }
func Go_UNIVERSAL() (uint8) { return  5                 }

type GetJointPropertiesResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_type uint8 `json:"type"`
    Go_damping []float64 `json:"damping"`
    Go_position []float64 `json:"position"`
    Go_rate []float64 `json:"rate"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetJointPropertiesResponse() (*GetJointPropertiesResponse) {
    newGetJointPropertiesResponse := new(GetJointPropertiesResponse)
    newGetJointPropertiesResponse.Go_type = 0
    newGetJointPropertiesResponse.Go_damping = []float64{}
    newGetJointPropertiesResponse.Go_position = []float64{}
    newGetJointPropertiesResponse.Go_rate = []float64{}
    newGetJointPropertiesResponse.Go_success = false
    newGetJointPropertiesResponse.Go_status_message = ""
    newGetJointPropertiesResponse.__id__ = 0
    return newGetJointPropertiesResponse
}

func (self *GetJointPropertiesResponse) Go_initialize() {
    self.Go_type = 0
    self.Go_damping = []float64{}
    self.Go_position = []float64{}
    self.Go_rate = []float64{}
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetJointPropertiesResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_type >> (8 * 0)) & 0xFF)
    offset += 1
    length_damping := len(self.Go_damping)
    buff[offset + 0] = byte((length_damping >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_damping >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_damping >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_damping >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_damping; i++ {
        bits_dampingi := math.Float64bits(self.Go_damping[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_dampingi)
        offset += 8
    }
    length_position := len(self.Go_position)
    buff[offset + 0] = byte((length_position >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_position >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_position >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_position >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_position; i++ {
        bits_positioni := math.Float64bits(self.Go_position[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_positioni)
        offset += 8
    }
    length_rate := len(self.Go_rate)
    buff[offset + 0] = byte((length_rate >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_rate >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_rate >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_rate >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_rate; i++ {
        bits_ratei := math.Float64bits(self.Go_rate[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_ratei)
        offset += 8
    }
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

func (self *GetJointPropertiesResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_type = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_damping := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_damping |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_damping |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_damping |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_damping = make([]float64, length_damping)
    for i := 0; i < length_damping; i++ {
        bits_dampingi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_damping[i] = math.Float64frombits(bits_dampingi)
        offset += 8
    }
    length_position := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_position |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_position |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_position |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_position = make([]float64, length_position)
    for i := 0; i < length_position; i++ {
        bits_positioni := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_position[i] = math.Float64frombits(bits_positioni)
        offset += 8
    }
    length_rate := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_rate |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_rate |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_rate |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_rate = make([]float64, length_rate)
    for i := 0; i < length_rate; i++ {
        bits_ratei := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_rate[i] = math.Float64frombits(bits_ratei)
        offset += 8
    }
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

func (self *GetJointPropertiesResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length += 4
    length_damping := len(self.Go_damping)
    for i := 0; i < length_damping; i++ {
        length += 8
    }
    length += 4
    length_position := len(self.Go_position)
    for i := 0; i < length_position; i++ {
        length += 8
    }
    length += 4
    length_rate := len(self.Go_rate)
    for i := 0; i < length_rate; i++ {
        length += 8
    }
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetJointPropertiesResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetJointPropertiesResponse) Go_getType() (string) { return "gazebo_msgs/GetJointProperties" }
func (self *GetJointPropertiesResponse) Go_getMD5() (string) { return "a60fbf691ac539e1355c979ca09b4573" }
func (self *GetJointPropertiesResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetJointPropertiesResponse) Go_setID(id uint32) { self.__id__ = id }


