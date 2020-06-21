package tinyros

import (
)
 
type NodeHandleBase interface {
    initNode(portName string) (bool)
    publish(id uint32, msg Msg, islog bool) (int)
    spin() (int)
    exit()
    ok() (bool)
}
 
type NodeHandle struct {
    hardware_ Hardware `json:"hardware"`
    loghd_ Hardware `json:"loghd"`
}

func (self *NodeHandle) publish(id uint32, msg Msg, islog bool) (int) {
    var chk int
    message_out := make([]byte, 65*1024)
    l := msg.Go_serialize(message_out[11:])
    message_out[0] = byte(0xff)
    message_out[1] = byte(0xb9)
    message_out[2] = byte((l >> 0) & 0xFF)
    message_out[3] = byte((l >> 8) & 0xFF)
    message_out[4] = byte((l >> 16) & 0xFF)
    message_out[5] = byte((l >> 24) & 0xFF)
    chk = int(message_out[2])
    chk += int(message_out[3])
    chk += int(message_out[4])
    chk += int(message_out[5])
    message_out[6] = byte(255 - (chk % 256))
    message_out[7] = byte((id >> 0) & 0xFF)
    message_out[8] = byte((id >> 8) & 0xFF)
    message_out[9] = byte((id >> 16) & 0xFF)
    message_out[10] = byte((id >> 24) & 0xFF)

    chk = 0
    for i := 7; i < l + 11; i++ {
      chk += int(message_out[i])
    }
    
    l += 11
    
    message_out[l] = byte(255 - (chk % 256))

    l += 1

    if !islog {
        l = self.hardware_.write(message_out[:l])
    } else {
        l = self.loghd_.write(message_out[:l])
    }
    return l
}

/////////////////////////////////////////////////////////

type NodeHandleTCP struct {
    NodeHandle
}

func (self *NodeHandleTCP) Go_publish(id uint32, msg Msg) (int) {
    return self.publish(id, msg, false);
}

func NewNodeHandleTCP() (*NodeHandleTCP) {
    newNodeHandleTCP := new(NodeHandleTCP)
    newNodeHandleTCP.hardware_ = NewHardwareTCP()
    newNodeHandleTCP.loghd_ = NewHardwareTCP()
    return newNodeHandleTCP
}


/////////////////////////////////////////////////////////

type NodeHandleUDP struct {
    NodeHandle
}

func (self *NodeHandleUDP) Go_publish(id uint32, msg Msg) (int) {
    return self.publish(id, msg, false);
}

func NewNodeHandleUDP() (*NodeHandleUDP) {
    newNodeHandleUDP := new(NodeHandleUDP)
    newNodeHandleUDP.hardware_ = NewHardwareUDP()
    newNodeHandleUDP.loghd_ = nil
    return newNodeHandleUDP
}

/////////////////////////////////////////////////////////
var (
   go_nh *NodeHandleTCP
   go_udp *NodeHandleUDP
)

func Go_nh() (*NodeHandleTCP) {
    if go_nh == nil {
        go_nh = NewNodeHandleTCP()
        go_nh.hardware_.init("127.0.0.1")
    }
    return go_nh
}

func Go_udp() (*NodeHandleUDP) {
    if go_udp == nil {
        go_udp = NewNodeHandleUDP()
        go_udp.hardware_.init("127.0.0.1")
    }
    return go_udp
}


