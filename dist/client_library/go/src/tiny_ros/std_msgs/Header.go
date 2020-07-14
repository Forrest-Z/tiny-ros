package std_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type Header struct {
    Go_seq uint32 `json:"seq"`
    Go_stamp *rostime.Time `json:"stamp"`
    Go_frame_id string `json:"frame_id"`
}

func NewHeader() (*Header) {
    newHeader := new(Header)
    newHeader.Go_seq = 0
    newHeader.Go_stamp = rostime.NewTime()
    newHeader.Go_frame_id = ""
    return newHeader
}

func (self *Header) Go_initialize() {
    self.Go_seq = 0
    self.Go_stamp = rostime.NewTime()
    self.Go_frame_id = ""
}

func (self *Header) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_seq >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_seq >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_seq >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_seq >> (8 * 3)) & 0xFF)
    offset += 4
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
    length_frame_id := len(self.Go_frame_id)
    buff[offset + 0] = byte((length_frame_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_frame_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_frame_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_frame_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_frame_id)], self.Go_frame_id)
    offset += length_frame_id
    return offset
}

func (self *Header) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_seq = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_seq |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_seq |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_seq |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
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
    length_frame_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_frame_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_frame_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_frame_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_frame_id = string(buff[offset:(offset+length_frame_id)])
    offset += length_frame_id
    return offset
}

func (self *Header) Go_serializedLength() (int) {
    length := 0
    length += 4
    length += 4
    length += 4
    length_frame_id := len(self.Go_frame_id)
    length += 4
    length += length_frame_id
    return length
}

func (self *Header) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Header) Go_getType() (string) { return "std_msgs/Header" }
func (self *Header) Go_getMD5() (string) { return "d33440e88be7b5b8255fc61ebbca06ad" }
func (self *Header) Go_getID() (uint32) { return 0 }
func (self *Header) Go_setID(id uint32) { }

