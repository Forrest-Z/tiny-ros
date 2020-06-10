import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import map_msgs.msg
import nav_msgs.msg

class ProjectedMap(tinyros.Message):
    __slots__ = ['map','min_z','max_z']
    _slot_types = ['nav_msgs.msg.OccupancyGrid','float64','float64']

    def __init__(self):
        super(ProjectedMap, self).__init__()
        self.map = nav_msgs.msg.OccupancyGrid()
        self.min_z = 0.
        self.max_z = 0.

    def serialize(self, buff):
        offset = 0
        offset += self.map.serialize(buff)
        try:
            struct_min_z = struct.Struct("<d")
            buff.write(struct_min_z.pack(self.min_z))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_max_z = struct.Struct("<d")
            buff.write(struct_max_z.pack(self.max_z))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.map.deserialize(buff[offset:])
        try:
            struct_min_z = struct.Struct("<d")
            (self.min_z,) = struct_min_z.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_max_z = struct.Struct("<d")
            (self.max_z,) = struct_max_z.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.map.serializedLength()
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"map": {"'
        string_echo += self.map.echo()
        string_echo += '}, '
        string_echo += '"min_z": %s'%min_z
        string_echo += ', '
        string_echo += '"max_z": %s'%max_z
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/ProjectedMap"

    def getMD5(self):
        return "cbd5598c259cc16f5aa07335587a7367"

_struct_I = struct.Struct('<I')

