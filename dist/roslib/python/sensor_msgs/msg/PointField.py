import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class PointField(tinyros.Message):
    __slots__ = ['name','offset','datatype','count']
    _slot_types = ['string','uint32','uint8','uint32']

    INT8 =  1
    UINT8 =  2
    INT16 =  3
    UINT16 =  4
    INT32 =  5
    UINT32 =  6
    FLOAT32 =  7
    FLOAT64 =  8

    def __init__(self):
        super(PointField, self).__init__()
        self.name = ''
        self.offset = 0
        self.datatype = 0
        self.count = 0

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
            buff.write(_struct_I.pack(self.offset))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_datatype = struct.Struct("<B")
            buff.write(struct_datatype.pack(self.datatype))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.count))
            offset += 4
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
            (self.offset,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_datatype = struct.Struct("<B")
            (self.datatype,) = struct_datatype.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.count,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
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
        length += 4
        length += 1
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"name": "%s"'%name
        string_echo += '", '
        string_echo += '"offset": %s'%offset
        string_echo += ', '
        string_echo += '"datatype": %s'%datatype
        string_echo += ', '
        string_echo += '"count": %s'%count
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/PointField"

    def getMD5(self):
        return "039974f05fdf0470d9dc865fd01dcc3e"

_struct_I = struct.Struct('<I')

