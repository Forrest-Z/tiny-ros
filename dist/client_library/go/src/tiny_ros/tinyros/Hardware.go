package tinyros

import (
    "net"
    "fmt"
)

type Hardware interface {
    init(ip string) (bool)
    read(data []byte) (int)
    write(data []byte) (int)
    connected() (bool)
    close()
}

type HardwareTCP struct {
    sockfd_ *net.TCPConn `json:"sockfd"`
    port_ string `json:"port"`
    connected_ bool `json:"connected"`
}

func NewHardwareTCP() (*HardwareTCP) {
    newHardwareTCP := new(HardwareTCP)
    newHardwareTCP.sockfd_ = nil
    newHardwareTCP.port_ = ":11315"
    newHardwareTCP.connected_ = false
    return newHardwareTCP
}

func (self *HardwareTCP) init(ip string) (bool) {
    target := ip + self.port_
    self.close()
    addr, err := net.ResolveTCPAddr("tcp", string(target))
    if err != nil {
        fmt.Println("HardwareTCP.init", err)
        return false
    }
    
    self.sockfd_, err = net.DialTCP("tcp", nil, addr)
    if err != nil {
        fmt.Println("HardwareTCP.init", err)
        return false
    }
    self.sockfd_.SetNoDelay(true)
    self.sockfd_.SetLinger(0)
    self.connected_ = true
    return self.connected_
}

func (self *HardwareTCP) read(data []byte) (int) {
    if self.connected_ {
        rv, err := self.sockfd_.Read(data)
        if err != nil {
            fmt.Println("HardwareTCP.read", err)
            self.close()
            rv = -1
        }
        return rv
    }
    return -1
}

func (self *HardwareTCP) write(data []byte) (int) {
    if self.connected_ {
        rv, err := self.sockfd_.Write(data)
        if err != nil {
            fmt.Println("HardwareTCP.write", err)
            self.close()
            rv = -1
        }
        return rv
    }
    return -1
}

func (self *HardwareTCP) connected() (bool) {
    return self.connected_
}

func (self *HardwareTCP) close() {
    self.connected_ = false
    if self.sockfd_ != nil {
        self.sockfd_.Close()
        self.sockfd_ = nil
    }
}


/////////////////////////////////////////////////////////


type HardwareUDP struct {
    conn_send_ *net.UDPConn `json:"conn_send"`
    conn_recv_ *net.UDPConn `json:"conn_recv"`
    addr_send_ *net.UDPAddr `json:"addr_send"`
    addr_recv_ *net.UDPAddr `json:"addr_recv"`
    client_port_ string `json:"client_port"`
    server_port_ string `json:"server_port"`
    connected_ bool `json:"connected"`
}

func NewHardwareUDP() (*HardwareUDP) {
    newHardwareUDP := new(HardwareUDP)
    newHardwareUDP.conn_send_ = nil
    newHardwareUDP.conn_recv_ = nil
    newHardwareUDP.addr_send_ = nil
    newHardwareUDP.addr_recv_ = nil
    newHardwareUDP.server_port_ = ":11316"
    newHardwareUDP.client_port_ = ":11317"
    newHardwareUDP.connected_ = false
    return newHardwareUDP
}

func (self *HardwareUDP) init(ip string) (bool) {
    target := ip + self.server_port_
    self.close()
    addr_send, err := net.ResolveUDPAddr("udp", string(target))
    if err != nil {
        fmt.Println("HardwareUDP.init", err)
        return false
    }
    self.addr_send_ = addr_send
    
    target = "" + self.client_port_
    self.addr_recv_, err = net.ResolveUDPAddr("udp", string(target))
    if err != nil {
        fmt.Println("HardwareUDP.init", err)
        return false
    }
    
    self.conn_send_, err = net.DialUDP("udp", nil, self.addr_send_)
    if err != nil {
        fmt.Println("HardwareUDP.init", err)
        return false
    }
    self.connected_ = true
    return self.connected_
}

func (self *HardwareUDP) read(data []byte) (int) {
    if self.connected_ {
        if self.conn_recv_ == nil {
            conn_recv, err := net.ListenUDP("udp", self.addr_recv_)
            if err != nil {
                fmt.Println("HardwareUDP.read", err)
                return -1
            }
            self.conn_recv_ = conn_recv
        }
        
        rv, _, err := self.conn_recv_.ReadFrom(data)
        if err != nil {
            fmt.Println("HardwareUDP.read", err)
            self.close()
            rv = -1
        }
        return rv
    }
    return -1
}

func (self *HardwareUDP) write(data []byte) (int) {
    if self.connected_ {
        rv, err := self.conn_send_.Write(data)
        if err != nil {
            fmt.Println("HardwareUDP.write", err)
            self.close()
            rv = -1
        }
        return rv
    }
    return -1
}

func (self *HardwareUDP) connected() (bool) {
    return self.connected_
}

func (self *HardwareUDP) close() {
    self.connected_ = false
    self.addr_send_ = nil
    self.addr_recv_ = nil
    if self.conn_send_ != nil {
        self.conn_send_.Close()
        self.conn_send_ = nil
    }
    if self.conn_recv_ != nil {
        self.conn_recv_.Close()
        self.conn_recv_ = nil
    }
}
