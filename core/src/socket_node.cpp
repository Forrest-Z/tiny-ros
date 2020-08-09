/*
 * File      : socket_node.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */
#include <signal.h>
#include "udp_stream.h"
#include "tcp_server.h"
#include "session.h"
#include "signals.h"
#include "common.h"

#define LOG_ROTATION_PATH    "tinyros_logs/"
#define LOG_ROTATION_SIZE    (2*1024*1024) //2MB
#define LOG_ROTATION_FILES   (5)

static void udp_service_run(int server_port, int client_port) {
  tinyros::UdpStream stream;
  tinyros::Session<tinyros::UdpStream> new_session(stream, tinyros::UDP_STREAM);
  while (1) {
    if (!new_session.is_active()) {
      if (new_session.socket().open(server_port, client_port, new_session.session_id_)) {
        new_session.start();
      }
    }
    sleep(1);
  }
}

#ifdef WITH_WEBSOCKETS
#include "uWS.h"
static void web_service_run(int web_server_port) {
  try {
    uWS::Server server(web_server_port);
    
    server.onConnection([](uWS::WebSocket socket) {
      tinyros::Session<uWS::WebSocket>* session = new tinyros::Session<uWS::WebSocket>(socket, tinyros::WEB_STREAM);
      socket.setData(session);
      session->start();
    });

    server.onMessage([](uWS::WebSocket socket, char *message, std::size_t length, uWS::OpCode opCode) {
      void *session = socket.getData();
      if (session && opCode == uWS::BINARY) {
        ((tinyros::Session<uWS::WebSocket>*)session)->consume_message((uint8_t*)message, (int)length);
      }
    });
    
    server.onDisconnection([](uWS::WebSocket socket, int code, char *message, std::size_t length) {
      void *session = socket.getData();
      if (session) {
        ((tinyros::Session<uWS::WebSocket>*)session)->stop();
        socket.setData(nullptr);
        delete ((tinyros::Session<uWS::WebSocket>*)session);
      }
    });
    
    server.run();
  } catch (...) {
    spdlog_error("web_service_run error: {0}(errno: {1})", strerror(errno), errno);
  }
}
#endif

int main(int argc, char* argv[]) {
#ifdef __linux__
  signal(SIGPIPE, SIG_IGN);
#endif
    
  int tcp_server_port = 11315;
  int udp_server_port = 11316;
  int udp_client_port = 11317;
  int web_server_port = 11318;

  auto stdout_sink = std::make_shared<spdlog::sinks::stdout_sink_mt>();
  auto rotating_file_sink = std::make_shared<spdlog::sinks::rotating_file_sink_mt>
    (LOG_ROTATION_PATH "tinyrosdds.log", LOG_ROTATION_SIZE, LOG_ROTATION_FILES);
  stdout_sink->set_level(spdlog::level::warn);
  rotating_file_sink->set_level(spdlog::level::trace);
  spdlog::sinks_init_list sinks = {rotating_file_sink, stdout_sink};
  auto logger = std::make_shared<spdlog::logger>("logger", sinks);
  logger->set_level(spdlog::level::trace);
  logger->flush_on(spdlog::level::trace);
  logger->set_pattern("[%Y-%m-%d %H:%M:%S.%e] [%l] %v");
  spdlog::register_logger(logger);

  std::thread tidudp(std::bind(udp_service_run, udp_server_port, udp_client_port));
  tidudp.detach();

#ifdef WITH_WEBSOCKETS
  std::thread tidws(std::bind(web_service_run, web_server_port));
  tidws.detach();
#endif

  tinyros::TcpServer tcp_server(tcp_server_port);
  tcp_server.start_accept();
  
  return 0;
}
