cmake_minimum_required(VERSION 2.6)

set(EXECUTABLE_OUTPUT_PATH ${OUTPUT_PATHS}/bin)
add_executable(ExampleServiceClient
  ExampleServiceClient.cpp
  ${ROSLIB_PATHS}/ros.cpp
  ${ROSLIB_PATHS}/time.cpp
  ${ROSLIB_PATHS}/duration.cpp
  ${ROSLIB_PATHS}/log.cpp
)

target_link_libraries(ExampleServiceClient
  libpthread.so
)


