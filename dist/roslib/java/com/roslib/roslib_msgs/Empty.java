package com.roslib.roslib_msgs;

import java.lang.*;

public class Empty implements com.roslib.ros.Msg {

    public Empty() {
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        return length;
    }

    public java.lang.String getType(){ return "roslib_msgs/Empty"; }
    public java.lang.String getMD5(){ return "a4758a8dc954913cf9a8cc8864ad5209"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
