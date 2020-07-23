package tinyros

import (
    "fmt"
    "time"
    "sync"
    "sync/atomic"
    "tiny_ros/std_msgs"
    "tiny_ros/tinyros_msgs"
    "tiny_ros/tinyros/time"
)

var (
    MAX_SUBSCRIBERS    int = 100
    MAX_PUBLISHERS     int = 100
    
    INPUT_SIZE         int = 64*1024
    OUTPUT_SIZE        int = 64*1024
    
    MODE_FIRST_FF      int = 0
    MODE_PROTOCOL_VER  int = 1
    MODE_SIZE_L        int = 2
    MODE_SIZE_L1       int = 3
    MODE_SIZE_H        int = 4
    MODE_SIZE_H1       int = 5
    MODE_SIZE_CHECKSUM int = 6
    MODE_TOPIC_L       int = 7
    MODE_TOPIC_L1      int = 8
    MODE_TOPIC_H       int = 9
    MODE_TOPIC_H1      int = 10
    MODE_MESSAGE       int = 11
    MODE_MSG_CHECKSUM  int = 12

    SYNC_TIME_SCOPE int64  = 10 // milliseconds

    TINYROS_LOG_TOPIC  string = "/tinyrosout"
)

type SpinObject struct {
    sub_ Subscriber_ `json:"sub"`
    message_in_ []byte `json:"message_in"`
}

type NodeHandleBase interface {
    initNode(node_name, ip string) (bool)
    keepalive()
    publish(id uint32, msg Msg, islog bool) (int)
    spin() (int)
    exit()
    ok() (bool)
    syncTime(data []byte)
}
 
type NodeHandle struct {
    ip_addr_ string `json:"ip_addr"`
    node_name_ string `json:"node_name"`
    spin_ bool `json:"spin"`
    keepalive_ bool `json:"keepalive"`
    hardware_ Hardware `json:"hardware"`
    loghd_ Hardware `json:"loghd"`
    publishers_ map[uint32]*Publisher `json:"publishers"`
    subscribers_ map[uint32]Subscriber_ `json:"subscribers"`
    message_out_ []byte `json:"message_out"`
    message_in_ []byte `json:"message_in"`
    mutex_ sync.Mutex `json:"mutex"`
    spin_thread_pool_ *ThreadPool `json:"spin_thread_pool"`
    spin_log_thread_pool_ *ThreadPool `json:"spin_log_thread_pool"`
    spin_srv_thread_pool_ *ThreadPool `json:"spin_srv_thread_pool"`
}

func (self *NodeHandle) publish(id uint32, msg Msg, islog bool) (int) {
    var chk int
    l := msg.Go_serializedLength() + 12
    if l > OUTPUT_SIZE {
        fmt.Println("NodeHandle.publish Overflow!")
        return -1
    }
    self.mutex_.Lock()
    l = msg.Go_serialize(self.message_out_[11:])
    self.message_out_[0] = byte(0xff)
    self.message_out_[1] = byte(0xb9)
    self.message_out_[2] = byte((l >> 0) & 0xFF)
    self.message_out_[3] = byte((l >> 8) & 0xFF)
    self.message_out_[4] = byte((l >> 16) & 0xFF)
    self.message_out_[5] = byte((l >> 24) & 0xFF)
    chk = int(self.message_out_[2])
    chk += int(self.message_out_[3])
    chk += int(self.message_out_[4])
    chk += int(self.message_out_[5])
    self.message_out_[6] = byte(255 - (chk % 256))
    self.message_out_[7] = byte((id >> 0) & 0xFF)
    self.message_out_[8] = byte((id >> 8) & 0xFF)
    self.message_out_[9] = byte((id >> 16) & 0xFF)
    self.message_out_[10] = byte((id >> 24) & 0xFF)

    chk = 0
    for i := 7; i < l + 11; i++ {
      chk += int(self.message_out_[i])
    }
    
    l += 11
    
    self.message_out_[l] = byte(255 - (chk % 256))

    l += 1

    if !islog {
        l = self.hardware_.write(self.message_out_[:l])
    } else {
        l = self.loghd_.write(self.message_out_[:l])
    }
    self.mutex_.Unlock()
    return l
}

func (self *NodeHandle) negotiateTopics_p(p *Publisher) {
    ti := tinyros_msgs.NewTopicInfo()
    ti.Go_topic_id = p.id_
    ti.Go_topic_name = p.topic_
    ti.Go_message_type = p.msg_.Go_getType()
    ti.Go_md5sum = p.msg_.Go_getMD5()
    ti.Go_buffer_size = int32(OUTPUT_SIZE)
    ti.Go_node = self.node_name_
    self.publish(p.Go_getEndpointType(), ti, false)
}
  
func (self *NodeHandle) negotiateTopics_s(s Subscriber_) {
    ti := tinyros_msgs.NewTopicInfo()
    ti.Go_topic_id = s.Go_id();
    ti.Go_topic_name = s.Go_topic();
    ti.Go_message_type = s.Go_getMsgType()
    ti.Go_md5sum = s.Go_getMsgMD5()
    ti.Go_buffer_size = int32(INPUT_SIZE)
    ti.Go_node = self.node_name_
    self.publish(s.Go_getEndpointType(), ti, false)
}

func (self *NodeHandle) negotiateTopics() {
    for _, p := range self.publishers_ {
        self.negotiateTopics_p(p)
    }
    for _, s :=  range self.subscribers_ {
        self.negotiateTopics_s(s)
    }
}

func (self *NodeHandle) ok() (bool) {
    return self.hardware_.connected()
}

func (self *NodeHandle) syncTime(data []byte) {
    t := tinyros_msgs.NewSyncTime()
    t.Go_deserialize(data)
    now := int64(rostime.TimeNow().Go_toMSec())
    rostime.Go_time_mutex.Lock()
    scope := now - rostime.Go_time_last - int64(t.Go_tick)
    if (rostime.Go_time_start == 0) || (scope >= 0 && scope <= SYNC_TIME_SCOPE) {
        rostime.Go_time_dds = int64(t.Go_data.Go_toMSec())
        rostime.Go_time_start = now
    }
    rostime.Go_time_last = now
    rostime.Go_time_mutex.Unlock()
}


//////////////////////////////////////////////////////////////////////////////////////
// NodeHandleTCP

type NodeHandleTCP struct {
    NodeHandle
}

func NewNodeHandleTCP() (*NodeHandleTCP) {
    newNodeHandleTCP := new(NodeHandleTCP)
    newNodeHandleTCP.hardware_ = NewHardwareTCP()
    newNodeHandleTCP.loghd_ = NewHardwareTCP()
    newNodeHandleTCP.keepalive_ = false
    newNodeHandleTCP.spin_ = true
    newNodeHandleTCP.ip_addr_ = ""
    newNodeHandleTCP.node_name_ = ""
    newNodeHandleTCP.publishers_ = make(map[uint32]*Publisher)
    newNodeHandleTCP.subscribers_ = make(map[uint32]Subscriber_)
    newNodeHandleTCP.spin_thread_pool_ = NewThreadPool(3)
    newNodeHandleTCP.spin_log_thread_pool_ = NewThreadPool(1)
    newNodeHandleTCP.spin_srv_thread_pool_ = NewThreadPool(3)
    newNodeHandleTCP.message_out_ = make([]byte, OUTPUT_SIZE)
    newNodeHandleTCP.message_in_ = make([]byte, INPUT_SIZE)
    return newNodeHandleTCP
}

func (self *NodeHandleTCP) keepalive() {
    buff := make([]byte, 200)
    for self.keepalive_ {
        if !self.loghd_.connected() {
            if self.loghd_.init(self.ip_addr_) {
                msg := std_msgs.NewString()
                msg.Go_data = self.node_name_ + "_log"
                self.publish(tinyros_msgs.Go_ID_SESSION_ID(), msg, true)
            }
            time.Sleep(time.Second)
        } else {
            self.loghd_.read(buff)
        }
    }
}

func (self *NodeHandleTCP) initNode(node_name string, ip string) (bool) {
    self.spin_ = true
    self.ip_addr_ = ip
    self.node_name_ = node_name

    msg := std_msgs.NewString()
    if (self.hardware_.init(self.ip_addr_)) {
        msg.Go_data = self.node_name_
        self.publish(tinyros_msgs.Go_ID_SESSION_ID(), msg, false)
    }

    if !self.keepalive_ {
        self.keepalive_ = true
        go self.keepalive()
    }
    
    return self.hardware_.connected()
}

func (self *NodeHandleTCP) spin() (int) {
    var index, bytes, total_bytes, rv, checksum, l, mode int = 0, 0, 0, 0, 0, 1, MODE_FIRST_FF
    var topic uint32 = 0
    var message_tmp []byte = make([]byte, INPUT_SIZE)
    self.mutex_.Lock()
    for pid := range self.publishers_ {
        self.publishers_[pid].negotiated_ = false
    }
    for sid := range self.subscribers_ {
        self.subscribers_[sid].Go_setnegotiated(false)
    }
    self.mutex_.Unlock()

    if !self.hardware_.connected() {
        return -1
    } else {
        self.negotiateTopics()
    }

    for self.spin_ && self.ok() {
        if l > INPUT_SIZE {
            l = 1
            mode = MODE_FIRST_FF
        }

        rv = self.hardware_.read(message_tmp[0:l])
        if rv < 0 {
            return -1
        }

        for i := 0; i < rv; i++ {
            checksum += int(message_tmp[i])
        }

        if mode == MODE_MESSAGE {
            for i := 0; i < rv; i++ {
                self.message_in_[index] = message_tmp[i]
                index += 1
                bytes -= 1
            }
            if bytes == 0 {
                l = 1
                mode = MODE_MSG_CHECKSUM
            } else {
                l = bytes
            }
        } else if mode == MODE_FIRST_FF {
            if message_tmp[0] == 0xFF {
                mode += 1
            }
        } else if mode == MODE_PROTOCOL_VER {
            if message_tmp[0] == 0xb9 {
                mode += 1
            } else {
                mode = MODE_FIRST_FF
            }
        } else if mode == MODE_SIZE_L {
            bytes = int(message_tmp[0] & 0xFF) << (8 * 0)
            index = 0
            mode += 1
            checksum = int(message_tmp[0])
        } else if mode == MODE_SIZE_L1 {
            bytes += int(message_tmp[0] & 0xFF) << (8 * 1)
            mode += 1
        } else if mode == MODE_SIZE_H {
            bytes += int(message_tmp[0] & 0xFF) << (8 * 2)
            mode += 1
        } else if mode == MODE_SIZE_H1 {
            bytes += int(message_tmp[0] & 0xFF) << (8 * 3)
            total_bytes = bytes
            if total_bytes <= 0 {
                total_bytes = 1
            }
            mode += 1
        } else if mode == MODE_SIZE_CHECKSUM {
            if (checksum % 256) == 255 {
                mode += 1
            } else {
                mode = MODE_FIRST_FF
            }
        } else if mode == MODE_TOPIC_L {
            topic = uint32(message_tmp[0] & 0xFF) << (8 * 0)
            mode += 1
            checksum = int(message_tmp[0])
        } else if mode == MODE_TOPIC_L1 {
            topic += uint32(message_tmp[0] & 0xFF) << (8 * 1)
            mode += 1
        } else if mode == MODE_TOPIC_H {
            topic += uint32(message_tmp[0] & 0xFF) << (8 * 2)
            mode += 1
        } else if mode == MODE_TOPIC_H1 {
            topic += uint32(message_tmp[0] & 0xFF) << (8 * 3)
            mode = MODE_MESSAGE
            if bytes == 0 {
                mode = MODE_MSG_CHECKSUM 
            }  else {
                l = bytes
            }
        } else if mode == MODE_MSG_CHECKSUM {
            mode = MODE_FIRST_FF
            if (checksum % 256) == 255 {
                if topic == tinyros_msgs.Go_ID_PUBLISHER() {
                    self.negotiateTopics()
                } else if topic == tinyros_msgs.Go_ID_ROSTOPIC_REQUEST() {
                
                } else if topic == tinyros_msgs.Go_ID_ROSSERVICE_REQUEST() {
                
                } else if topic == tinyros_msgs.Go_ID_NEGOTIATED() {
                    ti := tinyros_msgs.NewTopicInfo()
                    ti.Go_deserialize(self.message_in_)
                    for _, p := range self.publishers_ {
                        if p.id_ == ti.Go_topic_id {
                            p.negotiated_ = ti.Go_negotiated
                        }
                    }
                    for _, s := range self.subscribers_ {
                        if s.Go_id() == ti.Go_topic_id {
                            s.Go_setnegotiated(ti.Go_negotiated)
                        }
                    }
                } else if topic == tinyros_msgs.Go_ID_TIME() {
                    self.syncTime(self.message_in_)
                } else {
                    s, ok := self.subscribers_[topic]
                    if ok {
                        job := SpinObject {}
                        job.sub_ = s
                        job.message_in_ = make([]byte, total_bytes)
                        copy(job.message_in_, self.message_in_)
                        if s.Go_topic() == TINYROS_LOG_TOPIC {
                            self.spin_log_thread_pool_.schedule(job)
                        } else {
                            if s.Go_srv_flag() {
                                self.spin_srv_thread_pool_.schedule(job)
                            } else {
                                self.spin_thread_pool_.schedule(job)
                            }
                        }
                    }
                }
            }
        }
    }
    return 0
}

func (self *NodeHandleTCP) exit() {
    self.spin_ = false;
    self.keepalive_ = false;
    self.spin_thread_pool_.shutdown();
    self.spin_log_thread_pool_.shutdown();
    self.spin_srv_thread_pool_.shutdown();
    self.loghd_.close();
    self.hardware_.close();
}

func (self *NodeHandleTCP) Go_advertise(p *Publisher) (bool) {
    if len(self.publishers_) < MAX_PUBLISHERS {
        self.mutex_.Lock()
        p.id_ = uint32(len(self.publishers_) + 100 + MAX_SUBSCRIBERS)
        p.nh_ = self
        self.publishers_[p.id_] = p
        self.mutex_.Unlock()
        self.negotiateTopics_p(p)
        return true
    } else {
        return false
    }
}

func (self *NodeHandleTCP) Go_subscribe(s Subscriber_) (bool) {
    if len(self.subscribers_) < MAX_SUBSCRIBERS {
        self.mutex_.Lock()
        id := uint32(len(self.subscribers_) + 100)
        s.Go_setid(id )
        self.subscribers_[id] = s
        self.mutex_.Unlock()
        self.negotiateTopics_s(s)
        return true
    } else {
        return false
    }
}

func (self *NodeHandleTCP) Go_advertiseService(srv Subscriber_) (bool) {
    if len(self.subscribers_) < MAX_SUBSCRIBERS {
        self.mutex_.Lock()
        id := uint32(len(self.subscribers_) + 100)
        srv.Go_setid(id)
        self.subscribers_[id] = srv
        if len(self.publishers_) < MAX_PUBLISHERS {
            srv.Go_pub().id_ = id
            srv.Go_pub().nh_ = self
            self.publishers_[id] = srv.Go_pub()
            self.mutex_.Unlock()
            self.negotiateTopics_s(srv)
            self.negotiateTopics_p(srv.Go_pub())
            return true
        } else {
            self.mutex_.Unlock()
            return false
        }
    } else {
        return false
    }
}

func (self *NodeHandleTCP) Go_serviceClient(srv Subscriber_) (bool) {
    if len(self.subscribers_) < MAX_SUBSCRIBERS {
        self.mutex_.Lock() 
        id := uint32(len(self.subscribers_) + 100)
        srv.Go_setid(id)
        self.subscribers_[id] = srv
        if len(self.publishers_) < MAX_PUBLISHERS {
            srv.Go_pub().id_ = id
            srv.Go_pub().nh_ = self
            self.publishers_[id] = srv.Go_pub()
            self.mutex_.Unlock()
            self.negotiateTopics_s(srv)
            self.negotiateTopics_p(srv.Go_pub())
            return true
        } else {
            self.mutex_.Unlock()
            return false
        }
    } else {
        return false
    }
}

func (self *NodeHandleTCP) Go_log(level uint8, msg string) {
    if self.loghd_.connected() {
        l := tinyros_msgs.NewLog()
        l.Go_level = level
        l.Go_msg = "[" + self.node_name_ + "] " + msg
        self.publish(tinyros_msgs.Go_ID_LOG(), l, true)
    }
}
  
//////////////////////////////////////////////////////////////////////////////////////
// NodeHandleUDP

type NodeHandleUDP struct {
    NodeHandle
}

func NewNodeHandleUDP() (*NodeHandleUDP) {
    newNodeHandleUDP := new(NodeHandleUDP)
    newNodeHandleUDP.hardware_ = NewHardwareUDP()
    newNodeHandleUDP.loghd_ = nil
    newNodeHandleUDP.keepalive_ = false
    newNodeHandleUDP.spin_ = true
    newNodeHandleUDP.ip_addr_ = ""
    newNodeHandleUDP.node_name_ = ""
    newNodeHandleUDP.publishers_ = make(map[uint32]*Publisher)
    newNodeHandleUDP.subscribers_ = make(map[uint32]Subscriber_)
    newNodeHandleUDP.spin_thread_pool_ = NewThreadPool(3)
    newNodeHandleUDP.spin_log_thread_pool_ = nil
    newNodeHandleUDP.spin_srv_thread_pool_ = nil
    newNodeHandleUDP.message_out_ = make([]byte, OUTPUT_SIZE)
    newNodeHandleUDP.message_in_ = make([]byte, INPUT_SIZE)
    return newNodeHandleUDP
}

func (self *NodeHandleUDP) keepalive() {
    for self.keepalive_ {
        self.negotiateTopics()
        time.Sleep(time.Second)
    }
}

func (self *NodeHandleUDP) initNode(node_name string, ip string) (bool) {
    self.spin_ = true
    self.ip_addr_ = ip
    self.node_name_ = node_name
    self.hardware_.init(self.ip_addr_)
    if !self.keepalive_ {
        self.keepalive_ = true
        go self.keepalive()
    }
    return self.hardware_.connected()
}

func (self *NodeHandleUDP) spin() (int) {
    for self.spin_ && self.ok() {
        if len(self.subscribers_) <= 0 {
            time.Sleep(time.Second)
            continue
        }

        rv := self.hardware_.read(self.message_in_)
        if INPUT_SIZE >= rv && rv > 0 {
            var topic uint32 = 0
            var bytes, index, checksum int = 0, 0, 0
            if self.message_in_[index] != 0xff {
                continue
            }

            index += 1
            if self.message_in_[index] != 0xb9 {
                continue
            }

            index += 1
            bytes =  int(self.message_in_[index + 0] & 0xFF) << (8 * 0)
            bytes += int(self.message_in_[index + 1] & 0xFF) << (8 * 1)
            bytes += int(self.message_in_[index + 2] & 0xFF) << (8 * 2)
            bytes += int(self.message_in_[index + 3] & 0xFF) << (8 * 3)
            checksum =  int(self.message_in_[index + 0])
            checksum += int(self.message_in_[index + 1])
            checksum += int(self.message_in_[index + 2])
            checksum += int(self.message_in_[index + 3])
            checksum += int(self.message_in_[index + 4])
            index += 5
            if (checksum % 256) != 255 {
                continue
            }
            
            topic =  uint32(self.message_in_[index + 0] & 0xFF) << (8 * 0)
            topic += uint32(self.message_in_[index + 1] & 0xFF) << (8 * 1)
            topic += uint32(self.message_in_[index + 2] & 0xFF) << (8 * 2)
            topic += uint32(self.message_in_[index + 3] & 0xFF) << (8 * 3)
            checksum =  int(self.message_in_[index + 0])
            checksum += int(self.message_in_[index + 1])
            checksum += int(self.message_in_[index + 2])
            checksum += int(self.message_in_[index + 3])
            index += 4
            if INPUT_SIZE < (index + bytes + 1) {
                continue
            }

            if bytes > 0 {
                for i := 0; i < bytes + 1; i++ {
                    checksum += int(self.message_in_[index + i])
                }
            } else {
                checksum += int(self.message_in_[index])
            }

            if (checksum % 256) == 255 {
                s, ok := self.subscribers_[topic]
                if ok {
                    if bytes <= 0 {
                        bytes = 1
                    }
                    job := SpinObject {}
                    job.sub_ = s
                    job.message_in_ = make([]byte, bytes)
                    copy(job.message_in_, self.message_in_[index:])
                    self.spin_thread_pool_.schedule(job)
                }   
            }
        }
    }

    return 0
}


func (self *NodeHandleUDP) exit() {
    self.spin_ = false
    self.spin_thread_pool_.shutdown()
    self.hardware_.close()
}

func (self *NodeHandleUDP) generate_id() (uint32) {
    var timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()
    var hardwareAddr []byte
    var clockSeq uint32
    var h uint32 = 0
    
    var u []byte = make([]byte, 16)
    now := time.Now()
    tutc := now.In(time.UTC)

    t := uint64(tutc.Unix()-timeBase)*10000000 + uint64(tutc.Nanosecond()/100)
    u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
    u[4], u[5] = byte(t>>40), byte(t>>32)
    u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

    clock := atomic.AddUint32(&clockSeq, 1)
    u[8] = byte(clock >> 8)
    u[9] = byte(clock)

    copy(u[10:], hardwareAddr)

    u[6] |= 0x10
    u[8] &= 0x3F
    u[8] |= 0x80

    for _, b := range u {
        h = 31 * h + uint32(b & 0xff)
    }
    
    return h
}

func (self *NodeHandleUDP) Go_advertise(p *Publisher) (bool) {
    if len(self.publishers_) < MAX_PUBLISHERS {
        self.mutex_.Lock()
        p.id_ = self.generate_id()
        p.nh_ = self
        self.publishers_[p.id_] = p
        self.mutex_.Unlock()
        self.negotiateTopics_p(p)
        return true
    } else {
        return false
    }
}

func (self *NodeHandleUDP) Go_subscribe(s Subscriber_) (bool) {
    if len(self.subscribers_) < MAX_SUBSCRIBERS {
        self.mutex_.Lock()
        id := self.generate_id()
        s.Go_setid(id )
        self.subscribers_[id] = s
        self.mutex_.Unlock()
        self.negotiateTopics_s(s)
        return true
    } else {
        return false
    }
}

//////////////////////////////////////////////////////////////////////////////////////
// 
var (
    go_nh *NodeHandleTCP
    go_udp *NodeHandleUDP
    go_nh_mutex sync.Mutex
    go_udp_mutex sync.Mutex
    go_ipaddr string = "127.0.0.1"
    go_node_name string = ""
)

func Go_init(name string, ip ...string) {
    go_node_name = name
    if len(ip) > 0 {
        go_ipaddr = ip[0]
    }
    Go_nh()
}

func Go_nh() (*NodeHandleTCP) {
    if go_nh == nil {
        go_nh_mutex.Lock()
        if go_nh == nil {
            go_nh = NewNodeHandleTCP()
            go func() {
                for go_nh.spin_ {
                    if !go_nh.ok() {
                        go_nh.initNode(go_node_name, go_ipaddr)
                    }
                    if go_nh.ok() {
                        go_nh.spin()
                    }
                    time.Sleep(time.Second)
                }
            }()
            time.Sleep(200 * time.Millisecond)
        }
        go_nh_mutex.Unlock()
    }
    return go_nh
}

func Go_udp() (*NodeHandleUDP) {
    if go_udp == nil {
        go_udp_mutex.Lock()
        if go_udp == nil {
            go_udp = NewNodeHandleUDP()
            go func() {
                for go_udp.spin_ {
                    if !go_udp.ok() {
                        go_udp.initNode(go_node_name, go_ipaddr)
                    }
                    if go_udp.ok() {
                        go_udp.spin()
                    }
                    time.Sleep(time.Second)
                }
            }()
            time.Sleep(200 * time.Millisecond)
        }
        go_udp_mutex.Unlock()
    }
    return go_udp
}

func Go_logdebug(msg string) {
    Go_nh().Go_log(tinyros_msgs.Go_ROSDEBUG(), msg)
}

func Go_loginfo(msg string) {
    Go_nh().Go_log(tinyros_msgs.Go_ROSINFO(), msg)
}

func Go_logwarn(msg string) {
    Go_nh().Go_log(tinyros_msgs.Go_ROSWARN(), msg)
}

func Go_logerror(msg string) {
    Go_nh().Go_log(tinyros_msgs.Go_ROSERROR(), msg)
}

func Go_logfatal(msg string) {
    Go_nh().Go_log(tinyros_msgs.Go_ROSFATAL(), msg)
}
