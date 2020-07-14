import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import rosgraph_msgs.msg
import std_msgs.msg

class Log(tinyros.Message):
    __slots__ = ['header','level','name','msg','file','function','line','topics']
    _slot_types = ['std_msgs.msg.Header','byte','string','string','string','string','uint32','string[]']

    DEBUG = 1 
    INFO = 2  
    WARN = 4  
    ERROR = 8 
    FATAL = 16 

    def __init__(self):
        super(Log, self).__init__()
        self.header = std_msgs.msg.Header()
        self.level = 0
        self.name = ''
        self.msg = ''
        self.file = ''
        self.function = ''
        self.line = 0
        self.topics = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_level = struct.Struct("<b")
            buff.write(struct_level.pack(self.level))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.name
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
            _x = self.msg
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
            _x = self.file
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
            _x = self.function
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
            buff.write(_struct_I.pack(self.line))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            topics_length = len(self.topics)
            buff.write(_struct_I.pack(topics_length))
            offset += 4
            for i in range(0, topics_length):
                try:
                    _x = self.topics[i]
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
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_level = struct.Struct("<b")
            (self.level,) = struct_level.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.name = buff[offset:(offset + length_name)].decode('utf-8')
            else:
                self.name = buff[offset:(offset + length_name)]
            offset += length_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_msg,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.msg = buff[offset:(offset + length_msg)].decode('utf-8')
            else:
                self.msg = buff[offset:(offset + length_msg)]
            offset += length_msg
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_file,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.file = buff[offset:(offset + length_file)].decode('utf-8')
            else:
                self.file = buff[offset:(offset + length_file)]
            offset += length_file
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_function,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.function = buff[offset:(offset + length_function)].decode('utf-8')
            else:
                self.function = buff[offset:(offset + length_function)]
            offset += length_function
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.line,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (topics_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.topics = ['' for x in range(0, topics_length)]
            offset += 4
            for i in range(0, topics_length):
                try:
                    (length_topicsi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.topics[i] = buff[offset:(offset + length_topicsi)].decode('utf-8')
                    else:
                        self.topics[i] = buff[offset:(offset + length_topicsi)]
                    offset += length_topicsi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 1
        name_x = self.name
        name_length = len(name_x)
        if python3 or type(name_x) == unicode:
            name_x = name_x.encode('utf-8')
            name_length = len(name_x)
        length += 4
        length += name_length
        msg_x = self.msg
        msg_length = len(msg_x)
        if python3 or type(msg_x) == unicode:
            msg_x = msg_x.encode('utf-8')
            msg_length = len(msg_x)
        length += 4
        length += msg_length
        file_x = self.file
        file_length = len(file_x)
        if python3 or type(file_x) == unicode:
            file_x = file_x.encode('utf-8')
            file_length = len(file_x)
        length += 4
        length += file_length
        function_x = self.function
        function_length = len(function_x)
        if python3 or type(function_x) == unicode:
            function_x = function_x.encode('utf-8')
            function_length = len(function_x)
        length += 4
        length += function_length
        length += 4
        topics_length = len(self.topics)
        length += 4
        for i in range(0, topics_length):
            topicsi_x = self.topics[i]
            topicsi_length = len(topicsi_x)
            if python3 or type(topicsi_x) == unicode:
                topicsi_x = topicsi_x.encode('utf-8')
                topicsi_length = len(topicsi_x)
            length += 4
            length += topicsi_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"level": %s'%level
        string_echo += ', '
        string_echo += '"name": "%s"'%name
        string_echo += '", '
        string_echo += '"msg": "%s"'%msg
        string_echo += '", '
        string_echo += '"file": "%s"'%file
        string_echo += '", '
        string_echo += '"function": "%s"'%function
        string_echo += '", '
        string_echo += '"line": %s'%line
        string_echo += ', '
        string_echo += '"topics": ['
        topics_length = len(self.topics)
        for i in range(0, topics_length):
            if i == (topics_length - 1): 
                string_echo += '"topics[i]": "%s"'%topics[i]
                string_echo += '"'
            else:
                string_echo += '"topics[i]": "%s"'%topics[i]
                string_echo += '", '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "rosgraph_msgs/Log"

    def getMD5(self):
        return "2de9daf47e984009074d74dbdd492d49"

_struct_I = struct.Struct('<I')

