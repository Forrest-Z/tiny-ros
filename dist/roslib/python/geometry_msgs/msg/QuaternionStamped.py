import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class QuaternionStamped(tinyros.Message):
    __slots__ = ['header','quaternion']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Quaternion']

    def __init__(self):
        super(QuaternionStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.quaternion = geometry_msgs.msg.Quaternion()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.quaternion.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.quaternion.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.quaternion.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"quaternion": {"'
        string_echo += self.quaternion.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/QuaternionStamped"

    def getMD5(self):
        return "69e39922feb9ec6eaf93755f93fce2cf"

_struct_I = struct.Struct('<I')

