package nav_msgs

import (
    "encoding/json"
)


type GetMapResult struct {
    Go_map *OccupancyGrid `json:"map"`
}

func NewGetMapResult() (*GetMapResult) {
    newGetMapResult := new(GetMapResult)
    newGetMapResult.Go_map = NewOccupancyGrid()
    return newGetMapResult
}

func (self *GetMapResult) Go_initialize() {
    self.Go_map = NewOccupancyGrid()
}

func (self *GetMapResult) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_map.Go_serialize(buff[offset:])
    return offset
}

func (self *GetMapResult) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_map.Go_deserialize(buff[offset:])
    return offset
}

func (self *GetMapResult) Go_serializedLength() (int) {
    length := 0
    length += self.Go_map.Go_serializedLength()
    return length
}

func (self *GetMapResult) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GetMapResult) Go_getType() (string) { return "nav_msgs/GetMapResult" }
func (self *GetMapResult) Go_getMD5() (string) { return "dd8eb0759b1a400b141d7f3238732c4d" }
func (self *GetMapResult) Go_getID() (uint32) { return 0 }
func (self *GetMapResult) Go_setID(id uint32) { }

