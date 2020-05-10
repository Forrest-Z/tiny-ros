/*
 * File      : socket_node.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */
#include "udp_stream.h"
#include "tcp_server.h"
#include "session.h"
#include "signals.h"

static void udp_service_run(int server_port, int client_port) {
  tinyros::UdpStream stream;
  tinyros::Session<tinyros::UdpStream> new_session(stream, true);
  while (1) {
    if (!new_session.is_active()) {
      if (new_session.socket().open(server_port, client_port, new_session.session_id_)) {
        new_session.start();
      }
    }
    sleep(1);
  }
}

int main(int argc, char* argv[]) {
  int tcp_server_port = 11315;
  int udp_server_port = 11316;
  int udp_client_port = 11317;

  std::thread tidudp(std::bind(udp_service_run, udp_server_port, udp_client_port));
  tidudp.detach();

  tinyros::TcpServer tcp_server(tcp_server_port);
  tcp_server.start_accept();
  
  return 0;
}
