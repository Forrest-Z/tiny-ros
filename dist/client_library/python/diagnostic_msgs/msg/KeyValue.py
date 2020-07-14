import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import diagnostic_msgs.msg

class KeyValue(tinyros.Message):
    __slots__ = ['key','value']
    _slot_types = ['string','string']

    def __init__(self):
        super(KeyValue, self).__init__()
        self.key = ''
        self.value = ''

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.key
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
            _x = self.value
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
            (length_key,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.key = buff[offset:(offset + length_key)].decode('utf-8')
            else:
                self.key = buff[offset:(offset + length_key)]
            offset += length_key
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_value,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.value = buff[offset:(offset + length_value)].decode('utf-8')
            else:
                self.value = buff[offset:(offset + length_value)]
            offset += length_value
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        key_x = self.key
        key_length = len(key_x)
        if python3 or type(key_x) == unicode:
            key_x = key_x.encode('utf-8')
            key_length = len(key_x)
        length += 4
        length += key_length
        value_x = self.value
        value_length = len(value_x)
        if python3 or type(value_x) == unicode:
            value_x = value_x.encode('utf-8')
            value_length = len(value_x)
        length += 4
        length += value_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"key": "%s"'%key
        string_echo += '", '
        string_echo += '"value": "%s"'%value
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "diagnostic_msgs/KeyValue"

    def getMD5(self):
        return "1baa904b80c685c77d1a42a872ca1d07"

_struct_I = struct.Struct('<I')

