import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import map_msgs.msg

class SetMapProjectionsRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(SetMapProjectionsRequest, self).__init__()
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
        return "map_msgs/SetMapProjections"

    def getMD5(self):
        return "26dbba584adf9695d68b8667830be463"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetMapProjectionsResponse(tinyros.Message):
    __slots__ = ['__id__','projected_maps_info']
    _slot_types = ['uint32','map_msgs.msg.ProjectedMapInfo[]']

    def __init__(self):
        super(SetMapProjectionsResponse, self).__init__()
        self.__id__ = 0
        self.projected_maps_info = []

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            projected_maps_info_length = len(self.projected_maps_info)
            buff.write(_struct_I.pack(projected_maps_info_length))
            offset += 4
            for i in range(0, projected_maps_info_length):
                offset += self.projected_maps_info[i].serialize(buff)
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
            (projected_maps_info_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.projected_maps_info = [map_msgs.msg.ProjectedMapInfo() for x in range(0, projected_maps_info_length)]
            offset += 4
            for i in range(0, projected_maps_info_length):
                offset += self.projected_maps_info[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        projected_maps_info_length = len(self.projected_maps_info)
        length += 4
        for i in range(0, projected_maps_info_length):
            length += self.projected_maps_info[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"projected_maps_info": ['
        projected_maps_info_length = len(self.projected_maps_info)
        for i in range(0, projected_maps_info_length):
            if i == (projected_maps_info_length - 1): 
                string_echo += '{"projected_maps_info%s": {'%i
                string_echo += self.projected_maps_info[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"projected_maps_info%s": {'%i
                string_echo += self.projected_maps_info[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/SetMapProjections"

    def getMD5(self):
        return "47b7931263dc316e5b0f0264428853e9"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetMapProjections(object):
    Request = SetMapProjectionsRequest
    Response = SetMapProjectionsResponse

_struct_I = struct.Struct('<I')

