package tinyros

import (
    "fmt"
    "tiny_ros/tinyros_msgs"
)

type Publisher struct {
    id_ uint32  `json:"id"`
    topic_ string `json:"topic"`
    msg_ Msg `json:"msg"`
    nh_ NodeHandleBase `json:"nh"`
    negotiated_ bool `json:"negotiated"`
    endpoint_ uint32 `json:"endpoint"`
}

func NewPublisher(topic_name string, msg Msg) (*Publisher) {
    newPublisher := new(Publisher)
    newPublisher.topic_ = topic_name
    newPublisher.msg_ = msg
    newPublisher.nh_ = nil
    newPublisher.negotiated_ = false
    newPublisher.endpoint_ = tinyros_msgs.Go_ID_PUBLISHER()
    return newPublisher
}

func (self *Publisher) Go_publish(msg Msg) (int) {
    if self.nh_ != nil {
        return self.nh_.publish(self.id_, msg, false)
    } else {
        fmt.Println("Go_publish topic_name: ", self.topic_, ", nh is NULL, please advertise.")
        return -1
    }
}

func (self *Publisher) Go_getEndpointType() (uint32) {
    return self.endpoint_
}

func (self *Publisher) Go_getMsgType() (string) {
    return self.msg_.Go_getType()
}

func (self *Publisher) Go_getMsgMD5() (string) {
    return self.msg_.Go_getMD5()
}

func (self *Publisher) Go_negotiated() (bool) {
    return self.negotiated_
}

