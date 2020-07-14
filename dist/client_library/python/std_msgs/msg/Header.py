import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import std_msgs.msg
import tinyros

class Header(tinyros.Message):
    __slots__ = ['seq','stamp','frame_id']
    _slot_types = ['uint32','tinyros.Time','string']

    def __init__(self):
        super(Header, self).__init__()
        self.seq = 0
        self.stamp = tinyros.Time()
        self.frame_id = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.seq))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
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
            _x = self.frame_id
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
            (self.seq,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
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
            (length_frame_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.frame_id = buff[offset:(offset + length_frame_id)].decode('utf-8')
            else:
                self.frame_id = buff[offset:(offset + length_frame_id)]
            offset += length_frame_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 4
        length += 4
        frame_id_x = self.frame_id
        frame_id_length = len(frame_id_x)
        if python3 or type(frame_id_x) == unicode:
            frame_id_x = frame_id_x.encode('utf-8')
            frame_id_length = len(frame_id_x)
        length += 4
        length += frame_id_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"seq": %s'%seq
        string_echo += ', '
        string_echo += '"stamp.sec": %s, '%stamp.sec
        string_echo += '"stamp.nsec": %s'%stamp.nsec
        string_echo += ', '
        string_echo += '"frame_id": "%s"'%frame_id
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "std_msgs/Header"

    def getMD5(self):
        return "d33440e88be7b5b8255fc61ebbca06ad"

_struct_I = struct.Struct('<I')

