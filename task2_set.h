#pragma once

#include "task6_1_hash_tables.h"

#include <string>

struct Set {

    ChainTable table;
};

void initSet(Set& s);

bool setAdd(Set& s, int value);

bool setDelete(Set& s, int value);

bool setContains(const Set& s, int value);

void freeSet(Set& s);

void runSetFile(const std::string& path, const std::string& query);
