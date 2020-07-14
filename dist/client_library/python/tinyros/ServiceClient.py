import sys
import threading
import struct
import tinyros
import tinyros_msgs.msg

class ServiceClient(tinyros.Subscriber):
    __slots__ = ['data_req_class_', 'data_resp_class_', 'msg_req_', 'msg_resp_', 'pub_','call_req_', 'call_resp_', 'call_cond_', 'call_mutex_','call_wait_']

    gg_mutex_ = threading.Lock()

    gg_id_ = 1
 
    def __init__(self, topic_name, data_req_class, data_resp_class, \
        endpoint = tinyros_msgs.msg.TopicInfo.ID_SERVICE_CLIENT + tinyros_msgs.msg.TopicInfo.ID_SUBSCRIBER):
        super(ServiceClient, self).__init__(topic_name, None, data_resp_class)
        self.topic_ = topic_name
        self.endpoint_ = endpoint
        self.negotiated_ = False
        self.id_ = -1
        self.srv_flag_ = True
        self.call_mutex_ = threading.Lock()
        self.call_cond_ = threading.Condition()
        self.call_req_ = None
        self.call_resp_ = None
        self.call_wait_ = True
        self.data_req_class_ = data_req_class
        self.data_resp_class_ = data_resp_class
        self.msg_req_ = self.data_req_class_()
        self.msg_resp_ = self.data_resp_class_()
        self.pub_ = tinyros.Publisher(topic_name, data_req_class,  \
                   tinyros_msgs.msg.TopicInfo.ID_SERVICE_CLIENT + tinyros_msgs.msg.TopicInfo.ID_PUBLISHER)

    def call(self, request, response, duration = 3):
        self.call_mutex_.acquire()
        self.call_cond_.acquire()
        if not self.pub_.nh_.ok():
            self.call_req_ = self.call_resp_ = None
            self.call_cond_.release()
            self.call_mutex_.release()
            return False
        
        self.call_req_ = request
        self.call_resp_ = response
        self.call_wait_ = True

        ServiceClient.gg_mutex_.acquire()
        self.call_req_.setID(ServiceClient.gg_id_)
        ServiceClient.gg_id_ += 1
        ServiceClient.gg_mutex_.release()

        if self.pub_.publish(self.call_req_) <= 0:
            self.call_req_ = self.call_resp_ = None
            self.call_cond_.release()
            self.call_mutex_.release()
            return False
        self.call_cond_.wait(duration)
        self.call_req_ = self.call_resp_ = None
        result = not self.call_wait_
        self.call_cond_.release()
        self.call_mutex_.release()
        return result

    def callback(self, buff):
        self.call_cond_.acquire()
        if self.call_req_ != None and self.call_resp_ != None:
            resp_id = 0
            struct_I = struct.Struct('<I')
            (resp_id,) = struct_I.unpack(buff[0:4])
            req_id  = self.call_req_.getID()
            if resp_id == req_id:
                self.call_resp_.deserialize(buff)
                self.call_wait_ = False
                self.call_cond_.notify_all();
        self.call_cond_.release()

    def getMsgType(self):
        if self.msg_resp_ != None:
            return self.msg_resp_.getType()
        else:
            return ''

    def getMsgMD5(self):
        if self.msg_resp_ != None:
            return self.msg_resp_.getMD5()
        else:
            return ''
            
    def getEndpointType(self):
        return self.endpoint_
        
    def negotiated(self):
        return self.negotiated_ and self.pub_.negotiated_
     
