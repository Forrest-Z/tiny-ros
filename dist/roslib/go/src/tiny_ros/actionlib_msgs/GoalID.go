package actionlib_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type GoalID struct {
    Go_stamp *rostime.Time `json:"stamp"`
    Go_id string `json:"id"`
}

func NewGoalID() (*GoalID) {
    newGoalID := new(GoalID)
    newGoalID.Go_stamp = rostime.NewTime()
    newGoalID.Go_id = ""
    return newGoalID
}

func (self *GoalID) Go_initialize() {
    self.Go_stamp = rostime.NewTime()
    self.Go_id = ""
}

func (self *GoalID) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_stamp.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    length_id := len(self.Go_id)
    buff[offset + 0] = byte((length_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_id)], self.Go_id)
    offset += length_id
    return offset
}

func (self *GoalID) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_stamp.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_id = string(buff[offset:(offset+length_id)])
    offset += length_id
    return offset
}

func (self *GoalID) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length_id := len(self.Go_id)
    length += 4
    length += length_id
    return length
}

func (self *GoalID) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GoalID) Go_getType() (string) { return "actionlib_msgs/GoalID" }
func (self *GoalID) Go_getMD5() (string) { return "a6cee90e5a185f4cb050de49bc4fa1f4" }
func (self *GoalID) Go_getID() (uint32) { return 0 }
func (self *GoalID) Go_setID(id uint32) { }

