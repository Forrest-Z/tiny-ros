import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf.msg
import geometry_msgs.msg

class tfMessage(tinyros.Message):
    __slots__ = ['transforms']
    _slot_types = ['geometry_msgs.msg.TransformStamped[]']

    def __init__(self):
        super(tfMessage, self).__init__()
        self.transforms = []

    def serialize(self, buff):
        offset = 0
        try:
            transforms_length = len(self.transforms)
            buff.write(_struct_I.pack(transforms_length))
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (transforms_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.transforms = [geometry_msgs.msg.TransformStamped() for x in range(0, transforms_length)]
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        transforms_length = len(self.transforms)
        length += 4
        for i in range(0, transforms_length):
            length += self.transforms[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"transforms": ['
        transforms_length = len(self.transforms)
        for i in range(0, transforms_length):
            if i == (transforms_length - 1): 
                string_echo += '{"transforms%s": {'%i
                string_echo += self.transforms[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"transforms%s": {'%i
                string_echo += self.transforms[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf/tfMessage"

    def getMD5(self):
        return "0401e4f6583aa665321e471e02ec671b"

_struct_I = struct.Struct('<I')

