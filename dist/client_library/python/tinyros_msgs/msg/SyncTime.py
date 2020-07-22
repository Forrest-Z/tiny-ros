import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tinyros_msgs.msg
import tinyros

class SyncTime(tinyros.Message):
    __slots__ = ['tick','data']
    _slot_types = ['uint32','tinyros.Time']

    def __init__(self):
        super(SyncTime, self).__init__()
        self.tick = 0
        self.data = tinyros.Time()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.tick))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
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
            (self.tick,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
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
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"tick": %s'%tick
        string_echo += ', '
        string_echo += '"data.sec": %s, '%data.sec
        string_echo += '"data.nsec": %s'%data.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_msgs/SyncTime"

    def getMD5(self):
        return "45bf702585c65b1bb762993bdbb1de6f"

_struct_I = struct.Struct('<I')

