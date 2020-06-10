import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import stereo_msgs.msg
import std_msgs.msg
import sensor_msgs.msg

class DisparityImage(tinyros.Message):
    __slots__ = ['header','image','f','T','valid_window','min_disparity','max_disparity','delta_d']
    _slot_types = ['std_msgs.msg.Header','sensor_msgs.msg.Image','float32','float32','sensor_msgs.msg.RegionOfInterest','float32','float32','float32']

    def __init__(self):
        super(DisparityImage, self).__init__()
        self.header = std_msgs.msg.Header()
        self.image = sensor_msgs.msg.Image()
        self.f = 0.
        self.T = 0.
        self.valid_window = sensor_msgs.msg.RegionOfInterest()
        self.min_disparity = 0.
        self.max_disparity = 0.
        self.delta_d = 0.

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.image.serialize(buff)
        try:
            struct_f = struct.Struct("<f")
            buff.write(struct_f.pack(self.f))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_T = struct.Struct("<f")
            buff.write(struct_T.pack(self.T))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.valid_window.serialize(buff)
        try:
            struct_min_disparity = struct.Struct("<f")
            buff.write(struct_min_disparity.pack(self.min_disparity))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_max_disparity = struct.Struct("<f")
            buff.write(struct_max_disparity.pack(self.max_disparity))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_delta_d = struct.Struct("<f")
            buff.write(struct_delta_d.pack(self.delta_d))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.image.deserialize(buff[offset:])
        try:
            struct_f = struct.Struct("<f")
            (self.f,) = struct_f.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_T = struct.Struct("<f")
            (self.T,) = struct_T.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.valid_window.deserialize(buff[offset:])
        try:
            struct_min_disparity = struct.Struct("<f")
            (self.min_disparity,) = struct_min_disparity.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_max_disparity = struct.Struct("<f")
            (self.max_disparity,) = struct_max_disparity.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_delta_d = struct.Struct("<f")
            (self.delta_d,) = struct_delta_d.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.image.serializedLength()
        length += 4
        length += 4
        length += self.valid_window.serializedLength()
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"image": {"'
        string_echo += self.image.echo()
        string_echo += '}, '
        string_echo += '"f": %s'%f
        string_echo += ', '
        string_echo += '"T": %s'%T
        string_echo += ', '
        string_echo += '"valid_window": {"'
        string_echo += self.valid_window.echo()
        string_echo += '}, '
        string_echo += '"min_disparity": %s'%min_disparity
        string_echo += ', '
        string_echo += '"max_disparity": %s'%max_disparity
        string_echo += ', '
        string_echo += '"delta_d": %s'%delta_d
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "stereo_msgs/DisparityImage"

    def getMD5(self):
        return "03545cef8df8d20bea21fdbbf9482b4b"

_struct_I = struct.Struct('<I')

