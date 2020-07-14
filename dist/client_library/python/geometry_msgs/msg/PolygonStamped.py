import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class PolygonStamped(tinyros.Message):
    __slots__ = ['header','polygon']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Polygon']

    def __init__(self):
        super(PolygonStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.polygon = geometry_msgs.msg.Polygon()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.polygon.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.polygon.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.polygon.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"polygon": {"'
        string_echo += self.polygon.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/PolygonStamped"

    def getMD5(self):
        return "33bdf94066425e572879b25c9a51ed50"

_struct_I = struct.Struct('<I')

