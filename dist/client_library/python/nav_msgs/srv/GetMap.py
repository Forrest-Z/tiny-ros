import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg

class GetMapRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(GetMapRequest, self).__init__()
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
        return "nav_msgs/GetMap"

    def getMD5(self):
        return "8ea512c109a0b3a9eca8de407dd02d2a"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetMapResponse(tinyros.Message):
    __slots__ = ['__id__','map']
    _slot_types = ['uint32','nav_msgs.msg.OccupancyGrid']

    def __init__(self):
        super(GetMapResponse, self).__init__()
        self.__id__ = 0
        self.map = nav_msgs.msg.OccupancyGrid()

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
        return "nav_msgs/GetMap"

    def getMD5(self):
        return "5bf895a6cdaff312c69ca1cef20fed64"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetMap(object):
    Request = GetMapRequest
    Response = GetMapResponse

_struct_I = struct.Struct('<I')

