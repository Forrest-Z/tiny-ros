import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class CameraInfo(tinyros.Message):
    __slots__ = ['header','height','width','distortion_model','D','K','R','P','binning_x','binning_y','roi']
    _slot_types = ['std_msgs.msg.Header','uint32','uint32','string','float64[]','float64[9]','float64[9]','float64[12]','uint32','uint32','sensor_msgs.msg.RegionOfInterest']

    def __init__(self):
        super(CameraInfo, self).__init__()
        self.header = std_msgs.msg.Header()
        self.height = 0
        self.width = 0
        self.distortion_model = ''
        self.D = []
        self.K = [0. for x in range(0, 9)]
        self.R = [0. for x in range(0, 9)]
        self.P = [0. for x in range(0, 12)]
        self.binning_x = 0
        self.binning_y = 0
        self.roi = sensor_msgs.msg.RegionOfInterest()

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
            _x = self.distortion_model
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
            D_length = len(self.D)
            buff.write(_struct_I.pack(D_length))
            offset += 4
            for i in range(0, D_length):
                try:
                    struct_Di = struct.Struct("<d")
                    buff.write(struct_Di.pack(self.D[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            for i in range(0, 9):
                try:
                    struct_Ki = struct.Struct("<d")
                    buff.write(struct_Ki.pack(self.K[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            for i in range(0, 9):
                try:
                    struct_Ri = struct.Struct("<d")
                    buff.write(struct_Ri.pack(self.R[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            for i in range(0, 12):
                try:
                    struct_Pi = struct.Struct("<d")
                    buff.write(struct_Pi.pack(self.P[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.binning_x))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.binning_y))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        offset += self.roi.serialize(buff)
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
            (length_distortion_model,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.distortion_model = buff[offset:(offset + length_distortion_model)].decode('utf-8')
            else:
                self.distortion_model = buff[offset:(offset + length_distortion_model)]
            offset += length_distortion_model
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (D_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.D = [0. for x in range(0, D_length)]
            offset += 4
            for i in range(0, D_length):
                try:
                    struct_Di = struct.Struct("<d")
                    (self.D[i],) = struct_Di.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            self.K = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_Ki = struct.Struct("<d")
                    (self.K[i],) = struct_Ki.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            self.R = [0. for x in range(0, 9)]
            for i in range(0, 9):
                try:
                    struct_Ri = struct.Struct("<d")
                    (self.R[i],) = struct_Ri.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            self.P = [0. for x in range(0, 12)]
            for i in range(0, 12):
                try:
                    struct_Pi = struct.Struct("<d")
                    (self.P[i],) = struct_Pi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.binning_x,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.binning_y,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.roi.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        distortion_model_x = self.distortion_model
        distortion_model_length = len(distortion_model_x)
        if python3 or type(distortion_model_x) == unicode:
            distortion_model_x = distortion_model_x.encode('utf-8')
            distortion_model_length = len(distortion_model_x)
        length += 4
        length += distortion_model_length
        D_length = len(self.D)
        length += 4
        for i in range(0, D_length):
            length += 8
        for i in range(0, 9):
            length += 8
        for i in range(0, 9):
            length += 8
        for i in range(0, 12):
            length += 8
        length += 4
        length += 4
        length += self.roi.serializedLength()
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
        string_echo += '"distortion_model": "%s"'%distortion_model
        string_echo += '", '
        string_echo += '"D": ['
        D_length = len(self.D)
        for i in range(0, D_length):
            if i == (D_length - 1): 
                string_echo += '{"D%s": %s'%(i,D[i])
                string_echo += '}'
            else:
                string_echo += '{"D%s": %s'%(i,D[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"K": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"K%s": %s'%(i,K[i])
                string_echo += '}'
            else:
                string_echo += '{"K%s": %s'%(i,K[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"R": ['
        for i in range(0, 9):
            if i == (9 - 1): 
                string_echo += '{"R%s": %s'%(i,R[i])
                string_echo += '}'
            else:
                string_echo += '{"R%s": %s'%(i,R[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"P": ['
        for i in range(0, 12):
            if i == (12 - 1): 
                string_echo += '{"P%s": %s'%(i,P[i])
                string_echo += '}'
            else:
                string_echo += '{"P%s": %s'%(i,P[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"binning_x": %s'%binning_x
        string_echo += ', '
        string_echo += '"binning_y": %s'%binning_y
        string_echo += ', '
        string_echo += '"roi": {"'
        string_echo += self.roi.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/CameraInfo"

    def getMD5(self):
        return "57d2553deec0a7842f00837f40032798"

_struct_I = struct.Struct('<I')

