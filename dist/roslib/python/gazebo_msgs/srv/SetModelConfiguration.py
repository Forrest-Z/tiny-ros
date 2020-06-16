import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class SetModelConfigurationRequest(tinyros.Message):
    __slots__ = ['__id__','model_name','urdf_param_name','joint_names','joint_positions']
    _slot_types = ['uint32','string','string','string[]','float64[]']

    def __init__(self):
        super(SetModelConfigurationRequest, self).__init__()
        self.__id__ = 0
        self.model_name = ''
        self.urdf_param_name = ''
        self.joint_names = []
        self.joint_positions = []

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
            _x = self.urdf_param_name
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
            joint_names_length = len(self.joint_names)
            buff.write(_struct_I.pack(joint_names_length))
            offset += 4
            for i in range(0, joint_names_length):
                try:
                    _x = self.joint_names[i]
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
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            joint_positions_length = len(self.joint_positions)
            buff.write(_struct_I.pack(joint_positions_length))
            offset += 4
            for i in range(0, joint_positions_length):
                try:
                    struct_joint_positionsi = struct.Struct("<d")
                    buff.write(struct_joint_positionsi.pack(self.joint_positions[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
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
            (length_urdf_param_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.urdf_param_name = buff[offset:(offset + length_urdf_param_name)].decode('utf-8')
            else:
                self.urdf_param_name = buff[offset:(offset + length_urdf_param_name)]
            offset += length_urdf_param_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (joint_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.joint_names = ['' for x in range(0, joint_names_length)]
            offset += 4
            for i in range(0, joint_names_length):
                try:
                    (length_joint_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.joint_names[i] = buff[offset:(offset + length_joint_namesi)].decode('utf-8')
                    else:
                        self.joint_names[i] = buff[offset:(offset + length_joint_namesi)]
                    offset += length_joint_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (joint_positions_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.joint_positions = [0. for x in range(0, joint_positions_length)]
            offset += 4
            for i in range(0, joint_positions_length):
                try:
                    struct_joint_positionsi = struct.Struct("<d")
                    (self.joint_positions[i],) = struct_joint_positionsi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
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
        urdf_param_name_x = self.urdf_param_name
        urdf_param_name_length = len(urdf_param_name_x)
        if python3 or type(urdf_param_name_x) == unicode:
            urdf_param_name_x = urdf_param_name_x.encode('utf-8')
            urdf_param_name_length = len(urdf_param_name_x)
        length += 4
        length += urdf_param_name_length
        joint_names_length = len(self.joint_names)
        length += 4
        for i in range(0, joint_names_length):
            joint_namesi_x = self.joint_names[i]
            joint_namesi_length = len(joint_namesi_x)
            if python3 or type(joint_namesi_x) == unicode:
                joint_namesi_x = joint_namesi_x.encode('utf-8')
                joint_namesi_length = len(joint_namesi_x)
            length += 4
            length += joint_namesi_length
        joint_positions_length = len(self.joint_positions)
        length += 4
        for i in range(0, joint_positions_length):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"model_name": "%s"'%model_name
        string_echo += '", '
        string_echo += '"urdf_param_name": "%s"'%urdf_param_name
        string_echo += '", '
        string_echo += '"joint_names": ['
        joint_names_length = len(self.joint_names)
        for i in range(0, joint_names_length):
            if i == (joint_names_length - 1): 
                string_echo += '"joint_names[i]": "%s"'%joint_names[i]
                string_echo += '"'
            else:
                string_echo += '"joint_names[i]": "%s"'%joint_names[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"joint_positions": ['
        joint_positions_length = len(self.joint_positions)
        for i in range(0, joint_positions_length):
            if i == (joint_positions_length - 1): 
                string_echo += '{"joint_positions%s": %s'%(i,joint_positions[i])
                string_echo += '}'
            else:
                string_echo += '{"joint_positions%s": %s'%(i,joint_positions[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/SetModelConfiguration"

    def getMD5(self):
        return "74db6184ae83468b540d4c02d244ada7"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetModelConfigurationResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(SetModelConfigurationResponse, self).__init__()
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
        return "gazebo_msgs/SetModelConfiguration"

    def getMD5(self):
        return "6f12aefa315c8b37040d5d47471e39ee"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class SetModelConfiguration(object):
    Request = SetModelConfigurationRequest
    Response = SetModelConfigurationResponse

_struct_I = struct.Struct('<I')

