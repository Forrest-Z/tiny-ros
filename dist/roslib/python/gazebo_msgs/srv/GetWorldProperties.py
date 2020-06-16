import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class GetWorldPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(GetWorldPropertiesRequest, self).__init__()
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
        return "gazebo_msgs/GetWorldProperties"

    def getMD5(self):
        return "3aa5de7106eec5dae41ad1c9ae681123"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetWorldPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','sim_time','model_names','rendering_enabled','success','status_message']
    _slot_types = ['uint32','float64','string[]','bool','bool','string']

    def __init__(self):
        super(GetWorldPropertiesResponse, self).__init__()
        self.__id__ = 0
        self.sim_time = 0.
        self.model_names = []
        self.rendering_enabled = False
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
            struct_sim_time = struct.Struct("<d")
            buff.write(struct_sim_time.pack(self.sim_time))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            model_names_length = len(self.model_names)
            buff.write(_struct_I.pack(model_names_length))
            offset += 4
            for i in range(0, model_names_length):
                try:
                    _x = self.model_names[i]
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
            struct_rendering_enabled = struct.Struct("<B")
            buff.write(struct_rendering_enabled.pack(self.rendering_enabled))
            offset += 1
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
            struct_sim_time = struct.Struct("<d")
            (self.sim_time,) = struct_sim_time.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (model_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.model_names = ['' for x in range(0, model_names_length)]
            offset += 4
            for i in range(0, model_names_length):
                try:
                    (length_model_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.model_names[i] = buff[offset:(offset + length_model_namesi)].decode('utf-8')
                    else:
                        self.model_names[i] = buff[offset:(offset + length_model_namesi)]
                    offset += length_model_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_rendering_enabled = struct.Struct("<B")
            (self.rendering_enabled,) = struct_rendering_enabled.unpack(buff[offset:(offset + 1)])
            self.rendering_enabled = bool(self.rendering_enabled)
            offset += 1
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
        length += 8
        model_names_length = len(self.model_names)
        length += 4
        for i in range(0, model_names_length):
            model_namesi_x = self.model_names[i]
            model_namesi_length = len(model_namesi_x)
            if python3 or type(model_namesi_x) == unicode:
                model_namesi_x = model_namesi_x.encode('utf-8')
                model_namesi_length = len(model_namesi_x)
            length += 4
            length += model_namesi_length
        length += 1
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
        string_echo += '"sim_time": %s'%sim_time
        string_echo += ', '
        string_echo += '"model_names": ['
        model_names_length = len(self.model_names)
        for i in range(0, model_names_length):
            if i == (model_names_length - 1): 
                string_echo += '"model_names[i]": "%s"'%model_names[i]
                string_echo += '"'
            else:
                string_echo += '"model_names[i]": "%s"'%model_names[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"rendering_enabled": %s'%rendering_enabled
        string_echo += ', '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetWorldProperties"

    def getMD5(self):
        return "fe944c1c210919291ad14bc43b6c10cf"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetWorldProperties(object):
    Request = GetWorldPropertiesRequest
    Response = GetWorldPropertiesResponse

_struct_I = struct.Struct('<I')

