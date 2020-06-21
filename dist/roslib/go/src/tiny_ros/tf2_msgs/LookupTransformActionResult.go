package tf2_msgs

import (
    "tiny_ros/std_msgs"
    "actionlib_msgs/GoalStatus"
    "tf2_msgs/LookupTransformResult"
)

type LookupTransformActionResult struct {
    Go_header std_msgs.Header `json:"header"`
    Go_status actionlib_msgs.GoalStatus `json:"status"`
    Go_result tf2_msgs.LookupTransformResult `json:"result"`
}

func NewLookupTransformActionResult() (*LookupTransformActionResult) {
    newLookupTransformActionResult := new(LookupTransformActionResult)
    newLookupTransformActionResult.Go_header = std_msgs.NewHeader()
    newLookupTransformActionResult.Go_status = actionlib_msgs.NewGoalStatus()
    newLookupTransformActionResult.Go_result = tf2_msgs.NewLookupTransformResult()
    return newLookupTransformActionResult
}

func (self *LookupTransformActionResult) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_status.Go_serialize(buff[offset:])
    offset += self.Go_result.Go_serialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionResult) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_status.Go_deserialize(buff[offset:])
    offset += self.Go_result.Go_deserialize(buff[offset:])
    return offset
}

func (self *LookupTransformActionResult) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_status.Go_serializedLength()
    length += self.Go_result.Go_serializedLength()
    return length
}

func (self *LookupTransformActionResult) Go_echo() (string) { return "" }
func (self *LookupTransformActionResult) Go_getType() (string) { return "tf2_msgs/LookupTransformActionResult" }
func (self *LookupTransformActionResult) Go_getMD5() (string) { return "5a8abe079c2126ea9966563c5cae6d29" }
func (self *LookupTransformActionResult) Go_getID() (uint32) { return 0 }
func (self *LookupTransformActionResult) Go_setID(id uint32) { }

