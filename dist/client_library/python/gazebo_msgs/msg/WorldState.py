import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class WorldState(tinyros.Message):
    __slots__ = ['header','name','pose','twist','wrench']
    _slot_types = ['std_msgs.msg.Header','string[]','geometry_msgs.msg.Pose[]','geometry_msgs.msg.Twist[]','geometry_msgs.msg.Wrench[]']

    def __init__(self):
        super(WorldState, self).__init__()
        self.header = std_msgs.msg.Header()
        self.name = []
        self.pose = []
        self.twist = []
        self.wrench = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            name_length = len(self.name)
            buff.write(_struct_I.pack(name_length))
            offset += 4
            for i in range(0, name_length):
                try:
                    _x = self.name[i]
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
            pose_length = len(self.pose)
            buff.write(_struct_I.pack(pose_length))
            offset += 4
            for i in range(0, pose_length):
                offset += self.pose[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            twist_length = len(self.twist)
            buff.write(_struct_I.pack(twist_length))
            offset += 4
            for i in range(0, twist_length):
                offset += self.twist[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            wrench_length = len(self.wrench)
            buff.write(_struct_I.pack(wrench_length))
            offset += 4
            for i in range(0, wrench_length):
                offset += self.wrench[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (name_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.name = ['' for x in range(0, name_length)]
            offset += 4
            for i in range(0, name_length):
                try:
                    (length_namei,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.name[i] = buff[offset:(offset + length_namei)].decode('utf-8')
                    else:
                        self.name[i] = buff[offset:(offset + length_namei)]
                    offset += length_namei
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (pose_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.pose = [geometry_msgs.msg.Pose() for x in range(0, pose_length)]
            offset += 4
            for i in range(0, pose_length):
                offset += self.pose[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (twist_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.twist = [geometry_msgs.msg.Twist() for x in range(0, twist_length)]
            offset += 4
            for i in range(0, twist_length):
                offset += self.twist[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (wrench_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.wrench = [geometry_msgs.msg.Wrench() for x in range(0, wrench_length)]
            offset += 4
            for i in range(0, wrench_length):
                offset += self.wrench[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        name_length = len(self.name)
        length += 4
        for i in range(0, name_length):
            namei_x = self.name[i]
            namei_length = len(namei_x)
            if python3 or type(namei_x) == unicode:
                namei_x = namei_x.encode('utf-8')
                namei_length = len(namei_x)
            length += 4
            length += namei_length
        pose_length = len(self.pose)
        length += 4
        for i in range(0, pose_length):
            length += self.pose[i].serializedLength()
        twist_length = len(self.twist)
        length += 4
        for i in range(0, twist_length):
            length += self.twist[i].serializedLength()
        wrench_length = len(self.wrench)
        length += 4
        for i in range(0, wrench_length):
            length += self.wrench[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"name": ['
        name_length = len(self.name)
        for i in range(0, name_length):
            if i == (name_length - 1): 
                string_echo += '"name[i]": "%s"'%name[i]
                string_echo += '"'
            else:
                string_echo += '"name[i]": "%s"'%name[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"pose": ['
        pose_length = len(self.pose)
        for i in range(0, pose_length):
            if i == (pose_length - 1): 
                string_echo += '{"pose%s": {'%i
                string_echo += self.pose[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"pose%s": {'%i
                string_echo += self.pose[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"twist": ['
        twist_length = len(self.twist)
        for i in range(0, twist_length):
            if i == (twist_length - 1): 
                string_echo += '{"twist%s": {'%i
                string_echo += self.twist[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"twist%s": {'%i
                string_echo += self.twist[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"wrench": ['
        wrench_length = len(self.wrench)
        for i in range(0, wrench_length):
            if i == (wrench_length - 1): 
                string_echo += '{"wrench%s": {'%i
                string_echo += self.wrench[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"wrench%s": {'%i
                string_echo += self.wrench[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/WorldState"

    def getMD5(self):
        return "0e1997127271c4d021f99645c28f1c09"

_struct_I = struct.Struct('<I')

