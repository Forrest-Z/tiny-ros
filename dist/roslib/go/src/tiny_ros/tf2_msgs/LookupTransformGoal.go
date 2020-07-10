package tf2_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type LookupTransformGoal struct {
    Go_target_frame string `json:"target_frame"`
    Go_source_frame string `json:"source_frame"`
    Go_source_time *rostime.Time `json:"source_time"`
    Go_timeout *rostime.Duration `json:"timeout"`
    Go_target_time *rostime.Time `json:"target_time"`
    Go_fixed_frame string `json:"fixed_frame"`
    Go_advanced bool `json:"advanced"`
}

func NewLookupTransformGoal() (*LookupTransformGoal) {
    newLookupTransformGoal := new(LookupTransformGoal)
    newLookupTransformGoal.Go_target_frame = ""
    newLookupTransformGoal.Go_source_frame = ""
    newLookupTransformGoal.Go_source_time = rostime.NewTime()
    newLookupTransformGoal.Go_timeout = rostime.NewDuration()
    newLookupTransformGoal.Go_target_time = rostime.NewTime()
    newLookupTransformGoal.Go_fixed_frame = ""
    newLookupTransformGoal.Go_advanced = false
    return newLookupTransformGoal
}

func (self *LookupTransformGoal) Go_initialize() {
    self.Go_target_frame = ""
    self.Go_source_frame = ""
    self.Go_source_time = rostime.NewTime()
    self.Go_timeout = rostime.NewDuration()
    self.Go_target_time = rostime.NewTime()
    self.Go_fixed_frame = ""
    self.Go_advanced = false
}

func (self *LookupTransformGoal) Go_serialize(buff []byte) (int) {
    offset := 0
    length_target_frame := len(self.Go_target_frame)
    buff[offset + 0] = byte((length_target_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_target_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_target_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_target_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_target_frame)], self.Go_target_frame)
    offset += length_target_frame
    length_source_frame := len(self.Go_source_frame)
    buff[offset + 0] = byte((length_source_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_source_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_source_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_source_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_source_frame)], self.Go_source_frame)
    offset += length_source_frame
    buff[offset + 0] = byte((self.Go_source_time.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_source_time.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_source_time.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_source_time.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_source_time.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_source_time.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_source_time.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_source_time.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_timeout.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_timeout.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_timeout.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_timeout.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_timeout.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_timeout.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_timeout.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_timeout.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_target_time.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_target_time.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_target_time.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_target_time.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_target_time.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_target_time.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_target_time.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_target_time.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    length_fixed_frame := len(self.Go_fixed_frame)
    buff[offset + 0] = byte((length_fixed_frame >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_fixed_frame >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_fixed_frame >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_fixed_frame >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_fixed_frame)], self.Go_fixed_frame)
    offset += length_fixed_frame
    if self.Go_advanced {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    return offset
}

func (self *LookupTransformGoal) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_target_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_target_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_target_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_target_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_target_frame = string(buff[offset:(offset+length_target_frame)])
    offset += length_target_frame
    length_source_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_source_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_source_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_source_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_source_frame = string(buff[offset:(offset+length_source_frame)])
    offset += length_source_frame
    self.Go_source_time.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_source_time.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_source_time.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_source_time.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_source_time.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_source_time.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_source_time.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_source_time.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_timeout.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_timeout.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_timeout.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_timeout.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_timeout.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_timeout.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_timeout.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_timeout.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_target_time.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_target_time.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_target_time.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_target_time.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_target_time.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_target_time.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_target_time.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_target_time.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_fixed_frame := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_fixed_frame |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_fixed_frame |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_fixed_frame |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_fixed_frame = string(buff[offset:(offset+length_fixed_frame)])
    offset += length_fixed_frame
    if (buff[offset] & 0xFF) != 0 {
        self.Go_advanced = true
    } else {
        self.Go_advanced = false
    }
    offset += 1
    return offset
}

func (self *LookupTransformGoal) Go_serializedLength() (int) {
    length := 0
    length_target_frame := len(self.Go_target_frame)
    length += 4
    length += length_target_frame
    length_source_frame := len(self.Go_source_frame)
    length += 4
    length += length_source_frame
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length_fixed_frame := len(self.Go_fixed_frame)
    length += 4
    length += length_fixed_frame
    length += 1
    return length
}

func (self *LookupTransformGoal) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *LookupTransformGoal) Go_getType() (string) { return "tf2_msgs/LookupTransformGoal" }
func (self *LookupTransformGoal) Go_getMD5() (string) { return "677c84a912b788ccaaea5fbc38570182" }
func (self *LookupTransformGoal) Go_getID() (uint32) { return 0 }
func (self *LookupTransformGoal) Go_setID(id uint32) { }

