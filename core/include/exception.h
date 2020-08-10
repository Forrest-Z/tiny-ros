#ifndef TINY_ROS_EXCEPTION_H
#define TINY_ROS_EXCEPTION_H

#include <stdexcept>

namespace tinyros
{
class Exception : public std::runtime_error
{
public:
  Exception(const std::string& what)
  : std::runtime_error(what)
  {}
};

} // namespace ros

#endif

