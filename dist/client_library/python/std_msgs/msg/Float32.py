import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class Float32(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['float32']

    def __init__(self):
        super(Float32, self).__init__()
        self.data = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<f")
            buff.write(struct_data.pack(self.data))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<f")
            (self.data,) = struct_data.unpack(buff[offset:(offset + 4)])
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
        return "std_msgs/Float32"

    def getMD5(self):
        return "2aff5d2343e8e80ceea1362fc770035c"

_struct_I = struct.Struct('<I')

