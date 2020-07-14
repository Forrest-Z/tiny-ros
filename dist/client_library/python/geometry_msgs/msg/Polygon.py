import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class Polygon(tinyros.Message):
    __slots__ = ['points']
    _slot_types = ['geometry_msgs.msg.Point32[]']

    def __init__(self):
        super(Polygon, self).__init__()
        self.points = []

    def serialize(self, buff):
        offset = 0
        try:
            points_length = len(self.points)
            buff.write(_struct_I.pack(points_length))
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (points_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.points = [geometry_msgs.msg.Point32() for x in range(0, points_length)]
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        points_length = len(self.points)
        length += 4
        for i in range(0, points_length):
            length += self.points[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"points": ['
        points_length = len(self.points)
        for i in range(0, points_length):
            if i == (points_length - 1): 
                string_echo += '{"points%s": {'%i
                string_echo += self.points[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"points%s": {'%i
                string_echo += self.points[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/Polygon"

    def getMD5(self):
        return "f94a78a947b7879954bd14397db4bc9d"

_struct_I = struct.Struct('<I')

