import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class CompressedImage(tinyros.Message):
    __slots__ = ['header','format','data']
    _slot_types = ['std_msgs.msg.Header','string','uint8[]']

    def __init__(self):
        super(CompressedImage, self).__init__()
        self.header = std_msgs.msg.Header()
        self.format = ''
        self.data = ''

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            _x = self.format
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
            (length_format,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.format = buff[offset:(offset + length_format)].decode('utf-8')
            else:
                self.format = buff[offset:(offset + length_format)]
            offset += length_format
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
        format_x = self.format
        format_length = len(format_x)
        if python3 or type(format_x) == unicode:
            format_x = format_x.encode('utf-8')
            format_length = len(format_x)
        length += 4
        length += format_length
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
        string_echo += '"format": "%s"'%format
        string_echo += '", '
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
        return "sensor_msgs/CompressedImage"

    def getMD5(self):
        return "eed57d856457441995644e6294152301"

_struct_I = struct.Struct('<I')

