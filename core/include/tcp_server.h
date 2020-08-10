#ifndef TINY_ROS_TCP_SERVER_H
#define TINY_ROS_TCP_SERVER_H
#include <iostream>
#include <sys/socket.h>  
#include <netinet/in.h>  
#include <arpa/inet.h>
#include "common.h"
#include "tcp_stream.h"
#include "session.h"

namespace tinyros
{
class TcpServer: public TcpServer_
{
public:
  TcpServer(int port): port_(port) {
  }

  void start_accept() {
    int listen_fd;
    struct sockaddr_in server_addr;
    if ((listen_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
      spdlog_error("TcpServer::start_accept create socket error: {0}(errno: {1})", strerror(errno), errno);
      return;
    }
    
    int opt = 1;
    setsockopt(listen_fd, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));

    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    server_addr.sin_port = htons(port_);

    if (bind(listen_fd, (struct sockaddr *)&server_addr, sizeof(server_addr)) == -1) {
      spdlog_error("TcpServer::start_accept bind socket error: {0}(errno: {1})", strerror(errno), errno);
      return;
    }

    if (listen(listen_fd, 100) == -1) {
      spdlog_error("TcpServer::start_accept listen socket error: {0}(errno: {1})", strerror(errno), errno);
      return;
    }
    
    spdlog_info("TCP Listening for connections on port: {0:d}", port_);

    while (1) {
      int connect_fd;
      struct sockaddr_in client;
      socklen_t len = sizeof(client);
      if ((connect_fd = accept(listen_fd, (struct sockaddr *)&client, &len)) == -1) {
        spdlog_error("TcpServer::start_accept accept socket error: {0}(errno: {1})", strerror(errno), errno);
        continue;
      }
      
      std::unique_lock<std::mutex> sessions_lock(TcpServer::sessions_mutex_);
      TcpStream stream(connect_fd);
      TcpServer::sessions_[connect_fd] = SessionPtr(new Session<TcpStream>(stream, tinyros::TCP_STREAM));
      TcpServer::sessions_[connect_fd]->start();
    }
  }
  
  int port_;
};

}  // namespace

#endif  // TINY_ROS_TCP_SERVER_H
