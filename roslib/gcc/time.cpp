#include "tiny_ros/ros/time.h"
#include <limits.h>
#include <stdexcept>
#ifdef WIN32
  #include <windows.h>
  #include <sys/timeb.h>
#else
  #include <sys/time.h>
#endif

namespace tinyros
{
std::mutex Time::mutex_;

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

  if (sec_part < 0 || sec_part > UINT_MAX)
    throw std::runtime_error("Time is out of dual 32-bit range");

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

Time Time::dds() {
  std::unique_lock<std::mutex> lock(Time::mutex_);
  Time time = Time::now();
  int64_t offset = (int64_t)(time.toMSec());
  offset = offset > Time::time_start_ && Time::time_start_ > 0 ? offset - Time::time_start_ : 0;
  time.sec = (uint32_t)(offset / 1000);
  time.nsec = (uint32_t)((offset % 1000) * 1000000);
  time.sec += (uint32_t)(Time::time_dds_ / 1000);
  time.nsec += (uint32_t)((Time::time_dds_ % 1000) * 1000000);
  normalizeSecNSec(time.sec, time.nsec);
  return time;
}

Time Time::now()
{
  Time time;
#ifndef WIN32
#if HAS_CLOCK_GETTIME
    timespec start;
    clock_gettime(CLOCK_REALTIME, &start);
    time.sec  = start.tv_sec;
    time.nsec = start.tv_nsec;
#else
    struct timeval timeofday;
    gettimeofday(&timeofday,NULL);
    time.sec  = timeofday.tv_sec;
    time.nsec = timeofday.tv_usec * 1000;
#endif
#else
    static LARGE_INTEGER cpu_freq, init_cpu_time;
    static uint32_t start_sec = 0;
    static uint32_t start_nsec = 0;
    if (( start_sec == 0 ) && ( start_nsec == 0)) {
      QueryPerformanceFrequency(&cpu_freq);
      if (cpu_freq.QuadPart == 0) {
        throw std::runtime_error("This windows platform does not support the high-performance timing api.");
      }
      QueryPerformanceCounter(&init_cpu_time);
      FILETIME ft;
      GetSystemTimeAsFileTime(&ft);
      LARGE_INTEGER start_li;
      start_li.LowPart = ft.dwLowDateTime;
      start_li.HighPart = ft.dwHighDateTime;
#ifdef _MSC_VER
      start_li.QuadPart -= 116444736000000000Ui64;
#else
      start_li.QuadPart -= 116444736000000000ULL;
#endif
      start_sec = (uint32_t)(start_li.QuadPart / 10000000);
      start_nsec = (start_li.LowPart % 10000000) * 100;
    }
    LARGE_INTEGER cur_time;
    QueryPerformanceCounter(&cur_time);
    LARGE_INTEGER delta_cpu_time;
    delta_cpu_time.QuadPart = cur_time.QuadPart - init_cpu_time.QuadPart;
    double d_delta_cpu_time = delta_cpu_time.QuadPart / (double) cpu_freq.QuadPart;
    uint32_t delta_sec = (uint32_t) floor(d_delta_cpu_time);
    uint32_t delta_nsec = (uint32_t) round((d_delta_cpu_time-delta_sec) * 1e9);

    int64_t sec_sum  = (int64_t)start_sec  + (int64_t)delta_sec;
    int64_t nsec_sum = (int64_t)start_nsec + (int64_t)delta_nsec;

    normalizeSecNSecUnsigned(sec_sum, nsec_sum);

    time.sec = sec_sum;
    time.nsec = nsec_sum;
#endif
    return time;
}

void setNow(Time & new_now) {

}
}
