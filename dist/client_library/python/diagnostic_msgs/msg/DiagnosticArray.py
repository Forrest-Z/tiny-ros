import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import diagnostic_msgs.msg
import std_msgs.msg

class DiagnosticArray(tinyros.Message):
    __slots__ = ['header','status']
    _slot_types = ['std_msgs.msg.Header','diagnostic_msgs.msg.DiagnosticStatus[]']

    def __init__(self):
        super(DiagnosticArray, self).__init__()
        self.header = std_msgs.msg.Header()
        self.status = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            status_length = len(self.status)
            buff.write(_struct_I.pack(status_length))
            offset += 4
            for i in range(0, status_length):
                offset += self.status[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (status_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.status = [diagnostic_msgs.msg.DiagnosticStatus() for x in range(0, status_length)]
            offset += 4
            for i in range(0, status_length):
                offset += self.status[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        status_length = len(self.status)
        length += 4
        for i in range(0, status_length):
            length += self.status[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"status": ['
        status_length = len(self.status)
        for i in range(0, status_length):
            if i == (status_length - 1): 
                string_echo += '{"status%s": {'%i
                string_echo += self.status[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"status%s": {'%i
                string_echo += self.status[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "diagnostic_msgs/DiagnosticArray"

    def getMD5(self):
        return "79a87210f85eb6afbd600eb2ba49dd85"

_struct_I = struct.Struct('<I')

