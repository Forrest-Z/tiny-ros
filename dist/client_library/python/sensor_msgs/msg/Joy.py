import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class Joy(tinyros.Message):
    __slots__ = ['header','axes','buttons']
    _slot_types = ['std_msgs.msg.Header','float32[]','int32[]']

    def __init__(self):
        super(Joy, self).__init__()
        self.header = std_msgs.msg.Header()
        self.axes = []
        self.buttons = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            axes_length = len(self.axes)
            buff.write(_struct_I.pack(axes_length))
            offset += 4
            for i in range(0, axes_length):
                try:
                    struct_axesi = struct.Struct("<f")
                    buff.write(struct_axesi.pack(self.axes[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buttons_length = len(self.buttons)
            buff.write(_struct_I.pack(buttons_length))
            offset += 4
            for i in range(0, buttons_length):
                try:
                    struct_buttonsi = struct.Struct("<i")
                    buff.write(struct_buttonsi.pack(self.buttons[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            (axes_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.axes = [0. for x in range(0, axes_length)]
            offset += 4
            for i in range(0, axes_length):
                try:
                    struct_axesi = struct.Struct("<f")
                    (self.axes[i],) = struct_axesi.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (buttons_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.buttons = [0 for x in range(0, buttons_length)]
            offset += 4
            for i in range(0, buttons_length):
                try:
                    struct_buttonsi = struct.Struct("<i")
                    (self.buttons[i],) = struct_buttonsi.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        axes_length = len(self.axes)
        length += 4
        for i in range(0, axes_length):
            length += 4
        buttons_length = len(self.buttons)
        length += 4
        for i in range(0, buttons_length):
            length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"axes": ['
        axes_length = len(self.axes)
        for i in range(0, axes_length):
            if i == (axes_length - 1): 
                string_echo += '{"axes%s": %s'%(i,axes[i])
                string_echo += '}'
            else:
                string_echo += '{"axes%s": %s'%(i,axes[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"buttons": ['
        buttons_length = len(self.buttons)
        for i in range(0, buttons_length):
            if i == (buttons_length - 1): 
                string_echo += '{"buttons%s": %s'%(i,buttons[i])
                string_echo += '}'
            else:
                string_echo += '{"buttons%s": %s'%(i,buttons[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/Joy"

    def getMD5(self):
        return "da3438323dbe92a4d6e4658e06bf8da1"

_struct_I = struct.Struct('<I')

