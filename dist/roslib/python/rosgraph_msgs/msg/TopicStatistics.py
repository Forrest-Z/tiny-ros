import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import rosgraph_msgs.msg
import tinyros

class TopicStatistics(tinyros.Message):
    __slots__ = ['topic','node_pub','node_sub','window_start','window_stop','delivered_msgs','dropped_msgs','traffic','period_mean','period_stddev','period_max','stamp_age_mean','stamp_age_stddev','stamp_age_max']
    _slot_types = ['string','string','string','tinyros.Time','tinyros.Time','int32','int32','int32','tinyros.Duration','tinyros.Duration','tinyros.Duration','tinyros.Duration','tinyros.Duration','tinyros.Duration']

    def __init__(self):
        super(TopicStatistics, self).__init__()
        self.topic = ''
        self.node_pub = ''
        self.node_sub = ''
        self.window_start = tinyros.Time()
        self.window_stop = tinyros.Time()
        self.delivered_msgs = 0
        self.dropped_msgs = 0
        self.traffic = 0
        self.period_mean = tinyros.Duration()
        self.period_stddev = tinyros.Duration()
        self.period_max = tinyros.Duration()
        self.stamp_age_mean = tinyros.Duration()
        self.stamp_age_stddev = tinyros.Duration()
        self.stamp_age_max = tinyros.Duration()

    def serialize(self, buff):
        offset = 0
        try:
            _x = self.topic
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
            _x = self.node_pub
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
            _x = self.node_sub
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
            buff.write(_struct_I.pack(self.window_start.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.window_start.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.window_stop.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.window_stop.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_delivered_msgs = struct.Struct("<i")
            buff.write(struct_delivered_msgs.pack(self.delivered_msgs))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_dropped_msgs = struct.Struct("<i")
            buff.write(struct_dropped_msgs.pack(self.dropped_msgs))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_traffic = struct.Struct("<i")
            buff.write(struct_traffic.pack(self.traffic))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_mean.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_mean.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_stddev.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_stddev.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_max.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.period_max.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_mean.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_mean.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_stddev.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_stddev.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_max.sec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.stamp_age_max.nsec))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (length_topic,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.topic = buff[offset:(offset + length_topic)].decode('utf-8')
            else:
                self.topic = buff[offset:(offset + length_topic)]
            offset += length_topic
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_node_pub,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.node_pub = buff[offset:(offset + length_node_pub)].decode('utf-8')
            else:
                self.node_pub = buff[offset:(offset + length_node_pub)]
            offset += length_node_pub
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_node_sub,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.node_sub = buff[offset:(offset + length_node_sub)].decode('utf-8')
            else:
                self.node_sub = buff[offset:(offset + length_node_sub)]
            offset += length_node_sub
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.window_start.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.window_start.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.window_stop.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.window_stop.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_delivered_msgs = struct.Struct("<i")
            (self.delivered_msgs,) = struct_delivered_msgs.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_dropped_msgs = struct.Struct("<i")
            (self.dropped_msgs,) = struct_dropped_msgs.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_traffic = struct.Struct("<i")
            (self.traffic,) = struct_traffic.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_mean.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_mean.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_stddev.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_stddev.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_max.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.period_max.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_mean.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_mean.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_stddev.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_stddev.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_max.sec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.stamp_age_max.nsec,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        topic_x = self.topic
        topic_length = len(topic_x)
        if python3 or type(topic_x) == unicode:
            topic_x = topic_x.encode('utf-8')
            topic_length = len(topic_x)
        length += 4
        length += topic_length
        node_pub_x = self.node_pub
        node_pub_length = len(node_pub_x)
        if python3 or type(node_pub_x) == unicode:
            node_pub_x = node_pub_x.encode('utf-8')
            node_pub_length = len(node_pub_x)
        length += 4
        length += node_pub_length
        node_sub_x = self.node_sub
        node_sub_length = len(node_sub_x)
        if python3 or type(node_sub_x) == unicode:
            node_sub_x = node_sub_x.encode('utf-8')
            node_sub_length = len(node_sub_x)
        length += 4
        length += node_sub_length
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"topic": "%s"'%topic
        string_echo += '", '
        string_echo += '"node_pub": "%s"'%node_pub
        string_echo += '", '
        string_echo += '"node_sub": "%s"'%node_sub
        string_echo += '", '
        string_echo += '"window_start.sec": %s, '%window_start.sec
        string_echo += '"window_start.nsec": %s'%window_start.nsec
        string_echo += ', '
        string_echo += '"window_stop.sec": %s, '%window_stop.sec
        string_echo += '"window_stop.nsec": %s'%window_stop.nsec
        string_echo += ', '
        string_echo += '"delivered_msgs": %s'%delivered_msgs
        string_echo += ', '
        string_echo += '"dropped_msgs": %s'%dropped_msgs
        string_echo += ', '
        string_echo += '"traffic": %s'%traffic
        string_echo += ', '
        string_echo += '"period_mean.sec": %s, '%period_mean.sec
        string_echo += '"period_mean.nsec": %s'%period_mean.nsec
        string_echo += ', '
        string_echo += '"period_stddev.sec": %s, '%period_stddev.sec
        string_echo += '"period_stddev.nsec": %s'%period_stddev.nsec
        string_echo += ', '
        string_echo += '"period_max.sec": %s, '%period_max.sec
        string_echo += '"period_max.nsec": %s'%period_max.nsec
        string_echo += ', '
        string_echo += '"stamp_age_mean.sec": %s, '%stamp_age_mean.sec
        string_echo += '"stamp_age_mean.nsec": %s'%stamp_age_mean.nsec
        string_echo += ', '
        string_echo += '"stamp_age_stddev.sec": %s, '%stamp_age_stddev.sec
        string_echo += '"stamp_age_stddev.nsec": %s'%stamp_age_stddev.nsec
        string_echo += ', '
        string_echo += '"stamp_age_max.sec": %s, '%stamp_age_max.sec
        string_echo += '"stamp_age_max.nsec": %s'%stamp_age_max.nsec
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "rosgraph_msgs/TopicStatistics"

    def getMD5(self):
        return "8b30d3f22284a3bee7679b7194bd38a3"

_struct_I = struct.Struct('<I')

