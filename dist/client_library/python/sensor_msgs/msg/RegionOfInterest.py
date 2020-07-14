import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class RegionOfInterest(tinyros.Message):
    __slots__ = ['x_offset','y_offset','height','width','do_rectify']
    _slot_types = ['uint32','uint32','uint32','uint32','bool']

    def __init__(self):
        super(RegionOfInterest, self).__init__()
        self.x_offset = 0
        self.y_offset = 0
        self.height = 0
        self.width = 0
        self.do_rectify = False

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.x_offset))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.y_offset))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.height))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.width))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_do_rectify = struct.Struct("<B")
            buff.write(struct_do_rectify.pack(self.do_rectify))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.x_offset,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.y_offset,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.height,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.width,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_do_rectify = struct.Struct("<B")
            (self.do_rectify,) = struct_do_rectify.unpack(buff[offset:(offset + 1)])
            self.do_rectify = bool(self.do_rectify)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        length += 4
        length += 4
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"x_offset": %s'%x_offset
        string_echo += ', '
        string_echo += '"y_offset": %s'%y_offset
        string_echo += ', '
        string_echo += '"height": %s'%height
        string_echo += ', '
        string_echo += '"width": %s'%width
        string_echo += ', '
        string_echo += '"do_rectify": %s'%do_rectify
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/RegionOfInterest"

    def getMD5(self):
        return "8370dc286f915405c906299aef5bb442"

_struct_I = struct.Struct('<I')

