import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Pose2D(tinyros.Message):
    __slots__ = ['x','y','theta']
    _slot_types = ['float64','float64','float64']

    def __init__(self):
        super(Pose2D, self).__init__()
        self.x = 0.
        self.y = 0.
        self.theta = 0.

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
            struct_theta = struct.Struct("<d")
            buff.write(struct_theta.pack(self.theta))
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
            struct_theta = struct.Struct("<d")
            (self.theta,) = struct_theta.unpack(buff[offset:(offset + 8)])
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
        string_echo += '"theta": %s'%theta
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Pose2D"

    def getMD5(self):
        return "509f362ff66c4d3df21020fa7c01f8c6"

_struct_I = struct.Struct('<I')

