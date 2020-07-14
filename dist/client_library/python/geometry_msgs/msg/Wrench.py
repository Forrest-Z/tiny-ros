import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Wrench(tinyros.Message):
    __slots__ = ['force','torque']
    _slot_types = ['geometry_msgs.msg.Vector3','geometry_msgs.msg.Vector3']

    def __init__(self):
        super(Wrench, self).__init__()
        self.force = geometry_msgs.msg.Vector3()
        self.torque = geometry_msgs.msg.Vector3()

    def serialize(self, buff):
        offset = 0
        offset += self.force.serialize(buff)
        offset += self.torque.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.force.deserialize(buff[offset:])
        offset += self.torque.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.force.serializedLength()
        length += self.torque.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"force": {"'
        string_echo += self.force.echo()
        string_echo += '}, '
        string_echo += '"torque": {"'
        string_echo += self.torque.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Wrench"

    def getMD5(self):
        return "02d01d4a8dc253c7b42d4c9866201aee"

_struct_I = struct.Struct('<I')

