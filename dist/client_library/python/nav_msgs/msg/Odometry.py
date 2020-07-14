import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class Odometry(tinyros.Message):
    __slots__ = ['header','child_frame_id','pose','twist']
    _slot_types = ['std_msgs.msg.Header','string','geometry_msgs.msg.PoseWithCovariance','geometry_msgs.msg.TwistWithCovariance']

    def __init__(self):
        super(Odometry, self).__init__()
        self.header = std_msgs.msg.Header()
        self.child_frame_id = ''
        self.pose = geometry_msgs.msg.PoseWithCovariance()
        self.twist = geometry_msgs.msg.TwistWithCovariance()

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            _x = self.child_frame_id
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
        offset += self.pose.serialize(buff)
        offset += self.twist.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (length_child_frame_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.child_frame_id = buff[offset:(offset + length_child_frame_id)].decode('utf-8')
            else:
                self.child_frame_id = buff[offset:(offset + length_child_frame_id)]
            offset += length_child_frame_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.pose.deserialize(buff[offset:])
        offset += self.twist.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        child_frame_id_x = self.child_frame_id
        child_frame_id_length = len(child_frame_id_x)
        if python3 or type(child_frame_id_x) == unicode:
            child_frame_id_x = child_frame_id_x.encode('utf-8')
            child_frame_id_length = len(child_frame_id_x)
        length += 4
        length += child_frame_id_length
        length += self.pose.serializedLength()
        length += self.twist.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"child_frame_id": "%s"'%child_frame_id
        string_echo += '", '
        string_echo += '"pose": {"'
        string_echo += self.pose.echo()
        string_echo += '}, '
        string_echo += '"twist": {"'
        string_echo += self.twist.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/Odometry"

    def getMD5(self):
        return "8fbd8c2e0caeb7be9b30b66a3e735193"

_struct_I = struct.Struct('<I')

