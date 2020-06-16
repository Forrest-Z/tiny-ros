import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import std_msgs.msg

class PoseArray(tinyros.Message):
    __slots__ = ['header','poses']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Pose[]']

    def __init__(self):
        super(PoseArray, self).__init__()
        self.header = std_msgs.msg.Header()
        self.poses = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            poses_length = len(self.poses)
            buff.write(_struct_I.pack(poses_length))
            offset += 4
            for i in range(0, poses_length):
                offset += self.poses[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (poses_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.poses = [geometry_msgs.msg.Pose() for x in range(0, poses_length)]
            offset += 4
            for i in range(0, poses_length):
                offset += self.poses[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        poses_length = len(self.poses)
        length += 4
        for i in range(0, poses_length):
            length += self.poses[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"poses": ['
        poses_length = len(self.poses)
        for i in range(0, poses_length):
            if i == (poses_length - 1): 
                string_echo += '{"poses%s": {'%i
                string_echo += self.poses[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"poses%s": {'%i
                string_echo += self.poses[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/PoseArray"

    def getMD5(self):
        return "184f43246f3bc9cb5d0613694e6641a6"

_struct_I = struct.Struct('<I')

