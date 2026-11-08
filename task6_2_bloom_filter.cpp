#include "task6_2_bloom_filter.h"
#include <iostream>
#include <stdexcept>
#include <string>
using namespace std;
static int bloomPosition(const ChainTable& t, int key, int i) {
    return static_cast<int>(hashKey(key, i + 1) % (t.bloom.size * UINT64_C(8)));
}
void bloomAdd(ChainTable& t, int key) {
    for (int i = 0; i < 3; ++i) {
        int p = bloomPosition(t, key, i);
        t.bloom.data[p / 8] |= 1 << (p % 8);
    }
}
bool bloomMaybeContains(const ChainTable& t, int key) {
    for (int i = 0; i < 3; ++i) {
        int p = bloomPosition(t, key, i);
        if (!(t.bloom.data[p / 8] & (1 << (p % 8)))) {
            return false;
        }
    }
    return true;
}
void runHashTask() {
    int q;
    if (!(cin >> q) || q < 0) {
        throw runtime_error("Expected query count");
    }
    ChainTable t{};
    initChain(t);
    try {
        for (int i = 0; i < q; ++i) {
            string op;
            int key;
            if (!(cin >> op >> key)) {
                throw runtime_error("Expected operation and key");
            }
            if (op == "ADD") {
                chainAdd(t, key);
            } else if (op == "DEL") {
                chainDelete(t, key);
            } else if (op == "GET") {
                cout << chainContains(t, key) << '\n';
            } else if (op == "BLOOM") {
                cout << bloomMaybeContains(t, key) << '\n';
            } else {
                throw runtime_error("Unknown hash operation");
            }
        }
    } catch (...) {
        freeChain(t);
        throw;
    }
    freeChain(t);
}
