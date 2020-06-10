import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import gazebo_msgs.msg

class SetLinkPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__','link_name','com','gravity_mode','mass','ixx','ixy','ixz','iyy','iyz','izz']
    _slot_types = ['uint32','string','geometry_msgs.msg.Pose','bool','float64','float64','float64','float64','float64','float64','float64']

    def __init__(self):
        super(SetLinkPropertiesRequest, self).__init__()
        self.__id__ = 0
        self.link_name = ''
        self.com = geometry_msgs.msg.Pose()
        self.gravity_mode = False
        self.mass = 0.
        self.ixx = 0.
        self.ixy = 0.
        self.ixz = 0.
        self.iyy = 0.
        self.iyz = 0.
        self.izz = 0.

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.link_name
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
        offset += self.com.serialize(buff)
        try:
            struct_gravity_mode = struct.Struct("<B")
            buff.write(struct_gravity_mode.pack(self.gravity_mode))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_mass = struct.Struct("<d")
            buff.write(struct_mass.pack(self.mass))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_ixx = struct.Struct("<d")
            buff.write(struct_ixx.pack(self.ixx))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_ixy = struct.Struct("<d")
            buff.write(struct_ixy.pack(self.ixy))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_ixz = struct.Struct("<d")
            buff.write(struct_ixz.pack(self.ixz))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_iyy = struct.Struct("<d")
            buff.write(struct_iyy.pack(self.iyy))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_iyz = struct.Struct("<d")
            buff.write(struct_iyz.pack(self.iyz))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_izz = struct.Struct("<d")
            buff.write(struct_izz.pack(self.izz))
            offset += 8
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
            (length_link_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.link_name = buff[offset:(offset + length_link_name)].decode('utf-8')
            else:
                self.link_name = buff[offset:(offset + length_link_name)]
            offset += length_link_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.com.deserialize(buff[offset:])
        try:
            struct_gravity_mode = struct.Struct("<B")
            (self.gravity_mode,) = struct_gravity_mode.unpack(buff[offset:(offset + 1)])
            self.gravity_mode = bool(self.gravity_mode)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_mass = struct.Struct("<d")
            (self.mass,) = struct_mass.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_ixx = struct.Struct("<d")
            (self.ixx,) = struct_ixx.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_ixy = struct.Struct("<d")
            (self.ixy,) = struct_ixy.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_ixz = struct.Struct("<d")
            (self.ixz,) = struct_ixz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_iyy = struct.Struct("<d")
            (self.iyy,) = struct_iyy.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_iyz = struct.Struct("<d")
            (self.iyz,) = struct_iyz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_izz = struct.Struct("<d")
            (self.izz,) = struct_izz.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        link_name_x = self.link_name
        link_name_length = len(link_name_x)
        if python3 or type(link_name_x) == unicode:
            link_name_x = link_name_x.encode('utf-8')
            link_name_length = len(link_name_x)
        length += 4
        length += link_name_length
        length += self.com.serializedLength()
        length += 1
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"link_name": "%s"'%link_name
        string_echo += '", '
        string_echo += '"com": {"'
        string_echo += self.com.echo()
        string_echo += '}, '
        string_echo += '"gravity_mode": %s'%gravity_mode
        string_echo += ', '
        string_echo += '"mass": %s'%mass
        string_echo += ', '
        string_echo += '"ixx": %s'%ixx
        string_echo += ', '
        string_echo += '"ixy": %s'%ixy
        string_echo += ', '
        string_echo += '"ixz": %s'%ixz
        string_echo += ', '
        string_echo += '"iyy": %s'%iyy
        string_echo += ', '
        string_echo += '"iyz": %s'%iyz
        string_echo += ', '
        string_echo += '"izz": %s'%izz
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SetLinkProperties"

    def getMD5(self):
        return "03d7e308476601832a9e1ce9d7eab722"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetLinkPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(SetLinkPropertiesResponse, self).__init__()
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
        return "gazebo_msgs/SetLinkProperties"

    def getMD5(self):
        return "777dea05607e1bca37e3206f97801d89"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetLinkProperties(object):
    Request = SetLinkPropertiesRequest
    Response = SetLinkPropertiesResponse

_struct_I = struct.Struct('<I')

