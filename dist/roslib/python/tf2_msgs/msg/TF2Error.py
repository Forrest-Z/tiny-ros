import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg

class TF2Error(tinyros.Message):
    __slots__ = ['error','error_string']
    _slot_types = ['uint8','string']

    NO_ERROR =  0
    LOOKUP_ERROR =  1
    CONNECTIVITY_ERROR =  2
    EXTRAPOLATION_ERROR =  3
    INVALID_ARGUMENT_ERROR =  4
    TIMEOUT_ERROR =  5
    TRANSFORM_ERROR =  6

    def __init__(self):
        super(TF2Error, self).__init__()
        self.error = 0
        self.error_string = ''

    def serialize(self, buff):
        offset = 0
        try:
            struct_error = struct.Struct("<B")
            buff.write(struct_error.pack(self.error))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.error_string
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
            struct_error = struct.Struct("<B")
            (self.error,) = struct_error.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_error_string,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.error_string = buff[offset:(offset + length_error_string)].decode('utf-8')
            else:
                self.error_string = buff[offset:(offset + length_error_string)]
            offset += length_error_string
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        error_string_x = self.error_string
        error_string_length = len(error_string_x)
        if python3 or type(error_string_x) == unicode:
            error_string_x = error_string_x.encode('utf-8')
            error_string_length = len(error_string_x)
        length += 4
        length += error_string_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"error": %s'%error
        string_echo += ', '
        string_echo += '"error_string": "%s"'%error_string
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf2_msgs/TF2Error"

    def getMD5(self):
        return "ed32adf5a372962d977aea0e5630d1d6"

_struct_I = struct.Struct('<I')

