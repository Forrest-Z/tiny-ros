package com.roslib.roslib_msgs;

import java.lang.*;

public class UInt16 implements com.roslib.ros.Msg {
    public int data;

    public UInt16() {
        this.data = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.data >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.data >> (8 * 1)) & 0xFF);
        offset += 2;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.data   = (int)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.data |= (int)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        offset += 2;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 2;
        return length;
    }

    public java.lang.String getType(){ return "roslib_msgs/UInt16"; }
    public java.lang.String getMD5(){ return "f7fc05ba72b1e0b94d9852ae4a5c54cb"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
