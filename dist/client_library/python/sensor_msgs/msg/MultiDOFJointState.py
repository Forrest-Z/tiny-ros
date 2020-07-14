import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class MultiDOFJointState(tinyros.Message):
    __slots__ = ['header','joint_names','transforms','twist','wrench']
    _slot_types = ['std_msgs.msg.Header','string[]','geometry_msgs.msg.Transform[]','geometry_msgs.msg.Twist[]','geometry_msgs.msg.Wrench[]']

    def __init__(self):
        super(MultiDOFJointState, self).__init__()
        self.header = std_msgs.msg.Header()
        self.joint_names = []
        self.transforms = []
        self.twist = []
        self.wrench = []

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
            transforms_length = len(self.transforms)
            buff.write(_struct_I.pack(transforms_length))
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].serialize(buff)
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
            (transforms_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.transforms = [geometry_msgs.msg.Transform() for x in range(0, transforms_length)]
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].deserialize(buff[offset:])
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
        transforms_length = len(self.transforms)
        length += 4
        for i in range(0, transforms_length):
            length += self.transforms[i].serializedLength()
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
        string_echo += '"transforms": ['
        transforms_length = len(self.transforms)
        for i in range(0, transforms_length):
            if i == (transforms_length - 1): 
                string_echo += '{"transforms%s": {'%i
                string_echo += self.transforms[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"transforms%s": {'%i
                string_echo += self.transforms[i].echo()
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
        return "sensor_msgs/MultiDOFJointState"

    def getMD5(self):
        return "c1b0232d8e5071b24db76eb143f286eb"

_struct_I = struct.Struct('<I')

