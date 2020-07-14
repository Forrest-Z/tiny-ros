package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public class ServiceServer<MReq extends Msg, MRes extends Msg> extends SubscriberT {
    public MReq req;
    public MRes resp;
    public Publisher pub;

    private CallbackSrvT cb_;

    public ServiceServer(java.lang.String topic_name, CallbackSrvT cb, MReq req, MRes resp) {
        this.topic_ = topic_name;
        this.cb_ = cb;
        this.req = req;
        this.resp = resp;
        this.srv_flag_ = true;
        this.pub = new Publisher(this.topic_, resp);
        this.pub.endpoint_ = TopicInfo.ID_SERVICE_SERVER + TopicInfo.ID_PUBLISHER;
    }

    // these refer to the subscriber
    @Override
    public void callback(byte[] data) {
    	try {
        	Class<? extends Msg> clreq = this.req.getClass();
        	Class<? extends Msg> clresp = this.resp.getClass();
        	Msg treq = clreq.newInstance();
        	Msg tresp = clresp.newInstance();
            treq.deserialize(data, 0);
            this.cb_.callback(treq, tresp);
            tresp.setID(treq.getID());
            this.pub.publish(tresp);
		} catch (Exception e) {
			e.printStackTrace();
		}
    }

    @Override
    public java.lang.String getMsgType() {
        return this.req.getType();
    }

    @Override
    public java.lang.String getMsgMD5() {
        return this.req.getMD5();
    }

    @Override
    public long getEndpointType() {
        return TopicInfo.ID_SERVICE_SERVER + TopicInfo.ID_SUBSCRIBER;
    }

    @Override
    public boolean negotiated() {
        return (this.negotiated_ && this.pub.negotiated_); 
    }
}
