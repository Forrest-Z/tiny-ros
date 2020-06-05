import sys
import tinyros
import tinyros_msgs.msg

class ServiceServer(tinyros.Subscriber):
    __slots__ = ['data_req_class_', 'data_resp_class_', 'msg_req_', 'msg_resp_', 'pub_']
    def __init__(self, topic_name, cb, data_req_class, data_resp_class, \
        endpoint = tinyros_msgs.msg.TopicInfo.ID_SERVICE_SERVER + tinyros_msgs.msg.TopicInfo.ID_SUBSCRIBER):
        super(ServiceServer, self).__init__(topic_name, cb, data_req_class)
        self.topic_ = topic_name
        self.endpoint_ = endpoint
        self.negotiated_ = False
        self.id_ = -1
        self.cb_ = cb
        self.srv_flag_ = True
        self.data_req_class_ = data_req_class
        self.data_resp_class_ = data_resp_class
        self.msg_req_ = self.data_req_class_()
        self.msg_resp_ = self.data_resp_class_()
        self.pub_ = tinyros.Publisher(topic_name, data_resp_class,  \
                   tinyros_msgs.msg.TopicInfo.ID_SERVICE_SERVER + tinyros_msgs.msg.TopicInfo.ID_PUBLISHER)

    def callback(self, buff):
        msg_req =  self.data_req_class_()
        msg_resp = self.data_resp_class_()
        msg_req.deserialize(buff)
        self.cb_(msg_req, msg_resp)
        msg_resp.setID(msg_req.getID())
        self.pub_.publish(msg_resp)

    def getMsgType(self):
        if self.msg_req_ != None:
            return self.msg_req_.getType()
        else:
            return ''

    def getMsgMD5(self):
        if self.msg_req_ != None:
            return self.msg_req_.getMD5()
        else:
            return ''

    def getEndpointType(self):
        return self.endpoint_
  
    def negotiated(self):
        return (self.negotiated_ and self.pub.negotiated_)

