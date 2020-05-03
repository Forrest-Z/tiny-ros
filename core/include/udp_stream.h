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
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include "session.h"

namespace tinyros
{

using boost::asio::ip::udp;

class UdpStream : public udp::socket
{
public:
  explicit UdpStream(boost::asio::io_service& io_service) : udp::socket(io_service)
  {
  }

  void open(udp::endpoint server_endpoint, udp::endpoint client_endpoint)
  {
    boost::system::error_code ec;
    const protocol_type protocol = server_endpoint.protocol();

    udp::socket::open(protocol, ec);
    boost::asio::detail::throw_error(ec, "open");
    udp::socket::bind(server_endpoint, ec);
    boost::asio::detail::throw_error(ec, "bind");

    client_endpoint_ = client_endpoint;
  }
  
  template <typename ConstBufferSequence>
  std::size_t write_some(const ConstBufferSequence& buffers)
  {
    boost::system::error_code ec;
    std::size_t s = this->get_service().send_to(
        this->get_implementation(), buffers, client_endpoint_, 0, ec);
    boost::asio::detail::throw_error(ec, "send_to");
    return s;
  }  

  template <typename ConstBufferSequence>
  std::size_t write_some(const ConstBufferSequence& buffers,
      boost::system::error_code& ec)
  {
    std::size_t s = this->get_service().send_to(
        this->get_implementation(), buffers, client_endpoint_, 0, ec);
    boost::asio::detail::throw_error(ec, "send_to");
    return s;
  }

  template <typename MutableBufferSequence>
  std::size_t read_some(const MutableBufferSequence& buffers)
  {
    boost::system::error_code ec;
    std::size_t s = this->get_service().receive_from(
        this->get_implementation(), buffers, client_endpoint_, 0, ec);
    boost::asio::detail::throw_error(ec, "receive_from");
    return s;
  }

  template <typename MutableBufferSequence>
  std::size_t read_some(const MutableBufferSequence& buffers,
      boost::system::error_code& ec)
  {
    std::size_t s = this->get_service().receive_from(
        this->get_implementation(), buffers, client_endpoint_, 0, ec);
    boost::asio::detail::throw_error(ec, "receive_from");
    return s;
  }

private:
  udp::endpoint client_endpoint_;
};

}  // namespace

#endif  // TINY_ROS_UDP_STREAM_H
