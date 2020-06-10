import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class JointRequestRequest(tinyros.Message):
    __slots__ = ['__id__','joint_name']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(JointRequestRequest, self).__init__()
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
        return "gazebo_msgs/JointRequest"

    def getMD5(self):
        return "e0bdc37edb92be07f3069573364a169f"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class JointRequestResponse(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(JointRequestResponse, self).__init__()
        self.__id__ = 0

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
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
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/JointRequest"

    def getMD5(self):
        return "ac5876a62df51a76c2828bb62894779d"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class JointRequest(object):
    Request = JointRequestRequest
    Response = JointRequestResponse

_struct_I = struct.Struct('<I')

