import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import trajectory_msgs.msg
import std_msgs.msg

class JointTrajectory(tinyros.Message):
    __slots__ = ['header','joint_names','points']
    _slot_types = ['std_msgs.msg.Header','string[]','trajectory_msgs.msg.JointTrajectoryPoint[]']

    def __init__(self):
        super(JointTrajectory, self).__init__()
        self.header = std_msgs.msg.Header()
        self.joint_names = []
        self.points = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            joint_names_length = len(self.joint_names)
            buff.write(_struct_I.pack(joint_names_length))
            offset += 4
            for i in range(0, joint_names_length):
                try:
                    _x = self.joint_names[i]
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
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            points_length = len(self.points)
            buff.write(_struct_I.pack(points_length))
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (joint_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.joint_names = ['' for x in range(0, joint_names_length)]
            offset += 4
            for i in range(0, joint_names_length):
                try:
                    (length_joint_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.joint_names[i] = buff[offset:(offset + length_joint_namesi)].decode('utf-8')
                    else:
                        self.joint_names[i] = buff[offset:(offset + length_joint_namesi)]
                    offset += length_joint_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (points_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.points = [trajectory_msgs.msg.JointTrajectoryPoint() for x in range(0, points_length)]
            offset += 4
            for i in range(0, points_length):
                offset += self.points[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        joint_names_length = len(self.joint_names)
        length += 4
        for i in range(0, joint_names_length):
            joint_namesi_x = self.joint_names[i]
            joint_namesi_length = len(joint_namesi_x)
            if python3 or type(joint_namesi_x) == unicode:
                joint_namesi_x = joint_namesi_x.encode('utf-8')
                joint_namesi_length = len(joint_namesi_x)
            length += 4
            length += joint_namesi_length
        points_length = len(self.points)
        length += 4
        for i in range(0, points_length):
            length += self.points[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"joint_names": ['
        joint_names_length = len(self.joint_names)
        for i in range(0, joint_names_length):
            if i == (joint_names_length - 1): 
                string_echo += '"joint_names[i]": "%s"'%joint_names[i]
                string_echo += '"'
            else:
                string_echo += '"joint_names[i]": "%s"'%joint_names[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"points": ['
        points_length = len(self.points)
        for i in range(0, points_length):
            if i == (points_length - 1): 
                string_echo += '{"points%s": {'%i
                string_echo += self.points[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"points%s": {'%i
                string_echo += self.points[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "trajectory_msgs/JointTrajectory"

    def getMD5(self):
        return "33e07cd7f5a6df021dad1271b3770d66"

_struct_I = struct.Struct('<I')

