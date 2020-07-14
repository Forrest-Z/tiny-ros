import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class Empty(tinyros.Message):
    def __init__(self):
        super(Empty, self).__init__()

    def serialize(self, buff):
        offset = 0
        return offset

    def deserialize(self, buff):
        offset = 0
        return offset

    def serializedLength(self):
        length = 0
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/Empty"

    def getMD5(self):
        return "140bdcb7bbc50b4a936e90ff2350c8d3"

_struct_I = struct.Struct('<I')

