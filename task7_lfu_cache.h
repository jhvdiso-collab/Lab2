#pragma once

#include "lab1/DoublyLinkedList.h"

#include "lab1/DynamicArray.h"

struct Cache {

    DoublyList order;

    DynamicArray keys;

    DynamicArray values;

    long long* frequencies;

    int capacity;
};

void initCache(Cache& c, int capacity);

void cacheSet(Cache& c, int key, int value);

int cacheGet(Cache& c, int key);

void freeCache(Cache& c);

void runCache();
