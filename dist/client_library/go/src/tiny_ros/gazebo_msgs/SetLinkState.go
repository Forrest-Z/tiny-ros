package gazebo_msgs

import (
    "encoding/json"
)



type SetLinkStateRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_link_state *LinkState `json:"link_state"`
}

func NewSetLinkStateRequest() (*SetLinkStateRequest) {
    newSetLinkStateRequest := new(SetLinkStateRequest)
    newSetLinkStateRequest.Go_link_state = NewLinkState()
    newSetLinkStateRequest.__id__ = 0
    return newSetLinkStateRequest
}

func (self *SetLinkStateRequest) Go_initialize() {
    self.Go_link_state = NewLinkState()
    self.__id__ = 0
}

func (self *SetLinkStateRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_link_state.Go_serialize(buff[offset:])
    return offset
}

func (self *SetLinkStateRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_link_state.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetLinkStateRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_link_state.Go_serializedLength()
    return length
}

func (self *SetLinkStateRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetLinkStateRequest) Go_getType() (string) { return "gazebo_msgs/SetLinkState" }
func (self *SetLinkStateRequest) Go_getMD5() (string) { return "09e18d697b151a98db1e31fb5e89444f" }
func (self *SetLinkStateRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetLinkStateRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetLinkStateResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetLinkStateResponse() (*SetLinkStateResponse) {
    newSetLinkStateResponse := new(SetLinkStateResponse)
    newSetLinkStateResponse.Go_success = false
    newSetLinkStateResponse.Go_status_message = ""
    newSetLinkStateResponse.__id__ = 0
    return newSetLinkStateResponse
}

func (self *SetLinkStateResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetLinkStateResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetLinkStateResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetLinkStateResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetLinkStateResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetLinkStateResponse) Go_getType() (string) { return "gazebo_msgs/SetLinkState" }
func (self *SetLinkStateResponse) Go_getMD5() (string) { return "daf8f1df20722a4f504ceb200701e404" }
func (self *SetLinkStateResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetLinkStateResponse) Go_setID(id uint32) { self.__id__ = id }


