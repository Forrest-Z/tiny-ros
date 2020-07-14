import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import tinyros
import geometry_msgs.msg

class MapMetaData(tinyros.Message):
    __slots__ = ['map_load_time','resolution','width','height','origin']
    _slot_types = ['tinyros.Time','float32','uint32','uint32','geometry_msgs.msg.Pose']

    def __init__(self):
        super(MapMetaData, self).__init__()
        self.map_load_time = tinyros.Time()
        self.resolution = 0.
        self.width = 0
        self.height = 0
        self.origin = geometry_msgs.msg.Pose()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.map_load_time.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.map_load_time.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_resolution = struct.Struct("<f")
            buff.write(struct_resolution.pack(self.resolution))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.width))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.height))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.origin.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.map_load_time.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.map_load_time.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_resolution = struct.Struct("<f")
            (self.resolution,) = struct_resolution.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.width,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.height,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.origin.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += self.origin.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"map_load_time.sec": %s, '%map_load_time.sec
        string_echo += '"map_load_time.nsec": %s'%map_load_time.nsec
        string_echo += ', '
        string_echo += '"resolution": %s'%resolution
        string_echo += ', '
        string_echo += '"width": %s'%width
        string_echo += ', '
        string_echo += '"height": %s'%height
        string_echo += ', '
        string_echo += '"origin": {"'
        string_echo += self.origin.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/MapMetaData"

    def getMD5(self):
        return "328f5a1f2242fff4676d48189bd8b309"

_struct_I = struct.Struct('<I')

