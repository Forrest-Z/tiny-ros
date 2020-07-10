package rosgraph_msgs

import (
    "encoding/json"
    "tiny_ros/tinyros/time"
)


type TopicStatistics struct {
    Go_topic string `json:"topic"`
    Go_node_pub string `json:"node_pub"`
    Go_node_sub string `json:"node_sub"`
    Go_window_start *rostime.Time `json:"window_start"`
    Go_window_stop *rostime.Time `json:"window_stop"`
    Go_delivered_msgs int32 `json:"delivered_msgs"`
    Go_dropped_msgs int32 `json:"dropped_msgs"`
    Go_traffic int32 `json:"traffic"`
    Go_period_mean *rostime.Duration `json:"period_mean"`
    Go_period_stddev *rostime.Duration `json:"period_stddev"`
    Go_period_max *rostime.Duration `json:"period_max"`
    Go_stamp_age_mean *rostime.Duration `json:"stamp_age_mean"`
    Go_stamp_age_stddev *rostime.Duration `json:"stamp_age_stddev"`
    Go_stamp_age_max *rostime.Duration `json:"stamp_age_max"`
}

func NewTopicStatistics() (*TopicStatistics) {
    newTopicStatistics := new(TopicStatistics)
    newTopicStatistics.Go_topic = ""
    newTopicStatistics.Go_node_pub = ""
    newTopicStatistics.Go_node_sub = ""
    newTopicStatistics.Go_window_start = rostime.NewTime()
    newTopicStatistics.Go_window_stop = rostime.NewTime()
    newTopicStatistics.Go_delivered_msgs = 0
    newTopicStatistics.Go_dropped_msgs = 0
    newTopicStatistics.Go_traffic = 0
    newTopicStatistics.Go_period_mean = rostime.NewDuration()
    newTopicStatistics.Go_period_stddev = rostime.NewDuration()
    newTopicStatistics.Go_period_max = rostime.NewDuration()
    newTopicStatistics.Go_stamp_age_mean = rostime.NewDuration()
    newTopicStatistics.Go_stamp_age_stddev = rostime.NewDuration()
    newTopicStatistics.Go_stamp_age_max = rostime.NewDuration()
    return newTopicStatistics
}

func (self *TopicStatistics) Go_initialize() {
    self.Go_topic = ""
    self.Go_node_pub = ""
    self.Go_node_sub = ""
    self.Go_window_start = rostime.NewTime()
    self.Go_window_stop = rostime.NewTime()
    self.Go_delivered_msgs = 0
    self.Go_dropped_msgs = 0
    self.Go_traffic = 0
    self.Go_period_mean = rostime.NewDuration()
    self.Go_period_stddev = rostime.NewDuration()
    self.Go_period_max = rostime.NewDuration()
    self.Go_stamp_age_mean = rostime.NewDuration()
    self.Go_stamp_age_stddev = rostime.NewDuration()
    self.Go_stamp_age_max = rostime.NewDuration()
}

func (self *TopicStatistics) Go_serialize(buff []byte) (int) {
    offset := 0
    length_topic := len(self.Go_topic)
    buff[offset + 0] = byte((length_topic >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_topic >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_topic >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_topic >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_topic)], self.Go_topic)
    offset += length_topic
    length_node_pub := len(self.Go_node_pub)
    buff[offset + 0] = byte((length_node_pub >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_node_pub >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_node_pub >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_node_pub >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_node_pub)], self.Go_node_pub)
    offset += length_node_pub
    length_node_sub := len(self.Go_node_sub)
    buff[offset + 0] = byte((length_node_sub >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_node_sub >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_node_sub >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_node_sub >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_node_sub)], self.Go_node_sub)
    offset += length_node_sub
    buff[offset + 0] = byte((self.Go_window_start.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_window_start.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_window_start.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_window_start.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_window_start.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_window_start.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_window_start.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_window_start.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_window_stop.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_window_stop.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_window_stop.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_window_stop.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_window_stop.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_window_stop.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_window_stop.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_window_stop.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_delivered_msgs >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_delivered_msgs >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_delivered_msgs >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_delivered_msgs >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_dropped_msgs >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_dropped_msgs >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_dropped_msgs >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_dropped_msgs >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_traffic >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_traffic >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_traffic >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_traffic >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_mean.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_mean.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_mean.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_mean.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_mean.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_mean.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_mean.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_mean.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_stddev.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_stddev.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_stddev.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_stddev.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_stddev.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_stddev.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_stddev.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_stddev.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_max.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_max.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_max.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_max.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_period_max.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_period_max.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_period_max.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_period_max.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_mean.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_mean.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_mean.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_mean.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_mean.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_mean.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_mean.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_mean.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_stddev.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_stddev.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_stddev.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_stddev.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_stddev.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_stddev.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_stddev.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_stddev.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_max.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_max.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_max.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_max.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_stamp_age_max.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_stamp_age_max.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_stamp_age_max.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_stamp_age_max.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *TopicStatistics) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_topic := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_topic |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_topic |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_topic |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_topic = string(buff[offset:(offset+length_topic)])
    offset += length_topic
    length_node_pub := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_node_pub |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_node_pub |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_node_pub |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_node_pub = string(buff[offset:(offset+length_node_pub)])
    offset += length_node_pub
    length_node_sub := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_node_sub |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_node_sub |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_node_sub |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_node_sub = string(buff[offset:(offset+length_node_sub)])
    offset += length_node_sub
    self.Go_window_start.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_window_start.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_window_start.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_window_start.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_window_start.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_window_start.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_window_start.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_window_start.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_window_stop.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_window_stop.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_window_stop.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_window_stop.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_window_stop.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_window_stop.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_window_stop.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_window_stop.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_delivered_msgs = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_delivered_msgs |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_delivered_msgs |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_delivered_msgs |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_dropped_msgs = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_dropped_msgs |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_dropped_msgs |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_dropped_msgs |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_traffic = int32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_traffic |= int32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_traffic |= int32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_traffic |= int32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_mean.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_mean.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_mean.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_mean.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_mean.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_mean.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_mean.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_mean.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_stddev.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_stddev.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_stddev.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_stddev.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_stddev.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_stddev.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_stddev.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_stddev.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_max.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_max.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_max.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_max.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_period_max.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_period_max.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_period_max.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_period_max.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_mean.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_mean.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_mean.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_mean.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_mean.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_mean.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_mean.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_mean.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_stddev.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_stddev.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_stddev.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_stddev.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_stddev.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_stddev.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_stddev.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_stddev.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_max.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_max.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_max.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_max.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_stamp_age_max.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_stamp_age_max.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_stamp_age_max.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_stamp_age_max.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *TopicStatistics) Go_serializedLength() (int) {
    length := 0
    length_topic := len(self.Go_topic)
    length += 4
    length += length_topic
    length_node_pub := len(self.Go_node_pub)
    length += 4
    length += length_node_pub
    length_node_sub := len(self.Go_node_sub)
    length += 4
    length += length_node_sub
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    return length
}

func (self *TopicStatistics) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TopicStatistics) Go_getType() (string) { return "rosgraph_msgs/TopicStatistics" }
func (self *TopicStatistics) Go_getMD5() (string) { return "8b30d3f22284a3bee7679b7194bd38a3" }
func (self *TopicStatistics) Go_getID() (uint32) { return 0 }
func (self *TopicStatistics) Go_setID(id uint32) { }

