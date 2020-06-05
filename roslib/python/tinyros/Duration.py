import sys
import time
import tinyros.Time

if sys.version > '3': 
    long = int
        
class Duration(object):
    __slots__ = ['sec', 'nsec']
    def __init__(self, sec=0, nsec=0):
        self.sec = sec
        self.nsec = nsec
        if self.sec < 0:
            raise TypeError("Time sec values must be positive")
        if self.nsec < 0:
            raise TypeError("Time nsec values must be positive")

    def __floordiv__(self, val):
        t = type(val)
        if t in (int, long):
            return Duration(self.sec // val, self.nsec // val)
        elif t == float:
            return Duration.fromSec(self.toSec() // val)
        elif isinstance(val, Duration):
            return self.toSec() // val.toSec()
        else:
            return NotImplemented

    def __mul__(self, val):
        t = type(val)
        if t in (int, long):
            return Duration(self.sec * val, self.nsec * val)
        elif t == float:
            return Duration.fromSec(self.toSec() * val)
        else:
            return NotImplemented

    def __add__(self, other):
        if isinstance(other, Time):
            return other.__add__(self)
        elif isinstance(other, Duration):
            return Duration(self.sec+other.sec, self.nsec+other.nsec)
        else:
            return NotImplemented
    def __sub__(self, other):
        if not isinstance(other, Duration):
            return NotImplemented
        return Duration(self.sec-other.sec, self.nsec-other.nsec)        

    def __cmp__(self, other):
        if not isinstance(other, Duration):
            raise TypeError("Cannot compare to non-Duration")
        nanos = self.toNSec() - other.toNSec()
        if nanos > 0:
            return 1
        if nanos == 0:
            return 0
        return -1

    def __eq__(self, other):
        if not isinstance(other, Duration):
            return False
        return self.sec == other.sec and self.nsec == other.nsec

    def __hash__(self):
        return super(Duration, self).__hash__()

    def toSec(self):
        return float(self.sec) + float(self.nsec) / 1e9

    def toMSec(self):
        return self.sec * long(1e3) + float(self.nsec) / 1e6

    def toNSec(self):
        return self.sec * long(1e9) + self.nsec

    def fromSec(float_seconds):
        sec = int(float_seconds)
        nsec = int((float_seconds - sec) * 1000000000)
        return Duration(sec, nsec)
    fromSec = staticmethod(fromSec)


