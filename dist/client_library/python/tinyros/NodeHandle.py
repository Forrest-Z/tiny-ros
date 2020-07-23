import os
import sys
import socket
import threading
import time
import struct
import tinyros
import std_msgs.msg
import tinyros_msgs.msg
import sys
import signal
import uuid
python3 = True if sys.hexversion > 0x03000000 else False

try:
    from cStringIO import StringIO  # Python 2.x
except ImportError:
    from io import BytesIO as StringIO  # Python 3.x

class SpinObject(object):
    __slots__ = ['sub','message_in']

    def __init__(self):
        self.sub = None
        self.message_in = ''

class NodeHandleBase(object):
    PROTOCOL_VER        = 0xb9
    TINYROS_LOG_TOPIC   = '/tinyrosout'
    MAX_SUBSCRIBERS = 100
    MAX_PUBLISHERS  = 100
    INPUT_SIZE  = 64*1024
    OUTPUT_SIZE = 64*1024

    SYNC_TIME_SCOPE = 10 # milliseconds
    
    global_ip_addr = "127.0.0.1"
    
    global_node_name = ""
    
    __slots__ = ['ip_', 'name_', 'spin_', 'publishers_', 'subscribers_', 'hardware_', 'loghd_', 'keepalive_', 'keepalive_thread_', 'mutex_']
    
    def __init__(self):
        signal.signal(signal.SIGINT, self.signal_handler)
        self.ip_ = NodeHandleBase.global_ip_addr
        self.name_ = NodeHandleBase.global_node_name
        self.spin_ = True
        self.publishers_ = {}
        self.subscribers_ = {}
        self.hardware_ = None
        self.loghd_ = None
        self.keepalive_ = False
        self.keepalive_thread_ = None
        self.mutex_ = threading.Lock()

    def signal_handler(self, signal, frame):
        self.exit()
        os._exit(1)
        
    def initNode(self, name, ip):
        pass
    
    def spin(self):
        pass
    
    def exit(self):
        pass
    
    def keepalive(self):
        pass
   
    def ok(self):
        return self.hardware_.connected()
    
    def publish(self, id, msg, islog=False):
        buff = StringIO()
        buff.seek(11)
        l = msg.serialize(buff)
        buff.seek(0)
        struct_data = struct.Struct("<B")
        buff.write(struct_data.pack(0xFF))
        buff.write(struct_data.pack(NodeHandle.PROTOCOL_VER))
        buff.write(struct_data.pack((l >> 0) & 0xFF))
        buff.write(struct_data.pack((l >> 8) & 0xFF))
        buff.write(struct_data.pack((l >> 16) & 0xFF))
        buff.write(struct_data.pack((l >> 24) & 0xFF))
        chk = 255 - ((((l >> 0) & 0xFF) + ((l >> 8) & 0xFF) + ((l >> 16) & 0xFF) + ((l >> 24) & 0xFF)) % 256)
        buff.write(struct_data.pack(chk))
        buff.write(struct_data.pack((id >> 0) & 0xFF))
        buff.write(struct_data.pack((id >> 8) & 0xFF))
        buff.write(struct_data.pack((id >> 16) & 0xFF))
        buff.write(struct_data.pack((id >> 24) & 0xFF))
        
        buff.seek(7)
        x = buff.read()
        size = len(x)
        chk = 0
        for i in range(0, size):
            (data,) = struct_data.unpack(x[i:i+1])
            chk += data
        chk = 255 - (chk % 256)
        buff.seek(l + 11)
        buff.write(struct_data.pack(chk))
        buff.seek(0)

        if l <= NodeHandleBase.OUTPUT_SIZE: 
            l = -1
            if not islog:
                if self.hardware_ != None:
                    l = self.hardware_.write(buff.read())
            else:
                if self.loghd_ != None:
                    l = self.loghd_.write(buff.read())
            return l
        else:
            return -2

    def syncTime(self, buff):
        t = tinyros_msgs.msg.SyncTime()
        t.deserialize(buff)
        now = long(tinyros.Time.now().toMSec())
        tinyros.Time.mutex.acquire()
        scope = now - tinyros.Time.time_last - t.tick
        if (tinyros.Time.time_start == 0) or (scope >= 0 and scope <= NodeHandleBase.SYNC_TIME_SCOPE):
            tinyros.Time.time_dds = long(t.data.toMSec())
            tinyros.Time.time_start = now
        tinyros.Time.time_last = now
        tinyros.Time.mutex.release()

    def negotiateTopics(self, t):
        ti = tinyros_msgs.msg.TopicInfo()
        ti.topic_id = t.id_
        ti.topic_name = t.topic_
        ti.message_type = t.getMsgType()
        ti.md5sum = t.getMsgMD5()
        ti.node = self.name_
        if isinstance(t, tinyros.Publisher):
            ti.buffer_size = NodeHandleBase.OUTPUT_SIZE
        else:
            ti.buffer_size = NodeHandleBase.INPUT_SIZE
        self.publish(t.getEndpointType(), ti)

    def negotiateTopicsAll(self):
        for i in self.publishers_:
            self.negotiateTopics(self.publishers_[i])
        for i in self.subscribers_:
            self.negotiateTopics(self.subscribers_[i])
     
    def tinyros_main_loop(name, nh, ip):
        while nh.spin_ :
            if not nh.ok():
                nh.initNode(name, ip)
            if nh.ok():
                nh.spin()
            time.sleep(1)
    tinyros_main_loop = staticmethod(tinyros_main_loop)
 
class NodeHandle(NodeHandleBase):
    MODE_FIRST_FF       = 0
    MODE_PROTOCOL_VER   = 1
    MODE_SIZE_L         = 2
    MODE_SIZE_L1        = 3
    MODE_SIZE_H         = 4
    MODE_SIZE_H1        = 5
    MODE_SIZE_CHECKSUM  = 6
    MODE_TOPIC_L        = 7
    MODE_TOPIC_L1       = 8
    MODE_TOPIC_H        = 9
    MODE_TOPIC_H1       = 10
    MODE_MESSAGE        = 11
    MODE_MSG_CHECKSUM   = 12

    global_nh = None
    global_main_loop_init = False
    global_main_loop_mutex = threading.Lock()

    __slots__ = ['topic_list_', 'topic_list_recieved_', 'service_list_', 'service_list_recieved_', \
                'spin_thread_pool_', 'spin_log_thread_pool_', 'spin_srv_thread_pool_']

    def __init__(self):
        super(NodeHandle, self).__init__()
        self.topic_list_ = ''
        self.topic_list_recieved_ = False
        self.service_list_ = ''
        self.service_list_recieved_ = False
        self.spin_thread_pool_ = tinyros.ThreadPool(3)
        self.spin_log_thread_pool_ = tinyros.ThreadPool(1)
        self.spin_srv_thread_pool_ = tinyros.ThreadPool(3)
        self.hardware_ = tinyros.Hardware()
        self.loghd_ = tinyros.Hardware()
        
    def initNode(self, name, ip):
        self.name_ = name
        self.ip_ = ip

        if self.hardware_.init(self.ip_):
            msg = std_msgs.msg.String()
            msg.data = name
            self.publish(tinyros_msgs.msg.TopicInfo.ID_SESSION_ID, msg)

        if self.keepalive_ == False:
            self.keepalive_ = True
            self.keepalive_thread_ = threading.Thread(target=self.keepalive)
            self.keepalive_thread_.start()

        return self.hardware_.connected()

    def keepalive(self):
        while self.keepalive_:
            conn = self.loghd_.connected()
            if conn == False:
                if self.loghd_.init(self.ip_):
                    msg = std_msgs.msg.String()
                    msg.data = self.name_ + "_log"
                    self.publish(tinyros_msgs.msg.TopicInfo.ID_SESSION_ID, msg, True)
                time.sleep(1)
            else:
                self.loghd_.read(200)
    
    def spin(self):
        mode = self.MODE_FIRST_FF
        bytes = 0
        topic = 0
        need_len = 1
        checksum = 0
        message_in = b'' if python3 else ''
        
        struct_B = struct.Struct("<B")

        self.mutex_.acquire()
        for i in self.publishers_:
            self.publishers_[i].negotiated_ = False
        for i in self.subscribers_:
            self.subscribers_[i].negotiated_ = False
        self.mutex_.release()

        if not self.hardware_.connected():
            return -1
        else:
            self.negotiateTopicsAll()

        while self.spin_ and self.hardware_.connected():
            if need_len > NodeHandleBase.INPUT_SIZE:
                need_len = 1
                mode = NodeHandle.MODE_FIRST_FF
                continue
            
            data, rv = self.hardware_.read(need_len)
            
            if rv < 0:
                return -1
            
            for i in range(0, rv):
                (sum,) = struct_B.unpack(data[i:i+1])
                checksum += sum

            if mode == NodeHandle.MODE_MESSAGE:
                for i in range(0, rv):
                    message_in = message_in + data[i:i+1]
                    bytes -= 1
                if bytes == 0:
                    need_len = 1
                    mode = NodeHandle.MODE_MSG_CHECKSUM
                else:
                    need_len = bytes
            elif mode == NodeHandle.MODE_FIRST_FF:
                message_in = b'' if python3 else ''
                (sum,) = struct_B.unpack(data[0:1])
                if sum == 0xff:
                    mode += 1
            elif mode == NodeHandle.MODE_PROTOCOL_VER:
                (sum,) = struct_B.unpack(data[0:1])
                if sum == NodeHandleBase.PROTOCOL_VER:
                    mode += 1
                else:
                    mode = NodeHandle.MODE_FIRST_FF
            elif mode == self.MODE_SIZE_L:
                (checksum,) = struct_B.unpack(data[0:1])
                bytes = checksum
                mode += 1
            elif mode == NodeHandle.MODE_SIZE_L1:
                (sum,) = struct_B.unpack(data[0:1])
                bytes += sum << 8
                mode += 1
            elif mode == NodeHandle.MODE_SIZE_H:
                (sum,) = struct_B.unpack(data[0:1])
                bytes += sum << 16
                mode += 1
            elif mode == NodeHandle.MODE_SIZE_H1:
                (sum,) = struct_B.unpack(data[0:1])
                bytes += sum << 24
                mode += 1
            elif mode == NodeHandle.MODE_SIZE_CHECKSUM:
                if (checksum % 256) == 255:
                    mode += 1
                else:
                    mode = NodeHandle.MODE_FIRST_FF
            elif mode == NodeHandle.MODE_TOPIC_L:
                (topic,) = struct_B.unpack(data[0:1])
                mode += 1
                checksum = topic
            elif mode == NodeHandle.MODE_TOPIC_L1:
                (sum,) = struct_B.unpack(data[0:1])
                topic += sum << 8
                mode += 1
            elif mode == NodeHandle.MODE_TOPIC_H:
                (sum,) = struct_B.unpack(data[0:1])
                topic += sum << 16
                mode += 1
            elif mode == NodeHandle.MODE_TOPIC_H1:
                (sum,) = struct_B.unpack(data[0:1])
                topic += sum << 24
                mode = NodeHandle.MODE_MESSAGE
                if bytes == 0:
                    mode = NodeHandle.MODE_MSG_CHECKSUM
                else:
                    need_len = bytes
            elif mode == NodeHandle.MODE_MSG_CHECKSUM:
                mode = NodeHandle.MODE_FIRST_FF
                if (checksum % 256) == 255:
                    if topic == tinyros_msgs.msg.TopicInfo.ID_PUBLISHER:
                        self.negotiateTopicsAll()
                    elif topic == tinyros_msgs.msg.TopicInfo.ID_ROSTOPIC_REQUEST:
                        msg = std_msgs.msg.String()
                        msg.deserialize(message_in)
                        self.topic_list_ = msg.data
                        self.topic_list_recieved_ = True
                    elif topic == tinyros_msgs.msg.TopicInfo.ID_ROSSERVICE_REQUEST:
                        msg = std_msgs.msg.String()
                        msg.deserialize(message_in)
                        self.service_list_ = msg.data
                        self.service_list_recieved_ = True
                    elif topic == tinyros_msgs.msg.TopicInfo.ID_NEGOTIATED:
                        ti = tinyros_msgs.msg.TopicInfo()
                        ti.deserialize(message_in)
                        try:
                            self.publishers_[ti.topic_id].negotiated_ = ti.negotiated
                        except KeyError as ex: 
                            pass
                        try:
                            self.subscribers_[ti.topic_id].negotiated_ = ti.negotiated
                        except KeyError as ex: 
                            pass
                    elif topic == tinyros_msgs.msg.TopicInfo.ID_TIME:
                        self.syncTime(message_in)
                    else:
                        try:
                            obj = tinyros.SpinObject()
                            obj.sub = self.subscribers_[topic]
                            obj.message_in = message_in
                            if obj.sub.topic_ == NodeHandleBase.TINYROS_LOG_TOPIC:
                                self.spin_log_thread_pool_.schedule(obj)
                            else:
                                if obj.sub.srv_flag_ == True:
                                    self.spin_srv_thread_pool_.schedule(obj)
                                else:
                                    self.spin_thread_pool_.schedule(obj)
                        except: 
                            pass

    def exit(self):
        self.spin_ = False
        self.keepalive_ = False
        self.spin_thread_pool_.shutdown()
        self.spin_log_thread_pool_.shutdown()
        self.spin_srv_thread_pool_.shutdown()
        self.loghd_.close()
        self.hardware_.close()
   
    def advertise(self, p):
        if not isinstance(p, tinyros.Publisher):
            raise TypeError("Not tinyros.Publisher type")
        if len(self.publishers_) < NodeHandleBase.MAX_PUBLISHERS:
            self.mutex_.acquire() 
            p.id_ = len(self.publishers_) + 100 + NodeHandleBase.MAX_SUBSCRIBERS
            p.nh_ = self
            self.publishers_[p.id_] = p
            self.mutex_.release()
            self.negotiateTopics(p);
            return True
        else:
            return False

    def subscribe(self, s):
        if not isinstance(s, tinyros.Subscriber):
            raise TypeError("Not tinyros.Subscriber type")
        if len(self.subscribers_) < NodeHandleBase.MAX_SUBSCRIBERS:
            self.mutex_.acquire() 
            s.id_ = len(self.subscribers_) + 100
            self.subscribers_[s.id_] = s
            self.mutex_.release()
            self.negotiateTopics(s)
            return True
        else:
            return False
     
    def advertiseService(self, srv):
        if not isinstance(srv, tinyros.ServiceServer):
            raise TypeError("Not tinyros.ServiceServer type")
        if len(self.subscribers_) < NodeHandleBase.MAX_SUBSCRIBERS:
            self.mutex_.acquire() 
            srv.id_ = len(self.subscribers_) + 100
            self.subscribers_[srv.id_] = srv
            if len(self.publishers_) < NodeHandleBase.MAX_PUBLISHERS:
                srv.pub_.id_ = srv.id_
                srv.pub_.nh_ = self
                self.publishers_[srv.pub_.id_] = srv.pub_
                self.mutex_.release()
                self.negotiateTopics(srv)
                self.negotiateTopics(srv.pub_)
                return True
            else:
                self.mutex_.release()
                return False
        else:
            return False

    def serviceClient(self, srv):
        if not isinstance(srv, tinyros.ServiceClient):
            raise TypeError("Not tinyros.ServiceClient type")
        if len(self.subscribers_) < NodeHandleBase.MAX_SUBSCRIBERS:
            self.mutex_.acquire() 
            srv.id_ = len(self.subscribers_) + 100
            self.subscribers_[srv.id_] = srv
            if len(self.publishers_) < NodeHandleBase.MAX_PUBLISHERS:
                srv.pub_.id_ = srv.id_
                srv.pub_.nh_ = self
                self.publishers_[srv.pub_.id_] = srv.pub_
                self.mutex_.release()
                self.negotiateTopics(srv)
                self.negotiateTopics(srv.pub_)
                return True
            else:
                self.mutex_.release()
                return False
        else:
            return False
            
    def log(self, level, msg):
        if self.loghd_.connected():
            l = tinyros_msgs.msg.Log()
            l.level = level
            l.msg = "[" + self.name_ + "] " + msg
            self.publish(tinyros_msgs.msg.TopicInfo.ID_LOG, l, True)

    def getTopicList(self, timeout = 1000):
        msg = std_msgs.msg.String()
        self.topic_list_recieved_ = False
        self.publish(tinyros_msgs.msg.TopicInfo.ID_ROSTOPIC_REQUEST, msg)
        to = tinyros.Time.now().toMSec() + timeout
        while not self.topic_list_recieved_:
            now = tinyros.Time.now().toMSec()
            if now > to:
                print("Failed to get getTopicList: timeout expired")
                return ''
            time.sleep(0.1)
        return self.topic_list_

    def getServiceList(self, timeout = 1000):
        msg = std_msgs.msg.String()
        self.service_list_recieved_ = False
        self.publish(tinyros_msgs.msg.TopicInfo.ID_ROSSERVICE_REQUEST, msg)
        to = tinyros.Time.now().toMSec() + timeout
        while not self.service_list_recieved_:
            now = tinyros.Time.now().toMSec()
            if now > to:
                print("Failed to get getServiceList: timeout expired")
                return ''
            time.sleep(0.1)
        return self.service_list_

class NodeHandleUdp(NodeHandleBase):
    global_nh = None
    global_main_loop_init = False
    global_main_loop_mutex = threading.Lock()
    
    __slots__ = ['spin_thread_pool_']
    
    def __init__(self):
        super(NodeHandleUdp, self).__init__()
        self.hardware_ = tinyros.HardwareUdp()
        self.spin_thread_pool_ = tinyros.ThreadPool(3)

    def initNode(self, name, ip):
        self.name_ = name
        self.ip_ = ip
        if not self.keepalive_:
            self.keepalive_ = True
            self.keepalive_thread_ = threading.Thread(target=self.keepalive)
            self.keepalive_thread_.start()
        return self.hardware_.init(self.ip_)
    
    def spin(self):
        struct_B = struct.Struct("<B")
        while self.spin_ and self.hardware_.connected():
            if len(self.subscribers_) <= 0:
                time.sleep(1)
                continue
            (message_in, rv) = self.hardware_.read(NodeHandleBase.INPUT_SIZE)
            if NodeHandleBase.INPUT_SIZE >= rv and rv > 0:
                topic = index = bytes = checksum = 0
                (data,) = struct_B.unpack(message_in[index:index+1])
                if data != 0xff:
                    continue
                
                index += 1
                (data,) = struct_B.unpack(message_in[index:index+1])
                if data != NodeHandleBase.PROTOCOL_VER:
                    continue

                index += 1 
                (data,) = struct_B.unpack(message_in[index:index+1])
                bytes = checksum = data
                (data,) = struct_B.unpack(message_in[index + 1:index + 2])
                bytes += data << 8
                checksum += data
                (data,) = struct_B.unpack(message_in[index + 2:index + 3])
                bytes += data << 16
                checksum += data
                (data,) = struct_B.unpack(message_in[index + 3:index + 4])
                bytes += data << 24
                checksum += data
                (data,) = struct_B.unpack(message_in[index + 4:index + 5])
                checksum += data
                index += 5
                if (checksum % 256) != 255:
                    continue
                
                (data,) = struct_B.unpack(message_in[index:index + 1])
                topic = checksum = data
                (data,) = struct_B.unpack(message_in[index + 1:index + 2])
                topic += data << 8
                checksum += data
                (data,) = struct_B.unpack(message_in[index + 2:index + 3])
                topic += data << 16
                checksum += data
                (data,) = struct_B.unpack(message_in[index + 3:index + 4])
                topic += data << 24
                checksum += data
                index += 4

                if NodeHandleBase.INPUT_SIZE < (index + bytes + 1):
                    continue
                
                if bytes > 0:
                    for i in range(0, bytes + 1):
                        (data,) = struct_B.unpack(message_in[index + i:index + i + 1])
                        checksum += data
                else:
                    (data,) = struct_B.unpack(message_in[index:index + 1])
                    checksum += data
                if (checksum % 256) == 255:
                    try:
                        obj = tinyros.SpinObject()
                        obj.sub = self.subscribers_[topic]
                        obj.message_in = message_in[index:]
                        self.spin_thread_pool_.schedule(obj)
                    except: 
                        pass
        return True

    def exit(self):
        self.spin_ = False
        self.spin_thread_pool_.shutdown()
        self.hardware_.close()
    
    def keepalive(self):
        while self.keepalive_:
            self.negotiateTopicsAll()
            time.sleep(1)

    def generate_id(self):
        h = 0
        guid = str(uuid.uuid1()).replace("-", "")
        for i in guid:
            h = (31 * h + ord(i)) & 0xffffffff
        return h
    
    def advertise(self, p):
        if not isinstance(p, tinyros.Publisher):
            raise TypeError("Not tinyros.Publisher type")
        if len(self.publishers_) < NodeHandleBase.MAX_PUBLISHERS:
            self.mutex_.acquire() 
            p.id_ = self.generate_id()
            p.nh_ = self
            self.publishers_[p.id_] = p
            self.mutex_.release()
            self.negotiateTopics(p);
            return True
        else:
            return False

    def subscribe(self, s):
        if not isinstance(s, tinyros.Subscriber):
            raise TypeError("Not tinyros.Subscriber type")
        if len(self.subscribers_) < NodeHandleBase.MAX_SUBSCRIBERS:
            self.mutex_.acquire() 
            s.id_ = self.generate_id()
            self.subscribers_[s.id_] = s
            self.mutex_.release()
            self.negotiateTopics(s)
            return True
        else:
            return False

#####################################################################
# Global functions

NodeHandle.global_nh = NodeHandle()
NodeHandleUdp.global_nh = NodeHandleUdp()

def nh():
    if not NodeHandle.global_main_loop_init:
        NodeHandle.global_main_loop_mutex.acquire()
        if not NodeHandle.global_main_loop_init:
            thread = threading.Thread(target=NodeHandleBase.tinyros_main_loop, \
                     args = (NodeHandleBase.global_node_name, NodeHandle.global_nh, NodeHandleBase.global_ip_addr))
            thread.start()
            time.sleep(0.2)
            NodeHandle.global_main_loop_init = True
        NodeHandle.global_main_loop_mutex.release()
    return NodeHandle.global_nh

def udp():
    if not NodeHandleUdp.global_main_loop_init:
        NodeHandleUdp.global_main_loop_mutex.acquire()
        if not NodeHandleUdp.global_main_loop_init:
            thread = threading.Thread(target=NodeHandleBase.tinyros_main_loop, \
                     args = (NodeHandleBase.global_node_name, NodeHandleUdp.global_nh, NodeHandleBase.global_ip_addr))
            thread.start()
            time.sleep(0.2)
            NodeHandleUdp.global_main_loop_init = True
        NodeHandleUdp.global_main_loop_mutex.release()
    return NodeHandleUdp.global_nh
    
def init(name, ip = "127.0.0.1"):
    NodeHandleBase.global_node_name = name
    NodeHandleBase.global_ip_addr = ip
    tinyros.nh()
            
def logdebug(msg):
        nh().log(tinyros_msgs.msg.Log.ROSDEBUG, msg)
        
def loginfo(msg):
    nh().log(tinyros_msgs.msg.Log.ROSINFO, msg)
            
def logwarn(msg):
    nh().log(tinyros_msgs.msg.Log.ROSWARN, msg)
        
def logerror(msg):
    nh().log(tinyros_msgs.msg.Log.ROSERROR, msg)
        
def logfatal(msg):
    nh().log(tinyros_msgs.msg.Log.ROSFATAL, msg)
