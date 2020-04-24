package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public class Subscriber<MsgT extends Msg> extends SubscriberT {
    public int endpoint;
    private MsgT msg_;
    private CallbackSubT cb_;

    public Subscriber(java.lang.String topic_name, CallbackSubT cb, MsgT msg) {
        this.topic_ = topic_name;
        this.cb_ = cb;
        this.msg_ = msg;
        this.endpoint = TopicInfo.ID_SUBSCRIBER;
    }

    public void callback(byte[] data) {
    	try {
        	Class<? extends Msg> cl = this.msg_.getClass();
			Msg tmsg = cl.newInstance();
			tmsg.deserialize(data, 0);
	        this.cb_.callback(tmsg);
		} catch (Exception e) {
			e.printStackTrace();
		}
    }

    public java.lang.String getMsgType() {
        return this.msg_.getType();
    }

    public java.lang.String getMsgMD5() {
        return this.msg_.getMD5();
    }
    public int getEndpointType() {
        return this.endpoint;
    }
}
