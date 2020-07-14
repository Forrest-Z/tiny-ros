import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class Temperature(tinyros.Message):
    __slots__ = ['header','temperature','variance']
    _slot_types = ['std_msgs.msg.Header','float64','float64']

    def __init__(self):
        super(Temperature, self).__init__()
        self.header = std_msgs.msg.Header()
        self.temperature = 0.
        self.variance = 0.

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_temperature = struct.Struct("<d")
            buff.write(struct_temperature.pack(self.temperature))
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
            struct_temperature = struct.Struct("<d")
            (self.temperature,) = struct_temperature.unpack(buff[offset:(offset + 8)])
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
        string_echo += '"temperature": %s'%temperature
        string_echo += ', '
        string_echo += '"variance": %s'%variance
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/Temperature"

    def getMD5(self):
        return "898a0e5950c8e4073c0c3cf2d7e7bf26"

_struct_I = struct.Struct('<I')

