import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg
import tinyros

class LookupTransformGoal(tinyros.Message):
    __slots__ = ['target_frame','source_frame','source_time','timeout','target_time','fixed_frame','advanced']
    _slot_types = ['string','string','tinyros.Time','tinyros.Duration','tinyros.Time','string','bool']

    def __init__(self):
        super(LookupTransformGoal, self).__init__()
        self.target_frame = ''
        self.source_frame = ''
        self.source_time = tinyros.Time()
        self.timeout = tinyros.Duration()
        self.target_time = tinyros.Time()
        self.fixed_frame = ''
        self.advanced = False

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.target_frame
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
            _x = self.source_frame
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
            buff.write(_struct_I.pack(self.source_time.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.source_time.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.timeout.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.timeout.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.target_time.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.target_time.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.fixed_frame
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
            struct_advanced = struct.Struct("<B")
            buff.write(struct_advanced.pack(self.advanced))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_target_frame,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.target_frame = buff[offset:(offset + length_target_frame)].decode('utf-8')
            else:
                self.target_frame = buff[offset:(offset + length_target_frame)]
            offset += length_target_frame
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_source_frame,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.source_frame = buff[offset:(offset + length_source_frame)].decode('utf-8')
            else:
                self.source_frame = buff[offset:(offset + length_source_frame)]
            offset += length_source_frame
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.source_time.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.source_time.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.timeout.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.timeout.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.target_time.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.target_time.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_fixed_frame,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.fixed_frame = buff[offset:(offset + length_fixed_frame)].decode('utf-8')
            else:
                self.fixed_frame = buff[offset:(offset + length_fixed_frame)]
            offset += length_fixed_frame
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_advanced = struct.Struct("<B")
            (self.advanced,) = struct_advanced.unpack(buff[offset:(offset + 1)])
            self.advanced = bool(self.advanced)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        target_frame_x = self.target_frame
        target_frame_length = len(target_frame_x)
        if python3 or type(target_frame_x) == unicode:
            target_frame_x = target_frame_x.encode('utf-8')
            target_frame_length = len(target_frame_x)
        length += 4
        length += target_frame_length
        source_frame_x = self.source_frame
        source_frame_length = len(source_frame_x)
        if python3 or type(source_frame_x) == unicode:
            source_frame_x = source_frame_x.encode('utf-8')
            source_frame_length = len(source_frame_x)
        length += 4
        length += source_frame_length
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        fixed_frame_x = self.fixed_frame
        fixed_frame_length = len(fixed_frame_x)
        if python3 or type(fixed_frame_x) == unicode:
            fixed_frame_x = fixed_frame_x.encode('utf-8')
            fixed_frame_length = len(fixed_frame_x)
        length += 4
        length += fixed_frame_length
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"target_frame": "%s"'%target_frame
        string_echo += '", '
        string_echo += '"source_frame": "%s"'%source_frame
        string_echo += '", '
        string_echo += '"source_time.sec": %s, '%source_time.sec
        string_echo += '"source_time.nsec": %s'%source_time.nsec
        string_echo += ', '
        string_echo += '"timeout.sec": %s, '%timeout.sec
        string_echo += '"timeout.nsec": %s'%timeout.nsec
        string_echo += ', '
        string_echo += '"target_time.sec": %s, '%target_time.sec
        string_echo += '"target_time.nsec": %s'%target_time.nsec
        string_echo += ', '
        string_echo += '"fixed_frame": "%s"'%fixed_frame
        string_echo += '", '
        string_echo += '"advanced": %s'%advanced
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf2_msgs/LookupTransformGoal"

    def getMD5(self):
        return "677c84a912b788ccaaea5fbc38570182"

_struct_I = struct.Struct('<I')

