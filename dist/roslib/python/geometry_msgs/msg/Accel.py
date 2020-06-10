import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Accel(tinyros.Message):
    __slots__ = ['linear','angular']
    _slot_types = ['geometry_msgs.msg.Vector3','geometry_msgs.msg.Vector3']

    def __init__(self):
        super(Accel, self).__init__()
        self.linear = geometry_msgs.msg.Vector3()
        self.angular = geometry_msgs.msg.Vector3()

    def serialize(self, buff):
        offset = 0
        offset += self.linear.serialize(buff)
        offset += self.angular.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.linear.deserialize(buff[offset:])
        offset += self.angular.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.linear.serializedLength()
        length += self.angular.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"linear": {"'
        string_echo += self.linear.echo()
        string_echo += '}, '
        string_echo += '"angular": {"'
        string_echo += self.angular.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Accel"

    def getMD5(self):
        return "580cbad5f3bd2e9f0ca71e14b7ab1b0f"

_struct_I = struct.Struct('<I')

