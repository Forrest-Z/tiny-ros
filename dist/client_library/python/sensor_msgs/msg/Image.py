import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class Image(tinyros.Message):
    __slots__ = ['header','height','width','encoding','is_bigendian','step','data']
    _slot_types = ['std_msgs.msg.Header','uint32','uint32','string','uint8','uint32','uint8[]']

    def __init__(self):
        super(Image, self).__init__()
        self.header = std_msgs.msg.Header()
        self.height = 0
        self.width = 0
        self.encoding = ''
        self.is_bigendian = 0
        self.step = 0
        self.data = ''

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            buff.write(_struct_I.pack(self.height))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.width))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.encoding
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
            struct_is_bigendian = struct.Struct("<B")
            buff.write(struct_is_bigendian.pack(self.is_bigendian))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.step))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            data_length = len(self.data)
            buff.write(_struct_I.pack(data_length))
            offset += 4
            for i in range(0, data_length):
                try:
                    struct_datai = struct.Struct("<B")
                    buff.write(struct_datai.pack(self.data[i]))
                    offset += 1
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (self.height,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.width,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_encoding,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.encoding = buff[offset:(offset + length_encoding)].decode('utf-8')
            else:
                self.encoding = buff[offset:(offset + length_encoding)]
            offset += length_encoding
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_is_bigendian = struct.Struct("<B")
            (self.is_bigendian,) = struct_is_bigendian.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.step,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (data_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.data = [0 for x in range(0, data_length)]
            offset += 4
            for i in range(0, data_length):
                try:
                    struct_datai = struct.Struct("<B")
                    (self.data[i],) = struct_datai.unpack(buff[offset:(offset + 1)])
                    offset += 1
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        encoding_x = self.encoding
        encoding_length = len(encoding_x)
        if python3 or type(encoding_x) == unicode:
            encoding_x = encoding_x.encode('utf-8')
            encoding_length = len(encoding_x)
        length += 4
        length += encoding_length
        length += 1
        length += 4
        data_length = len(self.data)
        length += 4
        for i in range(0, data_length):
            length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"height": %s'%height
        string_echo += ', '
        string_echo += '"width": %s'%width
        string_echo += ', '
        string_echo += '"encoding": "%s"'%encoding
        string_echo += '", '
        string_echo += '"is_bigendian": %s'%is_bigendian
        string_echo += ', '
        string_echo += '"step": %s'%step
        string_echo += ', '
        string_echo += '"data": ['
        data_length = len(self.data)
        for i in range(0, data_length):
            if i == (data_length - 1): 
                string_echo += '{"data%s": %s'%(i,data[i])
                string_echo += '}'
            else:
                string_echo += '{"data%s": %s'%(i,data[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/Image"

    def getMD5(self):
        return "886f928dc81bf7f1496a8b452057c5b2"

_struct_I = struct.Struct('<I')

