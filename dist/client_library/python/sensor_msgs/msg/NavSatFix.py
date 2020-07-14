import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class NavSatFix(tinyros.Message):
    __slots__ = ['header','status','latitude','longitude','altitude','position_covariance','position_covariance_type']
    _slot_types = ['std_msgs.msg.Header','sensor_msgs.msg.NavSatStatus','float64','float64','float64','float64[9]','uint8']

    COVARIANCE_TYPE_UNKNOWN =  0
    COVARIANCE_TYPE_APPROXIMATED =  1
    COVARIANCE_TYPE_DIAGONAL_KNOWN =  2
    COVARIANCE_TYPE_KNOWN =  3

    def __init__(self):
        super(NavSatFix, self).__init__()
        self.header = std_msgs.msg.Header()
        self.status = sensor_msgs.msg.NavSatStatus()
        self.latitude = 0.
        self.longitude = 0.
        self.altitude = 0.
        self.position_covariance = [0. for x in range(0, 9)]
        self.position_covariance_type = 0

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        offset += self.status.serialize(buff)
        try:
            struct_latitude = struct.Struct("<d")
            buff.write(struct_latitude.pack(self.latitude))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_longitude = struct.Struct("<d")
            buff.write(struct_longitude.pack(self.longitude))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_altitude = struct.Struct("<d")
            buff.write(struct_altitude.pack(self.altitude))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            for i in range(0, 9):
                try:
                    struct_position_covariancei = struct.Struct("<d")
                    buff.write(struct_position_covariancei.pack(self.position_covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_position_covariance_type = struct.Struct("<B")
            buff.write(struct_position_covariance_type.pack(self.position_covariance_type))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        offset += self.status.deserialize(buff[offset:])
        try:
            struct_latitude = struct.Struct("<d")
            (self.latitude,) = struct_latitude.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_longitude = struct.Struct("<d")
            (self.longitude,) = struct_longitude.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_altitude = struct.Struct("<d")
            (self.altitude,) = struct_altitude.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            self.position_covariance = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_position_covariancei = struct.Struct("<d")
                    (self.position_covariance[i],) = struct_position_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_position_covariance_type = struct.Struct("<B")
            (self.position_covariance_type,) = struct_position_covariance_type.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += self.status.serializedLength()
        length += 8
        length += 8
        length += 8
        for i in range(0, 9):
            length += 8
        length += 1
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"status": {"'
        string_echo += self.status.echo()
        string_echo += '}, '
        string_echo += '"latitude": %s'%latitude
        string_echo += ', '
        string_echo += '"longitude": %s'%longitude
        string_echo += ', '
        string_echo += '"altitude": %s'%altitude
        string_echo += ', '
        string_echo += '"position_covariance": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"position_covariance%s": %s'%(i,position_covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"position_covariance%s": %s'%(i,position_covariance[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"position_covariance_type": %s'%position_covariance_type
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/NavSatFix"

    def getMD5(self):
        return "adc27ca9d8634ed087021b82f3c43576"

_struct_I = struct.Struct('<I')

