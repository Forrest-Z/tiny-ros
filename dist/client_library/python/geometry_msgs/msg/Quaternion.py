import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Quaternion(tinyros.Message):
    __slots__ = ['x','y','z','w']
    _slot_types = ['float64','float64','float64','float64']

    def __init__(self):
        super(Quaternion, self).__init__()
        self.x = 0.
        self.y = 0.
        self.z = 0.
        self.w = 0.

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
        try:
            struct_w = struct.Struct("<d")
            buff.write(struct_w.pack(self.w))
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
        try:
            struct_w = struct.Struct("<d")
            (self.w,) = struct_w.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 8
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
        string_echo += ', '
        string_echo += '"w": %s'%w
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Quaternion"

    def getMD5(self):
        return "175c1571887d10ebed42ba6c042ddd88"

_struct_I = struct.Struct('<I')

