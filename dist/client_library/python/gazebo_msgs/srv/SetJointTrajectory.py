import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import trajectory_msgs.msg
import gazebo_msgs.msg

class SetJointTrajectoryRequest(tinyros.Message):
    __slots__ = ['__id__','model_name','joint_trajectory','model_pose','set_model_pose','disable_physics_updates']
    _slot_types = ['uint32','string','trajectory_msgs.msg.JointTrajectory','geometry_msgs.msg.Pose','bool','bool']

    def __init__(self):
        super(SetJointTrajectoryRequest, self).__init__()
        self.__id__ = 0
        self.model_name = ''
        self.joint_trajectory = trajectory_msgs.msg.JointTrajectory()
        self.model_pose = geometry_msgs.msg.Pose()
        self.set_model_pose = False
        self.disable_physics_updates = False

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.model_name
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
        offset += self.joint_trajectory.serialize(buff)
        offset += self.model_pose.serialize(buff)
        try:
            struct_set_model_pose = struct.Struct("<B")
            buff.write(struct_set_model_pose.pack(self.set_model_pose))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_disable_physics_updates = struct.Struct("<B")
            buff.write(struct_disable_physics_updates.pack(self.disable_physics_updates))
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
            (length_model_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.model_name = buff[offset:(offset + length_model_name)].decode('utf-8')
            else:
                self.model_name = buff[offset:(offset + length_model_name)]
            offset += length_model_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.joint_trajectory.deserialize(buff[offset:])
        offset += self.model_pose.deserialize(buff[offset:])
        try:
            struct_set_model_pose = struct.Struct("<B")
            (self.set_model_pose,) = struct_set_model_pose.unpack(buff[offset:(offset + 1)])
            self.set_model_pose = bool(self.set_model_pose)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_disable_physics_updates = struct.Struct("<B")
            (self.disable_physics_updates,) = struct_disable_physics_updates.unpack(buff[offset:(offset + 1)])
            self.disable_physics_updates = bool(self.disable_physics_updates)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        model_name_x = self.model_name
        model_name_length = len(model_name_x)
        if python3 or type(model_name_x) == unicode:
            model_name_x = model_name_x.encode('utf-8')
            model_name_length = len(model_name_x)
        length += 4
        length += model_name_length
        length += self.joint_trajectory.serializedLength()
        length += self.model_pose.serializedLength()
        length += 1
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"model_name": "%s"'%model_name
        string_echo += '", '
        string_echo += '"joint_trajectory": {"'
        string_echo += self.joint_trajectory.echo()
        string_echo += '}, '
        string_echo += '"model_pose": {"'
        string_echo += self.model_pose.echo()
        string_echo += '}, '
        string_echo += '"set_model_pose": %s'%set_model_pose
        string_echo += ', '
        string_echo += '"disable_physics_updates": %s'%disable_physics_updates
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SetJointTrajectory"

    def getMD5(self):
        return "8230e1fcc0295d8d21fbd5df0ccb0af3"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetJointTrajectoryResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(SetJointTrajectoryResponse, self).__init__()
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
        return "gazebo_msgs/SetJointTrajectory"

    def getMD5(self):
        return "2f5fe47400272efd54556969ffe63e7e"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetJointTrajectory(object):
    Request = SetJointTrajectoryRequest
    Response = SetJointTrajectoryResponse

_struct_I = struct.Struct('<I')

