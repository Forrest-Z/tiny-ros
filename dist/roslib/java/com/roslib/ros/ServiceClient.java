package com.roslib.ros;

import com.roslib.tinyros_msgs.TopicInfo;

public class ServiceClient<MReq extends Msg, MRes extends Msg> extends SubscriberT {
    public MReq req = null;
    public MRes resp = null;
    public MReq call_req = null;
    public MRes call_resp = null;
    public Publisher pub;
    public static int gg_id = 1;
    public static Object gg_mutex = new Object();
    public Object mutex = new Object();
    public Object g_mutex = new Object();
  

    public ServiceClient(java.lang.String topic_name, MReq req, MRes resp) {
        this.resp = resp;
        this.req = req;
        this.topic_ = topic_name;
        this.pub = new Publisher(this.topic_, req);
        this.pub.endpoint_ = TopicInfo.ID_SERVICE_CLIENT + TopicInfo.ID_PUBLISHER;
        this.srv_flag_ = true;
    }

    public boolean call(MReq request, MRes response, int durationSec) {
        synchronized(g_mutex) {
            synchronized(mutex) {
                if (!this.pub.nh_.ok()) {
                    this.call_req = null;
                    this.call_resp = null;
                    return false;
                }
                
                this.call_req = request;
                this.call_resp = response;
                
                synchronized(gg_mutex) {
                    this.call_req.setID(gg_id++);
                }
                
                if (this.pub.publish(request) <= 0) {
                    this.call_req = null;
                    this.call_resp = null;
                    return false;
                }

                try {
					mutex.wait(durationSec * 1000);
				} catch (Exception e) {
					e.printStackTrace();
				}
                
                this.call_req = null;
                this.call_resp = null;
                
                return true;
            }
        }
    }

    // these refer to the subscriber
    public void callback(byte[] data) {
        if (this.call_resp != null && this.call_req != null) {
            synchronized(mutex) {
                if (this.call_resp != null && this.call_req != null) {
                    long req_id  = this.call_req.getID();
                    long resp_id =  (int) ((data[0] & 0xFF) << (8 * 0));
                    resp_id |= (int) ((data[1] & 0xFF) << (8 * 1));
                    resp_id |= (int) ((data[2] & 0xFF) << (8 * 2));
                    resp_id |= (int) ((data[3] & 0xFF) << (8 * 3));
                    if (req_id == resp_id) {
                        this.call_resp.deserialize(data, 0);
                        mutex.notifyAll();
                    }
                }
            }
        }
    }

    public java.lang.String getMsgType() {
        return this.resp.getType();
    }

    public java.lang.String getMsgMD5() {
        return this.resp.getMD5();
    }
    public int getEndpointType() {
        return TopicInfo.ID_SERVICE_CLIENT + TopicInfo.ID_SUBSCRIBER;
    }

    public boolean negotiated() {
        return (this.negotiated_ && this.pub.negotiated_); 
    }
}
