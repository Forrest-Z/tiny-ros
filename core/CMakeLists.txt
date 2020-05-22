cmake_minimum_required(VERSION 2.6)

set(EXECUTABLE_OUTPUT_PATH ${OUTPUT_PATHS}/bin)
aux_source_directory(src SRC_LIST)
include_directories(include)

add_executable(tinyroscore
  ${SRC_LIST}
  ${SPDLOG_SRC_LIST}
  ${ROSLIB_PATHS}/time.cpp
)

IF (CMAKE_SYSTEM_NAME MATCHES "CYGWIN")
  target_link_libraries(tinyroscore
    libpthread.a
  )
ENDIF()

IF (CMAKE_SYSTEM_NAME MATCHES "Linux")
  target_link_libraries(tinyroscore
    libpthread.so
  )
ENDIF()


