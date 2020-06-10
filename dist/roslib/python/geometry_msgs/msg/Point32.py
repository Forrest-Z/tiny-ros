import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Point32(tinyros.Message):
    __slots__ = ['x','y','z']
    _slot_types = ['float32','float32','float32']

    def __init__(self):
        super(Point32, self).__init__()
        self.x = 0.
        self.y = 0.
        self.z = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_x = struct.Struct("<f")
            buff.write(struct_x.pack(self.x))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<f")
            buff.write(struct_y.pack(self.y))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_z = struct.Struct("<f")
            buff.write(struct_z.pack(self.z))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_x = struct.Struct("<f")
            (self.x,) = struct_x.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<f")
            (self.y,) = struct_y.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_z = struct.Struct("<f")
            (self.z,) = struct_z.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        length += 4
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
        return "geometry_msgs/Point32"

    def getMD5(self):
        return "b17f2230f465fce816e3773d7d59a841"

_struct_I = struct.Struct('<I')

