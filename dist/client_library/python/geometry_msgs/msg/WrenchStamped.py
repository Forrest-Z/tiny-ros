import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class WrenchStamped(tinyros.Message):
    __slots__ = ['header','wrench']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Wrench']

    def __init__(self):
        super(WrenchStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.wrench = geometry_msgs.msg.Wrench()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.wrench.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.wrench.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.wrench.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"wrench": {"'
        string_echo += self.wrench.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/WrenchStamped"

    def getMD5(self):
        return "cf53874aa63609de4155ec8e9cf2c540"

_struct_I = struct.Struct('<I')

