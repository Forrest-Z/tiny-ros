/*
 * File      : tcp_server.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_TCP_SERVER_H
#define TINY_ROS_TCP_SERVER_H
#include <iostream>
#include <sys/socket.h>  
#include <netinet/in.h>  
#include <arpa/inet.h> 
#include "tcp_stream.h"
#include "session.h"

namespace tinyros
{
class TcpServer
{
public:
  TcpServer(int port): port_(port) {
  }

  void start_accept() {
    int listen_fd;
    struct sockaddr_in server_addr;
    if ((listen_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
      printf("TcpServer::start_accept create socket error: %s(errno: %d)\n", strerror(errno), errno);
      return;
    }
    
    int opt = 1;
    setsockopt(listen_fd, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));

    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    server_addr.sin_port = htons(port_);

    if (bind(listen_fd, (struct sockaddr *)&server_addr, sizeof(server_addr)) == -1) {
      printf("TcpServer::start_accept bind socket error: %s(errno: %d)\n", strerror(errno), errno);
      return;
    }

    if (listen(listen_fd, 100) == -1) {
      printf("TcpServer::start_accept listen socket error: %s(errno: %d)\n", strerror(errno), errno);
      return;
    }

    while (1) {
      int connect_fd;
      struct sockaddr_in client;
      socklen_t len = sizeof(client);
      if ((connect_fd = accept(listen_fd, (struct sockaddr *)&client, &len)) == -1) {
        printf("TcpServer::start_accept accept socket error: %s(errno: %d)\n", strerror(errno), errno);
        continue;
      }

      TcpStream stream(connect_fd);
      Session<TcpStream>* new_session = new Session<TcpStream>(stream);
      new_session->start();
    }
  }
  
  int port_;
};

}  // namespace

#endif  // TINY_ROS_TCP_SERVER_H