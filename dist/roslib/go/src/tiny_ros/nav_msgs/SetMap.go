package nav_msgs

import (
    "nav_msgs/OccupancyGrid"
    "geometry_msgs/PoseWithCovarianceStamped"
)


type SetMapRequest struct {
    __id__ uint32 `json:"__id__"`
    Go_map nav_msgs.OccupancyGrid `json:"map"`
    Go_initial_pose geometry_msgs.PoseWithCovarianceStamped `json:"initial_pose"`
}

func NewSetMapRequest() (*SetMapRequest) {
    newSetMapRequest := new(SetMapRequest)
    newSetMapRequest.Go_map = nav_msgs.NewOccupancyGrid()
    newSetMapRequest.Go_initial_pose = geometry_msgs.NewPoseWithCovarianceStamped()
    newSetMapRequest.__id__ = 0
    return newSetMapRequest
}

func (self *SetMapRequest) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    offset += self.Go_map.Go_serialize(buff[offset:])
    offset += self.Go_initial_pose.Go_serialize(buff[offset:])
    return offset
}

func (self *SetMapRequest) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32((buff[offset + 0] & 0xFF) << (8 * 0))
    self.__id__ |=  uint32((buff[offset + 1] & 0xFF) << (8 * 1))
    self.__id__ |=  uint32((buff[offset + 2] & 0xFF) << (8 * 2))
    self.__id__ |=  uint32((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    offset += self.Go_map.Go_deserialize(buff[offset:])
    offset += self.Go_initial_pose.Go_deserialize(buff[offset:])
    return offset
}

func (self *SetMapRequest) Go_serializedLength() (int) {
    length := 0
    length += self.Go_map.Go_serializedLength()
    length += self.Go_initial_pose.Go_serializedLength()
    return length
}

func (self *SetMapRequest) Go_echo() (string) { return "" }
func (self *SetMapRequest) Go_getType() (string) { return "nav_msgs/SetMap" }
func (self *SetMapRequest) Go_getMD5() (string) { return "946e1bd68c9db117a530a571e33d9e49" }
func (self *SetMapRequest) Go_getID() (uint32) { return self.__id__ }
func (self *SetMapRequest) Go_setID(id uint32) { self.__id__ = id }


///////////////////////////////////////////////////////////////////////////

type SetMapResponse struct {
    __id__ uint32 `json:"__id__"`
    Go_success bool `json:"success"`
}

func NewSetMapResponse() (*SetMapResponse) {
    newSetMapResponse := new(SetMapResponse)
    newSetMapResponse.Go_success = false
    newSetMapResponse.__id__ = 0
    return newSetMapResponse
}

func (self *SetMapResponse) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_success >> (8 * 0)) & 0xFF)
    offset += 1
    return offset
}

func (self *SetMapResponse) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.__id__ =  uint32((buff[offset + 0] & 0xFF) << (8 * 0))
    self.__id__ |=  uint32((buff[offset + 1] & 0xFF) << (8 * 1))
    self.__id__ |=  uint32((buff[offset + 2] & 0xFF) << (8 * 2))
    self.__id__ |=  uint32((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4
    self.Go_success = bool((buff[offset + 0] & 0xFF) << (8 * 0))
    offset += 1
    return offset
}

func (self *SetMapResponse) Go_serializedLength() (int) {
    length := 0
    length += 1
    return length
}

func (self *SetMapResponse) Go_echo() (string) { return "" }
func (self *SetMapResponse) Go_getType() (string) { return "nav_msgs/SetMap" }
func (self *SetMapResponse) Go_getMD5() (string) { return "1e32607e79013262dafbbac9044e9cda" }
func (self *SetMapResponse) Go_getID() (uint32) { return self.__id__ }
func (self *SetMapResponse) Go_setID(id uint32) { self.__id__ = id }


