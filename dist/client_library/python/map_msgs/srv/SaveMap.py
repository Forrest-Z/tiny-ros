import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg
import map_msgs.msg

class SaveMapRequest(tinyros.Message):
    __slots__ = ['__id__','filename']
    _slot_types = ['uint32','std_msgs.msg.String']

    def __init__(self):
        super(SaveMapRequest, self).__init__()
        self.__id__ = 0
        self.filename = std_msgs.msg.String()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.filename.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.filename.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.filename.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"filename": {"'
        string_echo += self.filename.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/SaveMap"

    def getMD5(self):
        return "6643d8ede81a23998690e6a3ff657316"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SaveMapResponse(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(SaveMapResponse, self).__init__()
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
        return "map_msgs/SaveMap"

    def getMD5(self):
        return "9cd07446fa1bd59b4758dadf19f196e9"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SaveMap(object):
    Request = SaveMapRequest
    Response = SaveMapResponse

_struct_I = struct.Struct('<I')

