import sys
import tinyros
import tinyros_msgs.msg

class Subscriber(object):
    __slots__ = ['topic_','srv_flag_', 'negotiated_', 'endpoint_', 'id_', 'cb_', 'data_class_', 'msg_']

    def __init__(self, topic_name, cb, data_class, \
        endpoint = tinyros_msgs.msg.TopicInfo.ID_SUBSCRIBER):
        self.topic_ = topic_name
        self.endpoint_ = endpoint
        self.negotiated_ = False
        self.id_ = -1
        self.cb_ = cb
        self.srv_flag_ = False
        self.data_class_ = data_class
        self.msg_ = self.data_class_()
   
    def callback(self, buff):
        msg = self.data_class_()
        msg.deserialize(buff)
        self.cb_(msg)

    def getMsgType(self):
        if self.msg_ != None:
            return self.msg_.getType()
        else:
            return ''
            
    def getMsgMD5(self):
        if self.msg_ != None:
            return self.msg_.getMD5()
        else:
            return ''
            
    def getEndpointType(self):
        return self.endpoint_

    def negotiated(self):
        return self.negotiated_


