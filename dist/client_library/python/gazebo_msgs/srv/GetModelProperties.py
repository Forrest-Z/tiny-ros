import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class GetModelPropertiesRequest(tinyros.Message):
    __slots__ = ['__id__','model_name']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(GetModelPropertiesRequest, self).__init__()
        self.__id__ = 0
        self.model_name = ''

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
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"model_name": "%s"'%model_name
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetModelProperties"

    def getMD5(self):
        return "fe0194bf75c917c89b820b09c12fe6c1"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetModelPropertiesResponse(tinyros.Message):
    __slots__ = ['__id__','parent_model_name','canonical_body_name','body_names','geom_names','joint_names','child_model_names','is_static','success','status_message']
    _slot_types = ['uint32','string','string','string[]','string[]','string[]','string[]','bool','bool','string']

    def __init__(self):
        super(GetModelPropertiesResponse, self).__init__()
        self.__id__ = 0
        self.parent_model_name = ''
        self.canonical_body_name = ''
        self.body_names = []
        self.geom_names = []
        self.joint_names = []
        self.child_model_names = []
        self.is_static = False
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
            _x = self.parent_model_name
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
            _x = self.canonical_body_name
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
            body_names_length = len(self.body_names)
            buff.write(_struct_I.pack(body_names_length))
            offset += 4
            for i in range(0, body_names_length):
                try:
                    _x = self.body_names[i]
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
            geom_names_length = len(self.geom_names)
            buff.write(_struct_I.pack(geom_names_length))
            offset += 4
            for i in range(0, geom_names_length):
                try:
                    _x = self.geom_names[i]
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
            child_model_names_length = len(self.child_model_names)
            buff.write(_struct_I.pack(child_model_names_length))
            offset += 4
            for i in range(0, child_model_names_length):
                try:
                    _x = self.child_model_names[i]
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
            struct_is_static = struct.Struct("<B")
            buff.write(struct_is_static.pack(self.is_static))
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
            (length_parent_model_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.parent_model_name = buff[offset:(offset + length_parent_model_name)].decode('utf-8')
            else:
                self.parent_model_name = buff[offset:(offset + length_parent_model_name)]
            offset += length_parent_model_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_canonical_body_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.canonical_body_name = buff[offset:(offset + length_canonical_body_name)].decode('utf-8')
            else:
                self.canonical_body_name = buff[offset:(offset + length_canonical_body_name)]
            offset += length_canonical_body_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (body_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.body_names = ['' for x in range(0, body_names_length)]
            offset += 4
            for i in range(0, body_names_length):
                try:
                    (length_body_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.body_names[i] = buff[offset:(offset + length_body_namesi)].decode('utf-8')
                    else:
                        self.body_names[i] = buff[offset:(offset + length_body_namesi)]
                    offset += length_body_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (geom_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.geom_names = ['' for x in range(0, geom_names_length)]
            offset += 4
            for i in range(0, geom_names_length):
                try:
                    (length_geom_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.geom_names[i] = buff[offset:(offset + length_geom_namesi)].decode('utf-8')
                    else:
                        self.geom_names[i] = buff[offset:(offset + length_geom_namesi)]
                    offset += length_geom_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
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
            (child_model_names_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.child_model_names = ['' for x in range(0, child_model_names_length)]
            offset += 4
            for i in range(0, child_model_names_length):
                try:
                    (length_child_model_namesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.child_model_names[i] = buff[offset:(offset + length_child_model_namesi)].decode('utf-8')
                    else:
                        self.child_model_names[i] = buff[offset:(offset + length_child_model_namesi)]
                    offset += length_child_model_namesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_is_static = struct.Struct("<B")
            (self.is_static,) = struct_is_static.unpack(buff[offset:(offset + 1)])
            self.is_static = bool(self.is_static)
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
        parent_model_name_x = self.parent_model_name
        parent_model_name_length = len(parent_model_name_x)
        if python3 or type(parent_model_name_x) == unicode:
            parent_model_name_x = parent_model_name_x.encode('utf-8')
            parent_model_name_length = len(parent_model_name_x)
        length += 4
        length += parent_model_name_length
        canonical_body_name_x = self.canonical_body_name
        canonical_body_name_length = len(canonical_body_name_x)
        if python3 or type(canonical_body_name_x) == unicode:
            canonical_body_name_x = canonical_body_name_x.encode('utf-8')
            canonical_body_name_length = len(canonical_body_name_x)
        length += 4
        length += canonical_body_name_length
        body_names_length = len(self.body_names)
        length += 4
        for i in range(0, body_names_length):
            body_namesi_x = self.body_names[i]
            body_namesi_length = len(body_namesi_x)
            if python3 or type(body_namesi_x) == unicode:
                body_namesi_x = body_namesi_x.encode('utf-8')
                body_namesi_length = len(body_namesi_x)
            length += 4
            length += body_namesi_length
        geom_names_length = len(self.geom_names)
        length += 4
        for i in range(0, geom_names_length):
            geom_namesi_x = self.geom_names[i]
            geom_namesi_length = len(geom_namesi_x)
            if python3 or type(geom_namesi_x) == unicode:
                geom_namesi_x = geom_namesi_x.encode('utf-8')
                geom_namesi_length = len(geom_namesi_x)
            length += 4
            length += geom_namesi_length
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
        child_model_names_length = len(self.child_model_names)
        length += 4
        for i in range(0, child_model_names_length):
            child_model_namesi_x = self.child_model_names[i]
            child_model_namesi_length = len(child_model_namesi_x)
            if python3 or type(child_model_namesi_x) == unicode:
                child_model_namesi_x = child_model_namesi_x.encode('utf-8')
                child_model_namesi_length = len(child_model_namesi_x)
            length += 4
            length += child_model_namesi_length
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
        string_echo += '"parent_model_name": "%s"'%parent_model_name
        string_echo += '", '
        string_echo += '"canonical_body_name": "%s"'%canonical_body_name
        string_echo += '", '
        string_echo += '"body_names": ['
        body_names_length = len(self.body_names)
        for i in range(0, body_names_length):
            if i == (body_names_length - 1): 
                string_echo += '"body_names[i]": "%s"'%body_names[i]
                string_echo += '"'
            else:
                string_echo += '"body_names[i]": "%s"'%body_names[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"geom_names": ['
        geom_names_length = len(self.geom_names)
        for i in range(0, geom_names_length):
            if i == (geom_names_length - 1): 
                string_echo += '"geom_names[i]": "%s"'%geom_names[i]
                string_echo += '"'
            else:
                string_echo += '"geom_names[i]": "%s"'%geom_names[i]
                string_echo += '", '
        string_echo += '], '
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
        string_echo += '"child_model_names": ['
        child_model_names_length = len(self.child_model_names)
        for i in range(0, child_model_names_length):
            if i == (child_model_names_length - 1): 
                string_echo += '"child_model_names[i]": "%s"'%child_model_names[i]
                string_echo += '"'
            else:
                string_echo += '"child_model_names[i]": "%s"'%child_model_names[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"is_static": %s'%is_static
        string_echo += ', '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetModelProperties"

    def getMD5(self):
        return "d8f16b08abaf0220a551cf9025748602"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetModelProperties(object):
    Request = GetModelPropertiesRequest
    Response = GetModelPropertiesResponse

_struct_I = struct.Struct('<I')

