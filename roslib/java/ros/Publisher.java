package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public class Publisher<MsgT extends Msg> {
    public java.lang.String topic_;
    public MsgT msg_;
    public long id_;
    public NodeHandleBase nh_;
    public boolean negotiated_;

    public long endpoint_;

    public Publisher(java.lang.String topic_name, MsgT msg) {
        this.endpoint_ = TopicInfo.ID_PUBLISHER;
        this.topic_ = topic_name;
        this.msg_ = msg;
		this.negotiated_ = false;
    }

    public int publish(MsgT msg) {
        if (this.nh_ != null) {
            return this.nh_.publish(this.id_, msg);
        }
        return 0;
    }

    public long getEndpointType() {
        return this.endpoint_;
    }

    public boolean negotiated() {
        return this.negotiated_;
    }
}
