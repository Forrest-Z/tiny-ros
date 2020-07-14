package gazebo_msgs

import (
    "encoding/json"
)



type DeleteModelRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
}

func NewDeleteModelRequest() (*DeleteModelRequest) {
    newDeleteModelRequest := new(DeleteModelRequest)
    newDeleteModelRequest.Go_model_name = ""
    newDeleteModelRequest.__id__ = 0
    return newDeleteModelRequest
}

func (self *DeleteModelRequest) Go_initialize() {
    self.Go_model_name = ""
    self.__id__ = 0
}

func (self *DeleteModelRequest) Go_serialize(buff []byte) (int) {
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

func (self *DeleteModelRequest) Go_deserialize(buff []byte) (int) {
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

func (self *DeleteModelRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    return length
}

func (self *DeleteModelRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *DeleteModelRequest) Go_getType() (string) { return "gazebo_msgs/DeleteModel" }
func (self *DeleteModelRequest) Go_getMD5() (string) { return "c4e25cd35d75c6c2f51ee0d986d8e556" }
func (self *DeleteModelRequest) Go_getID() (uint32) { return self.__id__ }
func (self *DeleteModelRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type DeleteModelResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewDeleteModelResponse() (*DeleteModelResponse) {
    newDeleteModelResponse := new(DeleteModelResponse)
    newDeleteModelResponse.Go_success = false
    newDeleteModelResponse.Go_status_message = ""
    newDeleteModelResponse.__id__ = 0
    return newDeleteModelResponse
}

func (self *DeleteModelResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *DeleteModelResponse) Go_serialize(buff []byte) (int) {
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

func (self *DeleteModelResponse) Go_deserialize(buff []byte) (int) {
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

func (self *DeleteModelResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *DeleteModelResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *DeleteModelResponse) Go_getType() (string) { return "gazebo_msgs/DeleteModel" }
func (self *DeleteModelResponse) Go_getMD5() (string) { return "3feb2eeea1c45bcf64067e4dd162726f" }
func (self *DeleteModelResponse) Go_getID() (uint32) { return self.__id__ }
func (self *DeleteModelResponse) Go_setID(id uint32) { self.__id__ = id }


