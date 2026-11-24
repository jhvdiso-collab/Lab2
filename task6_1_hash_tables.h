#pragma once
#include "lab1/DynamicArray.h"
#include "lab1/SinglyLinkedList.h"
#include <cstdint>
struct ChainTable {
    SinglyNode** buckets;
    int capacity;
    int size;
    DynamicArray bloom;
};
struct CuckooTable {
    DynamicArray keys;
    DynamicArray values;
    DynamicArray used;
    int size;
    std::uint64_t seed;
};
std::uint64_t hashKey(int key, std::uint64_t seed);
void initChain(ChainTable& t, int capacity = 16);
bool chainContains(const ChainTable& t, int key);
bool chainGet(const ChainTable& t, int key, int& value);
bool chainAdd(ChainTable& t, int key, int value);
bool chainDelete(ChainTable& t, int key);
void freeChain(ChainTable& t);
void initCuckoo(CuckooTable& t, int capacity = 16);
bool cuckooContains(const CuckooTable& t, int key);
bool cuckooGet(const CuckooTable& t, int key, int& value);
bool cuckooAdd(CuckooTable& t, int key, int value);
bool cuckooDelete(CuckooTable& t, int key);
void freeCuckoo(CuckooTable& t);
void runBenchmark();
