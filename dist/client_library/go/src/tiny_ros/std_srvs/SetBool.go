package std_srvs

import (
    "encoding/json"
)



type SetBoolRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_data bool `json:"data"`
}

func NewSetBoolRequest() (*SetBoolRequest) {
    newSetBoolRequest := new(SetBoolRequest)
    newSetBoolRequest.Go_data = false
    newSetBoolRequest.__id__ = 0
    return newSetBoolRequest
}

func (self *SetBoolRequest) Go_initialize() {
    self.Go_data = false
    self.__id__ = 0
}

func (self *SetBoolRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    if self.Go_data {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *SetBoolRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    if (buff[offset] & 0xFF) != 0 {
        self.Go_data = true
    } else {
        self.Go_data = false
    }
    offset += 1
    return offset
}

func (self *SetBoolRequest) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *SetBoolRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetBoolRequest) Go_getType() (string) { return "std_srvs/SetBool" }
func (self *SetBoolRequest) Go_getMD5() (string) { return "2600271ce244b6b0d236894ec6f04373" }
func (self *SetBoolRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetBoolRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetBoolResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_message string `json:"message"`
}

func NewSetBoolResponse() (*SetBoolResponse) {
    newSetBoolResponse := new(SetBoolResponse)
    newSetBoolResponse.Go_success = false
    newSetBoolResponse.Go_message = ""
    newSetBoolResponse.__id__ = 0
    return newSetBoolResponse
}

func (self *SetBoolResponse) Go_initialize() {
    self.Go_success = false
    self.Go_message = ""
    self.__id__ = 0
}

func (self *SetBoolResponse) Go_serialize(buff []byte) (int) {
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
    length_message := len(self.Go_message)
    buff[offset + 0] = byte((length_message >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_message >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_message >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_message >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_message)], self.Go_message)
    offset += length_message
    return offset
}

func (self *SetBoolResponse) Go_deserialize(buff []byte) (int) {
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
    length_message := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_message |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_message |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_message |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_message = string(buff[offset:(offset+length_message)])
    offset += length_message
    return offset
}

func (self *SetBoolResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_message := len(self.Go_message)
    length += 4
    length += length_message
    return length
}

func (self *SetBoolResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetBoolResponse) Go_getType() (string) { return "std_srvs/SetBool" }
func (self *SetBoolResponse) Go_getMD5() (string) { return "51cf1b5cb67d107350442299d694fd1d" }
func (self *SetBoolResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetBoolResponse) Go_setID(id uint32) { self.__id__ = id }


