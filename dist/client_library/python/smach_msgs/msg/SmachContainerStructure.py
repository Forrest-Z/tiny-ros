import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import smach_msgs.msg
import std_msgs.msg

class SmachContainerStructure(tinyros.Message):
    __slots__ = ['header','path','children','internal_outcomes','outcomes_from','outcomes_to','container_outcomes']
    _slot_types = ['std_msgs.msg.Header','string','string[]','string[]','string[]','string[]','string[]']

    def __init__(self):
        super(SmachContainerStructure, self).__init__()
        self.header = std_msgs.msg.Header()
        self.path = ''
        self.children = []
        self.internal_outcomes = []
        self.outcomes_from = []
        self.outcomes_to = []
        self.container_outcomes = []

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
            children_length = len(self.children)
            buff.write(_struct_I.pack(children_length))
            offset += 4
            for i in range(0, children_length):
                try:
                    _x = self.children[i]
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
            internal_outcomes_length = len(self.internal_outcomes)
            buff.write(_struct_I.pack(internal_outcomes_length))
            offset += 4
            for i in range(0, internal_outcomes_length):
                try:
                    _x = self.internal_outcomes[i]
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
            outcomes_from_length = len(self.outcomes_from)
            buff.write(_struct_I.pack(outcomes_from_length))
            offset += 4
            for i in range(0, outcomes_from_length):
                try:
                    _x = self.outcomes_from[i]
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
            outcomes_to_length = len(self.outcomes_to)
            buff.write(_struct_I.pack(outcomes_to_length))
            offset += 4
            for i in range(0, outcomes_to_length):
                try:
                    _x = self.outcomes_to[i]
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
            container_outcomes_length = len(self.container_outcomes)
            buff.write(_struct_I.pack(container_outcomes_length))
            offset += 4
            for i in range(0, container_outcomes_length):
                try:
                    _x = self.container_outcomes[i]
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
            (children_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.children = ['' for x in range(0, children_length)]
            offset += 4
            for i in range(0, children_length):
                try:
                    (length_childreni,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.children[i] = buff[offset:(offset + length_childreni)].decode('utf-8')
                    else:
                        self.children[i] = buff[offset:(offset + length_childreni)]
                    offset += length_childreni
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (internal_outcomes_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.internal_outcomes = ['' for x in range(0, internal_outcomes_length)]
            offset += 4
            for i in range(0, internal_outcomes_length):
                try:
                    (length_internal_outcomesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.internal_outcomes[i] = buff[offset:(offset + length_internal_outcomesi)].decode('utf-8')
                    else:
                        self.internal_outcomes[i] = buff[offset:(offset + length_internal_outcomesi)]
                    offset += length_internal_outcomesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (outcomes_from_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.outcomes_from = ['' for x in range(0, outcomes_from_length)]
            offset += 4
            for i in range(0, outcomes_from_length):
                try:
                    (length_outcomes_fromi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.outcomes_from[i] = buff[offset:(offset + length_outcomes_fromi)].decode('utf-8')
                    else:
                        self.outcomes_from[i] = buff[offset:(offset + length_outcomes_fromi)]
                    offset += length_outcomes_fromi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (outcomes_to_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.outcomes_to = ['' for x in range(0, outcomes_to_length)]
            offset += 4
            for i in range(0, outcomes_to_length):
                try:
                    (length_outcomes_toi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.outcomes_to[i] = buff[offset:(offset + length_outcomes_toi)].decode('utf-8')
                    else:
                        self.outcomes_to[i] = buff[offset:(offset + length_outcomes_toi)]
                    offset += length_outcomes_toi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (container_outcomes_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.container_outcomes = ['' for x in range(0, container_outcomes_length)]
            offset += 4
            for i in range(0, container_outcomes_length):
                try:
                    (length_container_outcomesi,) = _struct_I.unpack(buff[offset:(offset + 4)])
                    offset += 4
                    if python3:
                        self.container_outcomes[i] = buff[offset:(offset + length_container_outcomesi)].decode('utf-8')
                    else:
                        self.container_outcomes[i] = buff[offset:(offset + length_container_outcomesi)]
                    offset += length_container_outcomesi
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
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
        children_length = len(self.children)
        length += 4
        for i in range(0, children_length):
            childreni_x = self.children[i]
            childreni_length = len(childreni_x)
            if python3 or type(childreni_x) == unicode:
                childreni_x = childreni_x.encode('utf-8')
                childreni_length = len(childreni_x)
            length += 4
            length += childreni_length
        internal_outcomes_length = len(self.internal_outcomes)
        length += 4
        for i in range(0, internal_outcomes_length):
            internal_outcomesi_x = self.internal_outcomes[i]
            internal_outcomesi_length = len(internal_outcomesi_x)
            if python3 or type(internal_outcomesi_x) == unicode:
                internal_outcomesi_x = internal_outcomesi_x.encode('utf-8')
                internal_outcomesi_length = len(internal_outcomesi_x)
            length += 4
            length += internal_outcomesi_length
        outcomes_from_length = len(self.outcomes_from)
        length += 4
        for i in range(0, outcomes_from_length):
            outcomes_fromi_x = self.outcomes_from[i]
            outcomes_fromi_length = len(outcomes_fromi_x)
            if python3 or type(outcomes_fromi_x) == unicode:
                outcomes_fromi_x = outcomes_fromi_x.encode('utf-8')
                outcomes_fromi_length = len(outcomes_fromi_x)
            length += 4
            length += outcomes_fromi_length
        outcomes_to_length = len(self.outcomes_to)
        length += 4
        for i in range(0, outcomes_to_length):
            outcomes_toi_x = self.outcomes_to[i]
            outcomes_toi_length = len(outcomes_toi_x)
            if python3 or type(outcomes_toi_x) == unicode:
                outcomes_toi_x = outcomes_toi_x.encode('utf-8')
                outcomes_toi_length = len(outcomes_toi_x)
            length += 4
            length += outcomes_toi_length
        container_outcomes_length = len(self.container_outcomes)
        length += 4
        for i in range(0, container_outcomes_length):
            container_outcomesi_x = self.container_outcomes[i]
            container_outcomesi_length = len(container_outcomesi_x)
            if python3 or type(container_outcomesi_x) == unicode:
                container_outcomesi_x = container_outcomesi_x.encode('utf-8')
                container_outcomesi_length = len(container_outcomesi_x)
            length += 4
            length += container_outcomesi_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"path": "%s"'%path
        string_echo += '", '
        string_echo += '"children": ['
        children_length = len(self.children)
        for i in range(0, children_length):
            if i == (children_length - 1): 
                string_echo += '"children[i]": "%s"'%children[i]
                string_echo += '"'
            else:
                string_echo += '"children[i]": "%s"'%children[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"internal_outcomes": ['
        internal_outcomes_length = len(self.internal_outcomes)
        for i in range(0, internal_outcomes_length):
            if i == (internal_outcomes_length - 1): 
                string_echo += '"internal_outcomes[i]": "%s"'%internal_outcomes[i]
                string_echo += '"'
            else:
                string_echo += '"internal_outcomes[i]": "%s"'%internal_outcomes[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"outcomes_from": ['
        outcomes_from_length = len(self.outcomes_from)
        for i in range(0, outcomes_from_length):
            if i == (outcomes_from_length - 1): 
                string_echo += '"outcomes_from[i]": "%s"'%outcomes_from[i]
                string_echo += '"'
            else:
                string_echo += '"outcomes_from[i]": "%s"'%outcomes_from[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"outcomes_to": ['
        outcomes_to_length = len(self.outcomes_to)
        for i in range(0, outcomes_to_length):
            if i == (outcomes_to_length - 1): 
                string_echo += '"outcomes_to[i]": "%s"'%outcomes_to[i]
                string_echo += '"'
            else:
                string_echo += '"outcomes_to[i]": "%s"'%outcomes_to[i]
                string_echo += '", '
        string_echo += '], '
        string_echo += '"container_outcomes": ['
        container_outcomes_length = len(self.container_outcomes)
        for i in range(0, container_outcomes_length):
            if i == (container_outcomes_length - 1): 
                string_echo += '"container_outcomes[i]": "%s"'%container_outcomes[i]
                string_echo += '"'
            else:
                string_echo += '"container_outcomes[i]": "%s"'%container_outcomes[i]
                string_echo += '", '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "smach_msgs/SmachContainerStructure"

    def getMD5(self):
        return "0209663ab13753a56b6ac1d094d07fbe"

_struct_I = struct.Struct('<I')

