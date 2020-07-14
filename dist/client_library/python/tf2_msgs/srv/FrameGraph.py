import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg

class FrameGraphRequest(tinyros.Message):
    __slots__ = ['__id__']
    _slot_types = ['uint32']

    def __init__(self):
        super(FrameGraphRequest, self).__init__()
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
        return "tf2_msgs/FrameGraph"

    def getMD5(self):
        return "aa023d3c31410363e0583979223258c8"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class FrameGraphResponse(tinyros.Message):
    __slots__ = ['__id__','frame_yaml']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(FrameGraphResponse, self).__init__()
        self.__id__ = 0
        self.frame_yaml = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.frame_yaml
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
            (length_frame_yaml,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.frame_yaml = buff[offset:(offset + length_frame_yaml)].decode('utf-8')
            else:
                self.frame_yaml = buff[offset:(offset + length_frame_yaml)]
            offset += length_frame_yaml
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        frame_yaml_x = self.frame_yaml
        frame_yaml_length = len(frame_yaml_x)
        if python3 or type(frame_yaml_x) == unicode:
            frame_yaml_x = frame_yaml_x.encode('utf-8')
            frame_yaml_length = len(frame_yaml_x)
        length += 4
        length += frame_yaml_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"frame_yaml": "%s"'%frame_yaml
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf2_msgs/FrameGraph"

    def getMD5(self):
        return "97e361486f8cc8fb1a460cf17555126b"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class FrameGraph(object):
    Request = FrameGraphRequest
    Response = FrameGraphResponse

_struct_I = struct.Struct('<I')

