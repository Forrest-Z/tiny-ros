import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg

class LookupTransformFeedback(tinyros.Message):
    def __init__(self):
        super(LookupTransformFeedback, self).__init__()

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
        return "tf2_msgs/LookupTransformFeedback"

    def getMD5(self):
        return "e6217f8a8e77aa218a8d6f594d08ba08"

_struct_I = struct.Struct('<I')

