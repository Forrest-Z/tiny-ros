import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg
import tinyros

class Time(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['tinyros.Time']

    def __init__(self):
        super(Time, self).__init__()
        self.data = tinyros.Time()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.data.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.data.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.data.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.data.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
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
        string_echo += '"data.sec": %s, '%data.sec
        string_echo += '"data.nsec": %s'%data.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/Time"

    def getMD5(self):
        return "64602ed67393e1e61260ab68d6fa2045"

_struct_I = struct.Struct('<I')

