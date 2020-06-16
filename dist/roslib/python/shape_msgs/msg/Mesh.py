import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import shape_msgs.msg
import geometry_msgs.msg

class Mesh(tinyros.Message):
    __slots__ = ['triangles','vertices']
    _slot_types = ['shape_msgs.msg.MeshTriangle[]','geometry_msgs.msg.Point[]']

    def __init__(self):
        super(Mesh, self).__init__()
        self.triangles = []
        self.vertices = []

    def serialize(self, buff):
        offset = 0
        try:
            triangles_length = len(self.triangles)
            buff.write(_struct_I.pack(triangles_length))
            offset += 4
            for i in range(0, triangles_length):
                offset += self.triangles[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            vertices_length = len(self.vertices)
            buff.write(_struct_I.pack(vertices_length))
            offset += 4
            for i in range(0, vertices_length):
                offset += self.vertices[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (triangles_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.triangles = [shape_msgs.msg.MeshTriangle() for x in range(0, triangles_length)]
            offset += 4
            for i in range(0, triangles_length):
                offset += self.triangles[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (vertices_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.vertices = [geometry_msgs.msg.Point() for x in range(0, vertices_length)]
            offset += 4
            for i in range(0, vertices_length):
                offset += self.vertices[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        triangles_length = len(self.triangles)
        length += 4
        for i in range(0, triangles_length):
            length += self.triangles[i].serializedLength()
        vertices_length = len(self.vertices)
        length += 4
        for i in range(0, vertices_length):
            length += self.vertices[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"triangles": ['
        triangles_length = len(self.triangles)
        for i in range(0, triangles_length):
            if i == (triangles_length - 1): 
                string_echo += '{"triangles%s": {'%i
                string_echo += self.triangles[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"triangles%s": {'%i
                string_echo += self.triangles[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"vertices": ['
        vertices_length = len(self.vertices)
        for i in range(0, vertices_length):
            if i == (vertices_length - 1): 
                string_echo += '{"vertices%s": {'%i
                string_echo += self.vertices[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"vertices%s": {'%i
                string_echo += self.vertices[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "shape_msgs/Mesh"

    def getMD5(self):
        return "1579670b316f622b47d6700cd4f7e18d"

_struct_I = struct.Struct('<I')

