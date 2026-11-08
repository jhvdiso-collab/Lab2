#include "task6_1_hash_tables.h"

#include "task6_2_bloom_filter.h"

#include <algorithm>

#include <chrono>

#include <climits>

#include <iostream>

#include <stdexcept>

#include <string>

using namespace std;

uint64_t hashKey(int key, uint64_t seed) {

    uint64_t x = static_cast<uint32_t>(key) + seed;

    x = (x ^ (x >> 30)) * UINT64_C(0xbf58476d1ce4e5b9);

    x = (x ^ (x >> 27)) * UINT64_C(0x94d049bb133111eb);

    return x ^ (x >> 31);
}

static int doubled(int n) {

    if (n > INT_MAX / 2) {

        throw length_error("Table too large");
    }

    return n * 2;
}

static void fillArray(DynamicArray& a, int n) {

    initArray(a, n);

    for (int i = 0; i < n; ++i) {

        pushBack(a, 0);
    }
}

void initChain(ChainTable& t, int capacity) {

    t.capacity = max(16, capacity);

    t.buckets = new SinglyNode* [t.capacity] {
    };

    t.size = 0;

    fillArray(t.bloom, t.capacity);
}

bool chainContains(const ChainTable& t, int key) {

    if (!bloomMaybeContains(t, key)) {

        return false;
    }

    int p = static_cast<int>(hashKey(key, 0) % t.capacity);

    return findSingly(t.buckets[p], key) >= 0;
}

static void growChain(ChainTable& t) {

    ChainTable next{};

    initChain(next, doubled(t.capacity));

    for (int i = 0; i < t.capacity; ++i) {

        for (SinglyNode* p = t.buckets[i]; p; p = p->next) {

            chainAdd(next, p->data);
        }
    }

    freeChain(t);

    t = next;
}

bool chainAdd(ChainTable& t, int key) {

    if (chainContains(t, key)) {

        return false;
    }

    if (t.size >= t.capacity) {

        growChain(t);
    }

    int p = static_cast<int>(hashKey(key, 0) % t.capacity);

    pushFront(t.buckets[p], key);

    bloomAdd(t, key);

    ++t.size;

    return true;
}

bool chainDelete(ChainTable& t, int key) {

    int p = static_cast<int>(hashKey(key, 0) % t.capacity);

    if (!eraseSinglyValue(t.buckets[p], key)) {

        return false;
    }

    --t.size;

    return true;
}

void freeChain(ChainTable& t) {

    for (int i = 0; i < t.capacity; ++i) {

        freeList(t.buckets[i]);
    }

    delete[] t.buckets;

    freeArray(t.bloom);

    t.buckets = nullptr;

    t.capacity = t.size = 0;
}

void initCuckoo(CuckooTable& t, int capacity) {

    fillArray(t.keys, max(16, capacity));

    fillArray(t.used, t.keys.size);

    t.size = 0;

    t.seed = 1;
}

static int cuckooPosition(const CuckooTable& t, int key, int which) {

    int half = t.keys.size / 2;

    return static_cast<int>(hashKey(key, t.seed + which) % half) + which * half;
}

bool cuckooContains(const CuckooTable& t, int key) {

    for (int i = 0; i < 2; ++i) {

        int p = cuckooPosition(t, key, i);

        if (t.used.data[p] && t.keys.data[p] == key) {

            return true;
        }
    }

    return false;
}

static bool placeCuckoo(CuckooTable& t, int& key) {

    int which = 0;

    for (int step = 0; step < 128; ++step) {

        int p = cuckooPosition(t, key, which);

        if (!t.used.data[p]) {

            t.used.data[p] = 1;

            t.keys.data[p] = key;

            return true;
        }

        swap(key, t.keys.data[p]);

        which = 1 - which;
    }

    return false;
}

bool cuckooAdd(CuckooTable& t, int key) {

    if (cuckooContains(t, key)) {

        return false;
    }

    if (t.size < t.keys.size / 3 && placeCuckoo(t, key)) {

        ++t.size;

        return true;
    }

    DynamicArray all{};

    initArray(all);

    for (int i = 0; i < t.keys.size; ++i) {

        if (t.used.data[i]) {

            pushBack(all, t.keys.data[i]);
        }
    }

    pushBack(all, key);

    int capacity = t.keys.size;

    bool done = false;

    while (!done) {

        capacity = doubled(capacity);

        freeArray(t.keys);

        freeArray(t.used);

        fillArray(t.keys, capacity);

        fillArray(t.used, capacity);

        t.seed += 2;

        done = true;

        for (int i = 0; i < all.size; ++i) {

            int value = all.data[i];

            if (!placeCuckoo(t, value)) {

                done = false;

                break;
            }
        }
    }

    t.size = all.size;

    freeArray(all);

    return true;
}

bool cuckooDelete(CuckooTable& t, int key) {

    for (int i = 0; i < 2; ++i) {

        int p = cuckooPosition(t, key, i);

        if (t.used.data[p] && t.keys.data[p] == key) {

            t.used.data[p] = 0;

            --t.size;

            return true;
        }
    }

    return false;
}

void freeCuckoo(CuckooTable& t) {

    freeArray(t.keys);

    freeArray(t.used);

    t.size = 0;
}

void runBenchmark() {

    int n;

    if (!(cin >> n) || n < 1 || n > INT_MAX / 2) {

        throw runtime_error("Expected positive N <= INT_MAX/2");
    }

    ChainTable a{};

    CuckooTable b{};

    initChain(a);

    initCuckoo(b);

    using Clock = chrono::steady_clock;

    long long found = 0;

    for (int method = 0; method < 2; ++method) {

        auto start = Clock::now();

        for (int i = 0; i < n; ++i) {

            if (method == 0) {

                chainAdd(a, i);

            } else {

                cuckooAdd(b, i);
            }
        }

        auto inserted = Clock::now();

        for (int i = 0; i < n * 2; ++i) {

            found += method == 0 ? chainContains(a, i) : cuckooContains(b, i);
        }

        auto searched = Clock::now();

        for (int i = 0; i < n; ++i) {

            if (method == 0) {

                chainDelete(a, i);

            } else {

                cuckooDelete(b, i);
            }
        }

        auto deleted = Clock::now();

        cout

            << (method == 0 ? "CHAIN" : "CUCKOO") << " insert_us="

            << chrono::duration_cast<chrono::microseconds>(inserted - start).count()

            << " search_us="

            << chrono::duration_cast<chrono::microseconds>(searched - inserted).count()

            << " delete_us="

            << chrono::duration_cast<chrono::microseconds>(deleted - searched).count()

            << '\n';
    }

    cout << "Found: " << found << '\n';

    freeChain(a);

    freeCuckoo(b);
}
