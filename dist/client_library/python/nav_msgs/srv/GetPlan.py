import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import geometry_msgs.msg

class GetPlanRequest(tinyros.Message):
    __slots__ = ['__id__','start','goal','tolerance']
    _slot_types = ['uint32','geometry_msgs.msg.PoseStamped','geometry_msgs.msg.PoseStamped','float32']

    def __init__(self):
        super(GetPlanRequest, self).__init__()
        self.__id__ = 0
        self.start = geometry_msgs.msg.PoseStamped()
        self.goal = geometry_msgs.msg.PoseStamped()
        self.tolerance = 0.

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.start.serialize(buff)
        offset += self.goal.serialize(buff)
        try:
            struct_tolerance = struct.Struct("<f")
            buff.write(struct_tolerance.pack(self.tolerance))
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
        offset += self.start.deserialize(buff[offset:])
        offset += self.goal.deserialize(buff[offset:])
        try:
            struct_tolerance = struct.Struct("<f")
            (self.tolerance,) = struct_tolerance.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.start.serializedLength()
        length += self.goal.serializedLength()
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"start": {"'
        string_echo += self.start.echo()
        string_echo += '}, '
        string_echo += '"goal": {"'
        string_echo += self.goal.echo()
        string_echo += '}, '
        string_echo += '"tolerance": %s'%tolerance
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GetPlan"

    def getMD5(self):
        return "557d5ea947f7761284cf7abef1cd7227"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPlanResponse(tinyros.Message):
    __slots__ = ['__id__','plan']
    _slot_types = ['uint32','nav_msgs.msg.Path']

    def __init__(self):
        super(GetPlanResponse, self).__init__()
        self.__id__ = 0
        self.plan = nav_msgs.msg.Path()

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.plan.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.plan.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        length += self.plan.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"plan": {"'
        string_echo += self.plan.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GetPlan"

    def getMD5(self):
        return "67c62b8c931eabfe35c88aed4b8f1258"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPlan(object):
    Request = GetPlanRequest
    Response = GetPlanResponse

_struct_I = struct.Struct('<I')

