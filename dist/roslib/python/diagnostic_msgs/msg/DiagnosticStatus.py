import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import diagnostic_msgs.msg

class DiagnosticStatus(tinyros.Message):
    __slots__ = ['level','name','message','hardware_id','values']
    _slot_types = ['byte','string','string','string','diagnostic_msgs.msg.KeyValue[]']

    OK = 0
    WARN = 1
    ERROR = 2
    STALE = 3

    def __init__(self):
        super(DiagnosticStatus, self).__init__()
        self.level = 0
        self.name = ''
        self.message = ''
        self.hardware_id = ''
        self.values = []

    def serialize(self, buff):
        offset = 0
        try:
            struct_level = struct.Struct("<b")
            buff.write(struct_level.pack(self.level))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
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
            _x = self.message
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
            _x = self.hardware_id
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
                offset += self.values[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_level = struct.Struct("<b")
            (self.level,) = struct_level.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
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
            (length_message,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.message = buff[offset:(offset + length_message)].decode('utf-8')
            else:
                self.message = buff[offset:(offset + length_message)]
            offset += length_message
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_hardware_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.hardware_id = buff[offset:(offset + length_hardware_id)].decode('utf-8')
            else:
                self.hardware_id = buff[offset:(offset + length_hardware_id)]
            offset += length_hardware_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (values_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.values = [diagnostic_msgs.msg.KeyValue() for x in range(0, values_length)]
            offset += 4
            for i in range(0, values_length):
                offset += self.values[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        name_x = self.name
        name_length = len(name_x)
        if python3 or type(name_x) == unicode:
            name_x = name_x.encode('utf-8')
            name_length = len(name_x)
        length += 4
        length += name_length
        message_x = self.message
        message_length = len(message_x)
        if python3 or type(message_x) == unicode:
            message_x = message_x.encode('utf-8')
            message_length = len(message_x)
        length += 4
        length += message_length
        hardware_id_x = self.hardware_id
        hardware_id_length = len(hardware_id_x)
        if python3 or type(hardware_id_x) == unicode:
            hardware_id_x = hardware_id_x.encode('utf-8')
            hardware_id_length = len(hardware_id_x)
        length += 4
        length += hardware_id_length
        values_length = len(self.values)
        length += 4
        for i in range(0, values_length):
            length += self.values[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"level": %s'%level
        string_echo += ', '
        string_echo += '"name": "%s"'%name
        string_echo += '", '
        string_echo += '"message": "%s"'%message
        string_echo += '", '
        string_echo += '"hardware_id": "%s"'%hardware_id
        string_echo += '", '
        string_echo += '"values": ['
        values_length = len(self.values)
        for i in range(0, values_length):
            if i == (values_length - 1): 
                string_echo += '{"values%s": {'%i
                string_echo += self.values[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"values%s": {'%i
                string_echo += self.values[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "diagnostic_msgs/DiagnosticStatus"

    def getMD5(self):
        return "9ec892d2145f478061efd60bb1762361"

_struct_I = struct.Struct('<I')

