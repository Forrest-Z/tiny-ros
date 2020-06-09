package com.roslib.ros;

import java.net.Socket;

public class HardwareTCP implements Hardware {
    private final int DEFAULT_PORTNUM = 11315;
    private java.net.Socket client_;
    private java.io.OutputStream os_;
    private java.io.InputStream is_;
    private boolean connected_;

    public HardwareTCP() {
        this.client_ = null;
        this.os_ = null;
        this.is_ = null;
        this.connected_ = false;
    }
    
    @Override
    public boolean init(java.lang.String portName) {
        try {
            this.close();
            this.client_ = new Socket(portName, DEFAULT_PORTNUM);
            this.os_ = this.client_.getOutputStream();
            this.is_ = this.client_.getInputStream();
            this.connected_ = true;
        }catch(java.lang.Exception e){
            e.printStackTrace();
            this.connected_ = false;
            this.close();
        }
        return this.connected_;
    }
    
    @Override
    public int read(byte[] data, int length) {
        int rv = -1;
        if (this.connected_) {
            try {
                rv = this.is_.read(data, 0, length);
            } catch (java.lang.Exception e) {
                e.printStackTrace();
                this.close();
                rv = -1;
            }
        }
        return rv;
    }

    @Override
    public boolean write(byte[] data, int length) {
        if (this.connected_) {
            try {
                this.os_.write(data, 0, length);
                return true;
            } catch (java.lang.Exception e) {
                e.printStackTrace();
                this.close();
                return false;
            }
        } else {
            return false;
        }
    }
    
    @Override
    public boolean connected()  {
        return this.connected_;
    }

    @Override
    public void close()  {
        try {
            this.connected_ = false;
            if (this.os_ != null) {
                this.os_.close();
                this.os_ = null;
            }
            if (this.is_ != null) {
                this.is_.close();
                this.is_ = null;
            }
            if (this.client_ != null) {
                this.client_.close();
                this.client_ = null;
            }
        } catch (java.lang.Exception e) {
            e.printStackTrace();
        }
    }
}
