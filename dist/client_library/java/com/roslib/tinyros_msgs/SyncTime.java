package com.roslib.tinyros_msgs;

import java.lang.*;

public class SyncTime implements com.roslib.ros.Msg {
    public long tick;
    public com.roslib.ros.Time data;

    public SyncTime() {
        this.tick = 0;
        this.data = new com.roslib.ros.Time();
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.tick >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.tick >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.tick >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.tick >> (8 * 3)) & 0xFF);
        offset += 4;
        outbuffer[offset + 0] = (byte)((this.data.sec >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.data.sec >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.data.sec >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.data.sec >> (8 * 3)) & 0xFF);
        offset += 4;
        outbuffer[offset + 0] = (byte)((this.data.nsec >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.data.nsec >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.data.nsec >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.data.nsec >> (8 * 3)) & 0xFF);
        offset += 4;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.tick   = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.tick |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.tick |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.tick |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        this.data.sec   = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.data.sec |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.data.sec |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.data.sec |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        this.data.nsec   = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.data.nsec |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.data.nsec |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.data.nsec |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 4;
        length += 4;
        length += 4;
        return length;
    }

    public java.lang.String echo() { return ""; }
    public java.lang.String getType(){ return "tinyros_msgs/SyncTime"; }
    public java.lang.String getMD5(){ return "45bf702585c65b1bb762993bdbb1de6f"; }
    public long getID() { return 0; }
    public void setID(long id) { }
}
