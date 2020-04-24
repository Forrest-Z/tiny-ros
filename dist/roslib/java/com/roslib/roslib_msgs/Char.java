package com.roslib.roslib_msgs;

import java.lang.*;

public class Char implements com.roslib.ros.Msg {
    public char data;

    public Char() {
        this.data = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.data >> (8 * 0)) & 0xFF);
        offset += 1;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.data   = (char)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        offset += 1;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 1;
        return length;
    }

    public java.lang.String getType(){ return "roslib_msgs/Char"; }
    public java.lang.String getMD5(){ return "42e0bd733e1e77ec41439b18e22b9008"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
