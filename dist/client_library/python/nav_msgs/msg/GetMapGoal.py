import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg

class GetMapGoal(tinyros.Message):
    def __init__(self):
        super(GetMapGoal, self).__init__()

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
        return "nav_msgs/GetMapGoal"

    def getMD5(self):
        return "b39e6b705afaad0184bd2c87f4bd870f"

_struct_I = struct.Struct('<I')

