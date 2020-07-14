import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class MultiEchoLaserScan(tinyros.Message):
    __slots__ = ['header','angle_min','angle_max','angle_increment','time_increment','scan_time','range_min','range_max','ranges','intensities']
    _slot_types = ['std_msgs.msg.Header','float32','float32','float32','float32','float32','float32','float32','sensor_msgs.msg.LaserEcho[]','sensor_msgs.msg.LaserEcho[]']

    def __init__(self):
        super(MultiEchoLaserScan, self).__init__()
        self.header = std_msgs.msg.Header()
        self.angle_min = 0.
        self.angle_max = 0.
        self.angle_increment = 0.
        self.time_increment = 0.
        self.scan_time = 0.
        self.range_min = 0.
        self.range_max = 0.
        self.ranges = []
        self.intensities = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_angle_min = struct.Struct("<f")
            buff.write(struct_angle_min.pack(self.angle_min))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_angle_max = struct.Struct("<f")
            buff.write(struct_angle_max.pack(self.angle_max))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_angle_increment = struct.Struct("<f")
            buff.write(struct_angle_increment.pack(self.angle_increment))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_time_increment = struct.Struct("<f")
            buff.write(struct_time_increment.pack(self.time_increment))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_scan_time = struct.Struct("<f")
            buff.write(struct_scan_time.pack(self.scan_time))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_range_min = struct.Struct("<f")
            buff.write(struct_range_min.pack(self.range_min))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_range_max = struct.Struct("<f")
            buff.write(struct_range_max.pack(self.range_max))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            ranges_length = len(self.ranges)
            buff.write(_struct_I.pack(ranges_length))
            offset += 4
            for i in range(0, ranges_length):
                offset += self.ranges[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            intensities_length = len(self.intensities)
            buff.write(_struct_I.pack(intensities_length))
            offset += 4
            for i in range(0, intensities_length):
                offset += self.intensities[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_angle_min = struct.Struct("<f")
            (self.angle_min,) = struct_angle_min.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_angle_max = struct.Struct("<f")
            (self.angle_max,) = struct_angle_max.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_angle_increment = struct.Struct("<f")
            (self.angle_increment,) = struct_angle_increment.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_time_increment = struct.Struct("<f")
            (self.time_increment,) = struct_time_increment.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_scan_time = struct.Struct("<f")
            (self.scan_time,) = struct_scan_time.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_range_min = struct.Struct("<f")
            (self.range_min,) = struct_range_min.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_range_max = struct.Struct("<f")
            (self.range_max,) = struct_range_max.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (ranges_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.ranges = [sensor_msgs.msg.LaserEcho() for x in range(0, ranges_length)]
            offset += 4
            for i in range(0, ranges_length):
                offset += self.ranges[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (intensities_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.intensities = [sensor_msgs.msg.LaserEcho() for x in range(0, intensities_length)]
            offset += 4
            for i in range(0, intensities_length):
                offset += self.intensities[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        ranges_length = len(self.ranges)
        length += 4
        for i in range(0, ranges_length):
            length += self.ranges[i].serializedLength()
        intensities_length = len(self.intensities)
        length += 4
        for i in range(0, intensities_length):
            length += self.intensities[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"angle_min": %s'%angle_min
        string_echo += ', '
        string_echo += '"angle_max": %s'%angle_max
        string_echo += ', '
        string_echo += '"angle_increment": %s'%angle_increment
        string_echo += ', '
        string_echo += '"time_increment": %s'%time_increment
        string_echo += ', '
        string_echo += '"scan_time": %s'%scan_time
        string_echo += ', '
        string_echo += '"range_min": %s'%range_min
        string_echo += ', '
        string_echo += '"range_max": %s'%range_max
        string_echo += ', '
        string_echo += '"ranges": ['
        ranges_length = len(self.ranges)
        for i in range(0, ranges_length):
            if i == (ranges_length - 1): 
                string_echo += '{"ranges%s": {'%i
                string_echo += self.ranges[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"ranges%s": {'%i
                string_echo += self.ranges[i].echo()
                string_echo += '}}, '
        string_echo += '], '
        string_echo += '"intensities": ['
        intensities_length = len(self.intensities)
        for i in range(0, intensities_length):
            if i == (intensities_length - 1): 
                string_echo += '{"intensities%s": {'%i
                string_echo += self.intensities[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"intensities%s": {'%i
                string_echo += self.intensities[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/MultiEchoLaserScan"

    def getMD5(self):
        return "92f3933b4fa486e3889b461437899bf5"

_struct_I = struct.Struct('<I')

