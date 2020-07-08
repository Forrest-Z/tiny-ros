package tinyros

import (
  "sync"
  "time"
  "tiny_ros/tinyros_msgs"
)

var gg_id_ uint32 = 1

var gg_mutex_ *sync.Mutex = new(sync.Mutex)

type ServiceClient struct {
    Subscriber
    call_req_ Msg `json:"call_req"`
    call_resp_ Msg `json:"call_resp"`
    call_wait_ bool `json:"call_wait"`
    pub_ *Publisher
    cond_ *sync.Cond `json:"cond"`
    g_mutex_ *sync.Mutex `json:"g_mutex"`
}

func NewServiceClient(topic_name string, req Msg, resp Msg) (*ServiceClient) {
    newServiceClient := new(ServiceClient)
    newServiceClient.topic_ = topic_name
    newServiceClient.negotiated_ = false
    newServiceClient.srv_flag_ = true
    newServiceClient.msg_ = resp
    newServiceClient.g_mutex_ = new(sync.Mutex)
    newServiceClient.cond_ = sync.NewCond(new(sync.Mutex))
    newServiceClient.call_req_ = nil
    newServiceClient.call_resp_ = nil
    newServiceClient.call_wait_ = true
    newServiceClient.pub_ = NewPublisher(topic_name, req)
    newServiceClient.pub_.endpoint_ = tinyros_msgs.Go_ID_SERVICE_CLIENT() + tinyros_msgs.Go_ID_PUBLISHER()
    newServiceClient.endpoint_ = tinyros_msgs.Go_ID_SERVICE_CLIENT() + tinyros_msgs.Go_ID_SUBSCRIBER()
    return newServiceClient
}

func (self *ServiceClient) Go_call(request Msg, response Msg, args ...int) (bool) {
    var duration int
    if len(args) <=0 {
        duration = 3
    } else {
        duration = args[0]
    }
    
    self.g_mutex_.Lock()
    self.cond_.L.Lock()

    if !self.pub_.nh_.ok() {
      self.call_req_ = nil
      self.call_resp_ = nil
      self.cond_.L.Unlock()
      self.g_mutex_.Unlock()
      return false;
    }
    
    self.call_req_ = request
    self.call_resp_ = response

    gg_mutex_.Lock()
    self.call_req_.Go_setID(gg_id_)
    gg_id_ += 1
    gg_mutex_.Unlock()
    
    if self.pub_.Go_publish(request) <= 0 {
      self.call_req_ = nil
      self.call_resp_ = nil
      self.cond_.L.Unlock()
      self.g_mutex_.Unlock()
      return false
    }

    self.call_wait_ = true
    
    timer := time.AfterFunc(time.Duration(duration) * time.Second , func() {
		self.cond_.Signal()
	})
	defer timer.Stop()
	
	self.cond_.Wait()
	
    result := !self.call_wait_
    self.call_req_ = nil
    self.call_resp_ = nil
    self.cond_.L.Unlock()
    self.g_mutex_.Unlock()
    return result
}

func (self *ServiceClient) Go_callback(data []byte) {
    if self.call_resp_ != nil && self.call_req_ != nil {
        self.cond_.L.Lock()
        if self.call_resp_ != nil && self.call_req_ != nil {
            var req_id uint32 = self.call_req_.Go_getID()
            var resp_id uint32 = uint32(data[0] & 0xFF) << (8 * 0)
            resp_id |= uint32(data[1] & 0xFF) << (8 * 1)
            resp_id |= uint32(data[2] & 0xFF) << (8 * 2)
            resp_id |= uint32(data[3] & 0xFF) << (8 * 3)
            if req_id == resp_id {
                self.call_resp_.Go_deserialize(data)
                self.call_wait_ = false
                self.cond_.Signal()
            }
        }
        self.cond_.L.Unlock()
    }
}

func (self *ServiceClient) Go_negotiated() (bool) {
    return (self.negotiated_ && self.pub_.negotiated_)
}

func (self *ServiceClient) Go_pub() (*Publisher) {
    return self.pub_
}
