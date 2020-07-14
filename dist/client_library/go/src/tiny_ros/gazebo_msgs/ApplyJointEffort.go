package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
    "encoding/binary"
    "math"
)



type ApplyJointEffortRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_joint_name string `json:"joint_name"`
    Go_effort float64 `json:"effort"`
    Go_start_time *rostime.Time `json:"start_time"`
    Go_duration *rostime.Duration `json:"duration"`
}

func NewApplyJointEffortRequest() (*ApplyJointEffortRequest) {
    newApplyJointEffortRequest := new(ApplyJointEffortRequest)
    newApplyJointEffortRequest.Go_joint_name = ""
    newApplyJointEffortRequest.Go_effort = 0.0
    newApplyJointEffortRequest.Go_start_time = rostime.NewTime()
    newApplyJointEffortRequest.Go_duration = rostime.NewDuration()
    newApplyJointEffortRequest.__id__ = 0
    return newApplyJointEffortRequest
}

func (self *ApplyJointEffortRequest) Go_initialize() {
    self.Go_joint_name = ""
    self.Go_effort = 0.0
    self.Go_start_time = rostime.NewTime()
    self.Go_duration = rostime.NewDuration()
    self.__id__ = 0
}

func (self *ApplyJointEffortRequest) Go_serialize(buff []byte) (int) {
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
    bits_effort := math.Float64bits(self.Go_effort)
    binary.LittleEndian.PutUint64(buff[offset:], bits_effort)
    offset += 8
    buff[offset + 0] = byte((self.Go_start_time.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_start_time.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_start_time.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_start_time.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_start_time.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_start_time.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_start_time.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_start_time.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_duration.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_duration.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_duration.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_duration.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_duration.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_duration.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_duration.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_duration.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *ApplyJointEffortRequest) Go_deserialize(buff []byte) (int) {
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
    bits_effort := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_effort = math.Float64frombits(bits_effort)
    offset += 8
    self.Go_start_time.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_start_time.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_start_time.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_start_time.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_start_time.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_start_time.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_start_time.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_start_time.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_duration.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_duration.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_duration.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_duration.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_duration.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_duration.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_duration.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_duration.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *ApplyJointEffortRequest) Go_serializedLength() (int) {
    length := 0
    length_joint_name := len(self.Go_joint_name)
    length += 4
    length += length_joint_name
    length += 8
    length += 4
    length += 4
    length += 4
    length += 4
    return length
}

func (self *ApplyJointEffortRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ApplyJointEffortRequest) Go_getType() (string) { return "gazebo_msgs/ApplyJointEffort" }
func (self *ApplyJointEffortRequest) Go_getMD5() (string) { return "90595a46cf1fb4ee17158e2f7034a0eb" }
func (self *ApplyJointEffortRequest) Go_getID() (uint32) { return self.__id__ }
func (self *ApplyJointEffortRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type ApplyJointEffortResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewApplyJointEffortResponse() (*ApplyJointEffortResponse) {
    newApplyJointEffortResponse := new(ApplyJointEffortResponse)
    newApplyJointEffortResponse.Go_success = false
    newApplyJointEffortResponse.Go_status_message = ""
    newApplyJointEffortResponse.__id__ = 0
    return newApplyJointEffortResponse
}

func (self *ApplyJointEffortResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *ApplyJointEffortResponse) Go_serialize(buff []byte) (int) {
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

func (self *ApplyJointEffortResponse) Go_deserialize(buff []byte) (int) {
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

func (self *ApplyJointEffortResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *ApplyJointEffortResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ApplyJointEffortResponse) Go_getType() (string) { return "gazebo_msgs/ApplyJointEffort" }
func (self *ApplyJointEffortResponse) Go_getMD5() (string) { return "953194fc8ca726693bef2876cebb0438" }
func (self *ApplyJointEffortResponse) Go_getID() (uint32) { return self.__id__ }
func (self *ApplyJointEffortResponse) Go_setID(id uint32) { self.__id__ = id }


