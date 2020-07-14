import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class String(tinyros.Message):
    __slots__ = ['data']
    _slot_types = ['string']

    def __init__(self):
        super(String, self).__init__()
        self.data = ''

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.data
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
            (length_data,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.data = buff[offset:(offset + length_data)].decode('utf-8')
            else:
                self.data = buff[offset:(offset + length_data)]
            offset += length_data
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        data_x = self.data
        data_length = len(data_x)
        if python3 or type(data_x) == unicode:
            data_x = data_x.encode('utf-8')
            data_length = len(data_x)
        length += 4
        length += data_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"data": "%s"'%data
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/String"

    def getMD5(self):
        return "44b822292b4a9ed05e241aa225458f97"

_struct_I = struct.Struct('<I')

