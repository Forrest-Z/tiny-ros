import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import trajectory_msgs.msg
import tinyros

class JointTrajectoryPoint(tinyros.Message):
    __slots__ = ['positions','velocities','accelerations','effort','time_from_start']
    _slot_types = ['float64[]','float64[]','float64[]','float64[]','tinyros.Duration']

    def __init__(self):
        super(JointTrajectoryPoint, self).__init__()
        self.positions = []
        self.velocities = []
        self.accelerations = []
        self.effort = []
        self.time_from_start = tinyros.Duration()

    def serialize(self, buff):
        offset = 0
        try:
            positions_length = len(self.positions)
            buff.write(_struct_I.pack(positions_length))
            offset += 4
            for i in range(0, positions_length):
                try:
                    struct_positionsi = struct.Struct("<d")
                    buff.write(struct_positionsi.pack(self.positions[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            velocities_length = len(self.velocities)
            buff.write(_struct_I.pack(velocities_length))
            offset += 4
            for i in range(0, velocities_length):
                try:
                    struct_velocitiesi = struct.Struct("<d")
                    buff.write(struct_velocitiesi.pack(self.velocities[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            accelerations_length = len(self.accelerations)
            buff.write(_struct_I.pack(accelerations_length))
            offset += 4
            for i in range(0, accelerations_length):
                try:
                    struct_accelerationsi = struct.Struct("<d")
                    buff.write(struct_accelerationsi.pack(self.accelerations[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            effort_length = len(self.effort)
            buff.write(_struct_I.pack(effort_length))
            offset += 4
            for i in range(0, effort_length):
                try:
                    struct_efforti = struct.Struct("<d")
                    buff.write(struct_efforti.pack(self.effort[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
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
            (positions_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.positions = [0. for x in range(0, positions_length)]
            offset += 4
            for i in range(0, positions_length):
                try:
                    struct_positionsi = struct.Struct("<d")
                    (self.positions[i],) = struct_positionsi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (velocities_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.velocities = [0. for x in range(0, velocities_length)]
            offset += 4
            for i in range(0, velocities_length):
                try:
                    struct_velocitiesi = struct.Struct("<d")
                    (self.velocities[i],) = struct_velocitiesi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (accelerations_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.accelerations = [0. for x in range(0, accelerations_length)]
            offset += 4
            for i in range(0, accelerations_length):
                try:
                    struct_accelerationsi = struct.Struct("<d")
                    (self.accelerations[i],) = struct_accelerationsi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (effort_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.effort = [0. for x in range(0, effort_length)]
            offset += 4
            for i in range(0, effort_length):
                try:
                    struct_efforti = struct.Struct("<d")
                    (self.effort[i],) = struct_efforti.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
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
        positions_length = len(self.positions)
        length += 4
        for i in range(0, positions_length):
            length += 8
        velocities_length = len(self.velocities)
        length += 4
        for i in range(0, velocities_length):
            length += 8
        accelerations_length = len(self.accelerations)
        length += 4
        for i in range(0, accelerations_length):
            length += 8
        effort_length = len(self.effort)
        length += 4
        for i in range(0, effort_length):
            length += 8
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"positions": ['
        positions_length = len(self.positions)
        for i in range(0, positions_length):
            if i == (positions_length - 1): 
                string_echo += '{"positions%s": %s'%(i,positions[i])
                string_echo += '}'
            else:
                string_echo += '{"positions%s": %s'%(i,positions[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"velocities": ['
        velocities_length = len(self.velocities)
        for i in range(0, velocities_length):
            if i == (velocities_length - 1): 
                string_echo += '{"velocities%s": %s'%(i,velocities[i])
                string_echo += '}'
            else:
                string_echo += '{"velocities%s": %s'%(i,velocities[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"accelerations": ['
        accelerations_length = len(self.accelerations)
        for i in range(0, accelerations_length):
            if i == (accelerations_length - 1): 
                string_echo += '{"accelerations%s": %s'%(i,accelerations[i])
                string_echo += '}'
            else:
                string_echo += '{"accelerations%s": %s'%(i,accelerations[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"effort": ['
        effort_length = len(self.effort)
        for i in range(0, effort_length):
            if i == (effort_length - 1): 
                string_echo += '{"effort%s": %s'%(i,effort[i])
                string_echo += '}'
            else:
                string_echo += '{"effort%s": %s'%(i,effort[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"time_from_start.sec": %s, '%time_from_start.sec
        string_echo += '"time_from_start.nsec": %s'%time_from_start.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "trajectory_msgs/JointTrajectoryPoint"

    def getMD5(self):
        return "38679319321341510f6fde7d7f745eb0"

_struct_I = struct.Struct('<I')

