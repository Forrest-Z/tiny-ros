package diagnostic_msgs

import (
    "encoding/json"
)



type AddDiagnosticsRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_load_namespace string `json:"load_namespace"`
}

func NewAddDiagnosticsRequest() (*AddDiagnosticsRequest) {
    newAddDiagnosticsRequest := new(AddDiagnosticsRequest)
    newAddDiagnosticsRequest.Go_load_namespace = ""
    newAddDiagnosticsRequest.__id__ = 0
    return newAddDiagnosticsRequest
}

func (self *AddDiagnosticsRequest) Go_initialize() {
    self.Go_load_namespace = ""
    self.__id__ = 0
}

func (self *AddDiagnosticsRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_load_namespace := len(self.Go_load_namespace)
    buff[offset + 0] = byte((length_load_namespace >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_load_namespace >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_load_namespace >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_load_namespace >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_load_namespace)], self.Go_load_namespace)
    offset += length_load_namespace
    return offset
}

func (self *AddDiagnosticsRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_load_namespace := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_load_namespace |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_load_namespace |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_load_namespace |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_load_namespace = string(buff[offset:(offset+length_load_namespace)])
    offset += length_load_namespace
    return offset
}

func (self *AddDiagnosticsRequest) Go_serializedLength() (int) {
    length := 0
    length_load_namespace := len(self.Go_load_namespace)
    length += 4
    length += length_load_namespace
    return length
}

func (self *AddDiagnosticsRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *AddDiagnosticsRequest) Go_getType() (string) { return "diagnostic_msgs/AddDiagnostics" }
func (self *AddDiagnosticsRequest) Go_getMD5() (string) { return "005ba76b3cd04edebfe46acad928fbeb" }
func (self *AddDiagnosticsRequest) Go_getID() (uint32) { return self.__id__ }
func (self *AddDiagnosticsRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type AddDiagnosticsResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_message string `json:"message"`
}

func NewAddDiagnosticsResponse() (*AddDiagnosticsResponse) {
    newAddDiagnosticsResponse := new(AddDiagnosticsResponse)
    newAddDiagnosticsResponse.Go_success = false
    newAddDiagnosticsResponse.Go_message = ""
    newAddDiagnosticsResponse.__id__ = 0
    return newAddDiagnosticsResponse
}

func (self *AddDiagnosticsResponse) Go_initialize() {
    self.Go_success = false
    self.Go_message = ""
    self.__id__ = 0
}

func (self *AddDiagnosticsResponse) Go_serialize(buff []byte) (int) {
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

func (self *AddDiagnosticsResponse) Go_deserialize(buff []byte) (int) {
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

func (self *AddDiagnosticsResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_message := len(self.Go_message)
    length += 4
    length += length_message
    return length
}

func (self *AddDiagnosticsResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *AddDiagnosticsResponse) Go_getType() (string) { return "diagnostic_msgs/AddDiagnostics" }
func (self *AddDiagnosticsResponse) Go_getMD5() (string) { return "9bd37b30a2340a31743d1e80a2c52ed0" }
func (self *AddDiagnosticsResponse) Go_getID() (uint32) { return self.__id__ }
func (self *AddDiagnosticsResponse) Go_setID(id uint32) { self.__id__ = id }


