import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import map_msgs.msg

class GetPointMapRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(GetPointMapRequest, self).__init__()
        self.__id__ = 0

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
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
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/GetPointMap"

    def getMD5(self):
        return "418d4ee8c9d758b7ef1aab3e068c7568"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPointMapResponse(tinyros.Message):
    __slots__ = ['__id__','map']
    _slot_types = ['uint32','sensor_msgs.msg.PointCloud2']

    def __init__(self):
        super(GetPointMapResponse, self).__init__()
        self.__id__ = 0
        self.map = sensor_msgs.msg.PointCloud2()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.map.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.map.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.map.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"map": {"'
        string_echo += self.map.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/GetPointMap"

    def getMD5(self):
        return "abc97e6c5b25f536248da280bdf271d7"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPointMap(object):
    Request = GetPointMapRequest
    Response = GetPointMapResponse

_struct_I = struct.Struct('<I')

