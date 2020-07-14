package com.roslib.ros;
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;

public class HardwareUDP implements Hardware {
    private final int SERVER_PORTNUM = 11316;
    private final int CLIENT_PORTNUM = 11317;
    private DatagramSocket send_fd_;
    private DatagramSocket recv_fd_;
    private boolean connected_;
    private boolean binded_;
    private java.lang.String portName_ = "";

    public HardwareUDP() {
        this.send_fd_ = null;
        this.recv_fd_ = null;
        this.connected_ = false;
        this.binded_ = false;
        this.portName_ = "";
    }

    @Override
    public boolean init(String portName) {
        try {
            this.close();
            this.portName_ = portName;
            this.send_fd_ = new DatagramSocket();
            this.send_fd_.setReuseAddress(true);
            this.connected_ = true;
        } catch (SocketException e) {
            e.printStackTrace();
            this.close();
            this.connected_ = false;
        }
        return this.connected_;
    }

    @Override
    public int read(byte[] data, int length) {
        int rv = -1;
        if (this.connected_) {
            try {
                if (this.recv_fd_ == null){
                    this.recv_fd_ = new DatagramSocket(CLIENT_PORTNUM);
                    this.recv_fd_.setReuseAddress(true);
                }
                DatagramPacket packet = new DatagramPacket(data, length);
                this.recv_fd_.receive(packet);
                rv = packet.getLength();
            } catch (IOException e) {
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
                InetAddress address= InetAddress.getByName(this.portName_);
                DatagramPacket packet = new DatagramPacket(data, length, address, SERVER_PORTNUM);
                this.send_fd_.send(packet);
                return true;
            } catch (IOException e) {
                e.printStackTrace();
                this.close();
            }
        }
        return false;
    }

    @Override
    public boolean connected() {
        return this.connected_;
    }

    @Override
    public void close() {
        this.connected_ = false;
        if (this.send_fd_ != null) {
            this.send_fd_.close();
            this.send_fd_ = null;
        }
        if (this.recv_fd_ != null) {
            this.recv_fd_.close();
            this.recv_fd_ = null;
        }
    }
}
