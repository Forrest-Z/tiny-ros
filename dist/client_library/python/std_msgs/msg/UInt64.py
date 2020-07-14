import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class UInt64(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['uint64']

    def __init__(self):
        super(UInt64, self).__init__()
        self.data = 0

    def serialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<Q")
            buff.write(struct_data.pack(self.data))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<Q")
            (self.data,) = struct_data.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"data": %s'%data
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/UInt64"

    def getMD5(self):
        return "e815672a5006369d73f91f25ee5609ac"

_struct_I = struct.Struct('<I')

