import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class SetJointPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__','joint_name','ode_joint_config']
    _slot_types = ['uint32','string','gazebo_msgs.msg.ODEJointProperties']

    def __init__(self):
        super(SetJointPropertiesRequest, self).__init__()
        self.__id__ = 0
        self.joint_name = ''
        self.ode_joint_config = gazebo_msgs.msg.ODEJointProperties()

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
        offset += self.ode_joint_config.serialize(buff)
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
        offset += self.ode_joint_config.deserialize(buff[offset:])
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
        length += self.ode_joint_config.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"joint_name": "%s"'%joint_name
        string_echo += '", '
        string_echo += '"ode_joint_config": {"'
        string_echo += self.ode_joint_config.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SetJointProperties"

    def getMD5(self):
        return "e9063603bda4e7bdd2b5530ad7705661"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetJointPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(SetJointPropertiesResponse, self).__init__()
        self.__id__ = 0
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
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SetJointProperties"

    def getMD5(self):
        return "7e8650b70fd2dfc24b249dddf8f45cee"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetJointProperties(object):
    Request = SetJointPropertiesRequest
    Response = SetJointPropertiesResponse

_struct_I = struct.Struct('<I')

