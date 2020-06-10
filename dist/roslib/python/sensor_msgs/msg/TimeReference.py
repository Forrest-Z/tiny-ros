import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg
import tinyros

class TimeReference(tinyros.Message):
    __slots__ = ['header','time_ref','source']
    _slot_types = ['std_msgs.msg.Header','tinyros.Time','string']

    def __init__(self):
        super(TimeReference, self).__init__()
        self.header = std_msgs.msg.Header()
        self.time_ref = tinyros.Time()
        self.source = ''

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            buff.write(_struct_I.pack(self.time_ref.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.time_ref.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.source
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
        offset += self.header.deserialize(buff[offset:])
        try:
            (self.time_ref.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.time_ref.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_source,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.source = buff[offset:(offset + length_source)].decode('utf-8')
            else:
                self.source = buff[offset:(offset + length_source)]
            offset += length_source
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        source_x = self.source
        source_length = len(source_x)
        if python3 or type(source_x) == unicode:
            source_x = source_x.encode('utf-8')
            source_length = len(source_x)
        length += 4
        length += source_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"time_ref.sec": %s, '%time_ref.sec
        string_echo += '"time_ref.nsec": %s'%time_ref.nsec
        string_echo += ', '
        string_echo += '"source": "%s"'%source
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/TimeReference"

    def getMD5(self):
        return "8e1576e01de57cd0d55758112f0e84ec"

_struct_I = struct.Struct('<I')

