import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import geometry_msgs.msg

class PoseWithCovariance(tinyros.Message):
    __slots__ = ['pose','covariance']
    _slot_types = ['geometry_msgs.msg.Pose','float64[36]']

    def __init__(self):
        super(PoseWithCovariance, self).__init__()
        self.pose = geometry_msgs.msg.Pose()
        self.covariance = [0. for x in range(0, 36)]

    def serialize(self, buff):
        offset = 0
        offset += self.pose.serialize(buff)
        try:
            for i in range(0, 36):
                try:
                    struct_covariancei = struct.Struct("<d")
                    buff.write(struct_covariancei.pack(self.covariance[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.pose.deserialize(buff[offset:])
        try:
            self.covariance = [0. for x in range(0, 36)]
            for i in range(0, 36):
                try:
                    struct_covariancei = struct.Struct("<d")
                    (self.covariance[i],) = struct_covariancei.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.pose.serializedLength()
        for i in range(0, 36):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"pose": {"'
        string_echo += self.pose.echo()
        string_echo += '}, '
        string_echo += '"covariance": ['
        for i in range(0, 36):
            if i == (36 - 1): 
                string_echo += '{"covariance%s": %s'%(i,covariance[i])
                string_echo += '}'
            else:
                string_echo += '{"covariance%s": %s'%(i,covariance[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "geometry_msgs/PoseWithCovariance"

    def getMD5(self):
        return "054c6283d50e78f8d9358aaaee5f4c1b"

_struct_I = struct.Struct('<I')

