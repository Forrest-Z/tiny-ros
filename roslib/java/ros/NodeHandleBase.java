package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public abstract class NodeHandleBase {
    protected static final int MAX_SUBSCRIBERS = 100;
    protected static final int MAX_PUBLISHERS = 100;
    protected static final int INPUT_SIZE = 64*1024;
    protected static final int OUTPUT_SIZE = 64*1024;
    protected static final int SPIN_OK = 0;
    protected static final int SPIN_ERR = -1;
    protected static final int PROTOCOL_VER = 0xb9;
    protected static final int SYNC_TIME_SCOPE = 10; // milliseconds
    protected boolean spin_ = true;
    protected java.lang.String ip_addr_;
    protected java.lang.String node_name_;
    boolean initNode(java.lang.String node_name, java.lang.String ip_addr) { return false; }
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
        ti.node = node_name_;
        publish(p.getEndpointType(), ti);
    }

    void negotiateTopics(SubscriberT s) {
        TopicInfo ti = new TopicInfo();
        ti.topic_id = s.id_;
        ti.topic_name = s.topic_;
        ti.message_type = s.getMsgType();
        ti.md5sum = s.getMsgMD5();
        ti.buffer_size = INPUT_SIZE;
        ti.node = node_name_;
        publish(s.getEndpointType(), ti);
    }
    
    void syncTime(byte[] data) {
        com.roslib.tinyros_msgs.SyncTime t = new com.roslib.tinyros_msgs.SyncTime();
        t.deserialize(data, 0);
        long now =  (long)Time.now().toMSec();
        Time.lock.lock();
        long scope = now - Time.time_last - t.tick;
        if ((Time.time_start == 0) || (scope >=0 && scope <= SYNC_TIME_SCOPE)) {
            Time.time_dds = (long)t.data.toMSec();
            Time.time_start = now;
        }
        Time.time_last = now;
        Time.lock.unlock();  
    }
}

