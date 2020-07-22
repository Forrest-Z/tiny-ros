import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tinyros_msgs.msg

class TopicInfo(tinyros.Message):
    __slots__ = ['topic_id','topic_name','message_type','md5sum','buffer_size','negotiated','node']
    _slot_types = ['uint32','string','string','string','int32','bool','string']

    ID_PUBLISHER = 0
    ID_SUBSCRIBER = 1
    ID_SERVICE_SERVER = 2
    ID_SERVICE_CLIENT = 4
    ID_ROSTOPIC_REQUEST = 6
    ID_ROSSERVICE_REQUEST = 7
    ID_LOG = 8
    ID_TIME = 9
    ID_NEGOTIATED = 10
    ID_SESSION_ID = 11

    def __init__(self):
        super(TopicInfo, self).__init__()
        self.topic_id = 0
        self.topic_name = ''
        self.message_type = ''
        self.md5sum = ''
        self.buffer_size = 0
        self.negotiated = False
        self.node = ''

    def serialize(self, buff):
        offset = 0
        try:
            buff.write(_struct_I.pack(self.topic_id))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.topic_name
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
            _x = self.message_type
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
            _x = self.md5sum
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
            struct_buffer_size = struct.Struct("<i")
            buff.write(struct_buffer_size.pack(self.buffer_size))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_negotiated = struct.Struct("<B")
            buff.write(struct_negotiated.pack(self.negotiated))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.node
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
            (self.topic_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_topic_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.topic_name = buff[offset:(offset + length_topic_name)].decode('utf-8')
            else:
                self.topic_name = buff[offset:(offset + length_topic_name)]
            offset += length_topic_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_message_type,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.message_type = buff[offset:(offset + length_message_type)].decode('utf-8')
            else:
                self.message_type = buff[offset:(offset + length_message_type)]
            offset += length_message_type
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_md5sum,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.md5sum = buff[offset:(offset + length_md5sum)].decode('utf-8')
            else:
                self.md5sum = buff[offset:(offset + length_md5sum)]
            offset += length_md5sum
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_buffer_size = struct.Struct("<i")
            (self.buffer_size,) = struct_buffer_size.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_negotiated = struct.Struct("<B")
            (self.negotiated,) = struct_negotiated.unpack(buff[offset:(offset + 1)])
            self.negotiated = bool(self.negotiated)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_node,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.node = buff[offset:(offset + length_node)].decode('utf-8')
            else:
                self.node = buff[offset:(offset + length_node)]
            offset += length_node
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 4
        topic_name_x = self.topic_name
        topic_name_length = len(topic_name_x)
        if python3 or type(topic_name_x) == unicode:
            topic_name_x = topic_name_x.encode('utf-8')
            topic_name_length = len(topic_name_x)
        length += 4
        length += topic_name_length
        message_type_x = self.message_type
        message_type_length = len(message_type_x)
        if python3 or type(message_type_x) == unicode:
            message_type_x = message_type_x.encode('utf-8')
            message_type_length = len(message_type_x)
        length += 4
        length += message_type_length
        md5sum_x = self.md5sum
        md5sum_length = len(md5sum_x)
        if python3 or type(md5sum_x) == unicode:
            md5sum_x = md5sum_x.encode('utf-8')
            md5sum_length = len(md5sum_x)
        length += 4
        length += md5sum_length
        length += 4
        length += 1
        node_x = self.node
        node_length = len(node_x)
        if python3 or type(node_x) == unicode:
            node_x = node_x.encode('utf-8')
            node_length = len(node_x)
        length += 4
        length += node_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"topic_id": %s'%topic_id
        string_echo += ', '
        string_echo += '"topic_name": "%s"'%topic_name
        string_echo += '", '
        string_echo += '"message_type": "%s"'%message_type
        string_echo += '", '
        string_echo += '"md5sum": "%s"'%md5sum
        string_echo += '", '
        string_echo += '"buffer_size": %s'%buffer_size
        string_echo += ', '
        string_echo += '"negotiated": %s'%negotiated
        string_echo += ', '
        string_echo += '"node": "%s"'%node
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tinyros_msgs/TopicInfo"

    def getMD5(self):
        return "76d40676946fcde66f228def7575451a"

_struct_I = struct.Struct('<I')

