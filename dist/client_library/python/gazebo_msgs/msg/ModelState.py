import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg
import geometry_msgs.msg

class ModelState(tinyros.Message):
    __slots__ = ['model_name','pose','twist','reference_frame']
    _slot_types = ['string','geometry_msgs.msg.Pose','geometry_msgs.msg.Twist','string']

    def __init__(self):
        super(ModelState, self).__init__()
        self.model_name = ''
        self.pose = geometry_msgs.msg.Pose()
        self.twist = geometry_msgs.msg.Twist()
        self.reference_frame = ''

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.model_name
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
        offset += self.pose.serialize(buff)
        offset += self.twist.serialize(buff)
        try:
            _x = self.reference_frame
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
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_model_name,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.model_name = buff[offset:(offset + length_model_name)].decode('utf-8')
            else:
                self.model_name = buff[offset:(offset + length_model_name)]
            offset += length_model_name
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        offset += self.pose.deserialize(buff[offset:])
        offset += self.twist.deserialize(buff[offset:])
        try:
            (length_reference_frame,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.reference_frame = buff[offset:(offset + length_reference_frame)].decode('utf-8')
            else:
                self.reference_frame = buff[offset:(offset + length_reference_frame)]
            offset += length_reference_frame
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        model_name_x = self.model_name
        model_name_length = len(model_name_x)
        if python3 or type(model_name_x) == unicode:
            model_name_x = model_name_x.encode('utf-8')
            model_name_length = len(model_name_x)
        length += 4
        length += model_name_length
        length += self.pose.serializedLength()
        length += self.twist.serializedLength()
        reference_frame_x = self.reference_frame
        reference_frame_length = len(reference_frame_x)
        if python3 or type(reference_frame_x) == unicode:
            reference_frame_x = reference_frame_x.encode('utf-8')
            reference_frame_length = len(reference_frame_x)
        length += 4
        length += reference_frame_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"model_name": "%s"'%model_name
        string_echo += '", '
        string_echo += '"pose": {"'
        string_echo += self.pose.echo()
        string_echo += '}, '
        string_echo += '"twist": {"'
        string_echo += self.twist.echo()
        string_echo += '}, '
        string_echo += '"reference_frame": "%s"'%reference_frame
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ModelState"

    def getMD5(self):
        return "dee4d802363b4d6bd1ed61e20c2c4635"

_struct_I = struct.Struct('<I')

