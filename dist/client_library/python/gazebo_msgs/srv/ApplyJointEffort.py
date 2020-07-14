import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg
import tinyros

class ApplyJointEffortRequest(tinyros.Message):
    __slots__ = ['__id__','joint_name','effort','start_time','duration']
    _slot_types = ['uint32','string','float64','tinyros.Time','tinyros.Duration']

    def __init__(self):
        super(ApplyJointEffortRequest, self).__init__()
        self.__id__ = 0
        self.joint_name = ''
        self.effort = 0.
        self.start_time = tinyros.Time()
        self.duration = tinyros.Duration()

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
        try:
            struct_effort = struct.Struct("<d")
            buff.write(struct_effort.pack(self.effort))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.start_time.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.start_time.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.duration.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.duration.nsec))
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
        try:
            struct_effort = struct.Struct("<d")
            (self.effort,) = struct_effort.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.start_time.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.start_time.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.duration.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.duration.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
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
        length += 8
        length += 4
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"joint_name": "%s"'%joint_name
        string_echo += '", '
        string_echo += '"effort": %s'%effort
        string_echo += ', '
        string_echo += '"start_time.sec": %s, '%start_time.sec
        string_echo += '"start_time.nsec": %s'%start_time.nsec
        string_echo += ', '
        string_echo += '"duration.sec": %s, '%duration.sec
        string_echo += '"duration.nsec": %s'%duration.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ApplyJointEffort"

    def getMD5(self):
        return "90595a46cf1fb4ee17158e2f7034a0eb"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class ApplyJointEffortResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(ApplyJointEffortResponse, self).__init__()
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
        return "gazebo_msgs/ApplyJointEffort"

    def getMD5(self):
        return "953194fc8ca726693bef2876cebb0438"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class ApplyJointEffort(object):
    Request = ApplyJointEffortRequest
    Response = ApplyJointEffortResponse

_struct_I = struct.Struct('<I')

