import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class ChannelFloat32(tinyros.Message):
    __slots__ = ['name','values']
    _slot_types = ['string','float32[]']

    def __init__(self):
        super(ChannelFloat32, self).__init__()
        self.name = ''
        self.values = []

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.name
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
            values_length = len(self.values)
            buff.write(_struct_I.pack(values_length))
            offset += 4
            for i in range(0, values_length):
                try:
                    struct_valuesi = struct.Struct("<f")
                    buff.write(struct_valuesi.pack(self.values[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.name = buff[offset:(offset + length_name)].decode('utf-8')
            else:
                self.name = buff[offset:(offset + length_name)]
            offset += length_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (values_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.values = [0. for x in range(0, values_length)]
            offset += 4
            for i in range(0, values_length):
                try:
                    struct_valuesi = struct.Struct("<f")
                    (self.values[i],) = struct_valuesi.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        name_x = self.name
        name_length = len(name_x)
        if python3 or type(name_x) == unicode:
            name_x = name_x.encode('utf-8')
            name_length = len(name_x)
        length += 4
        length += name_length
        values_length = len(self.values)
        length += 4
        for i in range(0, values_length):
            length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"name": "%s"'%name
        string_echo += '", '
        string_echo += '"values": ['
        values_length = len(self.values)
        for i in range(0, values_length):
            if i == (values_length - 1): 
                string_echo += '{"values%s": %s'%(i,values[i])
                string_echo += '}'
            else:
                string_echo += '{"values%s": %s'%(i,values[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/ChannelFloat32"

    def getMD5(self):
        return "c4cf01c81334c609dca1afd3a227daff"

_struct_I = struct.Struct('<I')

