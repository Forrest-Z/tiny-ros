import sys
import socket

class Hardware(object):
    __slots__ = ['sockfd_', 'port_', 'connected_']
    
    def __init__(self):
        self.sockfd_ = None
        self.port_ = 11315
        self.connected_ = False

    def init(self, ip):
        try:
            self.close()
            self.sockfd_ = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
            self.sockfd_.setsockopt(socket.SOL_TCP, socket.TCP_NODELAY, 1)
            self.sockfd_.connect((ip, self.port_))
            self.connected_ = True
        except Exception as ex: 
            print('Hardware::init %s' % str(ex))
            self.connected_ = False
        return self.connected_
  
    def read(self, length):
        if self.connected_ == True:
            try:
                _x = self.sockfd_.recv(length)
                size = len(_x)
                return _x, size
            except Exception as ex: 
                print('Hardware::read %s' % str(ex))
                self.close()
                return '', -1

    def write(self, buff): 
        if self.connected_ == True:
            try:
                _x = buff
                size = len(_x)
                self.sockfd_.send(_x)
                return size
            except Exception as ex: 
                print('Hardware::write %s' % str(ex))
                self.close()
                return -1
        else:
            return -1 
    
    def connected(self):
        return self.connected_

    def close(self):
        self.connected_ = False
        if self.sockfd_ != None:
            self.sockfd_.close()
            self.sockfd_ = None

