import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import geometry_msgs.msg

class SetMapRequest(tinyros.Message):
    __slots__ = ['__id__','map','initial_pose']
    _slot_types = ['uint32','nav_msgs.msg.OccupancyGrid','geometry_msgs.msg.PoseWithCovarianceStamped']

    def __init__(self):
        super(SetMapRequest, self).__init__()
        self.__id__ = 0
        self.map = nav_msgs.msg.OccupancyGrid()
        self.initial_pose = geometry_msgs.msg.PoseWithCovarianceStamped()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.map.serialize(buff)
        offset += self.initial_pose.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.map.deserialize(buff[offset:])
        offset += self.initial_pose.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.map.serializedLength()
        length += self.initial_pose.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"map": {"'
        string_echo += self.map.echo()
        string_echo += '}, '
        string_echo += '"initial_pose": {"'
        string_echo += self.initial_pose.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/SetMap"

    def getMD5(self):
        return "946e1bd68c9db117a530a571e33d9e49"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetMapResponse(tinyros.Message):
    __slots__ = ['__id__','success']
    _slot_types = ['uint32','bool']

    def __init__(self):
        super(SetMapResponse, self).__init__()
        self.__id__ = 0
        self.success = False

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
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"success": %s'%success
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/SetMap"

    def getMD5(self):
        return "1e32607e79013262dafbbac9044e9cda"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetMap(object):
    Request = SetMapRequest
    Response = SetMapResponse

_struct_I = struct.Struct('<I')

