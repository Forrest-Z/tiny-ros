package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/tinyros/time"
)


type TimeReference struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_time_ref *rostime.Time `json:"time_ref"`
    Go_source string `json:"source"`
}

func NewTimeReference() (*TimeReference) {
    newTimeReference := new(TimeReference)
    newTimeReference.Go_header = std_msgs.NewHeader()
    newTimeReference.Go_time_ref = rostime.NewTime()
    newTimeReference.Go_source = ""
    return newTimeReference
}

func (self *TimeReference) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_time_ref = rostime.NewTime()
    self.Go_source = ""
}

func (self *TimeReference) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_time_ref.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_ref.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_ref.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_ref.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_time_ref.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_ref.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_ref.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_ref.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    length_source := len(self.Go_source)
    buff[offset + 0] = byte((length_source >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_source >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_source >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_source >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_source)], self.Go_source)
    offset += length_source
    return offset
}

func (self *TimeReference) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_time_ref.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_ref.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_ref.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_ref.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_time_ref.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_ref.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_ref.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_ref.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_source := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_source |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_source |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_source |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_source = string(buff[offset:(offset+length_source)])
    offset += length_source
    return offset
}

func (self *TimeReference) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length += 4
    length_source := len(self.Go_source)
    length += 4
    length += length_source
    return length
}

func (self *TimeReference) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TimeReference) Go_getType() (string) { return "sensor_msgs/TimeReference" }
func (self *TimeReference) Go_getMD5() (string) { return "8e1576e01de57cd0d55758112f0e84ec" }
func (self *TimeReference) Go_getID() (uint32) { return 0 }
func (self *TimeReference) Go_setID(id uint32) { }

