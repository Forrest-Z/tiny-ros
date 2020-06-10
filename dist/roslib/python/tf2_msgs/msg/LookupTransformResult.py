import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import tf2_msgs.msg
import geometry_msgs.msg

class LookupTransformResult(tinyros.Message):
    __slots__ = ['transform','error']
    _slot_types = ['geometry_msgs.msg.TransformStamped','tf2_msgs.msg.TF2Error']

    def __init__(self):
        super(LookupTransformResult, self).__init__()
        self.transform = geometry_msgs.msg.TransformStamped()
        self.error = tf2_msgs.msg.TF2Error()

    def serialize(self, buff):
        offset = 0
        offset += self.transform.serialize(buff)
        offset += self.error.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.transform.deserialize(buff[offset:])
        offset += self.error.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.transform.serializedLength()
        length += self.error.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"transform": {"'
        string_echo += self.transform.echo()
        string_echo += '}, '
        string_echo += '"error": {"'
        string_echo += self.error.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "tf2_msgs/LookupTransformResult"

    def getMD5(self):
        return "7be4fc6719f512bc94491db1ccda6aee"

_struct_I = struct.Struct('<I')

