import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class UInt32(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['uint32']

    def __init__(self):
        super(UInt32, self).__init__()
        self.data = 0

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.data))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.data,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"data": %s'%data
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/UInt32"

    def getMD5(self):
        return "d4e8dc9c9e9d5076e594a3e342c2d4e3"

_struct_I = struct.Struct('<I')

