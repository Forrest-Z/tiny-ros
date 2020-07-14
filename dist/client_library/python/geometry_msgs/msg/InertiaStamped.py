import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class InertiaStamped(tinyros.Message):
    __slots__ = ['header','inertia']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Inertia']

    def __init__(self):
        super(InertiaStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.inertia = geometry_msgs.msg.Inertia()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.inertia.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.inertia.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.inertia.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"inertia": {"'
        string_echo += self.inertia.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/InertiaStamped"

    def getMD5(self):
        return "2b3c9b263c59f65da44508cd041d18a0"

_struct_I = struct.Struct('<I')

