#include <math.h>
#include "tiny_ros/ros/duration.h"

namespace tinyros
{
void normalizeSecNSecSigned(int32_t &sec, int32_t &nsec)
{
  int32_t nsec_part = nsec;
  int32_t sec_part = sec;

  while (nsec_part > 1000000000L)
  {
    nsec_part -= 1000000000L;
    ++sec_part;
  }
  while (nsec_part < 0)
  {
    nsec_part += 1000000000L;
    --sec_part;
  }
  sec = sec_part;
  nsec = nsec_part;
}

Duration& Duration::operator+=(const Duration &rhs)
{
  sec += rhs.sec;
  nsec += rhs.nsec;
  normalizeSecNSecSigned(sec, nsec);
  return *this;
}

Duration& Duration::operator-=(const Duration &rhs)
{
  sec += -rhs.sec;
  nsec += -rhs.nsec;
  normalizeSecNSecSigned(sec, nsec);
  return *this;
}

Duration& Duration::operator*=(double scale)
{
  sec *= scale;
  nsec *= scale;
  normalizeSecNSecSigned(sec, nsec);
  return *this;
}

}

