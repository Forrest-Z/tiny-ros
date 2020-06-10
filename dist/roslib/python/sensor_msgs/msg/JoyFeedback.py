import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg

class JoyFeedback(tinyros.Message):
    __slots__ = ['type','id','intensity']
    _slot_types = ['uint8','uint8','float32']

    TYPE_LED =  0
    TYPE_RUMBLE =  1
    TYPE_BUZZER =  2

    def __init__(self):
        super(JoyFeedback, self).__init__()
        self.type = 0
        self.id = 0
        self.intensity = 0.

    def serialize(self, buff):
        offset = 0
        try:
            struct_type = struct.Struct("<B")
            buff.write(struct_type.pack(self.type))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_id = struct.Struct("<B")
            buff.write(struct_id.pack(self.id))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_intensity = struct.Struct("<f")
            buff.write(struct_intensity.pack(self.intensity))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_type = struct.Struct("<B")
            (self.type,) = struct_type.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_id = struct.Struct("<B")
            (self.id,) = struct_id.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_intensity = struct.Struct("<f")
            (self.intensity,) = struct_intensity.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        length += 1
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"type": %s'%type
        string_echo += ', '
        string_echo += '"id": %s'%id
        string_echo += ', '
        string_echo += '"intensity": %s'%intensity
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/JoyFeedback"

    def getMD5(self):
        return "206b65e86c8b195f816ccbe40b3568a2"

_struct_I = struct.Struct('<I')

