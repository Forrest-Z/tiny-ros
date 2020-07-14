import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg
import gazebo_msgs.msg
import tinyros

class ApplyBodyWrenchRequest(tinyros.Message):
    __slots__ = ['__id__','body_name','reference_frame','reference_point','wrench','start_time','duration']
    _slot_types = ['uint32','string','string','geometry_msgs.msg.Point','geometry_msgs.msg.Wrench','tinyros.Time','tinyros.Duration']

    def __init__(self):
        super(ApplyBodyWrenchRequest, self).__init__()
        self.__id__ = 0
        self.body_name = ''
        self.reference_frame = ''
        self.reference_point = geometry_msgs.msg.Point()
        self.wrench = geometry_msgs.msg.Wrench()
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
            _x = self.body_name
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
        offset += self.reference_point.serialize(buff)
        offset += self.wrench.serialize(buff)
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
            (length_body_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.body_name = buff[offset:(offset + length_body_name)].decode('utf-8')
            else:
                self.body_name = buff[offset:(offset + length_body_name)]
            offset += length_body_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
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
        offset += self.reference_point.deserialize(buff[offset:])
        offset += self.wrench.deserialize(buff[offset:])
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
        body_name_x = self.body_name
        body_name_length = len(body_name_x)
        if python3 or type(body_name_x) == unicode:
            body_name_x = body_name_x.encode('utf-8')
            body_name_length = len(body_name_x)
        length += 4
        length += body_name_length
        reference_frame_x = self.reference_frame
        reference_frame_length = len(reference_frame_x)
        if python3 or type(reference_frame_x) == unicode:
            reference_frame_x = reference_frame_x.encode('utf-8')
            reference_frame_length = len(reference_frame_x)
        length += 4
        length += reference_frame_length
        length += self.reference_point.serializedLength()
        length += self.wrench.serializedLength()
        length += 4
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"body_name": "%s"'%body_name
        string_echo += '", '
        string_echo += '"reference_frame": "%s"'%reference_frame
        string_echo += '", '
        string_echo += '"reference_point": {"'
        string_echo += self.reference_point.echo()
        string_echo += '}, '
        string_echo += '"wrench": {"'
        string_echo += self.wrench.echo()
        string_echo += '}, '
        string_echo += '"start_time.sec": %s, '%start_time.sec
        string_echo += '"start_time.nsec": %s'%start_time.nsec
        string_echo += ', '
        string_echo += '"duration.sec": %s, '%duration.sec
        string_echo += '"duration.nsec": %s'%duration.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ApplyBodyWrench"

    def getMD5(self):
        return "434adb4bdbb64c5610c7fadb31f0fa7d"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class ApplyBodyWrenchResponse(tinyros.Message):
    __slots__ = ['__id__','success','status_message']
    _slot_types = ['uint32','bool','string']

    def __init__(self):
        super(ApplyBodyWrenchResponse, self).__init__()
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
        return "gazebo_msgs/ApplyBodyWrench"

    def getMD5(self):
        return "f29b3c75e7d692065eda02aae6d3a3a0"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class ApplyBodyWrench(object):
    Request = ApplyBodyWrenchRequest
    Response = ApplyBodyWrenchResponse

_struct_I = struct.Struct('<I')

