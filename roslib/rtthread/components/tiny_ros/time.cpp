/*
 * File      : time.cpp
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#include <rtthread.h>
#include "tiny_ros/ros/time.h"

namespace tinyros
{
void normalizeSecNSec(uint32_t& sec, uint32_t& nsec)
{
  uint32_t nsec_part = nsec % 1000000000UL;
  uint32_t sec_part = nsec / 1000000000UL;
  sec += sec_part;
  nsec = nsec_part;
}

void normalizeSecNSecUnsigned(int64_t& sec, int64_t& nsec)
{
  int64_t nsec_part = nsec % 1000000000L;
  int64_t sec_part = sec + nsec / 1000000000L;
  if (nsec_part < 0)
  {
    nsec_part += 1000000000L;
    --sec_part;
  }

  sec = sec_part;
  nsec = nsec_part;
}

Time& Time::fromNSec(int32_t t)
{
  sec = t / 1000000000;
  nsec = t % 1000000000;
  normalizeSecNSec(sec, nsec);
  return *this;
}

Time& Time::operator +=(const Duration &rhs)
{
  sec += rhs.sec;
  nsec += rhs.nsec;
  normalizeSecNSec(sec, nsec);
  return *this;
}

Time& Time::operator -=(const Duration &rhs)
{
  sec += -rhs.sec;
  nsec += -rhs.nsec;
  normalizeSecNSec(sec, nsec);
  return *this;
}

Time Time::now()
{
  Time time;
  int64_t ms = rt_tick_get();
  time.sec = ms/1000;
  time.nsec = ms%1000 * 10000000;
  return time;
}

void setNow(Time & new_now) {

}
}

