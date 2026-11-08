#include "task7_lfu_cache.h"
#include <iostream>
#include <stdexcept>
#include <string>
using namespace std;
void initCache(Cache& c, int capacity) {
    c.capacity = capacity;
    initDoublyList(c.order);
    initArray(c.keys);
    initArray(c.values);
    c.frequencies = new long long[capacity]{};
}
static int findCache(const Cache& c, int key) {
    for (int i = 0; i < c.keys.size; ++i) {
        if (c.keys.data[i] == key) {
            return i;
        }
    }
    return -1;
}
static void touch(Cache& c, int index) {
    ++c.frequencies[index];
    eraseDoublyValue(c.order, index);
    pushBackDoubly(c.order, index);
}
int cacheGet(Cache& c, int key) {
    int index = findCache(c, key);
    if (index == -1) {
        return -1;
    }
    touch(c, index);
    return c.values.data[index];
}
void cacheSet(Cache& c, int key, int value) {
    if (c.capacity == 0) {
        return;
    }
    int index = findCache(c, key);
    if (index != -1) {
        c.values.data[index] = value;
        touch(c, index);
        return;
    }
    if (c.keys.size < c.capacity) {
        index = c.keys.size;
        pushBack(c.keys, key);
        pushBack(c.values, value);
    } else {
        index = c.order.head->data;
        for (DoubleNode* node = c.order.head; node; node = node->next) {
            if (c.frequencies[node->data] < c.frequencies[index]) {
                index = node->data;
            }
        }
        eraseDoublyValue(c.order, index);
        c.keys.data[index] = key;
        c.values.data[index] = value;
    }
    c.frequencies[index] = 1;
    pushBackDoubly(c.order, index);
}
void freeCache(Cache& c) {
    freeDoublyList(c.order);
    freeArray(c.keys);
    freeArray(c.values);
    delete[] c.frequencies;
    c.frequencies = nullptr;
}
void runCache() {
    int capacity, q;
    if (!(cin >> capacity >> q) || capacity < 0 || q < 0) {
        throw runtime_error("Expected nonnegative capacity and query count");
    }
    Cache c{};
    initCache(c, capacity);
    bool first = true;
    try {
        for (int i = 0; i < q; ++i) {
            string op;
            int key, value;
            if (!(cin >> op >> key)) {
                throw runtime_error("Expected cache operation");
            }
            if (op == "SET") {
                if (!(cin >> value)) {
                    throw runtime_error("Expected value");
                }
                cacheSet(c, key, value);
            } else if (op == "GET") {
                if (!first) {
                    cout << ' ';
                }
                first = false;
                cout << cacheGet(c, key);
            } else {
                throw runtime_error("Unknown cache operation");
            }
        }
        cout << '\n';
    } catch (...) {
        freeCache(c);
        throw;
    }
    freeCache(c);
}
