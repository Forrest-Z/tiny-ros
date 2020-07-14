import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class PointCloud2(tinyros.Message):
    __slots__ = ['header','height','width','fields','is_bigendian','point_step','row_step','data','is_dense']
    _slot_types = ['std_msgs.msg.Header','uint32','uint32','sensor_msgs.msg.PointField[]','bool','uint32','uint32','uint8[]','bool']

    def __init__(self):
        super(PointCloud2, self).__init__()
        self.header = std_msgs.msg.Header()
        self.height = 0
        self.width = 0
        self.fields = []
        self.is_bigendian = False
        self.point_step = 0
        self.row_step = 0
        self.data = ''
        self.is_dense = False

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            buff.write(_struct_I.pack(self.height))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.width))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            fields_length = len(self.fields)
            buff.write(_struct_I.pack(fields_length))
            offset += 4
            for i in range(0, fields_length):
                offset += self.fields[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_is_bigendian = struct.Struct("<B")
            buff.write(struct_is_bigendian.pack(self.is_bigendian))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.point_step))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.row_step))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            data_length = len(self.data)
            buff.write(_struct_I.pack(data_length))
            offset += 4
            for i in range(0, data_length):
                try:
                    struct_datai = struct.Struct("<B")
                    buff.write(struct_datai.pack(self.data[i]))
                    offset += 1
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_is_dense = struct.Struct("<B")
            buff.write(struct_is_dense.pack(self.is_dense))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (self.height,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.width,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (fields_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.fields = [sensor_msgs.msg.PointField() for x in range(0, fields_length)]
            offset += 4
            for i in range(0, fields_length):
                offset += self.fields[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_is_bigendian = struct.Struct("<B")
            (self.is_bigendian,) = struct_is_bigendian.unpack(buff[offset:(offset + 1)])
            self.is_bigendian = bool(self.is_bigendian)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.point_step,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.row_step,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (data_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.data = [0 for x in range(0, data_length)]
            offset += 4
            for i in range(0, data_length):
                try:
                    struct_datai = struct.Struct("<B")
                    (self.data[i],) = struct_datai.unpack(buff[offset:(offset + 1)])
                    offset += 1
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_is_dense = struct.Struct("<B")
            (self.is_dense,) = struct_is_dense.unpack(buff[offset:(offset + 1)])
            self.is_dense = bool(self.is_dense)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        fields_length = len(self.fields)
        length += 4
        for i in range(0, fields_length):
            length += self.fields[i].serializedLength()
        length += 1
        length += 4
        length += 4
        data_length = len(self.data)
        length += 4
        for i in range(0, data_length):
            length += 1
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"height": %s'%height
        string_echo += ', '
        string_echo += '"width": %s'%width
        string_echo += ', '
        string_echo += '"fields": ['
        fields_length = len(self.fields)
        for i in range(0, fields_length):
            if i == (fields_length - 1): 
                string_echo += '{"fields%s": {'%i
                string_echo += self.fields[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"fields%s": {'%i
                string_echo += self.fields[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"is_bigendian": %s'%is_bigendian
        string_echo += ', '
        string_echo += '"point_step": %s'%point_step
        string_echo += ', '
        string_echo += '"row_step": %s'%row_step
        string_echo += ', '
        string_echo += '"data": ['
        data_length = len(self.data)
        for i in range(0, data_length):
            if i == (data_length - 1): 
                string_echo += '{"data%s": %s'%(i,data[i])
                string_echo += '}'
            else:
                string_echo += '{"data%s": %s'%(i,data[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"is_dense": %s'%is_dense
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/PointCloud2"

    def getMD5(self):
        return "6aa926339b282463281af40546b3b3cf"

_struct_I = struct.Struct('<I')

