import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class TwistWithCovarianceStamped(tinyros.Message):
    __slots__ = ['header','twist']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.TwistWithCovariance']

    def __init__(self):
        super(TwistWithCovarianceStamped, self).__init__()
        self.header = std_msgs.msg.Header()
        self.twist = geometry_msgs.msg.TwistWithCovariance()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.twist.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.twist.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.twist.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"twist": {"'
        string_echo += self.twist.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/TwistWithCovarianceStamped"

    def getMD5(self):
        return "2cbcab62cac39de1d1d01785b99ba778"

_struct_I = struct.Struct('<I')

