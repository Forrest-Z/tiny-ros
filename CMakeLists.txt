cmake_minimum_required(VERSION 2.6)

set(OUTPUT_PATHS ${PROJECT_SOURCE_DIR}/build/output)
set(ROSLIB_PATHS ${OUTPUT_PATHS}/roslib/gcc/tiny_ros)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++11")

add_definitions(-DSPDLOG_COMPILED_LIB)

aux_source_directory(${PROJECT_SOURCE_DIR}/spdlog/src SPDLOG_SRC_LIST)

include_directories(
  ${OUTPUT_PATHS}/roslib/gcc
  ${PROJECT_SOURCE_DIR}/spdlog/include
)

add_subdirectory(core)
add_subdirectory(tools)
add_subdirectory(examples)

