#include "task4_subarray_sum.h"
#include "lab1/DynamicArray.h"
#include <iostream>
#include <stdexcept>
#include "input.h"
using namespace std;
void runSubarray() {
    int n = readCount();
    DynamicArray a{};
    initArray(a);
    try {
        for (int i = 0; i < n; ++i) {
            pushBack(a, readInt());
        }
        long long target;
        if (!(cin >> target)) {
            throw runtime_error("Expected target sum");
        }
        bool found = false;
        for (int left = 0; left < a.size; ++left) {
            long long sum = 0;
            for (int right = left; right < a.size; ++right) {
                sum += a.data[right];
                if (sum != target) {
                    continue;
                }
                found = true;
                cout << '[';
                for (int i = left; i <= right; ++i) {
                    if (i != left) {
                        cout << ", ";
                    }
                    cout << a.data[i];
                }
                cout << "]\n";
            }
        }
        if (!found) {
            cout << "Not found\n";
        }
    } catch (...) {
        freeArray(a);
        throw;
    }
    freeArray(a);
}
