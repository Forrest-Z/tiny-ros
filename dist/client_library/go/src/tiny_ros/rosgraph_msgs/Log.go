package rosgraph_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)

func Go_DEBUG() (byte) { return 1  }
func Go_INFO() (byte) { return 2   }
func Go_WARN() (byte) { return 4   }
func Go_ERROR() (byte) { return 8  }
func Go_FATAL() (byte) { return 16  }

type Log struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_level byte `json:"level"`
    Go_name string `json:"name"`
    Go_msg string `json:"msg"`
    Go_file string `json:"file"`
    Go_function string `json:"function"`
    Go_line uint32 `json:"line"`
    Go_topics []string `json:"topics"`
}

func NewLog() (*Log) {
    newLog := new(Log)
    newLog.Go_header = std_msgs.NewHeader()
    newLog.Go_level = 0
    newLog.Go_name = ""
    newLog.Go_msg = ""
    newLog.Go_file = ""
    newLog.Go_function = ""
    newLog.Go_line = 0
    newLog.Go_topics = []string{}
    return newLog
}

func (self *Log) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_level = 0
    self.Go_name = ""
    self.Go_msg = ""
    self.Go_file = ""
    self.Go_function = ""
    self.Go_line = 0
    self.Go_topics = []string{}
}

func (self *Log) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_level >> (8 * 0)) & 0xFF)
    offset += 1
    length_name := len(self.Go_name)
    buff[offset + 0] = byte((length_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_name)], self.Go_name)
    offset += length_name
    length_msg := len(self.Go_msg)
    buff[offset + 0] = byte((length_msg >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_msg >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_msg >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_msg >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_msg)], self.Go_msg)
    offset += length_msg
    length_file := len(self.Go_file)
    buff[offset + 0] = byte((length_file >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_file >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_file >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_file >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_file)], self.Go_file)
    offset += length_file
    length_function := len(self.Go_function)
    buff[offset + 0] = byte((length_function >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_function >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_function >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_function >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_function)], self.Go_function)
    offset += length_function
    buff[offset + 0] = byte((self.Go_line >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_line >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_line >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_line >> (8 * 3)) & 0xFF)
    offset += 4
    length_topics := len(self.Go_topics)
    buff[offset + 0] = byte((length_topics >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_topics >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_topics >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_topics >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_topics; i++ {
        length_topicsi := len(self.Go_topics[i])
        buff[offset + 0] = byte((length_topicsi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_topicsi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_topicsi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_topicsi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_topicsi)], self.Go_topics[i])
        offset += length_topicsi
    }
    return offset
}

func (self *Log) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    self.Go_level = byte(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_name = string(buff[offset:(offset+length_name)])
    offset += length_name
    length_msg := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_msg |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_msg |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_msg |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_msg = string(buff[offset:(offset+length_msg)])
    offset += length_msg
    length_file := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_file |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_file |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_file |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_file = string(buff[offset:(offset+length_file)])
    offset += length_file
    length_function := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_function |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_function |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_function |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_function = string(buff[offset:(offset+length_function)])
    offset += length_function
    self.Go_line = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_line |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_line |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_line |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    length_topics := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_topics |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_topics |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_topics |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_topics = make([]string, length_topics)
    for i := 0; i < length_topics; i++ {
        length_topicsi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_topicsi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_topicsi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_topicsi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_topics[i] = string(buff[offset:(offset+length_topicsi)])
        offset += length_topicsi
    }
    return offset
}

func (self *Log) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 1
    length_name := len(self.Go_name)
    length += 4
    length += length_name
    length_msg := len(self.Go_msg)
    length += 4
    length += length_msg
    length_file := len(self.Go_file)
    length += 4
    length += length_file
    length_function := len(self.Go_function)
    length += 4
    length += length_function
    length += 4
    length += 4
    length_topics := len(self.Go_topics)
    for i := 0; i < length_topics; i++ {
        length_topicsi := len(self.Go_topics[i])
        length += 4
        length += length_topicsi
    }
    return length
}

func (self *Log) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *Log) Go_getType() (string) { return "rosgraph_msgs/Log" }
func (self *Log) Go_getMD5() (string) { return "2de9daf47e984009074d74dbdd492d49" }
func (self *Log) Go_getID() (uint32) { return 0 }
func (self *Log) Go_setID(id uint32) { }

