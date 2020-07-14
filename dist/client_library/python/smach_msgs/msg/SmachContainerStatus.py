import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import smach_msgs.msg
import std_msgs.msg

class SmachContainerStatus(tinyros.Message):
    __slots__ = ['header','path','initial_states','active_states','local_data','info']
    _slot_types = ['std_msgs.msg.Header','string','string[]','string[]','string','string']

    def __init__(self):
        super(SmachContainerStatus, self).__init__()
        self.header = std_msgs.msg.Header()
        self.path = ''
        self.initial_states = []
        self.active_states = []
        self.local_data = ''
        self.info = ''

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            _x = self.path
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
            initial_states_length = len(self.initial_states)
            buff.write(_struct_I.pack(initial_states_length))
            offset += 4
            for i in range(0, initial_states_length):
                try:
                    _x = self.initial_states[i]
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
            active_states_length = len(self.active_states)
            buff.write(_struct_I.pack(active_states_length))
            offset += 4
            for i in range(0, active_states_length):
                try:
                    _x = self.active_states[i]
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
            _x = self.local_data
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
            _x = self.info
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
        offset += self.header.deserialize(buff[offset:])
        try:
            (length_path,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.path = buff[offset:(offset + length_path)].decode('utf-8')
            else:
                self.path = buff[offset:(offset + length_path)]
            offset += length_path
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (initial_states_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.initial_states = ['' for x in range(0, initial_states_length)]
            offset += 4
            for i in range(0, initial_states_length):
                try:
                    (length_initial_statesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.initial_states[i] = buff[offset:(offset + length_initial_statesi)].decode('utf-8')
                    else:
                        self.initial_states[i] = buff[offset:(offset + length_initial_statesi)]
                    offset += length_initial_statesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (active_states_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.active_states = ['' for x in range(0, active_states_length)]
            offset += 4
            for i in range(0, active_states_length):
                try:
                    (length_active_statesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.active_states[i] = buff[offset:(offset + length_active_statesi)].decode('utf-8')
                    else:
                        self.active_states[i] = buff[offset:(offset + length_active_statesi)]
                    offset += length_active_statesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_local_data,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.local_data = buff[offset:(offset + length_local_data)].decode('utf-8')
            else:
                self.local_data = buff[offset:(offset + length_local_data)]
            offset += length_local_data
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_info,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.info = buff[offset:(offset + length_info)].decode('utf-8')
            else:
                self.info = buff[offset:(offset + length_info)]
            offset += length_info
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        path_x = self.path
        path_length = len(path_x)
        if python3 or type(path_x) == unicode:
            path_x = path_x.encode('utf-8')
            path_length = len(path_x)
        length += 4
        length += path_length
        initial_states_length = len(self.initial_states)
        length += 4
        for i in range(0, initial_states_length):
            initial_statesi_x = self.initial_states[i]
            initial_statesi_length = len(initial_statesi_x)
            if python3 or type(initial_statesi_x) == unicode:
                initial_statesi_x = initial_statesi_x.encode('utf-8')
                initial_statesi_length = len(initial_statesi_x)
            length += 4
            length += initial_statesi_length
        active_states_length = len(self.active_states)
        length += 4
        for i in range(0, active_states_length):
            active_statesi_x = self.active_states[i]
            active_statesi_length = len(active_statesi_x)
            if python3 or type(active_statesi_x) == unicode:
                active_statesi_x = active_statesi_x.encode('utf-8')
                active_statesi_length = len(active_statesi_x)
            length += 4
            length += active_statesi_length
        local_data_x = self.local_data
        local_data_length = len(local_data_x)
        if python3 or type(local_data_x) == unicode:
            local_data_x = local_data_x.encode('utf-8')
            local_data_length = len(local_data_x)
        length += 4
        length += local_data_length
        info_x = self.info
        info_length = len(info_x)
        if python3 or type(info_x) == unicode:
            info_x = info_x.encode('utf-8')
            info_length = len(info_x)
        length += 4
        length += info_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"path": "%s"'%path
        string_echo += '", '
        string_echo += '"initial_states": ['
        initial_states_length = len(self.initial_states)
        for i in range(0, initial_states_length):
            if i == (initial_states_length - 1): 
                string_echo += '"initial_states[i]": "%s"'%initial_states[i]
                string_echo += '"'
            else:
                string_echo += '"initial_states[i]": "%s"'%initial_states[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"active_states": ['
        active_states_length = len(self.active_states)
        for i in range(0, active_states_length):
            if i == (active_states_length - 1): 
                string_echo += '"active_states[i]": "%s"'%active_states[i]
                string_echo += '"'
            else:
                string_echo += '"active_states[i]": "%s"'%active_states[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"local_data": "%s"'%local_data
        string_echo += '", '
        string_echo += '"info": "%s"'%info
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "smach_msgs/SmachContainerStatus"

    def getMD5(self):
        return "3e74cf72da18311be27e7a5970ea6415"

_struct_I = struct.Struct('<I')

