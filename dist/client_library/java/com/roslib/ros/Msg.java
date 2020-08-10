package com.roslib.ros;

/* Base Message Type */
public interface Msg {
    public int serialize(byte[] outbuffer, int start);
    public int deserialize(byte[] data, int start);
    public int serializedLength();
    public java.lang.String echo();
    public java.lang.String getType();
    public java.lang.String getMD5();
    public long getID();
    public void setID(long id);
}
