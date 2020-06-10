import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg

class GetMapResult(tinyros.Message):
    __slots__ = ['map']
    _slot_types = ['nav_msgs.msg.OccupancyGrid']

    def __init__(self):
        super(GetMapResult, self).__init__()
        self.map = nav_msgs.msg.OccupancyGrid()

    def serialize(self, buff):
        offset = 0
        offset += self.map.serialize(buff)
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.map.deserialize(buff[offset:])
        return offset

    def serializedLength(self):
        length = 0
        length += self.map.serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"map": {"'
        string_echo += self.map.echo()
        string_echo += '}'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GetMapResult"

    def getMD5(self):
        return "dd8eb0759b1a400b141d7f3238732c4d"

_struct_I = struct.Struct('<I')

