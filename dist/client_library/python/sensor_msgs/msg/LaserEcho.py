import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class LaserEcho(tinyros.Message):
    __slots__ = ['echoes']
    _slot_types = ['float32[]']

    def __init__(self):
        super(LaserEcho, self).__init__()
        self.echoes = []

    def serialize(self, buff):
        offset = 0
        try:
            echoes_length = len(self.echoes)
            buff.write(_struct_I.pack(echoes_length))
            offset += 4
            for i in range(0, echoes_length):
                try:
                    struct_echoesi = struct.Struct("<f")
                    buff.write(struct_echoesi.pack(self.echoes[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (echoes_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.echoes = [0. for x in range(0, echoes_length)]
            offset += 4
            for i in range(0, echoes_length):
                try:
                    struct_echoesi = struct.Struct("<f")
                    (self.echoes[i],) = struct_echoesi.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        echoes_length = len(self.echoes)
        length += 4
        for i in range(0, echoes_length):
            length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"echoes": ['
        echoes_length = len(self.echoes)
        for i in range(0, echoes_length):
            if i == (echoes_length - 1): 
                string_echo += '{"echoes%s": %s'%(i,echoes[i])
                string_echo += '}'
            else:
                string_echo += '{"echoes%s": %s'%(i,echoes[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/LaserEcho"

    def getMD5(self):
        return "a8537b388573845b3240b44db5bc4905"

_struct_I = struct.Struct('<I')

