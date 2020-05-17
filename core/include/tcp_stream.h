/*
 * File      : tcp_stream.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-03     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_TCP_STREAM_H
#define TINY_ROS_TCP_STREAM_H
#include <iostream>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include "common.h"

namespace tinyros
{
class TcpStream
{
public:
  TcpStream(int fd): sock_fd_(fd)  {
    int opt = 1;
    struct linger so_linger;
    so_linger.l_onoff = 1;
    so_linger.l_linger = 0;
    setsockopt(sock_fd_, IPPROTO_TCP, TCP_NODELAY, (const char *)&opt, sizeof(opt));
    setsockopt(sock_fd_, SOL_SOCKET, SO_LINGER, (const char *)&so_linger, sizeof(so_linger));
  }
  
  int write_some(uint8_t* data, int length, const std::string& session_id) {
   int rv, len = length, totalsent = 0;
    do {
      rv = ::write(sock_fd_, data + totalsent, len - totalsent);
      if (rv > 0) {
        totalsent += rv;
      } else if (rv == 0) {
        spdlog_error("[{0}] TcpStream::write_some socket close: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
        return -1;
      } else {
        if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR)) {
          spdlog_error("[{0}] TcpStream::write_some error: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
          return -1;
        }
      }
    } while (totalsent < len);

    return totalsent;
  } 

  int read_some(uint8_t* data, int length, const std::string& session_id) {
    int rv = ::read(sock_fd_, data, length);
    if (rv > 0) {
      return rv;
    } else if (rv == 0) {
      spdlog_error("[{0}] TcpStream::read_some socket close: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
      return -1;
    } else {
      if ((errno != EAGAIN) && (errno != EWOULDBLOCK) && (errno != EINTR))  {
        spdlog_error("[{0}] TcpStream::read_some error: {1}(errno: {2})", session_id.c_str(), strerror(errno), errno);
        return -1;
      }
      return 0;
    }
  }
  
  void close() {
    ::close(sock_fd_);
  }

private:
  int sock_fd_;
};
}  // namespace

#endif  // TINY_ROS_TCP_STREAM_H

