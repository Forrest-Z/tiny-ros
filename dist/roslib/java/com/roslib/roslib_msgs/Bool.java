package com.roslib.roslib_msgs;

import java.lang.*;

public class Bool implements com.roslib.ros.Msg {
    public boolean data;

    public Bool() {
        this.data = false;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset] = (byte)((data ? 0x01 : 0x00) & 0xFF);
        offset += 1;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.data = (boolean)((inbuffer[offset] & 0xFF) != 0 ? true : false);
        offset += 1;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 1;
        return length;
    }

    public java.lang.String getType(){ return "roslib_msgs/Bool"; }
    public java.lang.String getMD5(){ return "80a272e03c01d79e04a9cba6da3167d5"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
