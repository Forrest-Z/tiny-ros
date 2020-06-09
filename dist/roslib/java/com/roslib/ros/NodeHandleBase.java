package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public abstract class NodeHandleBase {
    protected static final int MAX_SUBSCRIBERS = 100;
    protected static final int MAX_PUBLISHERS = 100;
    protected static final int INPUT_SIZE = 65*1024;
    protected static final int OUTPUT_SIZE = 65*1024;
    protected static final int SPIN_OK = 0;
    protected static final int SPIN_ERR = -1;
    protected static final int PROTOCOL_VER = 0xb9;
    protected boolean spin_ = true;
    boolean initNode(java.lang.String portName) { return false; }
    int publish(long id, com.roslib.ros.Msg msg, boolean islog) { return -1; }
    int publish(long id, com.roslib.ros.Msg msg) { return -1; }
    int spin() { return -1; }
    void exit() {}
    boolean ok() { return false; }
    
    void negotiateTopics(Publisher p) {
        TopicInfo ti = new TopicInfo();
        ti.topic_id = p.id_;
        ti.topic_name = p.topic_;
        ti.message_type = p.msg_.getType();
        ti.md5sum = p.msg_.getMD5();
        ti.buffer_size = OUTPUT_SIZE;
        publish(p.getEndpointType(), ti);
    }

    void negotiateTopics(SubscriberT s) {
        TopicInfo ti = new TopicInfo();
        ti.topic_id = s.id_;
        ti.topic_name = s.topic_;
        ti.message_type = s.getMsgType();
        ti.md5sum = s.getMsgMD5();
        ti.buffer_size = INPUT_SIZE;
        publish(s.getEndpointType(), ti);
    }
}
