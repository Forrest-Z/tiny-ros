import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import diagnostic_msgs.msg

class SelfTestRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(SelfTestRequest, self).__init__()
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
        return "diagnostic_msgs/SelfTest"

    def getMD5(self):
        return "049f87742408b36b8ef5f7dd71e3ef5a"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SelfTestResponse(tinyros.Message):
    __slots__ = ['__id__','id','passed','status']
    _slot_types = ['uint32','string','byte','diagnostic_msgs.msg.DiagnosticStatus[]']

    def __init__(self):
        super(SelfTestResponse, self).__init__()
        self.__id__ = 0
        self.id = ''
        self.passed = 0
        self.status = []

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.id
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
        try:
            struct_passed = struct.Struct("<b")
            buff.write(struct_passed.pack(self.passed))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            status_length = len(self.status)
            buff.write(_struct_I.pack(status_length))
            offset += 4
            for i in range(0, status_length):
                offset += self.status[i].serialize(buff)
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
            (length_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.id = buff[offset:(offset + length_id)].decode('utf-8')
            else:
                self.id = buff[offset:(offset + length_id)]
            offset += length_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_passed = struct.Struct("<b")
            (self.passed,) = struct_passed.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (status_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.status = [diagnostic_msgs.msg.DiagnosticStatus() for x in range(0, status_length)]
            offset += 4
            for i in range(0, status_length):
                offset += self.status[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        id_x = self.id
        id_length = len(id_x)
        if python3 or type(id_x) == unicode:
            id_x = id_x.encode('utf-8')
            id_length = len(id_x)
        length += 4
        length += id_length
        length += 1
        status_length = len(self.status)
        length += 4
        for i in range(0, status_length):
            length += self.status[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"id": "%s"'%id
        string_echo += '", '
        string_echo += '"passed": %s'%passed
        string_echo += ', '
        string_echo += '"status": ['
        status_length = len(self.status)
        for i in range(0, status_length):
            if i == (status_length - 1): 
                string_echo += '{"status%s": {'%i
                string_echo += self.status[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"status%s": {'%i
                string_echo += self.status[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "diagnostic_msgs/SelfTest"

    def getMD5(self):
        return "70aaf2a851ccb5e946b2d112ea26f7b9"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SelfTest(object):
    Request = SelfTestRequest
    Response = SelfTestResponse

_struct_I = struct.Struct('<I')

