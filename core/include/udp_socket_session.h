/*
 * File      : udp_socket_session.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-02-09     Pinkie.Fu    initial version
 */

#ifndef TINY_ROS_UDP_SOCKET_SESSION_H
#define TINY_ROS_UDP_SOCKET_SESSION_H

#include <iostream>
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include "session.h"
#include "udp_stream.h"


namespace tinyros
{

using boost::asio::ip::udp;

class UdpSocketSession : public Session<UdpStream>
{
public:
  UdpSocketSession(boost::asio::io_service& io_service,
                   udp::endpoint server_endpoint,
                   udp::endpoint client_endpoint)
    : Session(io_service, true), timer_(io_service),
      server_endpoint_(server_endpoint), client_endpoint_(client_endpoint)
  {
    BOOST_LOG_TRIVIAL(info) << "tinyros: UDP created server: " << server_endpoint << ", client: " << client_endpoint;
    check_connection();
  }

private:
  void check_connection()
  {
    if (!is_active())
    {
      socket().open(server_endpoint_, client_endpoint_);
      start();
    }

    timer_.expires_from_now(boost::posix_time::milliseconds(2000));
    timer_.async_wait(boost::bind(&UdpSocketSession::check_connection, this));
  }

  boost::asio::deadline_timer timer_;
  udp::endpoint server_endpoint_;
  udp::endpoint client_endpoint_;
};

}  // namespace

#endif  // TINY_ROS_UDP_SOCKET_SESSION_H