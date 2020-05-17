/*
 * File      : hardware_udp.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_UDP_H_
#define TINYROS_HARDWARE_UDP_H_
#include <stdint.h>
#include <string>
#include <iostream>
#include "hardware.h"
#ifdef WIN32
#include <winsock2.h>
#include <ws2tcpip.h>
#include <Windows.h>
#include <tchar.h>
#pragma comment(lib, "Ws2_32.lib")
#else
#include <sys/socket.h>  
#include <netinet/in.h>  
#include <arpa/inet.h>
#endif

namespace tinyros {
#undef SERVER_PORTNUM
#undef CLIENT_PORTNUM
#define SERVER_PORTNUM 11316
#define CLIENT_PORTNUM 11317

#ifdef WIN32
#undef errno
#define errno (WSAGetLastError())
#define strerror(errno) ("")
#endif

class HardwareUdp: public Hardware
{
public:
  HardwareUdp()
    : sockfd_(-1)
    , sock_binded_(false)
    , connected_(false) {
#ifdef WIN32
    WSADATA wsaData;
    if (WSAStartup (MAKEWORD (2, 2), &wsaData) != 0) {
      std::cerr << "HardwareUdp::HardwareUdp Could not initialize windows socket: "
                << strerror(errno) << "(errno: " << errno <<")" << std::endl;
    }
#endif
  }

  ~HardwareUdp() {
#ifdef WIN32
    WSACleanup();
#endif
  }
  
  virtual bool init(std::string portName) {
    this->close();

    sockfd_ = socket(AF_INET, SOCK_DGRAM, 0);
    if (sockfd_ < 0) {
      std::cerr << "HardwareUdp::init() opening socket: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
      return false;
    }
    
    int opt = 1;
    struct linger so_linger;
    so_linger.l_onoff = 1;
    so_linger.l_linger = 0;
    setsockopt(sockfd_, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));
    setsockopt(sockfd_, SOL_SOCKET, SO_LINGER, (const char *)&so_linger, sizeof(so_linger));

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
          std::cerr << "HardwareUdp::read() bind socket: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
        } else {
          sock_binded_ = true;
        }
      }

      struct sockaddr_in from;
      socklen_t from_len = sizeof(from);
      int rv = recvfrom(sockfd_, (char*)data, length, 0, (struct sockaddr*)&from, &from_len);
      if (rv < 0) {
        std::cerr << "HardwareUdp::read() recvfrom: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
      }
      return (rv > 0 ? rv : 0);
    }
    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
      if(sendto(sockfd_, (const char*)data, length, 0, (struct sockaddr *)&server_endpoint_, sizeof(server_endpoint_)) <= 0) {
        printf("HardwareUdp::write() sendto: %s(errno: %d)\n", strerror(errno), errno);
        std::cerr << "HardwareUdp::write() sendto: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
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
#ifdef WIN32
      closesocket(sockfd_);
#else
      ::close(sockfd_);
#endif
      sockfd_ = -1;
    }
  }
  
private:
  int sockfd_;
  bool sock_binded_;
  bool connected_;
  struct sockaddr_in server_endpoint_;
  struct sockaddr_in client_endpoint_;
};
}
#endif //TINYROS_HARDWARE_LINUX_UDP_H_

