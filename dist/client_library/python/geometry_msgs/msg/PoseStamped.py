import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class PoseStamped(tinyros.Message):
    __slots__ = ['header','pose']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Pose']

    def __init__(self):
        super(PoseStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.pose = geometry_msgs.msg.Pose()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.pose.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.pose.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.pose.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"pose": {"'
        string_echo += self.pose.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/PoseStamped"

    def getMD5(self):
        return "c7084e6b27c3d6e62efd9bf6d2f6540f"

_struct_I = struct.Struct('<I')

