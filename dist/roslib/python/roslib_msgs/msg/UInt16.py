import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import roslib_msgs.msg

class UInt16(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['uint16']

    def __init__(self):
        super(UInt16, self).__init__()
        self.data = 0

    def serialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<H")
            buff.write(struct_data.pack(self.data))
            offset += 2
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_data = struct.Struct("<H")
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
        return "roslib_msgs/UInt16"

    def getMD5(self):
        return "f7fc05ba72b1e0b94d9852ae4a5c54cb"

_struct_I = struct.Struct('<I')

