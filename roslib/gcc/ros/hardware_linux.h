/*
 * File      : hardware_linux.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_LINUX_H_
#define TINYROS_HARDWARE_LINUX_H_
#if defined(linux) || defined(__linux__) || defined(__linux) || defined(__MACH__) || defined(__CYGWIN__)
#include <stdint.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <netdb.h>
#include <unistd.h>
#include "hardware.h"

namespace tinyros {
#undef DEFAULT_PORTNUM
#define DEFAULT_PORTNUM 11315
class HardwareLinux: public Hardware
{
public:
  HardwareLinux(): connected_(false), sockfd_(-1) {
    
  }
  virtual bool init(std::string portName) {
    struct sockaddr_in serv_addr;
    struct hostent *server;
    const char* pName = portName.c_str();
    const char* tcpPortNumString = strchr(pName, ':');
    if (!tcpPortNumString) {
      tcpPortNum_ = DEFAULT_PORTNUM;
      strncpy(ip, pName, 16);
    } else {
      tcpPortNum_ = strtol(tcpPortNumString + 1, NULL, 10);
      strncpy(ip, pName, tcpPortNumString - pName);
    }

    this->close();

    sockfd_ = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd_ < 0) {
      printf("ERROR opening socket.\n");
      return false;
    }

     // Disable the Nagle (TCP No Delay) algorithm.
    int flag = 1;
    if (setsockopt(sockfd_, IPPROTO_TCP, TCP_NODELAY, (char *)&flag, sizeof(flag)) < 0)
    {
      printf("Couldn't setsockopt(TCP_NODELAY)\n");
      this->close();
      return false;
    }

    // Connect to the server
    server = gethostbyname(ip);
    if (server == NULL) {
      printf("ERROR, no such host\n");
      this->close();
      return false;
    }
    bzero((char *) &serv_addr, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    bcopy((char *)server->h_addr, (char *)&serv_addr.sin_addr.s_addr, server->h_length);
    serv_addr.sin_port = htons(tcpPortNum_);
    if (connect(sockfd_, (struct sockaddr *) &serv_addr, sizeof(serv_addr)) < 0) {
      printf("ERROR connecting\n");
      this->close();
      return false;
    }

    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      int rv = ::read(sockfd_, data, length);
      if (rv > 0) {
        return rv;
      } else if (rv == 0) {
        printf("read(): socket close\n");
        this->close();
        return -1;
      } else {
        if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR))  {
          printf("read(): error %d:%s\n", errno, strerror(errno));
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
        rv = ::write(sockfd_, data + totalsent, len - totalsent);
        if (rv > 0) {
          totalsent += rv;
        } else if (rv == 0) {
          printf("write(): socket close\n");
          this->close();
          return false;
        } else {
          if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR)) {
            printf("write(): error %d:%s\n", errno, strerror(errno));
            this->close();
            return false;
          } else {
            printf("write(): error writing - trying again - \n");
          }
        }
      } while (totalsent < len);

      return true;
    }
    return true;
  }

  virtual bool connected() {
    return connected_;
  }

  virtual void close() {
    if (connected_) {
      connected_ = false;
      if (sockfd_ > 0) {
        ::close(sockfd_);
        sockfd_ = -1;
      }
    }
  }
  
private:
  char ip[16];
  int sockfd_;
  long int tcpPortNum_;
  bool connected_;
};
}
#endif
#endif //TINYROS_HARDWARE_LINUX_H_