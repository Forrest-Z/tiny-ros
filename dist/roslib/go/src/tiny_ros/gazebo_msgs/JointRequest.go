package gazebo_msgs

import (
    "encoding/json"
)



type JointRequestRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_joint_name string `json:"joint_name"`
}

func NewJointRequestRequest() (*JointRequestRequest) {
    newJointRequestRequest := new(JointRequestRequest)
    newJointRequestRequest.Go_joint_name = ""
    newJointRequestRequest.__id__ = 0
    return newJointRequestRequest
}

func (self *JointRequestRequest) Go_initialize() {
    self.Go_joint_name = ""
    self.__id__ = 0
}

func (self *JointRequestRequest) Go_serialize(buff []byte) (int) {
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

func (self *JointRequestRequest) Go_deserialize(buff []byte) (int) {
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

func (self *JointRequestRequest) Go_serializedLength() (int) {
    length := 0
    length_joint_name := len(self.Go_joint_name)
    length += 4
    length += length_joint_name
    return length
}

func (self *JointRequestRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JointRequestRequest) Go_getType() (string) { return "gazebo_msgs/JointRequest" }
func (self *JointRequestRequest) Go_getMD5() (string) { return "e0bdc37edb92be07f3069573364a169f" }
func (self *JointRequestRequest) Go_getID() (uint32) { return self.__id__ }
func (self *JointRequestRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type JointRequestResponse struct {
    __id__ uint32 `json:"__id__"`
}

func NewJointRequestResponse() (*JointRequestResponse) {
    newJointRequestResponse := new(JointRequestResponse)
    newJointRequestResponse.__id__ = 0
    return newJointRequestResponse
}

func (self *JointRequestResponse) Go_initialize() {
    self.__id__ = 0
}

func (self *JointRequestResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *JointRequestResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *JointRequestResponse) Go_serializedLength() (int) {
    length := 0
    return length
}

func (self *JointRequestResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JointRequestResponse) Go_getType() (string) { return "gazebo_msgs/JointRequest" }
func (self *JointRequestResponse) Go_getMD5() (string) { return "ac5876a62df51a76c2828bb62894779d" }
func (self *JointRequestResponse) Go_getID() (uint32) { return self.__id__ }
func (self *JointRequestResponse) Go_setID(id uint32) { self.__id__ = id }


