import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class PointStamped(tinyros.Message):
    __slots__ = ['header','point']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Point']

    def __init__(self):
        super(PointStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.point = geometry_msgs.msg.Point()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.point.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.point.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.point.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"point": {"'
        string_echo += self.point.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/PointStamped"

    def getMD5(self):
        return "d34e83bdbef7bf4b617a6293aab8390e"

_struct_I = struct.Struct('<I')

