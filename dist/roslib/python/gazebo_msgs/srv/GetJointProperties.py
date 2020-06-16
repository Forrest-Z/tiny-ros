import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class GetJointPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__','joint_name']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(GetJointPropertiesRequest, self).__init__()
        self.__id__ = 0
        self.joint_name = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.joint_name
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
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_joint_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.joint_name = buff[offset:(offset + length_joint_name)].decode('utf-8')
            else:
                self.joint_name = buff[offset:(offset + length_joint_name)]
            offset += length_joint_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        joint_name_x = self.joint_name
        joint_name_length = len(joint_name_x)
        if python3 or type(joint_name_x) == unicode:
            joint_name_x = joint_name_x.encode('utf-8')
            joint_name_length = len(joint_name_x)
        length += 4
        length += joint_name_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"joint_name": "%s"'%joint_name
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetJointProperties"

    def getMD5(self):
        return "b07a3ce5fb5aba1cfc56577c9cb3ecc6"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetJointPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','type','damping','position','rate','success','status_message']
    _slot_types = ['uint32','uint8','float64[]','float64[]','float64[]','bool','string']

    REVOLUTE =  0                
    CONTINUOUS =  1                
    PRISMATIC =  2                
    FIXED =  3                
    BALL =  4                
    UNIVERSAL =  5                

    def __init__(self):
        super(GetJointPropertiesResponse, self).__init__()
        self.__id__ = 0
        self.type = 0
        self.damping = []
        self.position = []
        self.rate = []
        self.success = False
        self.status_message = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_type = struct.Struct("<B")
            buff.write(struct_type.pack(self.type))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            damping_length = len(self.damping)
            buff.write(_struct_I.pack(damping_length))
            offset += 4
            for i in range(0, damping_length):
                try:
                    struct_dampingi = struct.Struct("<d")
                    buff.write(struct_dampingi.pack(self.damping[i]))
                    offset += 8
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
            rate_length = len(self.rate)
            buff.write(_struct_I.pack(rate_length))
            offset += 4
            for i in range(0, rate_length):
                try:
                    struct_ratei = struct.Struct("<d")
                    buff.write(struct_ratei.pack(self.rate[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_success = struct.Struct("<B")
            buff.write(struct_success.pack(self.success))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.status_message
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
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_type = struct.Struct("<B")
            (self.type,) = struct_type.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (damping_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.damping = [0. for x in range(0, damping_length)]
            offset += 4
            for i in range(0, damping_length):
                try:
                    struct_dampingi = struct.Struct("<d")
                    (self.damping[i],) = struct_dampingi.unpack(buff[offset:(offset + 8)])
                    offset += 8
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
            (rate_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.rate = [0. for x in range(0, rate_length)]
            offset += 4
            for i in range(0, rate_length):
                try:
                    struct_ratei = struct.Struct("<d")
                    (self.rate[i],) = struct_ratei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_success = struct.Struct("<B")
            (self.success,) = struct_success.unpack(buff[offset:(offset + 1)])
            self.success = bool(self.success)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_status_message,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.status_message = buff[offset:(offset + length_status_message)].decode('utf-8')
            else:
                self.status_message = buff[offset:(offset + length_status_message)]
            offset += length_status_message
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 1
        damping_length = len(self.damping)
        length += 4
        for i in range(0, damping_length):
            length += 8
        position_length = len(self.position)
        length += 4
        for i in range(0, position_length):
            length += 8
        rate_length = len(self.rate)
        length += 4
        for i in range(0, rate_length):
            length += 8
        length += 1
        status_message_x = self.status_message
        status_message_length = len(status_message_x)
        if python3 or type(status_message_x) == unicode:
            status_message_x = status_message_x.encode('utf-8')
            status_message_length = len(status_message_x)
        length += 4
        length += status_message_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"type": %s'%type
        string_echo += ', '
        string_echo += '"damping": ['
        damping_length = len(self.damping)
        for i in range(0, damping_length):
            if i == (damping_length - 1): 
                string_echo += '{"damping%s": %s'%(i,damping[i])
                string_echo += '}'
            else:
                string_echo += '{"damping%s": %s'%(i,damping[i])
                string_echo += '}, '
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
        string_echo += '"rate": ['
        rate_length = len(self.rate)
        for i in range(0, rate_length):
            if i == (rate_length - 1): 
                string_echo += '{"rate%s": %s'%(i,rate[i])
                string_echo += '}'
            else:
                string_echo += '{"rate%s": %s'%(i,rate[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetJointProperties"

    def getMD5(self):
        return "a60fbf691ac539e1355c979ca09b4573"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetJointProperties(object):
    Request = GetJointPropertiesRequest
    Response = GetJointPropertiesResponse

_struct_I = struct.Struct('<I')

