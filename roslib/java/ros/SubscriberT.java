package com.roslib.ros;

public abstract class SubscriberT {
    public void callback(byte[] data) { }
    public long getEndpointType() { return 0; }

    public java.lang.String getMsgType() { return ""; }
    public java.lang.String getMsgMD5() { return ""; }

    public java.lang.String topic_ = "";

    // id_ is set by NodeHandle when we advertise
    public long id_ = -1;

    // negotiated_ is set by NodeHandle when we negotiateTopics
    public boolean negotiated_ = false;
    public boolean srv_flag_ = false;
    public boolean negotiated() { return negotiated_; }
}
