import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import std_msgs.msg
import actionlib_msgs.msg

class GetMapActionGoal(tinyros.Message):
    __slots__ = ['header','goal_id','goal']
    _slot_types = ['std_msgs.msg.Header','actionlib_msgs.msg.GoalID','nav_msgs.msg.GetMapGoal']

    def __init__(self):
        super(GetMapActionGoal, self).__init__()
        self.header = std_msgs.msg.Header()
        self.goal_id = actionlib_msgs.msg.GoalID()
        self.goal = nav_msgs.msg.GetMapGoal()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.goal_id.serialize(buff)
        offset += self.goal.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.goal_id.deserialize(buff[offset:])
        offset += self.goal.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.goal_id.serializedLength()
        length += self.goal.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"goal_id": {"'
        string_echo += self.goal_id.echo()
        string_echo += '}, '
        string_echo += '"goal": {"'
        string_echo += self.goal.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GetMapActionGoal"

    def getMD5(self):
        return "8aea83336b4ee626241742bb14b14d90"

_struct_I = struct.Struct('<I')

