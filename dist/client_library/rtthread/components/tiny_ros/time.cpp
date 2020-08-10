#include <rtthread.h>
#include "tiny_ros/ros.h"
#include "tiny_ros/ros/time.h"

namespace tinyros
{
int64_t Time::time_start_ = 0;

int64_t Time::time_last_ = 0;

int64_t Time::time_dds_ = 0;

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

Time Time::dds()
{
  NodeHandle* nh = tinyros::nh();
  rt_mutex_take(&nh->sync_time_mutex_, RT_WAITING_FOREVER);
  Time time = Time::now();
  int64_t offset = (int64_t)(time.toMSec());
  offset = offset > Time::time_start_ && Time::time_start_ > 0 ? offset - Time::time_start_ : 0;
  time.sec = (uint32_t)(offset / 1000);
  time.nsec = (uint32_t)((offset % 1000) * 1000000);
  time.sec += (uint32_t)(Time::time_dds_ / 1000);
  time.nsec += (uint32_t)((Time::time_dds_ % 1000) * 1000000);
  normalizeSecNSec(time.sec, time.nsec);
  rt_mutex_release(&nh->sync_time_mutex_);
  return time;
}

Time Time::now()
{
  Time time;
  int64_t ms = rt_tick_get();
  time.sec = ms/1000;
  time.nsec = ms%1000 * 1000000;
  return time;
}

void setNow(Time & new_now) {

}
}

