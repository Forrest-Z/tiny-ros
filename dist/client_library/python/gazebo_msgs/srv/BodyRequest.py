import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class BodyRequestRequest(tinyros.Message):
    __slots__ = ['__id__','body_name']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(BodyRequestRequest, self).__init__()
        self.__id__ = 0
        self.body_name = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.body_name
            length = len(_x)
            if python3 or type(_x) == unicode:
                _x = _x.encode('utf-8')
                length = len(_x)
            if python3:
                buff.write(struct.pack('<I%sB'%length, length, *_x))
            else:
                buff.write(struct.pack('<I%ss'%length, length, _x))
            offset += 4
            offset += length
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
            (length_body_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.body_name = buff[offset:(offset + length_body_name)].decode('utf-8')
            else:
                self.body_name = buff[offset:(offset + length_body_name)]
            offset += length_body_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        body_name_x = self.body_name
        body_name_length = len(body_name_x)
        if python3 or type(body_name_x) == unicode:
            body_name_x = body_name_x.encode('utf-8')
            body_name_length = len(body_name_x)
        length += 4
        length += body_name_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"body_name": "%s"'%body_name
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/BodyRequest"

    def getMD5(self):
        return "d1c66fbceb0ee1b51e3b09ac030dedec"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class BodyRequestResponse(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(BodyRequestResponse, self).__init__()
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
        return "gazebo_msgs/BodyRequest"

    def getMD5(self):
        return "e0caf2eb9951542b962f95924c6eeb34"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class BodyRequest(object):
    Request = BodyRequestRequest
    Response = BodyRequestResponse

_struct_I = struct.Struct('<I')

