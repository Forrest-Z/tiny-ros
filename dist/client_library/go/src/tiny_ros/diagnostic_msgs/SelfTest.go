package diagnostic_msgs

import (
    "encoding/json"
)



type SelfTestRequest struct {
    __id__ uint32 `json:"__id__"`
}

func NewSelfTestRequest() (*SelfTestRequest) {
    newSelfTestRequest := new(SelfTestRequest)
    newSelfTestRequest.__id__ = 0
    return newSelfTestRequest
}

func (self *SelfTestRequest) Go_initialize() {
    self.__id__ = 0
}

func (self *SelfTestRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *SelfTestRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *SelfTestRequest) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *SelfTestRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SelfTestRequest) Go_getType() (string) { return "diagnostic_msgs/SelfTest" }
func (self *SelfTestRequest) Go_getMD5() (string) { return "049f87742408b36b8ef5f7dd71e3ef5a" }
func (self *SelfTestRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SelfTestRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SelfTestResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_id string `json:"id"`
    Go_passed byte `json:"passed"`
    Go_status []*DiagnosticStatus `json:"status"`
}

func NewSelfTestResponse() (*SelfTestResponse) {
    newSelfTestResponse := new(SelfTestResponse)
    newSelfTestResponse.Go_id = ""
    newSelfTestResponse.Go_passed = 0
    newSelfTestResponse.Go_status = []*DiagnosticStatus{}
    newSelfTestResponse.__id__ = 0
    return newSelfTestResponse
}

func (self *SelfTestResponse) Go_initialize() {
    self.Go_id = ""
    self.Go_passed = 0
    self.Go_status = []*DiagnosticStatus{}
    self.__id__ = 0
}

func (self *SelfTestResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_id := len(self.Go_id)
    buff[offset + 0] = byte((length_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_id)], self.Go_id)
    offset += length_id
    buff[offset + 0] = byte((self.Go_passed >> (8 * 0)) & 0xFF)
    offset += 1
    length_status := len(self.Go_status)
    buff[offset + 0] = byte((length_status >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_status >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_status >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_status >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_status; i++ {
        offset += self.Go_status[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *SelfTestResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_id = string(buff[offset:(offset+length_id)])
    offset += length_id
    self.Go_passed = byte(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_status := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_status |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_status |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_status |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_status = make([]*DiagnosticStatus, length_status)
    for i := 0; i < length_status; i++ {
        self.Go_status[i] = NewDiagnosticStatus()
    }
    for i := 0; i < length_status; i++ {
        offset += self.Go_status[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *SelfTestResponse) Go_serializedLength() (int) {
    length := 0
    length_id := len(self.Go_id)
    length += 4
    length += length_id
    length += 1
    length += 4
    length_status := len(self.Go_status)
    for i := 0; i < length_status; i++ {
        length += self.Go_status[i].Go_serializedLength()
    }
    return length
}

func (self *SelfTestResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SelfTestResponse) Go_getType() (string) { return "diagnostic_msgs/SelfTest" }
func (self *SelfTestResponse) Go_getMD5() (string) { return "70aaf2a851ccb5e946b2d112ea26f7b9" }
func (self *SelfTestResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SelfTestResponse) Go_setID(id uint32) { self.__id__ = id }


