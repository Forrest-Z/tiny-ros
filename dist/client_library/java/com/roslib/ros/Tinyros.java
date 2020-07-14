package com.roslib.ros;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Tinyros {
    private static java.lang.String ipaddr_ = "127.0.0.1";
    private static boolean main_loop_init_ = false;
    private static Lock main_loop_mutex_ = new ReentrantLock();

    private static boolean main_loop_udp_init_ = false;
    private static Lock main_loop_udp_mutex_ = new ReentrantLock();

    private static NodeHandle g_nh_ = new NodeHandle();

    private static NodeHandleUDP g_nh_udp_ = new NodeHandleUDP();

    private static class MainLoopThread extends Thread {
        private NodeHandleBase nh_ = null;
        public MainLoopThread(NodeHandleBase nh) {
            nh_ = nh;
        }

        @Override
        public void run() {
            while(nh_.spin_) {
                if (!nh_.ok()) {
                    nh_.initNode(Tinyros.ipaddr_);
                }
                if (nh_.ok()) {
                    nh_.spin();
                }
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    public static void init() {
        Tinyros.nh();
    }

    public static void init(java.lang.String ipaddr) {
        Tinyros.ipaddr_ = ipaddr;
        Tinyros.nh();
    }

    public static NodeHandle nh() {
        if (!Tinyros.main_loop_init_){
            Tinyros.main_loop_mutex_.lock();
            if (!Tinyros.main_loop_init_) {
                Tinyros.main_loop_init_ = true;
                new MainLoopThread(Tinyros.g_nh_).start();
                try {
                    Thread.sleep(200);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            Tinyros.main_loop_mutex_.unlock();
        }
        return Tinyros.g_nh_;
    }

    public static NodeHandleUDP udp() {
        if (!Tinyros.main_loop_udp_init_){
            Tinyros.main_loop_udp_mutex_.lock();
            if (!Tinyros.main_loop_udp_init_) {
                Tinyros.main_loop_udp_init_ = true;
                new MainLoopThread(Tinyros.g_nh_udp_).start();
                try {
                    Thread.sleep(200);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            Tinyros.main_loop_udp_mutex_.unlock();
        }
        return Tinyros.g_nh_udp_;
    }
}
