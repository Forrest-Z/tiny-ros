package gazebo_msgs

import (
    "gazebo_msgs/ModelState"
)


type SetModelStateRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_state gazebo_msgs.ModelState `json:"model_state"`
}

func NewSetModelStateRequest() (*SetModelStateRequest) {
    newSetModelStateRequest := new(SetModelStateRequest)
    newSetModelStateRequest.Go_model_state = gazebo_msgs.NewModelState()
    newSetModelStateRequest.__id__ = 0
    return newSetModelStateRequest
}

func (self *SetModelStateRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_model_state.Go_serialize(buff[offset:])
    return offset
}

func (self *SetModelStateRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32((buff[offset + 0] & 0xFF) << (8 * 0))
    self.__id__ |=  uint32((buff[offset + 1] & 0xFF) << (8 * 1))
    self.__id__ |=  uint32((buff[offset + 2] & 0xFF) << (8 * 2))
    self.__id__ |=  uint32((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    offset += self.Go_model_state.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetModelStateRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_model_state.Go_serializedLength()
    return length
}

func (self *SetModelStateRequest) Go_echo() (string) { return "" }
func (self *SetModelStateRequest) Go_getType() (string) { return "gazebo_msgs/SetModelState" }
func (self *SetModelStateRequest) Go_getMD5() (string) { return "328e9e434938c39f3fd8e958ad8d6dab" }
func (self *SetModelStateRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetModelStateRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////

type SetModelStateResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetModelStateResponse() (*SetModelStateResponse) {
    newSetModelStateResponse := new(SetModelStateResponse)
    newSetModelStateResponse.Go_success = false
    newSetModelStateResponse.Go_status_message = ""
    newSetModelStateResponse.__id__ = 0
    return newSetModelStateResponse
}

func (self *SetModelStateResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_success >> (8 * 0)) & 0xFF)
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

func (self *SetModelStateResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32((buff[offset + 0] & 0xFF) << (8 * 0))
    self.__id__ |=  uint32((buff[offset + 1] & 0xFF) << (8 * 1))
    self.__id__ |=  uint32((buff[offset + 2] & 0xFF) << (8 * 2))
    self.__id__ |=  uint32((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    self.Go_success = bool((buff[offset + 0] & 0xFF) << (8 * 0))
    offset += 1
    length_status_message := int((buff[offset + 0] & 0xFF) << (8 * 0))
    length_status_message |= int((buff[offset + 1] & 0xFF) << (8 * 1))
    length_status_message |= int((buff[offset + 2] & 0xFF) << (8 * 2))
    length_status_message |= int((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    self.Go_status_message = string(buff[offset:(offset+length_status_message)])
    offset += length_status_message
    return offset
}

func (self *SetModelStateResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetModelStateResponse) Go_echo() (string) { return "" }
func (self *SetModelStateResponse) Go_getType() (string) { return "gazebo_msgs/SetModelState" }
func (self *SetModelStateResponse) Go_getMD5() (string) { return "7fe6c3a1ea4a4df9bf5b6858fd028ee6" }
func (self *SetModelStateResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetModelStateResponse) Go_setID(id uint32) { self.__id__ = id }


