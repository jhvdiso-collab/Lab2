#pragma once
#include <iostream>
#include <stdexcept>
inline int readInt() {
    int value;
    if (!(std::cin >> value)) {
        throw std::runtime_error("Expected integer");
    }
    return value;
}
inline int readCount() {
    int count = readInt();
    if (count < 0) {
        throw std::runtime_error("Negative element count");
    }
    return count;
}
