package std_srvs

import (
    "encoding/json"
)



type TriggerRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewTriggerRequest() (*TriggerRequest) {
    newTriggerRequest := new(TriggerRequest)
    newTriggerRequest.__id__ = 0
    return newTriggerRequest
}

func (self *TriggerRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *TriggerRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *TriggerRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *TriggerRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *TriggerRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TriggerRequest) Go_getType() (string) { return "std_srvs/Trigger" }
func (self *TriggerRequest) Go_getMD5() (string) { return "23ce35174a691cb0bfbbdb21395edbac" }
func (self *TriggerRequest) Go_getID() (uint32) { return self.__id__ }
func (self *TriggerRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type TriggerResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_message string `json:"message"`
}

func NewTriggerResponse() (*TriggerResponse) {
    newTriggerResponse := new(TriggerResponse)
    newTriggerResponse.Go_success = false
    newTriggerResponse.Go_message = ""
    newTriggerResponse.__id__ = 0
    return newTriggerResponse
}

func (self *TriggerResponse) Go_initialize() {
    self.Go_success = false
    self.Go_message = ""
    self.__id__ = 0
}

func (self *TriggerResponse) Go_serialize(buff []byte) (int) {
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

func (self *TriggerResponse) Go_deserialize(buff []byte) (int) {
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

func (self *TriggerResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_message := len(self.Go_message)
    length += 4
    length += length_message
    return length
}

func (self *TriggerResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TriggerResponse) Go_getType() (string) { return "std_srvs/Trigger" }
func (self *TriggerResponse) Go_getMD5() (string) { return "08d111154ed595049573252ba929a6d8" }
func (self *TriggerResponse) Go_getID() (uint32) { return self.__id__ }
func (self *TriggerResponse) Go_setID(id uint32) { self.__id__ = id }


