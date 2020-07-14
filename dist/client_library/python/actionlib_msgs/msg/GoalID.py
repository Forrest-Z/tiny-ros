import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import actionlib_msgs.msg
import tinyros

class GoalID(tinyros.Message):
    __slots__ = ['stamp','id']
    _slot_types = ['tinyros.Time','string']

    def __init__(self):
        super(GoalID, self).__init__()
        self.stamp = tinyros.Time()
        self.id = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.stamp.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.id
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
            (self.stamp.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.id = buff[offset:(offset + length_id)].decode('utf-8')
            else:
                self.id = buff[offset:(offset + length_id)]
            offset += length_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        id_x = self.id
        id_length = len(id_x)
        if python3 or type(id_x) == unicode:
            id_x = id_x.encode('utf-8')
            id_length = len(id_x)
        length += 4
        length += id_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"stamp.sec": %s, '%stamp.sec
        string_echo += '"stamp.nsec": %s'%stamp.nsec
        string_echo += ', '
        string_echo += '"id": "%s"'%id
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "actionlib_msgs/GoalID"

    def getMD5(self):
        return "a6cee90e5a185f4cb050de49bc4fa1f4"

_struct_I = struct.Struct('<I')

