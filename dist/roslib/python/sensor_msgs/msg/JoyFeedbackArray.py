import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class JoyFeedbackArray(tinyros.Message):
    __slots__ = ['array']
    _slot_types = ['sensor_msgs.msg.JoyFeedback[]']

    def __init__(self):
        super(JoyFeedbackArray, self).__init__()
        self.array = []

    def serialize(self, buff):
        offset = 0
        try:
            array_length = len(self.array)
            buff.write(_struct_I.pack(array_length))
            offset += 4
            for i in range(0, array_length):
                offset += self.array[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (array_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.array = [sensor_msgs.msg.JoyFeedback() for x in range(0, array_length)]
            offset += 4
            for i in range(0, array_length):
                offset += self.array[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        array_length = len(self.array)
        length += 4
        for i in range(0, array_length):
            length += self.array[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"array": ['
        array_length = len(self.array)
        for i in range(0, array_length):
            if i == (array_length - 1): 
                string_echo += '{"array%s": {'%i
                string_echo += self.array[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"array%s": {'%i
                string_echo += self.array[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/JoyFeedbackArray"

    def getMD5(self):
        return "45361e458d526d5670706a9f083819b6"

_struct_I = struct.Struct('<I')

