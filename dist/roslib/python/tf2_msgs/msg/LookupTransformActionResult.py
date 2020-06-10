import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg
import std_msgs.msg
import actionlib_msgs.msg

class LookupTransformActionResult(tinyros.Message):
    __slots__ = ['header','status','result']
    _slot_types = ['std_msgs.msg.Header','actionlib_msgs.msg.GoalStatus','tf2_msgs.msg.LookupTransformResult']

    def __init__(self):
        super(LookupTransformActionResult, self).__init__()
        self.header = std_msgs.msg.Header()
        self.status = actionlib_msgs.msg.GoalStatus()
        self.result = tf2_msgs.msg.LookupTransformResult()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.status.serialize(buff)
        offset += self.result.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.status.deserialize(buff[offset:])
        offset += self.result.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.status.serializedLength()
        length += self.result.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"status": {"'
        string_echo += self.status.echo()
        string_echo += '}, '
        string_echo += '"result": {"'
        string_echo += self.result.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf2_msgs/LookupTransformActionResult"

    def getMD5(self):
        return "5a8abe079c2126ea9966563c5cae6d29"

_struct_I = struct.Struct('<I')

