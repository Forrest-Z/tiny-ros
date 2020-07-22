package tinyros_msgs

import (
    "encoding/json"
)

func Go_ID_PUBLISHER() (uint32) { return 0 }
func Go_ID_SUBSCRIBER() (uint32) { return 1 }
func Go_ID_SERVICE_SERVER() (uint32) { return 2 }
func Go_ID_SERVICE_CLIENT() (uint32) { return 4 }
func Go_ID_ROSTOPIC_REQUEST() (uint32) { return 6 }
func Go_ID_ROSSERVICE_REQUEST() (uint32) { return 7 }
func Go_ID_LOG() (uint32) { return 8 }
func Go_ID_TIME() (uint32) { return 9 }
func Go_ID_NEGOTIATED() (uint32) { return 10 }
func Go_ID_SESSION_ID() (uint32) { return 11 }

type TopicInfo struct {
    Go_topic_id uint32 `json:"topic_id"`
    Go_topic_name string `json:"topic_name"`
    Go_message_type string `json:"message_type"`
    Go_md5sum string `json:"md5sum"`
    Go_buffer_size int32 `json:"buffer_size"`
    Go_negotiated bool `json:"negotiated"`
    Go_node string `json:"node"`
}

func NewTopicInfo() (*TopicInfo) {
    newTopicInfo := new(TopicInfo)
    newTopicInfo.Go_topic_id = 0
    newTopicInfo.Go_topic_name = ""
    newTopicInfo.Go_message_type = ""
    newTopicInfo.Go_md5sum = ""
    newTopicInfo.Go_buffer_size = 0
    newTopicInfo.Go_negotiated = false
    newTopicInfo.Go_node = ""
    return newTopicInfo
}

func (self *TopicInfo) Go_initialize() {
    self.Go_topic_id = 0
    self.Go_topic_name = ""
    self.Go_message_type = ""
    self.Go_md5sum = ""
    self.Go_buffer_size = 0
    self.Go_negotiated = false
    self.Go_node = ""
}

func (self *TopicInfo) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_topic_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_topic_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_topic_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_topic_id >> (8 * 3)) & 0xFF)
    offset += 4
    length_topic_name := len(self.Go_topic_name)
    buff[offset + 0] = byte((length_topic_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_topic_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_topic_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_topic_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_topic_name)], self.Go_topic_name)
    offset += length_topic_name
    length_message_type := len(self.Go_message_type)
    buff[offset + 0] = byte((length_message_type >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_message_type >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_message_type >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_message_type >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_message_type)], self.Go_message_type)
    offset += length_message_type
    length_md5sum := len(self.Go_md5sum)
    buff[offset + 0] = byte((length_md5sum >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_md5sum >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_md5sum >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_md5sum >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_md5sum)], self.Go_md5sum)
    offset += length_md5sum
    buff[offset + 0] = byte((self.Go_buffer_size >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_buffer_size >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_buffer_size >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_buffer_size >> (8 * 3)) & 0xFF)
    offset += 4
    if self.Go_negotiated {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    length_node := len(self.Go_node)
    buff[offset + 0] = byte((length_node >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_node >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_node >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_node >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_node)], self.Go_node)
    offset += length_node
    return offset
}

func (self *TopicInfo) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_topic_id = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_topic_id |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_topic_id |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_topic_id |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_topic_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_topic_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_topic_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_topic_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_topic_name = string(buff[offset:(offset+length_topic_name)])
    offset += length_topic_name
    length_message_type := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_message_type |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_message_type |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_message_type |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_message_type = string(buff[offset:(offset+length_message_type)])
    offset += length_message_type
    length_md5sum := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_md5sum |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_md5sum |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_md5sum |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_md5sum = string(buff[offset:(offset+length_md5sum)])
    offset += length_md5sum
    self.Go_buffer_size = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_buffer_size |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_buffer_size |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_buffer_size |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    if (buff[offset] & 0xFF) != 0 {
        self.Go_negotiated = true
    } else {
        self.Go_negotiated = false
    }
    offset += 1
    length_node := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_node |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_node |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_node |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_node = string(buff[offset:(offset+length_node)])
    offset += length_node
    return offset
}

func (self *TopicInfo) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_topic_name := len(self.Go_topic_name)
    length += 4
    length += length_topic_name
    length_message_type := len(self.Go_message_type)
    length += 4
    length += length_message_type
    length_md5sum := len(self.Go_md5sum)
    length += 4
    length += length_md5sum
    length += 4
    length += 1
    length_node := len(self.Go_node)
    length += 4
    length += length_node
    return length
}

func (self *TopicInfo) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TopicInfo) Go_getType() (string) { return "tinyros_msgs/TopicInfo" }
func (self *TopicInfo) Go_getMD5() (string) { return "76d40676946fcde66f228def7575451a" }
func (self *TopicInfo) Go_getID() (uint32) { return 0 }
func (self *TopicInfo) Go_setID(id uint32) { }

