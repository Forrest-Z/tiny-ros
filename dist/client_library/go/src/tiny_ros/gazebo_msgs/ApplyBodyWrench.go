package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
    "tiny_ros/geometry_msgs"
)



type ApplyBodyWrenchRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_body_name string `json:"body_name"`
    Go_reference_frame string `json:"reference_frame"`
    Go_reference_point *geometry_msgs.Point `json:"reference_point"`
    Go_wrench *geometry_msgs.Wrench `json:"wrench"`
    Go_start_time *rostime.Time `json:"start_time"`
    Go_duration *rostime.Duration `json:"duration"`
}

func NewApplyBodyWrenchRequest() (*ApplyBodyWrenchRequest) {
    newApplyBodyWrenchRequest := new(ApplyBodyWrenchRequest)
    newApplyBodyWrenchRequest.Go_body_name = ""
    newApplyBodyWrenchRequest.Go_reference_frame = ""
    newApplyBodyWrenchRequest.Go_reference_point = geometry_msgs.NewPoint()
    newApplyBodyWrenchRequest.Go_wrench = geometry_msgs.NewWrench()
    newApplyBodyWrenchRequest.Go_start_time = rostime.NewTime()
    newApplyBodyWrenchRequest.Go_duration = rostime.NewDuration()
    newApplyBodyWrenchRequest.__id__ = 0
    return newApplyBodyWrenchRequest
}

func (self *ApplyBodyWrenchRequest) Go_initialize() {
    self.Go_body_name = ""
    self.Go_reference_frame = ""
    self.Go_reference_point = geometry_msgs.NewPoint()
    self.Go_wrench = geometry_msgs.NewWrench()
    self.Go_start_time = rostime.NewTime()
    self.Go_duration = rostime.NewDuration()
    self.__id__ = 0
}

func (self *ApplyBodyWrenchRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    length_body_name := len(self.Go_body_name)
    buff[offset + 0] = byte((length_body_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_body_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_body_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_body_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_body_name)], self.Go_body_name)
    offset += length_body_name
    length_reference_frame := len(self.Go_reference_frame)
    buff[offset + 0] = byte((length_reference_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_reference_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_reference_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_reference_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_reference_frame)], self.Go_reference_frame)
    offset += length_reference_frame
    offset += self.Go_reference_point.Go_serialize(buff[offset:])
    offset += self.Go_wrench.Go_serialize(buff[offset:])
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

func (self *ApplyBodyWrenchRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_body_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_body_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_body_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_body_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_body_name = string(buff[offset:(offset+length_body_name)])
    offset += length_body_name
    length_reference_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_reference_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_reference_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_reference_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_reference_frame = string(buff[offset:(offset+length_reference_frame)])
    offset += length_reference_frame
    offset += self.Go_reference_point.Go_deserialize(buff[offset:])
    offset += self.Go_wrench.Go_deserialize(buff[offset:])
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

func (self *ApplyBodyWrenchRequest) Go_serializedLength() (int) {
    length := 0
    length_body_name := len(self.Go_body_name)
    length += 4
    length += length_body_name
    length_reference_frame := len(self.Go_reference_frame)
    length += 4
    length += length_reference_frame
    length += self.Go_reference_point.Go_serializedLength()
    length += self.Go_wrench.Go_serializedLength()
    length += 4
    length += 4
    length += 4
    length += 4
    return length
}

func (self *ApplyBodyWrenchRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ApplyBodyWrenchRequest) Go_getType() (string) { return "gazebo_msgs/ApplyBodyWrench" }
func (self *ApplyBodyWrenchRequest) Go_getMD5() (string) { return "434adb4bdbb64c5610c7fadb31f0fa7d" }
func (self *ApplyBodyWrenchRequest) Go_getID() (uint32) { return self.__id__ }
func (self *ApplyBodyWrenchRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type ApplyBodyWrenchResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
    Go_status_message string `json:"status_message"`
}

func NewApplyBodyWrenchResponse() (*ApplyBodyWrenchResponse) {
    newApplyBodyWrenchResponse := new(ApplyBodyWrenchResponse)
    newApplyBodyWrenchResponse.Go_success = false
    newApplyBodyWrenchResponse.Go_status_message = ""
    newApplyBodyWrenchResponse.__id__ = 0
    return newApplyBodyWrenchResponse
}

func (self *ApplyBodyWrenchResponse) Go_initialize() {
    self.Go_success = false
    self.Go_status_message = ""
    self.__id__ = 0
}

func (self *ApplyBodyWrenchResponse) Go_serialize(buff []byte) (int) {
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

func (self *ApplyBodyWrenchResponse) Go_deserialize(buff []byte) (int) {
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

func (self *ApplyBodyWrenchResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_status_message := len(self.Go_status_message)
    length += 4
    length += length_status_message
    return length
}

func (self *ApplyBodyWrenchResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ApplyBodyWrenchResponse) Go_getType() (string) { return "gazebo_msgs/ApplyBodyWrench" }
func (self *ApplyBodyWrenchResponse) Go_getMD5() (string) { return "f29b3c75e7d692065eda02aae6d3a3a0" }
func (self *ApplyBodyWrenchResponse) Go_getID() (uint32) { return self.__id__ }
func (self *ApplyBodyWrenchResponse) Go_setID(id uint32) { self.__id__ = id }


