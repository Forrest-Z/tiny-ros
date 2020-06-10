import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class Float64(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['float64']

    def __init__(self):
        super(Float64, self).__init__()
        self.data = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<d")
            buff.write(struct_data.pack(self.data))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<d")
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
        return "std_msgs/Float64"

    def getMD5(self):
        return "3439fe880debfd63cf4e61e62e285345"

_struct_I = struct.Struct('<I')

