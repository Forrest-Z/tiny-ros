import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import map_msgs.msg

class GetMapROIRequest(tinyros.Message):
    __slots__ = ['__id__','x','y','l_x','l_y']
    _slot_types = ['uint32','float64','float64','float64','float64']

    def __init__(self):
        super(GetMapROIRequest, self).__init__()
        self.__id__ = 0
        self.x = 0.
        self.y = 0.
        self.l_x = 0.
        self.l_y = 0.

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_x = struct.Struct("<d")
            buff.write(struct_x.pack(self.x))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            buff.write(struct_y.pack(self.y))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_l_x = struct.Struct("<d")
            buff.write(struct_l_x.pack(self.l_x))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_l_y = struct.Struct("<d")
            buff.write(struct_l_y.pack(self.l_y))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_x = struct.Struct("<d")
            (self.x,) = struct_x.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            (self.y,) = struct_y.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_l_x = struct.Struct("<d")
            (self.l_x,) = struct_l_x.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_l_y = struct.Struct("<d")
            (self.l_y,) = struct_l_y.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 8
        length += 8
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"x": %s'%x
        string_echo += ', '
        string_echo += '"y": %s'%y
        string_echo += ', '
        string_echo += '"l_x": %s'%l_x
        string_echo += ', '
        string_echo += '"l_y": %s'%l_y
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/GetMapROI"

    def getMD5(self):
        return "f74ea8c8dc9b857aae7ea10033520d28"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetMapROIResponse(tinyros.Message):
    __slots__ = ['__id__','sub_map']
    _slot_types = ['uint32','nav_msgs.msg.OccupancyGrid']

    def __init__(self):
        super(GetMapROIResponse, self).__init__()
        self.__id__ = 0
        self.sub_map = nav_msgs.msg.OccupancyGrid()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.sub_map.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.sub_map.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.sub_map.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"sub_map": {"'
        string_echo += self.sub_map.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/GetMapROI"

    def getMD5(self):
        return "a178ec520c3d0d99d9d85c70ed4b535a"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetMapROI(object):
    Request = GetMapROIRequest
    Response = GetMapROIResponse

_struct_I = struct.Struct('<I')

