import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import actionlib_msgs.msg
import std_msgs.msg

class GoalStatusArray(tinyros.Message):
    __slots__ = ['header','status_list']
    _slot_types = ['std_msgs.msg.Header','actionlib_msgs.msg.GoalStatus[]']

    def __init__(self):
        super(GoalStatusArray, self).__init__()
        self.header = std_msgs.msg.Header()
        self.status_list = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            status_list_length = len(self.status_list)
            buff.write(_struct_I.pack(status_list_length))
            offset += 4
            for i in range(0, status_list_length):
                offset += self.status_list[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (status_list_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.status_list = [actionlib_msgs.msg.GoalStatus() for x in range(0, status_list_length)]
            offset += 4
            for i in range(0, status_list_length):
                offset += self.status_list[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        status_list_length = len(self.status_list)
        length += 4
        for i in range(0, status_list_length):
            length += self.status_list[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"status_list": ['
        status_list_length = len(self.status_list)
        for i in range(0, status_list_length):
            if i == (status_list_length - 1): 
                string_echo += '{"status_list%s": {'%i
                string_echo += self.status_list[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"status_list%s": {'%i
                string_echo += self.status_list[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "actionlib_msgs/GoalStatusArray"

    def getMD5(self):
        return "53f6501f7c14f5f3963638de4bbe3a71"

_struct_I = struct.Struct('<I')

