cmake_minimum_required(VERSION 3.16)
project(laba2 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
add_executable(laba2 main.cpp
    task1_asteroids.cpp task2_set.cpp task3_partition.cpp
    task4_subarray_sum.cpp task5_tree_depth.cpp
    task6_1_hash_tables.cpp task6_2_bloom_filter.cpp task7_lfu_cache.cpp
    lab1/Stack.cpp lab1/DynamicArray.cpp lab1/SinglyLinkedList.cpp
    lab1/DoublyLinkedList.cpp lab1/BST.cpp)
target_include_directories(laba2 PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})
target_compile_options(laba2 PRIVATE -Wall -Wextra -Wpedantic)
