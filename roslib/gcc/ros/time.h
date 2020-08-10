#ifndef TINYROS_TIME_H_
#define TINYROS_TIME_H_

#include "tiny_ros/ros/duration.h"
#include <math.h>
#include <stdint.h>
#include <mutex>

namespace tinyros
{
void normalizeSecNSec(uint32_t &sec, uint32_t &nsec);
void normalizeSecNSecUnsigned(int64_t& sec, int64_t& nsec);

class Time
{
public:
  uint32_t sec, nsec;

  Time() : sec(0), nsec(0) {}
  Time(uint32_t _sec, uint32_t _nsec) : sec(_sec), nsec(_nsec)
  {
    normalizeSecNSec(sec, nsec);
  }

  double toMSec() const
  {
    return (double)sec * 1000 + 1e-6 * (double)nsec;
  };

  double toSec() const
  {
    return (double)sec + 1e-9 * (double)nsec;
  };
  void fromSec(double t)
  {
    sec = (uint32_t) floor(t);
    nsec = (uint32_t) round((t - sec) * 1e9);
  };

  uint32_t toNsec()
  {
    return (uint32_t)sec * 1000000000ull + (uint32_t)nsec;
  };
  Time& fromNSec(int32_t t);

  Time& operator +=(const Duration &rhs);
  Time& operator -=(const Duration &rhs);

  static std::mutex mutex_;
  static int64_t time_start_;
  static int64_t time_last_;
  static int64_t time_dds_;
  static Time dds();
  static Time now();
  static void setNow(Time & new_now);
};

}

#endif
