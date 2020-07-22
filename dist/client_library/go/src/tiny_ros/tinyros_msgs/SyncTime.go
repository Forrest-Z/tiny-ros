package tinyros_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type SyncTime struct {
    Go_tick uint32 `json:"tick"`
    Go_data *rostime.Time `json:"data"`
}

func NewSyncTime() (*SyncTime) {
    newSyncTime := new(SyncTime)
    newSyncTime.Go_tick = 0
    newSyncTime.Go_data = rostime.NewTime()
    return newSyncTime
}

func (self *SyncTime) Go_initialize() {
    self.Go_tick = 0
    self.Go_data = rostime.NewTime()
}

func (self *SyncTime) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_tick >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_tick >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_tick >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_tick >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_data.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_data.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_data.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_data.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_data.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *SyncTime) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_tick = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_tick |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_tick |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_tick |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_data.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_data.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_data.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_data.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *SyncTime) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    return length
}

func (self *SyncTime) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SyncTime) Go_getType() (string) { return "tinyros_msgs/SyncTime" }
func (self *SyncTime) Go_getMD5() (string) { return "45bf702585c65b1bb762993bdbb1de6f" }
func (self *SyncTime) Go_getID() (uint32) { return 0 }
func (self *SyncTime) Go_setID(id uint32) { }

