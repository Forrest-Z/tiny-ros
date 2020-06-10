import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class Range(tinyros.Message):
    __slots__ = ['header','radiation_type','field_of_view','min_range','max_range','range']
    _slot_types = ['std_msgs.msg.Header','uint8','float32','float32','float32','float32']

    ULTRASOUND = 0
    INFRARED = 1

    def __init__(self):
        super(Range, self).__init__()
        self.header = std_msgs.msg.Header()
        self.radiation_type = 0
        self.field_of_view = 0.
        self.min_range = 0.
        self.max_range = 0.
        self.range = 0.

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_radiation_type = struct.Struct("<B")
            buff.write(struct_radiation_type.pack(self.radiation_type))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_field_of_view = struct.Struct("<f")
            buff.write(struct_field_of_view.pack(self.field_of_view))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_min_range = struct.Struct("<f")
            buff.write(struct_min_range.pack(self.min_range))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_max_range = struct.Struct("<f")
            buff.write(struct_max_range.pack(self.max_range))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_range = struct.Struct("<f")
            buff.write(struct_range.pack(self.range))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_radiation_type = struct.Struct("<B")
            (self.radiation_type,) = struct_radiation_type.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_field_of_view = struct.Struct("<f")
            (self.field_of_view,) = struct_field_of_view.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_min_range = struct.Struct("<f")
            (self.min_range,) = struct_min_range.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_max_range = struct.Struct("<f")
            (self.max_range,) = struct_max_range.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_range = struct.Struct("<f")
            (self.range,) = struct_range.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 1
        length += 4
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"radiation_type": %s'%radiation_type
        string_echo += ', '
        string_echo += '"field_of_view": %s'%field_of_view
        string_echo += ', '
        string_echo += '"min_range": %s'%min_range
        string_echo += ', '
        string_echo += '"max_range": %s'%max_range
        string_echo += ', '
        string_echo += '"range": %s'%range
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/Range"

    def getMD5(self):
        return "54d647e3a481f5b87ce39d1b97e84d53"

_struct_I = struct.Struct('<I')

