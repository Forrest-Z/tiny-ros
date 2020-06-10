import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import shape_msgs.msg

class Plane(tinyros.Message):
    __slots__ = ['coef']
    _slot_types = ['float64[4]']

    def __init__(self):
        super(Plane, self).__init__()
        self.coef = [0. for x in range(0, 4)]

    def serialize(self, buff):
        offset = 0
        try:
            for i in range(0, 4):
                try:
                    struct_coefi = struct.Struct("<d")
                    buff.write(struct_coefi.pack(self.coef[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            self.coef = [0. for x in range(0, 4)]
            for i in range(0, 4):
                try:
                    struct_coefi = struct.Struct("<d")
                    (self.coef[i],) = struct_coefi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        for i in range(0, 4):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"coef": ['
        for i in range(0, 4):
            if i == (4 - 1): 
                string_echo += '{"coef%s": %s'%(i,coef[i])
                string_echo += '}'
            else:
                string_echo += '{"coef%s": %s'%(i,coef[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "shape_msgs/Plane"

    def getMD5(self):
        return "770421286b7c90effe8aac9f1c37eac0"

_struct_I = struct.Struct('<I')

