#include "serialization.h"

namespace tinyros
{
namespace serialization
{
void throwStreamOverrun()
{
  throw StreamOverrunException("Buffer Overrun");
}
}
}
