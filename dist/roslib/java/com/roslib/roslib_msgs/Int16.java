package com.roslib.roslib_msgs;

import java.lang.*;

public class Int16 implements com.roslib.ros.Msg {
    public short data;

    public Int16() {
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
        this.data   = (short)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.data |= (short)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        offset += 2;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 2;
        return length;
    }

    public java.lang.String getType(){ return "roslib_msgs/Int16"; }
    public java.lang.String getMD5(){ return "78714f4d8eced3652112f4b2d18acff6"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
