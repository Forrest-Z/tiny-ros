import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import shape_msgs.msg

class SolidPrimitive(tinyros.Message):
    __slots__ = ['type','dimensions']
    _slot_types = ['uint8','float64[]']

    BOX = 1
    SPHERE = 2
    CYLINDER = 3
    CONE = 4
    BOX_X = 0
    BOX_Y = 1
    BOX_Z = 2
    SPHERE_RADIUS = 0
    CYLINDER_HEIGHT = 0
    CYLINDER_RADIUS = 1
    CONE_HEIGHT = 0
    CONE_RADIUS = 1

    def __init__(self):
        super(SolidPrimitive, self).__init__()
        self.type = 0
        self.dimensions = []

    def serialize(self, buff):
        offset = 0
        try:
            struct_type = struct.Struct("<B")
            buff.write(struct_type.pack(self.type))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            dimensions_length = len(self.dimensions)
            buff.write(_struct_I.pack(dimensions_length))
            offset += 4
            for i in range(0, dimensions_length):
                try:
                    struct_dimensionsi = struct.Struct("<d")
                    buff.write(struct_dimensionsi.pack(self.dimensions[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_type = struct.Struct("<B")
            (self.type,) = struct_type.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (dimensions_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.dimensions = [0. for x in range(0, dimensions_length)]
            offset += 4
            for i in range(0, dimensions_length):
                try:
                    struct_dimensionsi = struct.Struct("<d")
                    (self.dimensions[i],) = struct_dimensionsi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        dimensions_length = len(self.dimensions)
        length += 4
        for i in range(0, dimensions_length):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"type": %s'%type
        string_echo += ', '
        string_echo += '"dimensions": ['
        dimensions_length = len(self.dimensions)
        for i in range(0, dimensions_length):
            if i == (dimensions_length - 1): 
                string_echo += '{"dimensions%s": %s'%(i,dimensions[i])
                string_echo += '}'
            else:
                string_echo += '{"dimensions%s": %s'%(i,dimensions[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "shape_msgs/SolidPrimitive"

    def getMD5(self):
        return "f40805922065416be24909fd8fd751b5"

_struct_I = struct.Struct('<I')

