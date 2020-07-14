import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Inertia(tinyros.Message):
    __slots__ = ['m','com','ixx','ixy','ixz','iyy','iyz','izz']
    _slot_types = ['float64','geometry_msgs.msg.Vector3','float64','float64','float64','float64','float64','float64']

    def __init__(self):
        super(Inertia, self).__init__()
        self.m = 0.
        self.com = geometry_msgs.msg.Vector3()
        self.ixx = 0.
        self.ixy = 0.
        self.ixz = 0.
        self.iyy = 0.
        self.iyz = 0.
        self.izz = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_m = struct.Struct("<d")
            buff.write(struct_m.pack(self.m))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.com.serialize(buff)
        try:
            struct_ixx = struct.Struct("<d")
            buff.write(struct_ixx.pack(self.ixx))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_ixy = struct.Struct("<d")
            buff.write(struct_ixy.pack(self.ixy))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_ixz = struct.Struct("<d")
            buff.write(struct_ixz.pack(self.ixz))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_iyy = struct.Struct("<d")
            buff.write(struct_iyy.pack(self.iyy))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_iyz = struct.Struct("<d")
            buff.write(struct_iyz.pack(self.iyz))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_izz = struct.Struct("<d")
            buff.write(struct_izz.pack(self.izz))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_m = struct.Struct("<d")
            (self.m,) = struct_m.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.com.deserialize(buff[offset:])
        try:
            struct_ixx = struct.Struct("<d")
            (self.ixx,) = struct_ixx.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_ixy = struct.Struct("<d")
            (self.ixy,) = struct_ixy.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_ixz = struct.Struct("<d")
            (self.ixz,) = struct_ixz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_iyy = struct.Struct("<d")
            (self.iyy,) = struct_iyy.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_iyz = struct.Struct("<d")
            (self.iyz,) = struct_iyz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_izz = struct.Struct("<d")
            (self.izz,) = struct_izz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 8
        length += self.com.serializedLength()
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"m": %s'%m
        string_echo += ', '
        string_echo += '"com": {"'
        string_echo += self.com.echo()
        string_echo += '}, '
        string_echo += '"ixx": %s'%ixx
        string_echo += ', '
        string_echo += '"ixy": %s'%ixy
        string_echo += ', '
        string_echo += '"ixz": %s'%ixz
        string_echo += ', '
        string_echo += '"iyy": %s'%iyy
        string_echo += ', '
        string_echo += '"iyz": %s'%iyz
        string_echo += ', '
        string_echo += '"izz": %s'%izz
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Inertia"

    def getMD5(self):
        return "9116c935782bc29999dad1927624dff0"

_struct_I = struct.Struct('<I')

