import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tinyros_hello.msg

class TestRequest(tinyros.Message):
    __slots__ = ['__id__','input']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(TestRequest, self).__init__()
        self.__id__ = 0
        self.input = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.input
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
            (length_input,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.input = buff[offset:(offset + length_input)].decode('utf-8')
            else:
                self.input = buff[offset:(offset + length_input)]
            offset += length_input
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        input_x = self.input
        input_length = len(input_x)
        if python3 or type(input_x) == unicode:
            input_x = input_x.encode('utf-8')
            input_length = len(input_x)
        length += 4
        length += input_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"input": "%s"'%input
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_hello/Test"

    def getMD5(self):
        return "26ee7a44335f1f7b55a5a7490460807d"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class TestResponse(tinyros.Message):
    __slots__ = ['__id__','output']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(TestResponse, self).__init__()
        self.__id__ = 0
        self.output = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.output
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
            (length_output,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.output = buff[offset:(offset + length_output)].decode('utf-8')
            else:
                self.output = buff[offset:(offset + length_output)]
            offset += length_output
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        output_x = self.output
        output_length = len(output_x)
        if python3 or type(output_x) == unicode:
            output_x = output_x.encode('utf-8')
            output_length = len(output_x)
        length += 4
        length += output_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"output": "%s"'%output
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_hello/Test"

    def getMD5(self):
        return "e2f6296e6ea9df7406f4fac9fb52d44b"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class Test(object):
    Request = TestRequest
    Response = TestResponse

_struct_I = struct.Struct('<I')

