import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class Imu(tinyros.Message):
    __slots__ = ['header','orientation','orientation_covariance','angular_velocity','angular_velocity_covariance','linear_acceleration','linear_acceleration_covariance']
    _slot_types = ['std_msgs.msg.Header','geometry_msgs.msg.Quaternion','float64[9]','geometry_msgs.msg.Vector3','float64[9]','geometry_msgs.msg.Vector3','float64[9]']

    def __init__(self):
        super(Imu, self).__init__()
        self.header = std_msgs.msg.Header()
        self.orientation = geometry_msgs.msg.Quaternion()
        self.orientation_covariance = [0. for x in range(0, 9)]
        self.angular_velocity = geometry_msgs.msg.Vector3()
        self.angular_velocity_covariance = [0. for x in range(0, 9)]
        self.linear_acceleration = geometry_msgs.msg.Vector3()
        self.linear_acceleration_covariance = [0. for x in range(0, 9)]

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.orientation.serialize(buff)
        try:
            for i in range(0, 9):
                try:
                    struct_orientation_covariancei = struct.Struct("<d")
                    buff.write(struct_orientation_covariancei.pack(self.orientation_covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.angular_velocity.serialize(buff)
        try:
            for i in range(0, 9):
                try:
                    struct_angular_velocity_covariancei = struct.Struct("<d")
                    buff.write(struct_angular_velocity_covariancei.pack(self.angular_velocity_covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.linear_acceleration.serialize(buff)
        try:
            for i in range(0, 9):
                try:
                    struct_linear_acceleration_covariancei = struct.Struct("<d")
                    buff.write(struct_linear_acceleration_covariancei.pack(self.linear_acceleration_covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.orientation.deserialize(buff[offset:])
        try:
            self.orientation_covariance = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_orientation_covariancei = struct.Struct("<d")
                    (self.orientation_covariance[i],) = struct_orientation_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.angular_velocity.deserialize(buff[offset:])
        try:
            self.angular_velocity_covariance = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_angular_velocity_covariancei = struct.Struct("<d")
                    (self.angular_velocity_covariance[i],) = struct_angular_velocity_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.linear_acceleration.deserialize(buff[offset:])
        try:
            self.linear_acceleration_covariance = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_linear_acceleration_covariancei = struct.Struct("<d")
                    (self.linear_acceleration_covariance[i],) = struct_linear_acceleration_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.orientation.serializedLength()
        for i in range(0, 9):
            length += 8
        length += self.angular_velocity.serializedLength()
        for i in range(0, 9):
            length += 8
        length += self.linear_acceleration.serializedLength()
        for i in range(0, 9):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"orientation": {"'
        string_echo += self.orientation.echo()
        string_echo += '}, '
        string_echo += '"orientation_covariance": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"orientation_covariance%s": %s'%(i,orientation_covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"orientation_covariance%s": %s'%(i,orientation_covariance[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"angular_velocity": {"'
        string_echo += self.angular_velocity.echo()
        string_echo += '}, '
        string_echo += '"angular_velocity_covariance": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"angular_velocity_covariance%s": %s'%(i,angular_velocity_covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"angular_velocity_covariance%s": %s'%(i,angular_velocity_covariance[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"linear_acceleration": {"'
        string_echo += self.linear_acceleration.echo()
        string_echo += '}, '
        string_echo += '"linear_acceleration_covariance": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"linear_acceleration_covariance%s": %s'%(i,linear_acceleration_covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"linear_acceleration_covariance%s": %s'%(i,linear_acceleration_covariance[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/Imu"

    def getMD5(self):
        return "a42c1ab94665a5807834c0ea19a6d16a"

_struct_I = struct.Struct('<I')

