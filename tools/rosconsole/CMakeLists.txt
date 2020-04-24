cmake_minimum_required(VERSION 2.6)

set(EXECUTABLE_OUTPUT_PATH ${OUTPUT_PATHS}/bin)
add_executable(tinyrosconsole
  rosconsole.cpp
  ${ROSLIB_PATHS}/ros.cpp
  ${ROSLIB_PATHS}/time.cpp
  ${ROSLIB_PATHS}/duration.cpp
  ${ROSLIB_PATHS}/log.cpp
)

target_link_libraries(tinyrosconsole
  libboost_log.so
  libboost_log_setup.so
  libboost_atomic.so 
  libboost_chrono.so
  libboost_date_time.so
  libboost_regex.so
  libboost_system.so
  libboost_filesystem.so
  libboost_thread.so
  libpthread.so
)
