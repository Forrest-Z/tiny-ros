import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Transform(tinyros.Message):
    __slots__ = ['translation','rotation']
    _slot_types = ['geometry_msgs.msg.Vector3','geometry_msgs.msg.Quaternion']

    def __init__(self):
        super(Transform, self).__init__()
        self.translation = geometry_msgs.msg.Vector3()
        self.rotation = geometry_msgs.msg.Quaternion()

    def serialize(self, buff):
        offset = 0
        offset += self.translation.serialize(buff)
        offset += self.rotation.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.translation.deserialize(buff[offset:])
        offset += self.rotation.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.translation.serializedLength()
        length += self.rotation.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"translation": {"'
        string_echo += self.translation.echo()
        string_echo += '}, '
        string_echo += '"rotation": {"'
        string_echo += self.rotation.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Transform"

    def getMD5(self):
        return "2526ee1b1cc2e723e386c3c1b048ba72"

_struct_I = struct.Struct('<I')

