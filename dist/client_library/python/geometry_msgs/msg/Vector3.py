import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Vector3(tinyros.Message):
    __slots__ = ['x','y','z']
    _slot_types = ['float64','float64','float64']

    def __init__(self):
        super(Vector3, self).__init__()
        self.x = 0.
        self.y = 0.
        self.z = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_x = struct.Struct("<d")
            buff.write(struct_x.pack(self.x))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            buff.write(struct_y.pack(self.y))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_z = struct.Struct("<d")
            buff.write(struct_z.pack(self.z))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_x = struct.Struct("<d")
            (self.x,) = struct_x.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            (self.y,) = struct_y.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_z = struct.Struct("<d")
            (self.z,) = struct_z.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 8
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"x": %s'%x
        string_echo += ', '
        string_echo += '"y": %s'%y
        string_echo += ', '
        string_echo += '"z": %s'%z
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Vector3"

    def getMD5(self):
        return "b5c8c5b484ec7d5e36a4d9de9124c561"

_struct_I = struct.Struct('<I')

