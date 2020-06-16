import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class JointState(tinyros.Message):
    __slots__ = ['header','name','position','velocity','effort']
    _slot_types = ['std_msgs.msg.Header','string[]','float64[]','float64[]','float64[]']

    def __init__(self):
        super(JointState, self).__init__()
        self.header = std_msgs.msg.Header()
        self.name = []
        self.position = []
        self.velocity = []
        self.effort = []

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
            position_length = len(self.position)
            buff.write(_struct_I.pack(position_length))
            offset += 4
            for i in range(0, position_length):
                try:
                    struct_positioni = struct.Struct("<d")
                    buff.write(struct_positioni.pack(self.position[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            velocity_length = len(self.velocity)
            buff.write(_struct_I.pack(velocity_length))
            offset += 4
            for i in range(0, velocity_length):
                try:
                    struct_velocityi = struct.Struct("<d")
                    buff.write(struct_velocityi.pack(self.velocity[i]))
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
            (position_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.position = [0. for x in range(0, position_length)]
            offset += 4
            for i in range(0, position_length):
                try:
                    struct_positioni = struct.Struct("<d")
                    (self.position[i],) = struct_positioni.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (velocity_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.velocity = [0. for x in range(0, velocity_length)]
            offset += 4
            for i in range(0, velocity_length):
                try:
                    struct_velocityi = struct.Struct("<d")
                    (self.velocity[i],) = struct_velocityi.unpack(buff[offset:(offset + 8)])
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
        position_length = len(self.position)
        length += 4
        for i in range(0, position_length):
            length += 8
        velocity_length = len(self.velocity)
        length += 4
        for i in range(0, velocity_length):
            length += 8
        effort_length = len(self.effort)
        length += 4
        for i in range(0, effort_length):
            length += 8
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
        string_echo += '"position": ['
        position_length = len(self.position)
        for i in range(0, position_length):
            if i == (position_length - 1): 
                string_echo += '{"position%s": %s'%(i,position[i])
                string_echo += '}'
            else:
                string_echo += '{"position%s": %s'%(i,position[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"velocity": ['
        velocity_length = len(self.velocity)
        for i in range(0, velocity_length):
            if i == (velocity_length - 1): 
                string_echo += '{"velocity%s": %s'%(i,velocity[i])
                string_echo += '}'
            else:
                string_echo += '{"velocity%s": %s'%(i,velocity[i])
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
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/JointState"

    def getMD5(self):
        return "6df7130a6d6a4c2f2037ce4a6e061fb9"

_struct_I = struct.Struct('<I')

