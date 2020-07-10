package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/geometry_msgs"
    "tiny_ros/trajectory_msgs"
)



type SetJointTrajectoryRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_model_name string `json:"model_name"`
    Go_joint_trajectory *trajectory_msgs.JointTrajectory `json:"joint_trajectory"`
    Go_model_pose *geometry_msgs.Pose `json:"model_pose"`
    Go_set_model_pose bool `json:"set_model_pose"`
    Go_disable_physics_updates bool `json:"disable_physics_updates"`
}

func NewSetJointTrajectoryRequest() (*SetJointTrajectoryRequest) {
    newSetJointTrajectoryRequest := new(SetJointTrajectoryRequest)
    newSetJointTrajectoryRequest.Go_model_name = ""
    newSetJointTrajectoryRequest.Go_joint_trajectory = trajectory_msgs.NewJointTrajectory()
    newSetJointTrajectoryRequest.Go_model_pose = geometry_msgs.NewPose()
    newSetJointTrajectoryRequest.Go_set_model_pose = false
    newSetJointTrajectoryRequest.Go_disable_physics_updates = false
    newSetJointTrajectoryRequest.__id__ = 0
    return newSetJointTrajectoryRequest
}

func (self *SetJointTrajectoryRequest) Go_initialize() {
    self.Go_model_name = ""
    self.Go_joint_trajectory = trajectory_msgs.NewJointTrajectory()
    self.Go_model_pose = geometry_msgs.NewPose()
    self.Go_set_model_pose = false
    self.Go_disable_physics_updates = false
    self.__id__ = 0
}

func (self *SetJointTrajectoryRequest) Go_serialize(buff []byte) (int) {
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
    offset += self.Go_joint_trajectory.Go_serialize(buff[offset:])
    offset += self.Go_model_pose.Go_serialize(buff[offset:])
    if self.Go_set_model_pose {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    if self.Go_disable_physics_updates {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *SetJointTrajectoryRequest) Go_deserialize(buff []byte) (int) {
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
    offset += self.Go_joint_trajectory.Go_deserialize(buff[offset:])
    offset += self.Go_model_pose.Go_deserialize(buff[offset:])
    if (buff[offset] & 0xFF) != 0 {
        self.Go_set_model_pose = true
    } else {
        self.Go_set_model_pose = false
    }
    offset += 1
    if (buff[offset] & 0xFF) != 0 {
        self.Go_disable_physics_updates = true
    } else {
        self.Go_disable_physics_updates = false
    }
    offset += 1
    return offset
}

func (self *SetJointTrajectoryRequest) Go_serializedLength() (int) {
    length := 0
    length_model_name := len(self.Go_model_name)
    length += 4
    length += length_model_name
    length += self.Go_joint_trajectory.Go_serializedLength()
    length += self.Go_model_pose.Go_serializedLength()
    length += 1
    length += 1
    return length
}

func (self *SetJointTrajectoryRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetJointTrajectoryRequest) Go_getType() (string) { return "gazebo_msgs/SetJointTrajectory" }
func (self *SetJointTrajectoryRequest) Go_getMD5() (string) { return "8230e1fcc0295d8d21fbd5df0ccb0af3" }
func (self *SetJointTrajectoryRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetJointTrajectoryRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type SetJointTrajectoryResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewSetJointTrajectoryResponse() (*SetJointTrajectoryResponse) {
    newSetJointTrajectoryResponse := new(SetJointTrajectoryResponse)
    newSetJointTrajectoryResponse.Go_success = false
    newSetJointTrajectoryResponse.Go_status_message = ""
    newSetJointTrajectoryResponse.__id__ = 0
    return newSetJointTrajectoryResponse
}

func (self *SetJointTrajectoryResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *SetJointTrajectoryResponse) Go_serialize(buff []byte) (int) {
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

func (self *SetJointTrajectoryResponse) Go_deserialize(buff []byte) (int) {
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

func (self *SetJointTrajectoryResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *SetJointTrajectoryResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SetJointTrajectoryResponse) Go_getType() (string) { return "gazebo_msgs/SetJointTrajectory" }
func (self *SetJointTrajectoryResponse) Go_getMD5() (string) { return "2f5fe47400272efd54556969ffe63e7e" }
func (self *SetJointTrajectoryResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetJointTrajectoryResponse) Go_setID(id uint32) { self.__id__ = id }


