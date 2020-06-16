import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg
import geometry_msgs.msg

class ContactState(tinyros.Message):
    __slots__ = ['info','collision1_name','collision2_name','wrenches','total_wrench','contact_positions','contact_normals','depths']
    _slot_types = ['string','string','string','geometry_msgs.msg.Wrench[]','geometry_msgs.msg.Wrench','geometry_msgs.msg.Vector3[]','geometry_msgs.msg.Vector3[]','float64[]']

    def __init__(self):
        super(ContactState, self).__init__()
        self.info = ''
        self.collision1_name = ''
        self.collision2_name = ''
        self.wrenches = []
        self.total_wrench = geometry_msgs.msg.Wrench()
        self.contact_positions = []
        self.contact_normals = []
        self.depths = []

    def serialize(self, buff):
        offset = 0
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
        try:
            _x = self.collision1_name
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
            _x = self.collision2_name
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
            wrenches_length = len(self.wrenches)
            buff.write(_struct_I.pack(wrenches_length))
            offset += 4
            for i in range(0, wrenches_length):
                offset += self.wrenches[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.total_wrench.serialize(buff)
        try:
            contact_positions_length = len(self.contact_positions)
            buff.write(_struct_I.pack(contact_positions_length))
            offset += 4
            for i in range(0, contact_positions_length):
                offset += self.contact_positions[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            contact_normals_length = len(self.contact_normals)
            buff.write(_struct_I.pack(contact_normals_length))
            offset += 4
            for i in range(0, contact_normals_length):
                offset += self.contact_normals[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            depths_length = len(self.depths)
            buff.write(_struct_I.pack(depths_length))
            offset += 4
            for i in range(0, depths_length):
                try:
                    struct_depthsi = struct.Struct("<d")
                    buff.write(struct_depthsi.pack(self.depths[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
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
        try:
            (length_collision1_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.collision1_name = buff[offset:(offset + length_collision1_name)].decode('utf-8')
            else:
                self.collision1_name = buff[offset:(offset + length_collision1_name)]
            offset += length_collision1_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_collision2_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.collision2_name = buff[offset:(offset + length_collision2_name)].decode('utf-8')
            else:
                self.collision2_name = buff[offset:(offset + length_collision2_name)]
            offset += length_collision2_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (wrenches_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.wrenches = [geometry_msgs.msg.Wrench() for x in range(0, wrenches_length)]
            offset += 4
            for i in range(0, wrenches_length):
                offset += self.wrenches[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.total_wrench.deserialize(buff[offset:])
        try:
            (contact_positions_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.contact_positions = [geometry_msgs.msg.Vector3() for x in range(0, contact_positions_length)]
            offset += 4
            for i in range(0, contact_positions_length):
                offset += self.contact_positions[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (contact_normals_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.contact_normals = [geometry_msgs.msg.Vector3() for x in range(0, contact_normals_length)]
            offset += 4
            for i in range(0, contact_normals_length):
                offset += self.contact_normals[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (depths_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.depths = [0. for x in range(0, depths_length)]
            offset += 4
            for i in range(0, depths_length):
                try:
                    struct_depthsi = struct.Struct("<d")
                    (self.depths[i],) = struct_depthsi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        info_x = self.info
        info_length = len(info_x)
        if python3 or type(info_x) == unicode:
            info_x = info_x.encode('utf-8')
            info_length = len(info_x)
        length += 4
        length += info_length
        collision1_name_x = self.collision1_name
        collision1_name_length = len(collision1_name_x)
        if python3 or type(collision1_name_x) == unicode:
            collision1_name_x = collision1_name_x.encode('utf-8')
            collision1_name_length = len(collision1_name_x)
        length += 4
        length += collision1_name_length
        collision2_name_x = self.collision2_name
        collision2_name_length = len(collision2_name_x)
        if python3 or type(collision2_name_x) == unicode:
            collision2_name_x = collision2_name_x.encode('utf-8')
            collision2_name_length = len(collision2_name_x)
        length += 4
        length += collision2_name_length
        wrenches_length = len(self.wrenches)
        length += 4
        for i in range(0, wrenches_length):
            length += self.wrenches[i].serializedLength()
        length += self.total_wrench.serializedLength()
        contact_positions_length = len(self.contact_positions)
        length += 4
        for i in range(0, contact_positions_length):
            length += self.contact_positions[i].serializedLength()
        contact_normals_length = len(self.contact_normals)
        length += 4
        for i in range(0, contact_normals_length):
            length += self.contact_normals[i].serializedLength()
        depths_length = len(self.depths)
        length += 4
        for i in range(0, depths_length):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"info": "%s"'%info
        string_echo += '", '
        string_echo += '"collision1_name": "%s"'%collision1_name
        string_echo += '", '
        string_echo += '"collision2_name": "%s"'%collision2_name
        string_echo += '", '
        string_echo += '"wrenches": ['
        wrenches_length = len(self.wrenches)
        for i in range(0, wrenches_length):
            if i == (wrenches_length - 1): 
                string_echo += '{"wrenches%s": {'%i
                string_echo += self.wrenches[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"wrenches%s": {'%i
                string_echo += self.wrenches[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"total_wrench": {"'
        string_echo += self.total_wrench.echo()
        string_echo += '}, '
        string_echo += '"contact_positions": ['
        contact_positions_length = len(self.contact_positions)
        for i in range(0, contact_positions_length):
            if i == (contact_positions_length - 1): 
                string_echo += '{"contact_positions%s": {'%i
                string_echo += self.contact_positions[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"contact_positions%s": {'%i
                string_echo += self.contact_positions[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"contact_normals": ['
        contact_normals_length = len(self.contact_normals)
        for i in range(0, contact_normals_length):
            if i == (contact_normals_length - 1): 
                string_echo += '{"contact_normals%s": {'%i
                string_echo += self.contact_normals[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"contact_normals%s": {'%i
                string_echo += self.contact_normals[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"depths": ['
        depths_length = len(self.depths)
        for i in range(0, depths_length):
            if i == (depths_length - 1): 
                string_echo += '{"depths%s": %s'%(i,depths[i])
                string_echo += '}'
            else:
                string_echo += '{"depths%s": %s'%(i,depths[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ContactState"

    def getMD5(self):
        return "d82d0f0cae88aebf6b2cc86caea33a2b"

_struct_I = struct.Struct('<I')

