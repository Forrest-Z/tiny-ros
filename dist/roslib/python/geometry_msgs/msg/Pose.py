import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Pose(tinyros.Message):
    __slots__ = ['position','orientation']
    _slot_types = ['geometry_msgs.msg.Point','geometry_msgs.msg.Quaternion']

    def __init__(self):
        super(Pose, self).__init__()
        self.position = geometry_msgs.msg.Point()
        self.orientation = geometry_msgs.msg.Quaternion()

    def serialize(self, buff):
        offset = 0
        offset += self.position.serialize(buff)
        offset += self.orientation.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.position.deserialize(buff[offset:])
        offset += self.orientation.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.position.serializedLength()
        length += self.orientation.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"position": {"'
        string_echo += self.position.echo()
        string_echo += '}, '
        string_echo += '"orientation": {"'
        string_echo += self.orientation.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Pose"

    def getMD5(self):
        return "0b42fb88be8cac0efa6e446e13befcae"

_struct_I = struct.Struct('<I')

