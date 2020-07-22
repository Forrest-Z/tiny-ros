package com.roslib.ros;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Time {
    public static Lock lock = new ReentrantLock();
    public static long time_start = 0;
    public static long time_dds = 0;
    public static long time_last = 0;
    public long sec;
    public long nsec;

    public Time() {
        this.sec = 0;
        this.nsec = 0;
    }

    public Time(long _sec, long _nsec) {
        this.sec = sec;
        this.nsec = nsec;
        normalize();
    }

    public double toSec() {
        return sec + 1e-9 * nsec;
    }

    public double toMSec() {
        return sec * 1000 + 1e-6 * nsec;
    }

    public void fromSec(double t) {
        sec = (int) Math.floor(t);
        nsec = (int) Math.round((t - sec) * 1e9);
    }

    public long toNsec() {
        return (sec) * 1000000000 + nsec;
    }

    public Time fromNSec(long t) {
        int secs = (int) (t / 1000000000);
        int nsecs = (int) (t % 1000000000);
        return new Time(secs, nsecs);
    }
    public Time add(Duration d) {
        return new Time(sec + d.sec, nsec + d.nsec);
    }

    public Time subtract(Duration d) {
        return new Time(sec - d.sec, nsec - d.nsec);
    }

    public void normalize() {
        while (nsec < 0) {
            nsec += 1000000000;
            sec -= 1;
        }
        while (nsec >= 1000000000) {
            nsec -= 1000000000;
            sec += 1;
        }
    }

    public static Time dds() {
        Time.lock.lock();
        Time t = Time.now();
        long offset = (long)t.toMSec();
        offset = offset > Time.time_start && Time.time_start > 0 ? offset - Time.time_start : 0;
        t.sec = offset / 1000;
        t.nsec = (offset % 1000) * 1000000;
        t.sec += Time.time_dds / 1000;
        t.nsec += (Time.time_dds % 1000) * 1000000;
        Time.normalizeSecNSec(t);
        Time.lock.unlock();
        return t;
    }

    public static Time now() {
        Time time = new Time();
        long millis = java.lang.System.currentTimeMillis();
        time.sec = millis / 1000;
        time.nsec = (millis % 1000) * 1000000;
        Time.normalizeSecNSec(time);
        return time;
    }

    public static com.roslib.ros.Time normalizeSecNSec(long sec, long nsec) {
        com.roslib.ros.Time time = new com.roslib.ros.Time();
        long nsec_part = nsec % 1000000000;
        long sec_part = nsec / 1000000000;
        time.sec = sec + sec_part;
        time.nsec = nsec_part;
        return time;
    }

    public static void normalizeSecNSec(com.roslib.ros.Time time) {
        long nsec_part = time.nsec % 1000000000;
        long sec_part = time.nsec / 1000000000;
        time.sec += sec_part;
        time.nsec = nsec_part;
    }

    public static  com.roslib.ros.Time normalizeSecNSecUnsigned(long sec, long nsec) {
        com.roslib.ros.Time time = new com.roslib.ros.Time();
        long nsec_part = nsec % 1000000000;
        long sec_part = sec + nsec / 1000000000;
        if (nsec_part < 0) {
            nsec_part += 1000000000;
            --sec_part;
        }

        time.sec = sec_part;
        time.nsec = nsec_part;
        return time;
    }

    public static void normalizeSecNSecUnsigned(com.roslib.ros.Time time) {
        long nsec_part = time.nsec % 1000000000;
        long sec_part = time.sec + time.nsec / 1000000000;
        if (nsec_part < 0) {
            nsec_part += 1000000000;
            --sec_part;
        }

        time.sec = sec_part;
        time.nsec = nsec_part;
    }
}

