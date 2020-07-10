package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
)



type SpawnModelRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
    Go_model_xml string `json:"model_xml"`
    Go_robot_namespace string `json:"robot_namespace"`
    Go_initial_pose *geometry_msgs.Pose `json:"initial_pose"`
    Go_reference_frame string `json:"reference_frame"`
}

func NewSpawnModelRequest() (*SpawnModelRequest) {
    newSpawnModelRequest := new(SpawnModelRequest)
    newSpawnModelRequest.Go_model_name = ""
    newSpawnModelRequest.Go_model_xml = ""
    newSpawnModelRequest.Go_robot_namespace = ""
    newSpawnModelRequest.Go_initial_pose = geometry_msgs.NewPose()
    newSpawnModelRequest.Go_reference_frame = ""
    newSpawnModelRequest.__id__ = 0
    return newSpawnModelRequest
}

func (self *SpawnModelRequest) Go_initialize() {
    self.Go_model_name = ""
    self.Go_model_xml = ""
    self.Go_robot_namespace = ""
    self.Go_initial_pose = geometry_msgs.NewPose()
    self.Go_reference_frame = ""
    self.__id__ = 0
}

func (self *SpawnModelRequest) Go_serialize(buff []byte) (int) {
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
    length_model_xml := len(self.Go_model_xml)
    buff[offset + 0] = byte((length_model_xml >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_model_xml >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_model_xml >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_model_xml >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_model_xml)], self.Go_model_xml)
    offset += length_model_xml
    length_robot_namespace := len(self.Go_robot_namespace)
    buff[offset + 0] = byte((length_robot_namespace >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_robot_namespace >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_robot_namespace >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_robot_namespace >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_robot_namespace)], self.Go_robot_namespace)
    offset += length_robot_namespace
    offset += self.Go_initial_pose.Go_serialize(buff[offset:])
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

func (self *SpawnModelRequest) Go_deserialize(buff []byte) (int) {
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
    length_model_xml := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_model_xml |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_model_xml |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_model_xml |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_model_xml = string(buff[offset:(offset+length_model_xml)])
    offset += length_model_xml
    length_robot_namespace := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_robot_namespace |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_robot_namespace |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_robot_namespace |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_robot_namespace = string(buff[offset:(offset+length_robot_namespace)])
    offset += length_robot_namespace
    offset += self.Go_initial_pose.Go_deserialize(buff[offset:])
    length_reference_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_reference_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_reference_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_reference_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_reference_frame = string(buff[offset:(offset+length_reference_frame)])
    offset += length_reference_frame
    return offset
}

func (self *SpawnModelRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    length_model_xml := len(self.Go_model_xml)
    length += 4
    length += length_model_xml
    length_robot_namespace := len(self.Go_robot_namespace)
    length += 4
    length += length_robot_namespace
    length += self.Go_initial_pose.Go_serializedLength()
    length_reference_frame := len(self.Go_reference_frame)
    length += 4
    length += length_reference_frame
    return length
}

func (self *SpawnModelRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SpawnModelRequest) Go_getType() (string) { return "gazebo_msgs/SpawnModel" }
func (self *SpawnModelRequest) Go_getMD5() (string) { return "da34e61c8813e52ac159e5f31fbf55be" }
func (self *SpawnModelRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SpawnModelRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SpawnModelResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSpawnModelResponse() (*SpawnModelResponse) {
    newSpawnModelResponse := new(SpawnModelResponse)
    newSpawnModelResponse.Go_success = false
    newSpawnModelResponse.Go_status_message = ""
    newSpawnModelResponse.__id__ = 0
    return newSpawnModelResponse
}

func (self *SpawnModelResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SpawnModelResponse) Go_serialize(buff []byte) (int) {
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

func (self *SpawnModelResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SpawnModelResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SpawnModelResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SpawnModelResponse) Go_getType() (string) { return "gazebo_msgs/SpawnModel" }
func (self *SpawnModelResponse) Go_getMD5() (string) { return "d59d46cc4e5a64f978a429dd7c306d30" }
func (self *SpawnModelResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SpawnModelResponse) Go_setID(id uint32) { self.__id__ = id }


