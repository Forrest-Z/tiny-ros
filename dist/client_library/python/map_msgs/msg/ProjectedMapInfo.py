import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import map_msgs.msg

class ProjectedMapInfo(tinyros.Message):
    __slots__ = ['frame_id','x','y','width','height','min_z','max_z']
    _slot_types = ['string','float64','float64','float64','float64','float64','float64']

    def __init__(self):
        super(ProjectedMapInfo, self).__init__()
        self.frame_id = ''
        self.x = 0.
        self.y = 0.
        self.width = 0.
        self.height = 0.
        self.min_z = 0.
        self.max_z = 0.

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.frame_id
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
            struct_x = struct.Struct("<d")
            buff.write(struct_x.pack(self.x))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            buff.write(struct_y.pack(self.y))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_width = struct.Struct("<d")
            buff.write(struct_width.pack(self.width))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_height = struct.Struct("<d")
            buff.write(struct_height.pack(self.height))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_min_z = struct.Struct("<d")
            buff.write(struct_min_z.pack(self.min_z))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_max_z = struct.Struct("<d")
            buff.write(struct_max_z.pack(self.max_z))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_frame_id,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.frame_id = buff[offset:(offset + length_frame_id)].decode('utf-8')
            else:
                self.frame_id = buff[offset:(offset + length_frame_id)]
            offset += length_frame_id
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_x = struct.Struct("<d")
            (self.x,) = struct_x.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_y = struct.Struct("<d")
            (self.y,) = struct_y.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_width = struct.Struct("<d")
            (self.width,) = struct_width.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_height = struct.Struct("<d")
            (self.height,) = struct_height.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_min_z = struct.Struct("<d")
            (self.min_z,) = struct_min_z.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_max_z = struct.Struct("<d")
            (self.max_z,) = struct_max_z.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        frame_id_x = self.frame_id
        frame_id_length = len(frame_id_x)
        if python3 or type(frame_id_x) == unicode:
            frame_id_x = frame_id_x.encode('utf-8')
            frame_id_length = len(frame_id_x)
        length += 4
        length += frame_id_length
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"frame_id": "%s"'%frame_id
        string_echo += '", '
        string_echo += '"x": %s'%x
        string_echo += ', '
        string_echo += '"y": %s'%y
        string_echo += ', '
        string_echo += '"width": %s'%width
        string_echo += ', '
        string_echo += '"height": %s'%height
        string_echo += ', '
        string_echo += '"min_z": %s'%min_z
        string_echo += ', '
        string_echo += '"max_z": %s'%max_z
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "map_msgs/ProjectedMapInfo"

    def getMD5(self):
        return "f661365637fb759e63cb5d179a4461e1"

_struct_I = struct.Struct('<I')

