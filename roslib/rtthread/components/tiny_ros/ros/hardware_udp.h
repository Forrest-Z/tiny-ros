#ifndef TINYROS_HARDWARE_UDP_H_
#define TINYROS_HARDWARE_UDP_H_
#include <stdint.h>
#include <stdio.h>
#include <rthw.h>
#include <rtthread.h>
#if defined(RT_USING_SAL)
#include <netdb.h>
#include <sys/socket.h>
#else
#include <lwip/netdb.h>
#include <lwip/sockets.h>
#endif /* RT_USING_SAL */
#include "tiny_ros/ros/string.h"

namespace tinyros {
#undef SERVER_PORTNUM
#undef CLIENT_PORTNUM
#define SERVER_PORTNUM 11316
#define CLIENT_PORTNUM 11317
class HardwareUdp
{
public:
  HardwareUdp(): sockfd_(-1), connected_(false), sock_binded_(false) {

  }

  virtual bool init(tinyros::string portName) {
    this->close();

    sockfd_ = socket(AF_INET, SOCK_DGRAM, 0);
    if (sockfd_ < 0) {
      rt_kprintf("HardwareUdp::init error opening socket: %s(errno: %d)\n", strerror(errno), errno);
      return false;
    }
    
    int opt = 1;
    setsockopt(sockfd_, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));

    memset(&server_endpoint_, 0, sizeof(struct sockaddr_in));
    server_endpoint_.sin_family = AF_INET;
    server_endpoint_.sin_port = htons(SERVER_PORTNUM);
    server_endpoint_.sin_addr.s_addr = inet_addr(portName.c_str());

    memset(&client_endpoint_, 0, sizeof(struct sockaddr_in));
    client_endpoint_.sin_family = AF_INET;
    client_endpoint_.sin_port = htons(CLIENT_PORTNUM);
    client_endpoint_.sin_addr.s_addr = htonl(INADDR_ANY);

    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      if (!sock_binded_) {
        if (bind(sockfd_, (struct sockaddr*)&client_endpoint_, sizeof(client_endpoint_)) < 0) {
          printf("HardwareUdp::read bind socket: %s(errno: %d)\n", strerror(errno), errno);
        } else {
          sock_binded_ = true;
        }
      }
      
      struct sockaddr_in from;
      socklen_t from_len = sizeof(from);
      int rv = recvfrom(sockfd_, data, length, 0, (struct sockaddr*)&from, &from_len);
      if (rv < 0) {
        rt_kprintf("HardwareUdp::read error: %s(errno: %d)\n", strerror(errno), errno);
      }
      return (rv > 0 ? rv : 0);
    }
    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
      if (sendto(sockfd_, data, length, 0, (struct sockaddr *)&server_endpoint_, sizeof(server_endpoint_)) < 0) {
        rt_kprintf("HardwareUdp::write error sendto: %s(errno: %d)\n", strerror(errno), errno);
        return false;
      } else {
        return true;
      }
    }
    return false;
  }

  virtual bool connected() {
    return connected_;
  }

  virtual void close() {
    connected_ = false;
    sock_binded_ = false;
    if (sockfd_ > 0) {
      closesocket(sockfd_);
      sockfd_ = -1;
    }
  }

private:
  int sockfd_;
  bool connected_;
  bool sock_binded_;
  struct sockaddr_in server_endpoint_;
  struct sockaddr_in client_endpoint_;
};
}

#endif /* TINYROS_HARDWARE_UDP_H_ */
