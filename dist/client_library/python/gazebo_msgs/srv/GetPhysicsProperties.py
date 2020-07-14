import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import gazebo_msgs.msg

class GetPhysicsPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(GetPhysicsPropertiesRequest, self).__init__()
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
        return "gazebo_msgs/GetPhysicsProperties"

    def getMD5(self):
        return "cba4b83a6824644ef787d2ac8bb7aa60"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPhysicsPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','time_step','pause','max_update_rate','gravity','ode_config','success','status_message']
    _slot_types = ['uint32','float64','bool','float64','geometry_msgs.msg.Vector3','gazebo_msgs.msg.ODEPhysics','bool','string']

    def __init__(self):
        super(GetPhysicsPropertiesResponse, self).__init__()
        self.__id__ = 0
        self.time_step = 0.
        self.pause = False
        self.max_update_rate = 0.
        self.gravity = geometry_msgs.msg.Vector3()
        self.ode_config = gazebo_msgs.msg.ODEPhysics()
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
            struct_time_step = struct.Struct("<d")
            buff.write(struct_time_step.pack(self.time_step))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_pause = struct.Struct("<B")
            buff.write(struct_pause.pack(self.pause))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_max_update_rate = struct.Struct("<d")
            buff.write(struct_max_update_rate.pack(self.max_update_rate))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.gravity.serialize(buff)
        offset += self.ode_config.serialize(buff)
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
            struct_time_step = struct.Struct("<d")
            (self.time_step,) = struct_time_step.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_pause = struct.Struct("<B")
            (self.pause,) = struct_pause.unpack(buff[offset:(offset + 1)])
            self.pause = bool(self.pause)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_max_update_rate = struct.Struct("<d")
            (self.max_update_rate,) = struct_max_update_rate.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.gravity.deserialize(buff[offset:])
        offset += self.ode_config.deserialize(buff[offset:])
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
        length += 8
        length += 1
        length += 8
        length += self.gravity.serializedLength()
        length += self.ode_config.serializedLength()
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
        string_echo += '"time_step": %s'%time_step
        string_echo += ', '
        string_echo += '"pause": %s'%pause
        string_echo += ', '
        string_echo += '"max_update_rate": %s'%max_update_rate
        string_echo += ', '
        string_echo += '"gravity": {"'
        string_echo += self.gravity.echo()
        string_echo += '}, '
        string_echo += '"ode_config": {"'
        string_echo += self.ode_config.echo()
        string_echo += '}, '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetPhysicsProperties"

    def getMD5(self):
        return "8f17f751935ff418006bf6b24982cb08"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetPhysicsProperties(object):
    Request = GetPhysicsPropertiesRequest
    Response = GetPhysicsPropertiesResponse

_struct_I = struct.Struct('<I')

