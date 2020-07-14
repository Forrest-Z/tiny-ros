package gazebo_msgs

import (
    "encoding/json"
)



type GetLinkStateRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_link_name string `json:"link_name"`
    Go_reference_frame string `json:"reference_frame"`
}

func NewGetLinkStateRequest() (*GetLinkStateRequest) {
    newGetLinkStateRequest := new(GetLinkStateRequest)
    newGetLinkStateRequest.Go_link_name = ""
    newGetLinkStateRequest.Go_reference_frame = ""
    newGetLinkStateRequest.__id__ = 0
    return newGetLinkStateRequest
}

func (self *GetLinkStateRequest) Go_initialize() {
    self.Go_link_name = ""
    self.Go_reference_frame = ""
    self.__id__ = 0
}

func (self *GetLinkStateRequest) Go_serialize(buff []byte) (int) {
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
    length_reference_frame := len(self.Go_reference_frame)
    buff[offset + 0] = byte((length_reference_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_reference_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_reference_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_reference_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_reference_frame)], self.Go_reference_frame)
    offset += length_reference_frame
    return offset
}

func (self *GetLinkStateRequest) Go_deserialize(buff []byte) (int) {
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
    length_reference_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_reference_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_reference_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_reference_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_reference_frame = string(buff[offset:(offset+length_reference_frame)])
    offset += length_reference_frame
    return offset
}

func (self *GetLinkStateRequest) Go_serializedLength() (int) {
    length := 0
    length_link_name := len(self.Go_link_name)
    length += 4
    length += length_link_name
    length_reference_frame := len(self.Go_reference_frame)
    length += 4
    length += length_reference_frame
    return length
}

func (self *GetLinkStateRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetLinkStateRequest) Go_getType() (string) { return "gazebo_msgs/GetLinkState" }
func (self *GetLinkStateRequest) Go_getMD5() (string) { return "b9de4ed1795bda93c873763a2e55e76b" }
func (self *GetLinkStateRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetLinkStateRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetLinkStateResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_link_state *LinkState `json:"link_state"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewGetLinkStateResponse() (*GetLinkStateResponse) {
    newGetLinkStateResponse := new(GetLinkStateResponse)
    newGetLinkStateResponse.Go_link_state = NewLinkState()
    newGetLinkStateResponse.Go_success = false
    newGetLinkStateResponse.Go_status_message = ""
    newGetLinkStateResponse.__id__ = 0
    return newGetLinkStateResponse
}

func (self *GetLinkStateResponse) Go_initialize() {
    self.Go_link_state = NewLinkState()
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *GetLinkStateResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_link_state.Go_serialize(buff[offset:])
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

func (self *GetLinkStateResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_link_state.Go_deserialize(buff[offset:])
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

func (self *GetLinkStateResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_link_state.Go_serializedLength()
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *GetLinkStateResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetLinkStateResponse) Go_getType() (string) { return "gazebo_msgs/GetLinkState" }
func (self *GetLinkStateResponse) Go_getMD5() (string) { return "4d4305d53d97f8edc3b3ce04bcb94ed0" }
func (self *GetLinkStateResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetLinkStateResponse) Go_setID(id uint32) { self.__id__ = id }


