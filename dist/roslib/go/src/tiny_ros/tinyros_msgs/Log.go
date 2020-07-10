package tinyros_msgs

import (
    "encoding/json"
)

func Go_ROSDEBUG() (uint8) { return 0 }
func Go_ROSINFO() (uint8) { return 1 }
func Go_ROSWARN() (uint8) { return 2 }
func Go_ROSERROR() (uint8) { return 3 }
func Go_ROSFATAL() (uint8) { return 4 }

type Log struct {
    Go_level uint8 `json:"level"`
    Go_msg string `json:"msg"`
}

func NewLog() (*Log) {
    newLog := new(Log)
    newLog.Go_level = 0
    newLog.Go_msg = ""
    return newLog
}

func (self *Log) Go_initialize() {
    self.Go_level = 0
    self.Go_msg = ""
}

func (self *Log) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_level >> (8 * 0)) & 0xFF)
    offset += 1
    length_msg := len(self.Go_msg)
    buff[offset + 0] = byte((length_msg >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_msg >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_msg >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_msg >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_msg)], self.Go_msg)
    offset += length_msg
    return offset
}

func (self *Log) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_level = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_msg := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_msg |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_msg |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_msg |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_msg = string(buff[offset:(offset+length_msg)])
    offset += length_msg
    return offset
}

func (self *Log) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_msg := len(self.Go_msg)
    length += 4
    length += length_msg
    return length
}

func (self *Log) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Log) Go_getType() (string) { return "tinyros_msgs/Log" }
func (self *Log) Go_getMD5() (string) { return "0bd74339b4d77cb15766d831a3d15eeb" }
func (self *Log) Go_getID() (uint32) { return 0 }
func (self *Log) Go_setID(id uint32) { }

