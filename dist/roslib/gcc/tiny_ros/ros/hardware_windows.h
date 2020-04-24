/*
 * File      : hardware_windows.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_WINDOWS_H_
#define TINYROS_HARDWARE_WINDOWS_H_
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
#include <string>
#include <iostream>
#include <winsock2.h>
#include <ws2tcpip.h>
#include "hardware.h"
#pragma comment(lib, "Ws2_32.lib")

namespace tinyros {
#undef DEFAULT_PORTNUM
#define DEFAULT_PORTNUM "11315"
class HardwareWindows: public Hardware
{
public:
  HardwareWindows(): connected_(false), mySocket(INVALID_SOCKET) {

  }
  
  virtual bool init(std::string portName) {
    WSADATA wsaData;

    this->close();

    int result = WSAStartup (MAKEWORD (2, 2), &wsaData);
    if (result) {
      std::cerr << "Could not initialize windows socket (" << result << ")" << std::endl;
      return false;
    }

    struct addrinfo *servers = get_server_addr(portName);
    if (NULL == servers) {
      WSACleanup ();
      return false;
    }

    connect_to_server(servers);

    freeaddrinfo(servers);

    if (INVALID_SOCKET == mySocket) {
      std::cerr << "Could not connect to server" << std::endl;
      WSACleanup ();
      return false;
    }
    connected_ = true;
    return connected_;
  }

  virtual int read(uint8_t* data, int length) {
    if (connected_) {
      int result = recv (mySocket, (char *)data, length, 0);
      if (result > 0) {
        return result;
      } else if (result < 0) {
        int error = WSAGetLastError();
        if ((error != WSAEWOULDBLOCK) && (error != WSAEINTR)) {
          std::cerr << "Failed to receive data from server " << WSAGetLastError() << std::endl;
          this->close();
          return -1;
        }
      } else {
        std::cerr << "Connection to server closed" << std::endl;
        this->close();
        return -1;
      }
    }
    return -1;
  }

  virtual bool write(uint8_t* data, int length) {
    if (connected_) {
      int result = send(mySocket, (const char *)data, length, 0);
      if (SOCKET_ERROR == result) {
        std::cerr << "Send failed with error " << WSAGetLastError () << std::endl;
        this->close();
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
      if (mySocket!= INVALID_SOCKET) {
        closesocket (mySocket);
        mySocket = INVALID_SOCKET;
        WSACleanup ();
      }
    }
  }

protected:
  struct addrinfo *get_server_addr (const std::string & hostname) {
    int result;
    struct addrinfo *ai_output = NULL;
    struct addrinfo ai_input;

    // split off the port number if given
    int c = hostname.find_last_of (':');
    std::string host = hostname.substr (0, c);
    std::string port = (c < 0) ? DEFAULT_PORTNUM : hostname.substr (c + 1);

    ZeroMemory (&ai_input, sizeof (ai_input));
    ai_input.ai_family = AF_UNSPEC;
    ai_input.ai_socktype = SOCK_STREAM;
    ai_input.ai_protocol = IPPROTO_TCP;

    // Resolve the server address and port
    result = getaddrinfo (host.c_str (), port.c_str (), &ai_input, &ai_output);
    if (result != 0) {
      std::cerr << "Could not resolve server address (" << result << ")" << std::endl;
      return NULL;
    }
    return ai_output;
  }

  void connect_to_server (struct addrinfo *servers) {
    int result;
    for (struct addrinfo * ptr = servers; ptr != NULL; ptr = ptr->ai_next) {
      mySocket = socket (ptr->ai_family, ptr->ai_socktype, ptr->ai_protocol);
      if (INVALID_SOCKET == mySocket) {
        std::cerr << "Could not great socket " << WSAGetLastError ();
        return;
      }

      result = connect (mySocket, ptr->ai_addr, (int) ptr->ai_addrlen);
      if (SOCKET_ERROR == result){
        closesocket (mySocket);
        mySocket = INVALID_SOCKET;
      } else {
        break;
      }
    }

    // disable nagle's algorithm
    char value = 1;
    setsockopt (mySocket, IPPROTO_TCP, TCP_NODELAY, &value, sizeof (value));
  }

private:
  SOCKET mySocket;
  bool connected_;
};
}
#endif
#endif // TINYROS_HARDWARE_WINDOWS_H_
