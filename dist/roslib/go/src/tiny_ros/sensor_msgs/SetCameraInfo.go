package sensor_msgs

import (
    "encoding/json"
)



type SetCameraInfoRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_camera_info *CameraInfo `json:"camera_info"`
}

func NewSetCameraInfoRequest() (*SetCameraInfoRequest) {
    newSetCameraInfoRequest := new(SetCameraInfoRequest)
    newSetCameraInfoRequest.Go_camera_info = NewCameraInfo()
    newSetCameraInfoRequest.__id__ = 0
    return newSetCameraInfoRequest
}

func (self *SetCameraInfoRequest) Go_initialize() {
    self.Go_camera_info = NewCameraInfo()
    self.__id__ = 0
}

func (self *SetCameraInfoRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_camera_info.Go_serialize(buff[offset:])
    return offset
}

func (self *SetCameraInfoRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_camera_info.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetCameraInfoRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_camera_info.Go_serializedLength()
    return length
}

func (self *SetCameraInfoRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetCameraInfoRequest) Go_getType() (string) { return "sensor_msgs/SetCameraInfo" }
func (self *SetCameraInfoRequest) Go_getMD5() (string) { return "7c09688f428450a1ac49eacdbb57c012" }
func (self *SetCameraInfoRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetCameraInfoRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetCameraInfoResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetCameraInfoResponse() (*SetCameraInfoResponse) {
    newSetCameraInfoResponse := new(SetCameraInfoResponse)
    newSetCameraInfoResponse.Go_success = false
    newSetCameraInfoResponse.Go_status_message = ""
    newSetCameraInfoResponse.__id__ = 0
    return newSetCameraInfoResponse
}

func (self *SetCameraInfoResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetCameraInfoResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetCameraInfoResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetCameraInfoResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetCameraInfoResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetCameraInfoResponse) Go_getType() (string) { return "sensor_msgs/SetCameraInfo" }
func (self *SetCameraInfoResponse) Go_getMD5() (string) { return "e03fdc9555d1e3c7347a728e913775f6" }
func (self *SetCameraInfoResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetCameraInfoResponse) Go_setID(id uint32) { self.__id__ = id }


