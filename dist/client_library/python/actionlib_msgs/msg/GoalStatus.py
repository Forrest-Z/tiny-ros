import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import actionlib_msgs.msg

class GoalStatus(tinyros.Message):
    __slots__ = ['goal_id','status','text']
    _slot_types = ['actionlib_msgs.msg.GoalID','uint8','string']

    PENDING =  0   
    ACTIVE =  1   
    PREEMPTED =  2   
    SUCCEEDED =  3   
    ABORTED =  4   
    REJECTED =  5   
    PREEMPTING =  6   
    RECALLING =  7   
    RECALLED =  8   
    LOST =  9   

    def __init__(self):
        super(GoalStatus, self).__init__()
        self.goal_id = actionlib_msgs.msg.GoalID()
        self.status = 0
        self.text = ''

    def serialize(self, buff):
        offset = 0
        offset += self.goal_id.serialize(buff)
        try:
            struct_status = struct.Struct("<B")
            buff.write(struct_status.pack(self.status))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.text
            length = len(_x)
            if python3 or type(_x) == unicode:
                _x = _x.encode('utf-8')
                length = len(_x)
            if python3:
                buff.write(struct.pack('<I%sB'%length, length, *_x))
            else:
                buff.write(struct.pack('<I%ss'%length, length, _x))
            offset += 4
            offset += length
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.goal_id.deserialize(buff[offset:])
        try:
            struct_status = struct.Struct("<B")
            (self.status,) = struct_status.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_text,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.text = buff[offset:(offset + length_text)].decode('utf-8')
            else:
                self.text = buff[offset:(offset + length_text)]
            offset += length_text
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.goal_id.serializedLength()
        length += 1
        text_x = self.text
        text_length = len(text_x)
        if python3 or type(text_x) == unicode:
            text_x = text_x.encode('utf-8')
            text_length = len(text_x)
        length += 4
        length += text_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"goal_id": {"'
        string_echo += self.goal_id.echo()
        string_echo += '}, '
        string_echo += '"status": %s'%status
        string_echo += ', '
        string_echo += '"text": "%s"'%text
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "actionlib_msgs/GoalStatus"

    def getMD5(self):
        return "086be35ea957e692de83fc3477e4ef0b"

_struct_I = struct.Struct('<I')

