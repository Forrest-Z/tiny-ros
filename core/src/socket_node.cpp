/*
 * File      : socket_node.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */

#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <boost/filesystem.hpp>
#include <boost/log/trivial.hpp>
#include <boost/log/utility/setup/common_attributes.hpp>
#include <boost/log/utility/setup/formatter_parser.hpp>
#include <boost/log/utility/setup/filter_parser.hpp>
#include <boost/log/utility/setup/settings.hpp>
#include <boost/log/utility/setup/from_settings.hpp>
#include "udp_socket_session.h"
#include "tcp_server.h"

#define LOG_ROTATION_SIZE  2*1024*1024 //2MB

#define LOG_MAX_SIZE       10*1024*1024 //10MB

static void init_log_environment() {
  namespace logging = boost::log;
  using namespace logging::trivial;

  logging::settings setts;
  logging::add_common_attributes();
  logging::register_simple_formatter_factory<severity_level, char>("Severity");
  logging::register_simple_filter_factory<severity_level, char>("Severity");

  setts["Core"]["Filter"] = "%Severity% >= info";
  setts["Core"]["DisableLogging"] = false;

  // Sinks.console
  setts["Sinks.console.Format"] = "[%TimeStamp%] [%Severity%] %Message%";
  setts["Sinks.console.Destination"] = "Console";
  setts["Sinks.console.Filter"] = "%Severity% >= info";
  setts["Sinks.console.Asynchronous"] = false;
  setts["Sinks.console.AutoFlush"] = true;

  // Sinks.trace
  setts["Sinks.trace.Filter"] = "%Severity% >= info";
  setts["Sinks.trace.Format"] = "[%TimeStamp%] [%Severity%] %Message%";
  setts["Sinks.trace.Destination"] = "TextFile";
  setts["Sinks.trace.FileName"] = "tinyros-logs/%Y%m%d_%H%M%S.log";
  setts["Sinks.trace.Target"] = "tinyros-logs";
  setts["Sinks.trace.Asynchronous"] = false;
  setts["Sinks.trace.AutoFlush"] = true;
  setts["Sinks.trace.RotationSize"] = LOG_ROTATION_SIZE;
  setts["Sinks.trace.MaxSize"] = LOG_MAX_SIZE;
  setts["Sinks.trace.ScanForFiles"] = "All";

  logging::init_from_settings(setts);
}

using boost::asio::ip::udp;
using boost::asio::ip::address_v4;

static void udp_service_run(int server_port, int client_port) {
  boost::asio::io_service io_service;
  tinyros::UdpSocketSession udp_socket_session(
      io_service,
      udp::endpoint(address_v4::from_string("127.0.0.1"), server_port),
      udp::endpoint(udp::v4(), client_port));
  io_service.run();
}

int main(int argc, char* argv[]) {
  int tcp_server_port = 11315;
  int udp_server_port = 11316;
  int udp_client_port = 11317;

  init_log_environment();

  boost::thread tid(boost::bind(udp_service_run, udp_server_port, udp_client_port));
  tid.detach();

  boost::asio::io_service io_service;
  tinyros::TcpServer<> tcp_server(io_service, tcp_server_port);
  BOOST_LOG_TRIVIAL(info) << "tinyros: TCP listening for connections on port " << tcp_server_port;
  io_service.run();
  return 0;
}
