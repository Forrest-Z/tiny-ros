cmake_minimum_required(VERSION 2.6)

set(OUTPUT_PATHS ${PROJECT_SOURCE_DIR}/build/output)
set(ROSLIB_PATHS ${OUTPUT_PATHS}/roslib/gcc/tiny_ros)
set(CMAKE_C_FLAGS "-std=gnu99 -O3 -fexceptions -fPIC -Wl,-v -Wall -Wno-strict-aliasing -Wno-unknown-pragmas -Wno-deprecated-declarations")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-unused-function -Wno-unused-variable -Wno-unused-parameter -Wno-switch")
set(CMAKE_CXX_FLAGS "-std=c++11 -O3 -fexceptions -fPIC -Wl,-v -Wall -Wno-strict-aliasing -Wno-unknown-pragmas -Wno-switch")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wno-reorder -Wno-unused-function -Wno-unused-variable -Wno-unused-parameter -Wno-deprecated-declarations")

add_definitions(-DBOOST_LOG_DYN_LINK)

include_directories(
  ${OUTPUT_PATHS}/roslib/gcc
  ${Boost_INCLUDE_DIRS}
)

add_subdirectory(core)
add_subdirectory(tools)
add_subdirectory(examples)

