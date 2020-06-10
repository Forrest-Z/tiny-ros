import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg

class LookupTransformAction(tinyros.Message):
    __slots__ = ['action_goal','action_result','action_feedback']
    _slot_types = ['tf2_msgs.msg.LookupTransformActionGoal','tf2_msgs.msg.LookupTransformActionResult','tf2_msgs.msg.LookupTransformActionFeedback']

    def __init__(self):
        super(LookupTransformAction, self).__init__()
        self.action_goal = tf2_msgs.msg.LookupTransformActionGoal()
        self.action_result = tf2_msgs.msg.LookupTransformActionResult()
        self.action_feedback = tf2_msgs.msg.LookupTransformActionFeedback()

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
        return "tf2_msgs/LookupTransformAction"

    def getMD5(self):
        return "49655e848adf6c64870a8eb763b94484"

_struct_I = struct.Struct('<I')

