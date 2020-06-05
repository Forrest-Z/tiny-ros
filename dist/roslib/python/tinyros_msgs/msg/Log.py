import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tinyros_msgs.msg

class Log(tinyros.Message):
    __slots__ = ['level','msg']
    _slot_types = ['uint8','string']

    ROSDEBUG = 0
    ROSINFO = 1
    ROSWARN = 2
    ROSERROR = 3
    ROSFATAL = 4

    def __init__(self):
        super(Log, self).__init__()
        self.level = 0
        self.msg = ''

    def serialize(self, buff):
        offset = 0
        try:
            struct_level = struct.Struct("<B")
            buff.write(struct_level.pack(self.level))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.msg
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
            struct_level = struct.Struct("<B")
            (self.level,) = struct_level.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_msg,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.msg = buff[offset:(offset + length_msg)].decode('utf-8')
            else:
                self.msg = buff[offset:(offset + length_msg)]
            offset += length_msg
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        msg_x = self.msg
        msg_length = len(msg_x)
        if python3 or type(msg_x) == unicode:
            msg_x = msg_x.encode('utf-8')
            msg_length = len(msg_x)
        length += 4
        length += msg_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"level": %s'%level
        string_echo += ', '
        string_echo += '"msg": "%s"'%msg
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_msgs/Log"

    def getMD5(self):
        return "0bd74339b4d77cb15766d831a3d15eeb"

_struct_I = struct.Struct('<I')

