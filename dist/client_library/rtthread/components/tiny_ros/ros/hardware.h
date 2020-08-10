#ifndef TINYROS_HARDWARE_H_
#define TINYROS_HARDWARE_H_
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
#undef DEFAULT_PORTNUM
#define DEFAULT_PORTNUM 11315
class Hardware
{
public:
  Hardware(): connected_(false), sockfd_(-1) {

  }

  ~Hardware() {
    this->close();
  }

  bool init(tinyros::string portName) {
    this->close();

    // Create the socket.
    sockfd_ = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd_ < 0) {
      rt_kprintf("Hardware::init opening socket error %d:%s\n", errno, strerror(errno));
      return false;
    }

    int opt = 1;
    setsockopt(sockfd_, IPPROTO_TCP, TCP_NODELAY, (const char *)&opt, sizeof(opt));

    struct hostent *host;
    struct sockaddr_in serv_addr;
    host = gethostbyname(portName.c_str());
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port =  htons(DEFAULT_PORTNUM);
    serv_addr.sin_addr = *((struct in_addr *)host->h_addr);
    memset(&(serv_addr.sin_zero), 0, sizeof(serv_addr.sin_zero));
    if (connect(sockfd_, (struct sockaddr *) &serv_addr, sizeof(struct sockaddr)) < 0) {
      rt_kprintf("Hardware::init Couldn't connecting %d:%s\n", errno, strerror(errno));
      closesocket(sockfd_);
      sockfd_ = -1;
      return false;
    }

    connected_ = true;
    return connected_;
  }

  int read(uint8_t* data, int length) {
    if (connected_) {
      int rv = recv(sockfd_, data, length, 0);
      if (rv > 0) {
        return rv;
      } else if (rv == 0) {
        rt_kprintf("Hardware::read socket close %d:%s\n", errno, strerror(errno));
        this->close();
        return -1;
      } else {
        if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR))  {
          rt_kprintf("Hardware::read error %d:%s\n", errno, strerror(errno));
          this->close();
          return -1;
        }
      }
    }

    return -1;
  }

  bool write(uint8_t* data, int length) {
    if (connected_) {
      int rv, len = length, totalsent = 0;
      do {
        rv = send(sockfd_, data + totalsent, len - totalsent, 0);
        if (rv > 0) {
          totalsent += rv;
        } else if (rv == 0) {
          rt_kprintf("Hardware::write socket close %d:%s\n", errno, strerror(errno));
          this->close();
          return false;
        } else {
          if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR)) {
            rt_kprintf("Hardware::write error %d:%s\n", errno, strerror(errno));
            this->close();
            return false;
          } else {
            rt_kprintf("Hardware::write: error writing - trying again - %d:%s\n", errno, strerror(errno));
          }
        }
      } while (totalsent < len);

      return true;
    }
    return false;
  }

  bool connected() {
    return connected_;
  }

  void close() {
    connected_ = false;
    if(sockfd_ > 0){
      closesocket(sockfd_);
      sockfd_ = -1;
    }
  }

private:
  char ip[16];
  bool connected_;
  int sockfd_;
  long int tcpPortNum_;
};
}

#endif //TINYROS_HARDWARE_H_

