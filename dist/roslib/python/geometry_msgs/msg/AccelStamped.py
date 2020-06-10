import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class AccelStamped(tinyros.Message):
    __slots__ = ['header','accel']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Accel']

    def __init__(self):
        super(AccelStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.accel = geometry_msgs.msg.Accel()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.accel.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.accel.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.accel.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"accel": {"'
        string_echo += self.accel.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/AccelStamped"

    def getMD5(self):
        return "fa35432963826361a1073b1df905a559"

_struct_I = struct.Struct('<I')

