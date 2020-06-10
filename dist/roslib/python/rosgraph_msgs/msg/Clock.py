import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import rosgraph_msgs.msg
import tinyros

class Clock(tinyros.Message):
    __slots__ = ['clock']
    _slot_types = ['tinyros.Time']

    def __init__(self):
        super(Clock, self).__init__()
        self.clock = tinyros.Time()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.clock.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.clock.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.clock.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.clock.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"clock.sec": %s, '%clock.sec
        string_echo += '"clock.nsec": %s'%clock.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "rosgraph_msgs/Clock"

    def getMD5(self):
        return "d3bedbe03b904b8181e3fef4bbe0a73e"

_struct_I = struct.Struct('<I')

