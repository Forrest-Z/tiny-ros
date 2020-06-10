import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class NavSatStatus(tinyros.Message):
    __slots__ = ['status','service']
    _slot_types = ['int8','uint16']

    STATUS_NO_FIX =   -1        
    STATUS_FIX =       0        
    STATUS_SBAS_FIX =  1        
    STATUS_GBAS_FIX =  2        
    SERVICE_GPS =      1
    SERVICE_GLONASS =  2
    SERVICE_COMPASS =  4      
    SERVICE_GALILEO =  8

    def __init__(self):
        super(NavSatStatus, self).__init__()
        self.status = 0
        self.service = 0

    def serialize(self, buff):
        offset = 0
        try:
            struct_status = struct.Struct("<b")
            buff.write(struct_status.pack(self.status))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_service = struct.Struct("<H")
            buff.write(struct_service.pack(self.service))
            offset += 2
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_status = struct.Struct("<b")
            (self.status,) = struct_status.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_service = struct.Struct("<H")
            (self.service,) = struct_service.unpack(buff[offset:(offset + 2)])
            offset += 2
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        length += 2
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"status": %s'%status
        string_echo += ', '
        string_echo += '"service": %s'%service
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/NavSatStatus"

    def getMD5(self):
        return "85ed59cf80532c1c15a2cf735d06279b"

_struct_I = struct.Struct('<I')

