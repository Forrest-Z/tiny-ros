#ifndef TINYROS_HARDWARE_TCP_H_
#define TINYROS_HARDWARE_TCP_H_
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
#include <netinet/tcp.h>
#include <arpa/inet.h> 
#include <unistd.h>
#endif

namespace tinyros {
#undef SERVER_PORTNUM
#define SERVER_PORTNUM 11315

#ifdef WIN32
#undef errno
#define errno (WSAGetLastError())
#define strerror(errno) ("")
#endif

class HardwareTcp: public Hardware
{
public:
  HardwareTcp(): connected_(false), sockfd_(-1) {
#ifdef WIN32
    WSADATA wsaData;
    if (WSAStartup (MAKEWORD (2, 2), &wsaData) != 0) {
      std::cerr << "HardwareTcp::HardwareTcp Could not initialize windows socket: "
                << strerror(errno) << "(errno: " << errno <<")" << std::endl;
    }
#endif
  }

  ~HardwareTcp() {
#ifdef WIN32
    WSACleanup();
#endif
  }

  virtual bool init(std::string portName) {
    this->close();

    sockfd_ = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd_ < 0) {
      std::cerr << "HardwareTcp::init() opening socket: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
      return false;
    }

    int opt = 1;
    struct linger so_linger;
    so_linger.l_onoff = 1;
    so_linger.l_linger = 0;
    setsockopt(sockfd_, IPPROTO_TCP, TCP_NODELAY, (const char*)&opt, sizeof(opt));
    setsockopt(sockfd_, SOL_SOCKET, SO_LINGER, (const char *)&so_linger, sizeof(so_linger));

    struct sockaddr_in serv_addr;
    memset(&serv_addr, 0, sizeof(struct sockaddr_in));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(SERVER_PORTNUM);
    serv_addr.sin_addr.s_addr = inet_addr(portName.c_str());

    if (connect(sockfd_, (struct sockaddr *) &serv_addr, sizeof(serv_addr)) < 0) {
      std::cerr << "HardwareTcp::init() connecting: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
      this->close();
      return false;
    }

    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      int rv = recv(sockfd_, (char*)data, length, 0);
      if (rv > 0) {
        return rv;
      } else if (rv == 0) {
        std::cerr << "HardwareTcp::read() socket close: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
        this->close();
        return -1;
      } else {
#ifdef WIN32
        int error = WSAGetLastError();
        if ((error != WSAEWOULDBLOCK) && (error != WSAEINTR)) {
#else
        if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR))  {
#endif
          std::cerr << "HardwareTcp::read() error: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
          this->close();
          return -1;
        }
      }
    }

    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
      int rv, len = length, totalsent = 0;
      do {
        rv = send(sockfd_, (const char *)(data + totalsent), len - totalsent, 0);
        if (rv > 0) {
          totalsent += rv;
        } else if (rv == 0) {
          std::cerr << "HardwareTcp::write() socket close: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
          this->close();
          return false;
        } else {
#ifdef WIN32
          if ((errno != WSAEWOULDBLOCK) && (errno != WSAEINTR)) {
#else
          if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR)) {
#endif
            std::cerr << "HardwareTcp::write() error: " << strerror(errno) << "(errno: " << errno <<")" << std::endl;
            this->close();
            return false;
          } else {
            std::cerr << "HardwareTcp::write() error writing - trying again - : "
                      << strerror(errno) << "(errno: " << errno <<")" << std::endl;
          }
        }
      } while (totalsent < len);

      return true;
    }
    return false;
  }

  virtual bool connected() {
    return connected_;
  }

  virtual void close() {
    connected_ = false;
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
  char ip[16];
  int sockfd_;
  long int tcpPortNum_;
  bool connected_;
};
}
#endif //TINYROS_HARDWARE_TCP_H_

