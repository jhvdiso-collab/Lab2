#pragma once
#include "task6_1_hash_tables.h"
void bloomAdd(ChainTable& table, int key);
bool bloomMaybeContains(const ChainTable& table, int key);
void runHashTask();
