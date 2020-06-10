import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg

class GetMapAction(tinyros.Message):
    __slots__ = ['action_goal','action_result','action_feedback']
    _slot_types = ['nav_msgs.msg.GetMapActionGoal','nav_msgs.msg.GetMapActionResult','nav_msgs.msg.GetMapActionFeedback']

    def __init__(self):
        super(GetMapAction, self).__init__()
        self.action_goal = nav_msgs.msg.GetMapActionGoal()
        self.action_result = nav_msgs.msg.GetMapActionResult()
        self.action_feedback = nav_msgs.msg.GetMapActionFeedback()

    def serialize(self, buff):
        offset = 0
        offset += self.action_goal.serialize(buff)
        offset += self.action_result.serialize(buff)
        offset += self.action_feedback.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.action_goal.deserialize(buff[offset:])
        offset += self.action_result.deserialize(buff[offset:])
        offset += self.action_feedback.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.action_goal.serializedLength()
        length += self.action_result.serializedLength()
        length += self.action_feedback.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"action_goal": {"'
        string_echo += self.action_goal.echo()
        string_echo += '}, '
        string_echo += '"action_result": {"'
        string_echo += self.action_result.echo()
        string_echo += '}, '
        string_echo += '"action_feedback": {"'
        string_echo += self.action_feedback.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GetMapAction"

    def getMD5(self):
        return "10a4e277d7b8e53bfc3df54d98b3edb1"

_struct_I = struct.Struct('<I')

