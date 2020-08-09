cmake_minimum_required(VERSION 2.6)

set(OUTPUT_PATHS ${PROJECT_SOURCE_DIR}/build/output)
set(ROSLIB_PATHS ${OUTPUT_PATHS}/client_library/gcc/include/tiny_ros)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++11 -O3")

add_definitions(-DSPDLOG_COMPILED_LIB)

aux_source_directory(${PROJECT_SOURCE_DIR}/spdlog/src SPDLOG_SRC_LIST)

aux_source_directory(${PROJECT_SOURCE_DIR}/uWebSockets UWS_SRC_LIST)

include_directories(
  ${OUTPUT_PATHS}/client_library/gcc/include
  ${PROJECT_SOURCE_DIR}/spdlog/include
  ${PROJECT_SOURCE_DIR}/uWebSockets
)

add_subdirectory(core)
add_subdirectory(tools)
add_subdirectory(examples)
