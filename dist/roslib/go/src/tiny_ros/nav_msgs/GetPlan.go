package nav_msgs

import (
    "encoding/json"
    "encoding/binary"
    "tiny_ros/geometry_msgs"
    "math"
)



type GetPlanRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_start *geometry_msgs.PoseStamped `json:"start"`
    Go_goal *geometry_msgs.PoseStamped `json:"goal"`
    Go_tolerance float32 `json:"tolerance"`
}

func NewGetPlanRequest() (*GetPlanRequest) {
    newGetPlanRequest := new(GetPlanRequest)
    newGetPlanRequest.Go_start = geometry_msgs.NewPoseStamped()
    newGetPlanRequest.Go_goal = geometry_msgs.NewPoseStamped()
    newGetPlanRequest.Go_tolerance = 0.0
    newGetPlanRequest.__id__ = 0
    return newGetPlanRequest
}

func (self *GetPlanRequest) Go_initialize() {
    self.Go_start = geometry_msgs.NewPoseStamped()
    self.Go_goal = geometry_msgs.NewPoseStamped()
    self.Go_tolerance = 0.0
    self.__id__ = 0
}

func (self *GetPlanRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_start.Go_serialize(buff[offset:])
    offset += self.Go_goal.Go_serialize(buff[offset:])
    bits_tolerance := math.Float32bits(self.Go_tolerance)
    binary.LittleEndian.PutUint32(buff[offset:], bits_tolerance)
    offset += 4
    return offset
}

func (self *GetPlanRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_start.Go_deserialize(buff[offset:])
    offset += self.Go_goal.Go_deserialize(buff[offset:])
    bits_tolerance := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_tolerance = math.Float32frombits(bits_tolerance)
    offset += 4
    return offset
}

func (self *GetPlanRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_start.Go_serializedLength()
    length += self.Go_goal.Go_serializedLength()
    length += 4
    return length
}

func (self *GetPlanRequest) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPlanRequest) Go_getType() (string) { return "nav_msgs/GetPlan" }
func (self *GetPlanRequest) Go_getMD5() (string) { return "557d5ea947f7761284cf7abef1cd7227" }
func (self *GetPlanRequest) Go_getID() (uint32) { return self.__id__ }
func (self *GetPlanRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////


type GetPlanResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_plan *Path `json:"plan"`
}

func NewGetPlanResponse() (*GetPlanResponse) {
    newGetPlanResponse := new(GetPlanResponse)
    newGetPlanResponse.Go_plan = NewPath()
    newGetPlanResponse.__id__ = 0
    return newGetPlanResponse
}

func (self *GetPlanResponse) Go_initialize() {
    self.Go_plan = NewPath()
    self.__id__ = 0
}

func (self *GetPlanResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_plan.Go_serialize(buff[offset:])
    return offset
}

func (self *GetPlanResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    offset += self.Go_plan.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetPlanResponse) Go_serializedLength() (int) {
    length := 0
    length += self.Go_plan.Go_serializedLength()
    return length
}

func (self *GetPlanResponse) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetPlanResponse) Go_getType() (string) { return "nav_msgs/GetPlan" }
func (self *GetPlanResponse) Go_getMD5() (string) { return "67c62b8c931eabfe35c88aed4b8f1258" }
func (self *GetPlanResponse) Go_getID() (uint32) { return self.__id__ }
func (self *GetPlanResponse) Go_setID(id uint32) { self.__id__ = id }


