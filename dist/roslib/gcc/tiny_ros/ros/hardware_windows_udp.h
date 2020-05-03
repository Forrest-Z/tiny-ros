/*
 * File      : hardware_windows_udp.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_WINDOWS_UDP_H_
#define TINYROS_HARDWARE_WINDOWS_UDP_H_
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
#include <string>
#include <iostream>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <Windows.h>
#include <tchar.h>
#include "hardware.h"
#pragma comment(lib, "Ws2_32.lib")

namespace tinyros {
#undef SERVER_PORTNUM
#undef CLIENT_PORTNUM
#define SERVER_PORTNUM 11316
#define CLIENT_PORTNUM 11317
class HardwareWindowsUdp: public Hardware
{
public:
  HardwareWindowsUdp(): connected_(false), sockfd_(INVALID_SOCKET) {

  }

  virtual bool init(std::string portName) {
    WSADATA wsaData;

    this->close();

    int result = WSAStartup (MAKEWORD (2, 2), &wsaData);
    if (result) {
      std::cerr << "Could not initialize windows socket (" << result << ")" << std::endl;
      return false;
    }

    SOCKET sockfd_ = socket(AF_INET, SOCK_DGRAM, IPPROTO_IP);
    if (sockfd_ == INVALID_SOCKET) {
      WSACleanup();
      std::cerr << "Failed to create socket!" << std::endl;
      return false;
    }

    int opt = 1;
    setsockopt(sockfd_, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));

    memset(&server_endpoint_, 0, sizeof(SOCKADDR_IN));
    server_endpoint_.sin_family = AF_INET;
    server_endpoint_.sin_port = htons(SERVER_PORTNUM);
    server_endpoint_.sin_addr.S_un.S_addr = inet_addr(portName.c_str());

    memset(&client_endpoint_, 0, sizeof(SOCKADDR_IN));
    client_endpoint_.sin_family = AF_INET;
    client_endpoint_.sin_port = htons(CLIENT_PORTNUM);
    client_endpoint_.sin_addr.S_un.S_addr = htonl(INADDR_ANY);
    if (bind(sockfd_, (sockaddr*)&client_endpoint_, sizeof(client_endpoint_)) != 0) {
      std::cerr << "ERROR bind socket: (" << errno <<":"<<strerror(errno)<<")"<<std::endl;
    }

    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      SOCKADDR_IN from;
      int from_len = sizeof(SOCKADDR_IN);
      int rv = recvfrom(sockfd_, (char*)data, length, 0, (SOCKADDR*)&from, &from_len);
      return (rv > 0 ? rv : 0);
    }
    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
      std::cerr << "ERROR sendto: (" << length <<std::endl;
      if (sendto(sockfd_, (const char *)data, length, 0,
                 (SOCKADDR *)&server_endpoint_, sizeof(server_endpoint_)) <= 0) {
        std::cerr << "ERROR sendto: (" << errno <<":"<<strerror(errno)<<")"<<std::endl;
        return false;
      }
    }
    return true;
  }

  virtual bool connected() {
    return connected_;
  }

  virtual void close() {
    if (connected_) {
      connected_ = false;
      if (sockfd_ != INVALID_SOCKET) {
        closesocket (sockfd_);
        sockfd_ = INVALID_SOCKET;
        WSACleanup ();
      }
    }
  }

private:
  SOCKET sockfd_;
  bool connected_;
  SOCKADDR_IN server_endpoint_;
  SOCKADDR_IN client_endpoint_;
};
}
#endif
#endif // TINYROS_HARDWARE_WINDOWS_UDP_H_
