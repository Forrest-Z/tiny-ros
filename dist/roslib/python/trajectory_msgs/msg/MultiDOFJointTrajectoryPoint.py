import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import trajectory_msgs.msg
import geometry_msgs.msg
import tinyros

class MultiDOFJointTrajectoryPoint(tinyros.Message):
    __slots__ = ['transforms','velocities','accelerations','time_from_start']
    _slot_types = ['geometry_msgs.msg.Transform[]','geometry_msgs.msg.Twist[]','geometry_msgs.msg.Twist[]','tinyros.Duration']

    def __init__(self):
        super(MultiDOFJointTrajectoryPoint, self).__init__()
        self.transforms = []
        self.velocities = []
        self.accelerations = []
        self.time_from_start = tinyros.Duration()

    def serialize(self, buff):
        offset = 0
        try:
            transforms_length = len(self.transforms)
            buff.write(_struct_I.pack(transforms_length))
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            velocities_length = len(self.velocities)
            buff.write(_struct_I.pack(velocities_length))
            offset += 4
            for i in range(0, velocities_length):
                offset += self.velocities[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            accelerations_length = len(self.accelerations)
            buff.write(_struct_I.pack(accelerations_length))
            offset += 4
            for i in range(0, accelerations_length):
                offset += self.accelerations[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.time_from_start.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.time_from_start.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (transforms_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.transforms = [geometry_msgs.msg.Transform() for x in range(0, transforms_length)]
            offset += 4
            for i in range(0, transforms_length):
                offset += self.transforms[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (velocities_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.velocities = [geometry_msgs.msg.Twist() for x in range(0, velocities_length)]
            offset += 4
            for i in range(0, velocities_length):
                offset += self.velocities[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (accelerations_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.accelerations = [geometry_msgs.msg.Twist() for x in range(0, accelerations_length)]
            offset += 4
            for i in range(0, accelerations_length):
                offset += self.accelerations[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.time_from_start.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.time_from_start.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        transforms_length = len(self.transforms)
        length += 4
        for i in range(0, transforms_length):
            length += self.transforms[i].serializedLength()
        velocities_length = len(self.velocities)
        length += 4
        for i in range(0, velocities_length):
            length += self.velocities[i].serializedLength()
        accelerations_length = len(self.accelerations)
        length += 4
        for i in range(0, accelerations_length):
            length += self.accelerations[i].serializedLength()
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
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
        string_echo += '"velocities": ['
        velocities_length = len(self.velocities)
        for i in range(0, velocities_length):
            if i == (velocities_length - 1): 
                string_echo += '{"velocities%s": {'%i
                string_echo += self.velocities[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"velocities%s": {'%i
                string_echo += self.velocities[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"accelerations": ['
        accelerations_length = len(self.accelerations)
        for i in range(0, accelerations_length):
            if i == (accelerations_length - 1): 
                string_echo += '{"accelerations%s": {'%i
                string_echo += self.accelerations[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"accelerations%s": {'%i
                string_echo += self.accelerations[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"time_from_start.sec": %s, '%time_from_start.sec
        string_echo += '"time_from_start.nsec": %s'%time_from_start.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "trajectory_msgs/MultiDOFJointTrajectoryPoint"

    def getMD5(self):
        return "f8b4a74b416279b5c5d565029308ff08"

_struct_I = struct.Struct('<I')

