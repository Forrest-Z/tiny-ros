import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class Vector3Stamped(tinyros.Message):
    __slots__ = ['header','vector']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Vector3']

    def __init__(self):
        super(Vector3Stamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.vector = geometry_msgs.msg.Vector3()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.vector.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.vector.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.vector.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"vector": {"'
        string_echo += self.vector.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Vector3Stamped"

    def getMD5(self):
        return "4b85025eb6f70f6b1e0cefbb75f69ac2"

_struct_I = struct.Struct('<I')

