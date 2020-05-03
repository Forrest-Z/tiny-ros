/*
 * File      : hardware_linux.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_LINUX_UDP_H_
#define TINYROS_HARDWARE_LINUX_UDP_H_
#if defined(linux) || defined(__linux__) || defined(__linux) || defined(__MACH__) || defined(__CYGWIN__)
#include <sys/socket.h>  
#include <netinet/in.h>  
#include <arpa/inet.h> 
#include "hardware.h"

namespace tinyros {
#undef SERVER_PORTNUM
#undef CLIENT_PORTNUM
#define SERVER_PORTNUM 11316
#define CLIENT_PORTNUM 11317
class HardwareLinuxUdp: public Hardware
{
public:
  HardwareLinuxUdp(): connected_(false), sockfd_(-1) {
    
  }
  
  virtual bool init(std::string portName) {
    this->close();
    
    sockfd_ = socket(AF_INET, SOCK_DGRAM, 0);
    if (sockfd_ < 0) {
      printf("ERROR opening socket: (%d:%s)\n", errno, strerror(errno));
      return false;
    }
    
    int opt = 1;
    setsockopt(sockfd_, SOL_SOCKET, SO_REUSEADDR, (const void *)&opt, sizeof(opt));

    memset(&server_endpoint_, 0, sizeof(struct sockaddr_in));
	  server_endpoint_.sin_family = AF_INET;
	  server_endpoint_.sin_port = htons(SERVER_PORTNUM);
    server_endpoint_.sin_addr.s_addr = inet_addr(portName.c_str());

	  memset(&client_endpoint_, 0, sizeof(struct sockaddr_in));
	  client_endpoint_.sin_family = AF_INET;
	  client_endpoint_.sin_port = htons(CLIENT_PORTNUM);
	  client_endpoint_.sin_addr.s_addr = htonl(INADDR_ANY);
    if (bind(sockfd_, (struct sockaddr*)&client_endpoint_, sizeof(client_endpoint_)) < 0) {
      printf("ERROR bind socket: (%d:%s)\n", errno, strerror(errno));
    }

    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      struct sockaddr_in from;
	    socklen_t from_len = sizeof(from);
      int rv = recvfrom(sockfd_, data, length, 0, (struct sockaddr*)&from, &from_len);
      return (rv > 0 ? rv : 0);
    }
    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
	    if (sendto(sockfd_, data, length, 0, (struct sockaddr *)&server_endpoint_, sizeof(server_endpoint_)) <= 0) {
        printf("ERROR sendto: (%d:%s)\n", errno, strerror(errno));
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
      if (sockfd_ > 0) {
        ::close(sockfd_);
        sockfd_ = -1;
      }
    }
  }
  
private:
  int sockfd_;
  bool connected_;
  struct sockaddr_in server_endpoint_;
  struct sockaddr_in client_endpoint_;
};
}
#endif
#endif //TINYROS_HARDWARE_LINUX_UDP_H_


