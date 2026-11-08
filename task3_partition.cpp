#include "task3_partition.h"
#include "task2_set.h"
#include "lab1/DynamicArray.h"
#include "input.h"
#include <iostream>
#include <stdexcept>
using namespace std;
struct Partition {
    DynamicArray values;
    DynamicArray chosen;
    DynamicArray best;
    long long target;
    long long bestSum;
};
static void chooseSubset(Partition& p, int index, long long sum) {
    if (sum > p.target || p.bestSum == p.target) {
        return;
    }
    if (sum > p.bestSum) {
        p.bestSum = sum;
        for (int i = 0; i < p.values.size; ++i) {
            p.best.data[i] = p.chosen.data[i];
        }
    }
    if (index == p.values.size) {
        return;
    }
    p.chosen.data[index] = 1;
    chooseSubset(p, index + 1, sum + p.values.data[index]);
    p.chosen.data[index] = 0;
    chooseSubset(p, index + 1, sum);
}
static void printSubset(const Partition& p, int group) {
    cout << '{';
    bool first = true;
    for (int i = 0; i < p.values.size; ++i) {
        if (p.best.data[i] == group) {
            if (!first) {
                cout << ", ";
            }
            cout << p.values.data[i];
            first = false;
        }
    }
    cout << "}\n";
}
void runPartition() {
    int n = readCount();
    Set numbers{};
    Partition p{};
    initSet(numbers);
    initArray(p.values);
    initArray(p.chosen);
    initArray(p.best);
    try {
        long long total = 0;
        for (int i = 0; i < n; ++i) {
            int value = readInt();
            if (value <= 0) {
                throw runtime_error("Expected natural numbers > 0");
            }
            if (setAdd(numbers, value)) {
                pushBack(p.values, value);
                pushBack(p.chosen, 0);
                pushBack(p.best, 0);
                total += value;
            }
        }
        if (p.values.size < 2) {
            throw runtime_error("Need at least two distinct elements");
        }
        p.target = total / 2;
        chooseSubset(p, 0, 0);
        printSubset(p, 1);
        printSubset(p, 0);
        cout << total - 2 * p.bestSum << '\n';
    } catch (...) {
        freeSet(numbers);
        freeArray(p.values);
        freeArray(p.chosen);
        freeArray(p.best);
        throw;
    }
    freeSet(numbers);
    freeArray(p.values);
    freeArray(p.chosen);
    freeArray(p.best);
}
