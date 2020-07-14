import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class PointCloud(tinyros.Message):
    __slots__ = ['header','points','channels']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Point32[]','sensor_msgs.msg.ChannelFloat32[]']

    def __init__(self):
        super(PointCloud, self).__init__()
        self.header = std_msgs.msg.Header()
        self.points = []
        self.channels = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            points_length = len(self.points)
            buff.write(_struct_I.pack(points_length))
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            channels_length = len(self.channels)
            buff.write(_struct_I.pack(channels_length))
            offset += 4
            for i in range(0, channels_length):
                offset += self.channels[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (points_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.points = [geometry_msgs.msg.Point32() for x in range(0, points_length)]
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (channels_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.channels = [sensor_msgs.msg.ChannelFloat32() for x in range(0, channels_length)]
            offset += 4
            for i in range(0, channels_length):
                offset += self.channels[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        points_length = len(self.points)
        length += 4
        for i in range(0, points_length):
            length += self.points[i].serializedLength()
        channels_length = len(self.channels)
        length += 4
        for i in range(0, channels_length):
            length += self.channels[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
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
        string_echo += '], '
        string_echo += '"channels": ['
        channels_length = len(self.channels)
        for i in range(0, channels_length):
            if i == (channels_length - 1): 
                string_echo += '{"channels%s": {'%i
                string_echo += self.channels[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"channels%s": {'%i
                string_echo += self.channels[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/PointCloud"

    def getMD5(self):
        return "b01249148cae0106a561ab36cd1e48a8"

_struct_I = struct.Struct('<I')

