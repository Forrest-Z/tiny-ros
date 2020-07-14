import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class MultiArrayLayout(tinyros.Message):
    __slots__ = ['dim','data_offset']
    _slot_types = ['std_msgs.msg.MultiArrayDimension[]','uint32']

    def __init__(self):
        super(MultiArrayLayout, self).__init__()
        self.dim = []
        self.data_offset = 0

    def serialize(self, buff):
        offset = 0
        try:
            dim_length = len(self.dim)
            buff.write(_struct_I.pack(dim_length))
            offset += 4
            for i in range(0, dim_length):
                offset += self.dim[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.data_offset))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (dim_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.dim = [std_msgs.msg.MultiArrayDimension() for x in range(0, dim_length)]
            offset += 4
            for i in range(0, dim_length):
                offset += self.dim[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.data_offset,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        dim_length = len(self.dim)
        length += 4
        for i in range(0, dim_length):
            length += self.dim[i].serializedLength()
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"dim": ['
        dim_length = len(self.dim)
        for i in range(0, dim_length):
            if i == (dim_length - 1): 
                string_echo += '{"dim%s": {'%i
                string_echo += self.dim[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"dim%s": {'%i
                string_echo += self.dim[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"data_offset": %s'%data_offset
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/MultiArrayLayout"

    def getMD5(self):
        return "f40f0b5b285a93ca167c98c1012a989a"

_struct_I = struct.Struct('<I')

