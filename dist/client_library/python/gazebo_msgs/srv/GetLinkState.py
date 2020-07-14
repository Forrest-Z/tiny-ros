import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class GetLinkStateRequest(tinyros.Message):
    __slots__ = ['__id__','link_name','reference_frame']
    _slot_types = ['uint32','string','string']

    def __init__(self):
        super(GetLinkStateRequest, self).__init__()
        self.__id__ = 0
        self.link_name = ''
        self.reference_frame = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.link_name
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
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (self.__id__,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_link_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.link_name = buff[offset:(offset + length_link_name)].decode('utf-8')
            else:
                self.link_name = buff[offset:(offset + length_link_name)]
            offset += length_link_name
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
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        link_name_x = self.link_name
        link_name_length = len(link_name_x)
        if python3 or type(link_name_x) == unicode:
            link_name_x = link_name_x.encode('utf-8')
            link_name_length = len(link_name_x)
        length += 4
        length += link_name_length
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
        string_echo += '"link_name": "%s"'%link_name
        string_echo += '", '
        string_echo += '"reference_frame": "%s"'%reference_frame
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetLinkState"

    def getMD5(self):
        return "b9de4ed1795bda93c873763a2e55e76b"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetLinkStateResponse(tinyros.Message):
    __slots__ = ['__id__','link_state','success','status_message']
    _slot_types = ['uint32','gazebo_msgs.msg.LinkState','bool','string']

    def __init__(self):
        super(GetLinkStateResponse, self).__init__()
        self.__id__ = 0
        self.link_state = gazebo_msgs.msg.LinkState()
        self.success = False
        self.status_message = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.link_state.serialize(buff)
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
        offset += self.link_state.deserialize(buff[offset:])
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
        length += self.link_state.serializedLength()
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
        string_echo += '"link_state": {"'
        string_echo += self.link_state.echo()
        string_echo += '}, '
        string_echo += '"success": %s'%success
        string_echo += ', '
        string_echo += '"status_message": "%s"'%status_message
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/GetLinkState"

    def getMD5(self):
        return "4d4305d53d97f8edc3b3ce04bcb94ed0"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class GetLinkState(object):
    Request = GetLinkStateRequest
    Response = GetLinkStateResponse

_struct_I = struct.Struct('<I')

