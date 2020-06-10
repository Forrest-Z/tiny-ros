import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf.msg

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
        return "tf/FrameGraph"

    def getMD5(self):
        return "d66179e20d21cee31291f40363976e1d"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class FrameGraphResponse(tinyros.Message):
    __slots__ = ['__id__','dot_graph']
    _slot_types = ['uint32','string']

    def __init__(self):
        super(FrameGraphResponse, self).__init__()
        self.__id__ = 0
        self.dot_graph = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.__id__))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.dot_graph
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
            (length_dot_graph,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.dot_graph = buff[offset:(offset + length_dot_graph)].decode('utf-8')
            else:
                self.dot_graph = buff[offset:(offset + length_dot_graph)]
            offset += length_dot_graph
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        dot_graph_x = self.dot_graph
        dot_graph_length = len(dot_graph_x)
        if python3 or type(dot_graph_x) == unicode:
            dot_graph_x = dot_graph_x.encode('utf-8')
            dot_graph_length = len(dot_graph_x)
        length += 4
        length += dot_graph_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"__id__": %s'%__id__
        string_echo += ', '
        string_echo += '"dot_graph": "%s"'%dot_graph
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf/FrameGraph"

    def getMD5(self):
        return "8528671f80a5c0815f9700a446efbc36"

    def getID(self):
        return self.__id__

    def setID(self, id):
        self.__id__ = id

class FrameGraph(object):
    Request = FrameGraphRequest
    Response = FrameGraphResponse

_struct_I = struct.Struct('<I')

