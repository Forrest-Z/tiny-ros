package tinyros

import (
  "reflect"
  "tiny_ros/tinyros_msgs"
)

type ServiceServer struct {
    Subscriber
    resp_ Msg `json:"resp"`
    pub_ *Publisher
    callback_ func(req Msg, resp Msg) `json:"callback"`
}

func NewServiceServer(topic_name string, cb func(req, resp Msg), req Msg, resp Msg) (*ServiceServer) {
    newServiceServer := new(ServiceServer)
    newServiceServer.topic_ = topic_name
    newServiceServer.negotiated_ = false
    newServiceServer.srv_flag_ = true
    newServiceServer.msg_ = req
    newServiceServer.resp_ = resp
    newServiceServer.callback_ = cb
    newServiceServer.pub_ = NewPublisher(topic_name, resp)
    newServiceServer.pub_.endpoint_ = tinyros_msgs.Go_ID_SERVICE_SERVER() + tinyros_msgs.Go_ID_PUBLISHER()
    newServiceServer.endpoint_ = tinyros_msgs.Go_ID_SERVICE_SERVER() + tinyros_msgs.Go_ID_SUBSCRIBER()
    return newServiceServer
}

func (self *ServiceServer) Go_callback(data []byte) {
    t := reflect.New(reflect.TypeOf(self.msg_).Elem()).Interface()
    treq := t.(Msg)
    treq.Go_initialize()
    treq.Go_deserialize(data)

    t = reflect.New(reflect.TypeOf(self.resp_).Elem()).Interface()
    tresp := t.(Msg)
    tresp.Go_initialize()
    self.callback_(treq, tresp)
    tresp.Go_setID(treq.Go_getID())
    self.pub_.Go_publish(tresp)
}

func (self *ServiceServer) Go_negotiated() (bool) {
    return (self.negotiated_ && self.pub_.negotiated_)
}

func (self *ServiceServer) Go_pub() (*Publisher) {
    return self.pub_
}

