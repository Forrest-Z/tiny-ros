package com.roslib.nav_msgs;

import java.lang.*;

public class GetMap {

public static final java.lang.String GETMAP = "nav_msgs/GetMap";

public static class GetMapRequest implements com.roslib.ros.Msg {
    private long __id__;

    public GetMapRequest() {
        this.__id__ = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.__id__ >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.__id__ >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.__id__ >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.__id__ >> (8 * 3)) & 0xFF);
        offset += 4;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.__id__  = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.__id__ |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.__id__ |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.__id__ |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        return length;
    }

    public java.lang.String echo(){ return ""; }
    public java.lang.String getType() { return GETMAP; }
    public java.lang.String getMD5(){ return "8ea512c109a0b3a9eca8de407dd02d2a"; }
    public long getID() { return this.__id__; }
    public void setID(long id) { this.__id__ = id; }
}

public static class GetMapResponse implements com.roslib.ros.Msg {
    private long __id__;
    public com.roslib.nav_msgs.OccupancyGrid map;

    public GetMapResponse() {
        this.map = new com.roslib.nav_msgs.OccupancyGrid();
        this.__id__ = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.__id__ >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.__id__ >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.__id__ >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.__id__ >> (8 * 3)) & 0xFF);
        offset += 4;
        offset = this.map.serialize(outbuffer, offset);
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.__id__  = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.__id__ |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.__id__ |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.__id__ |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        offset = this.map.deserialize(inbuffer, offset);
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += this.map.serializedLength();
        return length;
    }

    public java.lang.String echo(){ return ""; }
    public java.lang.String getType() { return GETMAP; }
    public java.lang.String getMD5(){ return "5bf895a6cdaff312c69ca1cef20fed64"; }
    public long getID() { return this.__id__; }
    public void setID(long id) { this.__id__ = id; }
}

}
