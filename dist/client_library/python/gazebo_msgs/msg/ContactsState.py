import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg
import std_msgs.msg

class ContactsState(tinyros.Message):
    __slots__ = ['header','states']
    _slot_types = ['std_msgs.msg.Header','gazebo_msgs.msg.ContactState[]']

    def __init__(self):
        super(ContactsState, self).__init__()
        self.header = std_msgs.msg.Header()
        self.states = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            states_length = len(self.states)
            buff.write(_struct_I.pack(states_length))
            offset += 4
            for i in range(0, states_length):
                offset += self.states[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (states_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.states = [gazebo_msgs.msg.ContactState() for x in range(0, states_length)]
            offset += 4
            for i in range(0, states_length):
                offset += self.states[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        states_length = len(self.states)
        length += 4
        for i in range(0, states_length):
            length += self.states[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"states": ['
        states_length = len(self.states)
        for i in range(0, states_length):
            if i == (states_length - 1): 
                string_echo += '{"states%s": {'%i
                string_echo += self.states[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"states%s": {'%i
                string_echo += self.states[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ContactsState"

    def getMD5(self):
        return "d19cd2a086cbd43da4252eb8d5cc64f5"

_struct_I = struct.Struct('<I')

