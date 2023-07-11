cmake_minimum_required(VERSION 3.13)
project(gtest_test)

# GoogleTest requires at least C++14
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/refs/tags/release-1.12.1.zip
  #URL https://github.com/google/googletest/archive/03597a01ee50ed33e9dfd640b249b4be3799d395.zip
)

enable_testing()

add_executable(
    hello_test
    hello_test.cpp
)

target_link_libraries(
    hello_test
    gtest_main
)

add_test(
    NAME hello_test
    COMMAND hello_test
    )
