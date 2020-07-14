import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg

class MultiArrayDimension(tinyros.Message):
    __slots__ = ['label','size','stride']
    _slot_types = ['string','uint32','uint32']

    def __init__(self):
        super(MultiArrayDimension, self).__init__()
        self.label = ''
        self.size = 0
        self.stride = 0

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.label
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
            buff.write(_struct_I.pack(self.size))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stride))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_label,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.label = buff[offset:(offset + length_label)].decode('utf-8')
            else:
                self.label = buff[offset:(offset + length_label)]
            offset += length_label
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.size,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stride,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        label_x = self.label
        label_length = len(label_x)
        if python3 or type(label_x) == unicode:
            label_x = label_x.encode('utf-8')
            label_length = len(label_x)
        length += 4
        length += label_length
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"label": "%s"'%label
        string_echo += '", '
        string_echo += '"size": %s'%size
        string_echo += ', '
        string_echo += '"stride": %s'%stride
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/MultiArrayDimension"

    def getMD5(self):
        return "c2aacf83d49c7aa4a8504bd32158e990"

_struct_I = struct.Struct('<I')

