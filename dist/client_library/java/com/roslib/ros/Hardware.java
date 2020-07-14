package com.roslib.ros;

public interface Hardware {
    boolean init(java.lang.String portName);
    int read(byte[] data, int length);
    boolean write(byte[] data, int length);
    boolean connected();
    void close();
}
