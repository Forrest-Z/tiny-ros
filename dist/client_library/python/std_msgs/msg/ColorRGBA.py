import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class ColorRGBA(tinyros.Message):
    __slots__ = ['r','g','b','a']
    _slot_types = ['float32','float32','float32','float32']

    def __init__(self):
        super(ColorRGBA, self).__init__()
        self.r = 0.
        self.g = 0.
        self.b = 0.
        self.a = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_r = struct.Struct("<f")
            buff.write(struct_r.pack(self.r))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_g = struct.Struct("<f")
            buff.write(struct_g.pack(self.g))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_b = struct.Struct("<f")
            buff.write(struct_b.pack(self.b))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_a = struct.Struct("<f")
            buff.write(struct_a.pack(self.a))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_r = struct.Struct("<f")
            (self.r,) = struct_r.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_g = struct.Struct("<f")
            (self.g,) = struct_g.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_b = struct.Struct("<f")
            (self.b,) = struct_b.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_a = struct.Struct("<f")
            (self.a,) = struct_a.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"r": %s'%r
        string_echo += ', '
        string_echo += '"g": %s'%g
        string_echo += ', '
        string_echo += '"b": %s'%b
        string_echo += ', '
        string_echo += '"a": %s'%a
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/ColorRGBA"

    def getMD5(self):
        return "3c740aa311a337bfa4133c69405e0aed"

_struct_I = struct.Struct('<I')

