import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import std_msgs.msg
import actionlib_msgs.msg

class GetMapActionResult(tinyros.Message):
    __slots__ = ['header','status','result']
    _slot_types = ['std_msgs.msg.Header','actionlib_msgs.msg.GoalStatus','nav_msgs.msg.GetMapResult']

    def __init__(self):
        super(GetMapActionResult, self).__init__()
        self.header = std_msgs.msg.Header()
        self.status = actionlib_msgs.msg.GoalStatus()
        self.result = nav_msgs.msg.GetMapResult()

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
        return "nav_msgs/GetMapActionResult"

    def getMD5(self):
        return "9c9f64758f2627a010c16b17ea745028"

_struct_I = struct.Struct('<I')

