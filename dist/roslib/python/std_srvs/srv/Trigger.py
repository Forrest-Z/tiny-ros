import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_srvs.msg

class TriggerRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(TriggerRequest, self).__init__()
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
        return "std_srvs/Trigger"

    def getMD5(self):
        return "23ce35174a691cb0bfbbdb21395edbac"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class TriggerResponse(tinyros.Message):
    __slots__ = ['__id__','success','message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(TriggerResponse, self).__init__()
        self.__id__ = 0
        self.success = False
        self.message = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_success = struct.Struct("<B")
            buff.write(struct_success.pack(self.success))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.message
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
            struct_success = struct.Struct("<B")
            (self.success,) = struct_success.unpack(buff[offset:(offset + 1)])
            self.success = bool(self.success)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_message,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.message = buff[offset:(offset + length_message)].decode('utf-8')
            else:
                self.message = buff[offset:(offset + length_message)]
            offset += length_message
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 1
        message_x = self.message
        message_length = len(message_x)
        if python3 or type(message_x) == unicode:
            message_x = message_x.encode('utf-8')
            message_length = len(message_x)
        length += 4
        length += message_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"message": "%s"'%message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_srvs/Trigger"

    def getMD5(self):
        return "08d111154ed595049573252ba929a6d8"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class Trigger(object):
    Request = TriggerRequest
    Response = TriggerResponse

_struct_I = struct.Struct('<I')

