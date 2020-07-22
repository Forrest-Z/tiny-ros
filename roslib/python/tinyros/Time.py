import sys
import time
import threading
import tinyros.Duration

if sys.version > '3': 
    long = int

class Time(object):
    __slots__ = ['sec', 'nsec']
    time_start = 0
    time_dds = 0
    time_last = 0
    mutex = threading.Lock()
    def __init__(self, sec=0, nsec=0):
        self.sec = sec
        self.nsec = nsec
        if self.sec < 0:
            raise TypeError("Time sec values must be positive")
        if self.nsec < 0:
            raise TypeError("Time nsec values must be positive")

    def __add__(self, other):
        if not isinstance(other, Duration):
            return NotImplemented
        return Time(self.sec + other.sec, self.nsec + other.nsec)

    def __sub__(self, other):
        if isinstance(other, Time):
            return Duration(self.sec - other.sec, self.nsec - other.nsec)
        elif isinstance(other, Duration):
            return Time(self.sec - other.sec, self.nsec - other.nsec)
        else:
            return NotImplemented

    def __cmp__(self, other):
        if not isinstance(other, Time):
            raise TypeError("cannot compare to non-Time")
        nanos = self.toNsec() - other.toNsec()
        if nanos > 0:
            return 1
        if nanos == 0:
            return 0
        return -1

    def __eq__(self, other):
        if not isinstance(other, Time):
            return False
        return self.sec == other.sec and self.nsec == other.nsec

    def toSec(self):
        return float(self.sec) + float(self.nsec) / 1e9

    def toMSec(self):
        return self.sec * long(1e3) + float(self.nsec) / 1e6

    def toNSec(self):
        return self.sec * long(1e9) + self.nsec
    
    def fromSec(float_sec):
        sec = int(float_sec)
        nsec = int((float_sec - sec) * 1000000000)
        return Time(sec, nsec)
    fromSec = staticmethod(fromSec)

    def dds():
        Time.mutex.acquire()
        t = Time.now()
        offset = long(t.toMSec());
        offset = (offset - Time.time_start) if ((offset > Time.time_start) and (Time.time_start > 0)) else 0
        t.sec = int(offset / 1000)
        t.nsec = int((offset % 1000) * 1000000)
        t.sec += int(Time.time_dds / 1000)
        t.nsec += int((Time.time_dds % 1000) * 1000000)
        t.sec += int(t.nsec / 1e9)
        t.nsec = int(t.nsec % 1e9)
        Time.mutex.release()
        return t
    dds = staticmethod(dds)

    def now():
        ticks = time.time()
        sec = int(ticks)
        nsec = int((ticks - sec) * 1000000000)
        return Time(sec, nsec)
    now = staticmethod(now)
    
    def setNow(time_now):
        if isinstance(time_now, Time):
            return Time(time_now.sec, time_now.nsec)
        else:
            return NotImplemented
    setNow = staticmethod(setNow)

