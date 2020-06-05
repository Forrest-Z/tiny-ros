import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tinyros_hello.msg

class TinyrosHello(tinyros.Message):
    __slots__ = ['hello']
    _slot_types = ['string']

    def __init__(self):
        super(TinyrosHello, self).__init__()
        self.hello = ''

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.hello
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
            (length_hello,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.hello = buff[offset:(offset + length_hello)].decode('utf-8')
            else:
                self.hello = buff[offset:(offset + length_hello)]
            offset += length_hello
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        hello_x = self.hello
        hello_length = len(hello_x)
        if python3 or type(hello_x) == unicode:
            hello_x = hello_x.encode('utf-8')
            hello_length = len(hello_x)
        length += 4
        length += hello_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"hello": "%s"'%hello
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_hello/TinyrosHello"

    def getMD5(self):
        return "0c68e66a217802ad0c9f648b7a69d090"

_struct_I = struct.Struct('<I')

