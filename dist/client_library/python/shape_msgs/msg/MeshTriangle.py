import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import shape_msgs.msg

class MeshTriangle(tinyros.Message):
    __slots__ = ['vertex_indices']
    _slot_types = ['uint32[3]']

    def __init__(self):
        super(MeshTriangle, self).__init__()
        self.vertex_indices = [0 for x in range(0, 3)]

    def serialize(self, buff):
        offset = 0
        try:
            for i in range(0, 3):
                try:
                    buff.write(_struct_I.pack(self.vertex_indices[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            self.vertex_indices = [0 for x in range(0, 3)]
            for i in range(0, 3):
                try:
                    (self.vertex_indices[i],) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        for i in range(0, 3):
            length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"vertex_indices": ['
        for i in range(0, 3):
            if i == (3 - 1): 
                string_echo += '{"vertex_indices%s": %s'%(i,vertex_indices[i])
                string_echo += '}'
            else:
                string_echo += '{"vertex_indices%s": %s'%(i,vertex_indices[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "shape_msgs/MeshTriangle"

    def getMD5(self):
        return "01020cfeb9ad7679dd18bbd7149962ba"

_struct_I = struct.Struct('<I')

