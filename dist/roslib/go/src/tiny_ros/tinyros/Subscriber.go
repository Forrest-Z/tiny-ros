package tinyros

import (
  "reflect"
  "tiny_ros/tinyros_msgs"
)
 
type Subscriber_ interface {
    Go_callback(data []byte)
    Go_getEndpointType() (uint32)
    Go_getMsgType() (string)
    Go_getMsgMD5() (string)
    Go_negotiated() (bool)
    Go_setnegotiated(v bool)
    Go_id() (uint32)
    Go_setid(id uint32)
    Go_topic() (string)
    Go_srv_flag() (bool)
    Go_pub() (*Publisher)
}

type Subscriber struct {
    id_ uint32  `json:"id"`
    topic_ string `json:"topic"`
    negotiated_ bool `json:"negotiated"`
    srv_flag_ bool `json:"srv_flag"`
    endpoint_ uint32 `json:"endpoint"`
    cb_ func(msg Msg) `json:"cb"`
    msg_ Msg `json:"msg"`
}

func NewSubscriber(topic_name string, cb func(msg Msg), msg Msg) (*Subscriber) {
    newSubscriber := new(Subscriber)
    newSubscriber.topic_ = topic_name
    newSubscriber.cb_ = cb
    newSubscriber.negotiated_ = false
    newSubscriber.srv_flag_ = false
    newSubscriber.msg_ = msg
    newSubscriber.endpoint_ = tinyros_msgs.Go_ID_SUBSCRIBER()
    return newSubscriber
}

func (self *Subscriber) Go_callback(data []byte) {
    t := reflect.New(reflect.TypeOf(self.msg_).Elem()).Interface()
    tmsg := t.(Msg)
    tmsg.Go_initialize()
    tmsg.Go_deserialize(data)
    self.cb_(tmsg)
}

func (self *Subscriber) Go_getEndpointType() (uint32) {
    return self.endpoint_
}

func (self *Subscriber) Go_getMsgType() (string) {
    return self.msg_.Go_getType()
}

func (self *Subscriber) Go_getMsgMD5() (string) {
    return self.msg_.Go_getMD5()
}

func (self *Subscriber) Go_negotiated() (bool) {
    return self.negotiated_
}

func (self *Subscriber) Go_setnegotiated(v bool)  {
    self.negotiated_ = v
}

func (self *Subscriber) Go_id() (uint32) {
    return self.id_
}

func (self *Subscriber) Go_setid(id uint32) {
    self.id_ = id
}

func (self *Subscriber) Go_topic() (string) {
    return self.topic_
}

func (self *Subscriber) Go_srv_flag() (bool) {
    return self.srv_flag_
}

func (self *Subscriber) Go_pub() (*Publisher) {
    return nil
}
