import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import gazebo_msgs.msg

class SpawnModelRequest(tinyros.Message):
    __slots__ = ['__id__','model_name','model_xml','robot_namespace','initial_pose','reference_frame']
    _slot_types = ['uint32','string','string','string','geometry_msgs.msg.Pose','string']

    def __init__(self):
        super(SpawnModelRequest, self).__init__()
        self.__id__ = 0
        self.model_name = ''
        self.model_xml = ''
        self.robot_namespace = ''
        self.initial_pose = geometry_msgs.msg.Pose()
        self.reference_frame = ''

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
        try:
            _x = self.model_xml
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
            _x = self.robot_namespace
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
        offset += self.initial_pose.serialize(buff)
        try:
            _x = self.reference_frame
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
            (length_model_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.model_name = buff[offset:(offset + length_model_name)].decode('utf-8')
            else:
                self.model_name = buff[offset:(offset + length_model_name)]
            offset += length_model_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_model_xml,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.model_xml = buff[offset:(offset + length_model_xml)].decode('utf-8')
            else:
                self.model_xml = buff[offset:(offset + length_model_xml)]
            offset += length_model_xml
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_robot_namespace,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.robot_namespace = buff[offset:(offset + length_robot_namespace)].decode('utf-8')
            else:
                self.robot_namespace = buff[offset:(offset + length_robot_namespace)]
            offset += length_robot_namespace
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.initial_pose.deserialize(buff[offset:])
        try:
            (length_reference_frame,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.reference_frame = buff[offset:(offset + length_reference_frame)].decode('utf-8')
            else:
                self.reference_frame = buff[offset:(offset + length_reference_frame)]
            offset += length_reference_frame
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
        model_xml_x = self.model_xml
        model_xml_length = len(model_xml_x)
        if python3 or type(model_xml_x) == unicode:
            model_xml_x = model_xml_x.encode('utf-8')
            model_xml_length = len(model_xml_x)
        length += 4
        length += model_xml_length
        robot_namespace_x = self.robot_namespace
        robot_namespace_length = len(robot_namespace_x)
        if python3 or type(robot_namespace_x) == unicode:
            robot_namespace_x = robot_namespace_x.encode('utf-8')
            robot_namespace_length = len(robot_namespace_x)
        length += 4
        length += robot_namespace_length
        length += self.initial_pose.serializedLength()
        reference_frame_x = self.reference_frame
        reference_frame_length = len(reference_frame_x)
        if python3 or type(reference_frame_x) == unicode:
            reference_frame_x = reference_frame_x.encode('utf-8')
            reference_frame_length = len(reference_frame_x)
        length += 4
        length += reference_frame_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"model_name": "%s"'%model_name
        string_echo += '", '
        string_echo += '"model_xml": "%s"'%model_xml
        string_echo += '", '
        string_echo += '"robot_namespace": "%s"'%robot_namespace
        string_echo += '", '
        string_echo += '"initial_pose": {"'
        string_echo += self.initial_pose.echo()
        string_echo += '}, '
        string_echo += '"reference_frame": "%s"'%reference_frame
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SpawnModel"

    def getMD5(self):
        return "da34e61c8813e52ac159e5f31fbf55be"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SpawnModelResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(SpawnModelResponse, self).__init__()
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
        return "gazebo_msgs/SpawnModel"

    def getMD5(self):
        return "d59d46cc4e5a64f978a429dd7c306d30"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SpawnModel(object):
    Request = SpawnModelRequest
    Response = SpawnModelResponse

_struct_I = struct.Struct('<I')

