/*
 * File      : udp_stream.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-03     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_UDP_STREAM_H
#define TINY_ROS_UDP_STREAM_H
#include <iostream>
#include <fcntl.h>
#include <sys/socket.h>  
#include <netinet/in.h>  
#include <arpa/inet.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include "common.h"

namespace tinyros
{
class UdpStream
{
public:
  UdpStream() : sock_fd_(-1) { }

  bool open(int server_port, int client_port, const std::string& session_id) {
    sock_fd_ = socket(AF_INET, SOCK_DGRAM, 0);
    if (sock_fd_ < 0) {
      spdlog_error("[{0}] UdpStream::open opening socket: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
      return false;
    }

    int opt = 1;
    struct linger so_linger;
    so_linger.l_onoff = 1;
    so_linger.l_linger = 0;
    setsockopt(sock_fd_, SOL_SOCKET, SO_BROADCAST, (const char *)&opt, sizeof(opt));
    setsockopt(sock_fd_, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));
    setsockopt(sock_fd_, SOL_SOCKET, SO_LINGER, (const char *)&so_linger, sizeof(so_linger));

    memset(&client_endpoint_, 0, sizeof(struct sockaddr_in));
    client_endpoint_.sin_family = AF_INET;
    client_endpoint_.sin_port = htons(client_port);
    client_endpoint_.sin_addr.s_addr = htonl(INADDR_BROADCAST);

    memset(&server_endpoint_, 0, sizeof(struct sockaddr_in));
    server_endpoint_.sin_family = AF_INET;
    server_endpoint_.sin_port = htons(server_port);
    server_endpoint_.sin_addr.s_addr = htonl(INADDR_ANY);
    if (bind(sock_fd_, (struct sockaddr*)&server_endpoint_, sizeof(server_endpoint_)) < 0) {
      spdlog_error("[{0}] UdpStream::open bind socket: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
      ::close(sock_fd_);
      sock_fd_ = -1;
      return false;
    }
    return true;
  }
  
  int read_some(uint8_t* data, int length, const std::string& session_id) {
     struct sockaddr_in from;
     socklen_t from_len = sizeof(from);
     int rv = recvfrom(sock_fd_, data, length, 0, (struct sockaddr*)&from, &from_len);
     if (rv < 0) {
      spdlog_error("[{0}] UdpStream::read_some: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
     }
     return (rv > 0 ? rv : 0);
  } 

  int write_some(uint8_t* data, int length, const std::string& session_id) {
    int s = sendto(sock_fd_, data, length, 0, (struct sockaddr *)&client_endpoint_, sizeof(client_endpoint_));
    if(s <= 0) {
      spdlog_error("[{0}] UdpStream::write_some: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
      return 0;
    }
    return s;
  }

  void close() {
    if (sock_fd_ > 0) {
      ::close(sock_fd_);
      sock_fd_ = -1;
    }
  }

private:
  int sock_fd_;
  struct sockaddr_in server_endpoint_;
  struct sockaddr_in client_endpoint_;
};
}  // namespace

#endif  // TINY_ROS_UDP_STREAM_H
