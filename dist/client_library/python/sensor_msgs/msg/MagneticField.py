import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class MagneticField(tinyros.Message):
    __slots__ = ['header','magnetic_field','magnetic_field_covariance']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Vector3','float64[9]']

    def __init__(self):
        super(MagneticField, self).__init__()
        self.header = std_msgs.msg.Header()
        self.magnetic_field = geometry_msgs.msg.Vector3()
        self.magnetic_field_covariance = [0. for x in range(0, 9)]

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.magnetic_field.serialize(buff)
        try:
            for i in range(0, 9):
                try:
                    struct_magnetic_field_covariancei = struct.Struct("<d")
                    buff.write(struct_magnetic_field_covariancei.pack(self.magnetic_field_covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.magnetic_field.deserialize(buff[offset:])
        try:
            self.magnetic_field_covariance = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_magnetic_field_covariancei = struct.Struct("<d")
                    (self.magnetic_field_covariance[i],) = struct_magnetic_field_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.magnetic_field.serializedLength()
        for i in range(0, 9):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"magnetic_field": {"'
        string_echo += self.magnetic_field.echo()
        string_echo += '}, '
        string_echo += '"magnetic_field_covariance": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"magnetic_field_covariance%s": %s'%(i,magnetic_field_covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"magnetic_field_covariance%s": %s'%(i,magnetic_field_covariance[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/MagneticField"

    def getMD5(self):
        return "f8e051d776de1349146122759c66db92"

_struct_I = struct.Struct('<I')

