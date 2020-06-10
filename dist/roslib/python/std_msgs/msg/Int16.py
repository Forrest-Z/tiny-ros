import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class Int16(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['int16']

    def __init__(self):
        super(Int16, self).__init__()
        self.data = 0

    def serialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<h")
            buff.write(struct_data.pack(self.data))
            offset += 2
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<h")
            (self.data,) = struct_data.unpack(buff[offset:(offset + 2)])
            offset += 2
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 2
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"data": %s'%data
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/Int16"

    def getMD5(self):
        return "156735359f7b6d347f113a1090134af3"

_struct_I = struct.Struct('<I')

