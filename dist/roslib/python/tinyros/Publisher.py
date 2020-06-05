import sys
import tinyros
import tinyros_msgs.msg

class Publisher(object):
    __slots__ = ['topic_','msg_', 'data_class_','nh_', 'negotiated_', 'endpoint_', 'id_']

    def __init__(self, topic_name, data_class, \
        endpoint = tinyros_msgs.msg.TopicInfo.ID_PUBLISHER):
        self.topic_ = topic_name
        self.endpoint_ = endpoint
        self.negotiated_ = False
        self.nh_ = None
        self.id_ = -1
        try:  
            self.data_class_ = data_class
            self.msg_ = self.data_class_()
        except:  
            self.data_class_ = self.msg_ = None
            raise RuntimeError('Publisher::init: no module of: %s found'%(self.data_class_)) 

    def publish(self, msg, islog = False):
        if self.nh_ != None:
            return self.nh_.publish(self.id_, msg, islog)
        else:
            self.nh_.logerror("Publisher::publish topic_name: %s, nh is None, please advertise." % self.topic_)
            return -1

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
        return self.negotiated_;

