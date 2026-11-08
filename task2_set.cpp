#include "task2_set.h"
#include <fstream>
#include <iostream>
#include <sstream>
#include <stdexcept>
using namespace std;
void initSet(Set& s) {
    initChain(s.table);
}
bool setAdd(Set& s, int v) {
    return chainAdd(s.table, v);
}
bool setDelete(Set& s, int v) {
    return chainDelete(s.table, v);
}
bool setContains(const Set& s, int v) {
    return chainContains(s.table, v);
}
void freeSet(Set& s) {
    freeChain(s.table);
}
void runSetFile(const string& path, const string& query) {
    istringstream command(query);
    string op, extra;
    int value;
    if (!(command >> op >> value) || (command >> extra)) {
        throw runtime_error("Query: SETADD|SETDEL|SET_AT integer");
    }
    if (op != "SETADD" && op != "SETDEL" && op != "SET_AT") {
        throw runtime_error("Unknown set operation");
    }
    Set s{};
    initSet(s);
    try {
        ifstream input(path);
        if (!input) {
            throw runtime_error("Cannot open data file; create it first");
        }
        int x;
        while (input >> ws && !input.eof()) {
            if (!(input >> x)) {
                throw runtime_error("Data file must contain integers");
            }
            setAdd(s, x);
        }
        input.close();
        if (op == "SET_AT") {
            cout << setContains(s, value) << '\n';
        } else {
            if (op == "SETADD") {
                setAdd(s, value);
            } else {
                setDelete(s, value);
            }
            ofstream output(path);
            if (!output) {
                throw runtime_error("Cannot write data file");
            }
            for (int i = 0; i < s.table.capacity; ++i) {
                for (SinglyNode* p = s.table.buckets[i]; p; p = p->next) {
                    output << p->data << '\n';
                }
            }
            output.close();
            if (!output) {
                throw runtime_error("Failed to save data file");
            }
        }
    } catch (...) {
        freeSet(s);
        throw;
    }
    freeSet(s);
}
