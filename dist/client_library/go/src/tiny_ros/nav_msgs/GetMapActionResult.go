package nav_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "tiny_ros/actionlib_msgs"
)


type GetMapActionResult struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_status *actionlib_msgs.GoalStatus `json:"status"`
    Go_result *GetMapResult `json:"result"`
}

func NewGetMapActionResult() (*GetMapActionResult) {
    newGetMapActionResult := new(GetMapActionResult)
    newGetMapActionResult.Go_header = std_msgs.NewHeader()
    newGetMapActionResult.Go_status = actionlib_msgs.NewGoalStatus()
    newGetMapActionResult.Go_result = NewGetMapResult()
    return newGetMapActionResult
}

func (self *GetMapActionResult) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_status = actionlib_msgs.NewGoalStatus()
    self.Go_result = NewGetMapResult()
}

func (self *GetMapActionResult) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    offset += self.Go_status.Go_serialize(buff[offset:])
    offset += self.Go_result.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapActionResult) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    offset += self.Go_status.Go_deserialize(buff[offset:])
    offset += self.Go_result.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapActionResult) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += self.Go_status.Go_serializedLength()
    length += self.Go_result.Go_serializedLength()
    return length
}

func (self *GetMapActionResult) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapActionResult) Go_getType() (string) { return "nav_msgs/GetMapActionResult" }
func (self *GetMapActionResult) Go_getMD5() (string) { return "9c9f64758f2627a010c16b17ea745028" }
func (self *GetMapActionResult) Go_getID() (uint32) { return 0 }
func (self *GetMapActionResult) Go_setID(id uint32) { }

