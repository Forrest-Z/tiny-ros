cmake_minimum_required(VERSION 2.6)

set(EXECUTABLE_OUTPUT_PATH output)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++11 -O3")

aux_source_directory(include/tiny_ros TINYROS_SRC_LIST)

include_directories(include)

add_executable(ExamplePublisher
  src/examples/publisher/ExamplePublisher.cpp
  ${TINYROS_SRC_LIST}
)

add_executable(ExampleSubscriber
  src/examples/subscriber/ExampleSubscriber.cpp
  ${TINYROS_SRC_LIST}
)

add_executable(ExampleService
  src/examples/service/ExampleService.cpp
  ${TINYROS_SRC_LIST}
)

add_executable(ExampleServiceClient
  src/examples/service_client/ExampleServiceClient.cpp
  ${TINYROS_SRC_LIST}
)

IF (CMAKE_SYSTEM_NAME MATCHES "CYGWIN")
  target_link_libraries(ExamplePublisher libpthread.a)
  target_link_libraries(ExampleSubscriber libpthread.a)
  target_link_libraries(ExampleService libpthread.a)
  target_link_libraries(ExampleServiceClient libpthread.a)
ENDIF()

IF (CMAKE_SYSTEM_NAME MATCHES "Linux")
  target_link_libraries(ExamplePublisher libpthread.so)
  target_link_libraries(ExampleSubscriber libpthread.so)
  target_link_libraries(ExampleService libpthread.so)
  target_link_libraries(ExampleServiceClient libpthread.so)
ENDIF()

