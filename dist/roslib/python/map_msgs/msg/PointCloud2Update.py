import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import map_msgs.msg
import std_msgs.msg
import sensor_msgs.msg

class PointCloud2Update(tinyros.Message):
    __slots__ = ['header','type','points']
    _slot_types = ['std_msgs.msg.Header','uint32','sensor_msgs.msg.PointCloud2']

    ADD = 0
    DELETE = 1

    def __init__(self):
        super(PointCloud2Update, self).__init__()
        self.header = std_msgs.msg.Header()
        self.type = 0
        self.points = sensor_msgs.msg.PointCloud2()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            buff.write(_struct_I.pack(self.type))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.points.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (self.type,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.points.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += self.points.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"type": %s'%type
        string_echo += ', '
        string_echo += '"points": {"'
        string_echo += self.points.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/PointCloud2Update"

    def getMD5(self):
        return "e79dfbefd7336861352e1bc7148491c4"

_struct_I = struct.Struct('<I')

