package com.roslib.sensor_msgs;

import java.lang.*;

public class SetCameraInfo {

public static final java.lang.String SETCAMERAINFO = "sensor_msgs/SetCameraInfo";

public static class SetCameraInfoRequest implements com.roslib.ros.Msg {
    private long __id__;
    public com.roslib.sensor_msgs.CameraInfo camera_info;

    public SetCameraInfoRequest() {
        this.camera_info = new com.roslib.sensor_msgs.CameraInfo();
        this.__id__ = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.__id__ >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.__id__ >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.__id__ >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.__id__ >> (8 * 3)) & 0xFF);
        offset += 4;
        offset = this.camera_info.serialize(outbuffer, offset);
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.__id__  = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.__id__ |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.__id__ |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.__id__ |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        offset = this.camera_info.deserialize(inbuffer, offset);
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += this.camera_info.serializedLength();
        return length;
    }

    public java.lang.String echo() { return ""; }
    public java.lang.String getType() { return SETCAMERAINFO; }
    public java.lang.String getMD5(){ return "7c09688f428450a1ac49eacdbb57c012"; }
    public long getID() { return this.__id__; }
    public void setID(long id) { this.__id__ = id; }
}

public static class SetCameraInfoResponse implements com.roslib.ros.Msg {
    private long __id__;
    public boolean success;
    public java.lang.String status_message;

    public SetCameraInfoResponse() {
        this.success = false;
        this.status_message = "";
        this.__id__ = 0;
    }

    public int serialize(byte[] outbuffer, int start) {
        int offset = start;
        outbuffer[offset + 0] = (byte)((this.__id__ >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((this.__id__ >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((this.__id__ >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((this.__id__ >> (8 * 3)) & 0xFF);
        offset += 4;
        outbuffer[offset] = (byte)((success ? 0x01 : 0x00) & 0xFF);
        offset += 1;
        int length_status_message = this.status_message.getBytes().length;
        outbuffer[offset + 0] = (byte)((length_status_message >> (8 * 0)) & 0xFF);
        outbuffer[offset + 1] = (byte)((length_status_message >> (8 * 1)) & 0xFF);
        outbuffer[offset + 2] = (byte)((length_status_message >> (8 * 2)) & 0xFF);
        outbuffer[offset + 3] = (byte)((length_status_message >> (8 * 3)) & 0xFF);
        offset += 4;
        for (int k=0; k<length_status_message; k++) {
            outbuffer[offset + k] = (byte)((this.status_message.getBytes())[k] & 0xFF);
        }
        offset += length_status_message;
        return offset;
    }

    public int deserialize(byte[] inbuffer, int start) {
        int offset = start;
        this.__id__  = (long)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        this.__id__ |= (long)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        this.__id__ |= (long)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        this.__id__ |= (long)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        this.success = (boolean)((inbuffer[offset] & 0xFF) != 0 ? true : false);
        offset += 1;
        int length_status_message = (int)((inbuffer[offset + 0] & 0xFF) << (8 * 0));
        length_status_message |= (int)((inbuffer[offset + 1] & 0xFF) << (8 * 1));
        length_status_message |= (int)((inbuffer[offset + 2] & 0xFF) << (8 * 2));
        length_status_message |= (int)((inbuffer[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
        byte[] bytes_status_message = new byte[length_status_message];
        for(int k= 0; k< length_status_message; k++){
            bytes_status_message[k] = (byte)(inbuffer[k + offset] & 0xFF);
        }
        this.status_message = new java.lang.String(bytes_status_message);
        offset += length_status_message;
        return offset;
    }

    public int serializedLength() {
        int length = 0;
        length += 1;
        int length_status_message = this.status_message.getBytes().length;
        length += 4;
        length += length_status_message;
        return length;
    }

    public java.lang.String echo() { return ""; }
    public java.lang.String getType() { return SETCAMERAINFO; }
    public java.lang.String getMD5(){ return "e03fdc9555d1e3c7347a728e913775f6"; }
    public long getID() { return this.__id__; }
    public void setID(long id) { this.__id__ = id; }
}

}
