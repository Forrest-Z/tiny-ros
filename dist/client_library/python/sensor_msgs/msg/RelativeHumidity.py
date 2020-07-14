import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class RelativeHumidity(tinyros.Message):
    __slots__ = ['header','relative_humidity','variance']
    _slot_types = ['std_msgs.msg.Header','float64','float64']

    def __init__(self):
        super(RelativeHumidity, self).__init__()
        self.header = std_msgs.msg.Header()
        self.relative_humidity = 0.
        self.variance = 0.

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_relative_humidity = struct.Struct("<d")
            buff.write(struct_relative_humidity.pack(self.relative_humidity))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_variance = struct.Struct("<d")
            buff.write(struct_variance.pack(self.variance))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_relative_humidity = struct.Struct("<d")
            (self.relative_humidity,) = struct_relative_humidity.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_variance = struct.Struct("<d")
            (self.variance,) = struct_variance.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"relative_humidity": %s'%relative_humidity
        string_echo += ', '
        string_echo += '"variance": %s'%variance
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/RelativeHumidity"

    def getMD5(self):
        return "d9a3a4b2c3c0c55eede767d38b460110"

_struct_I = struct.Struct('<I')

