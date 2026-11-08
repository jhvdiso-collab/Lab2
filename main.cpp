#include "task1_asteroids.h"

#include "task3_partition.h"

#include "task4_subarray_sum.h"

#include "task5_tree_depth.h"

#include "task2_set.h"

#include "task6_1_hash_tables.h"

#include "task6_2_bloom_filter.h"

#include "task7_lfu_cache.h"

#include <exception>

#include <iostream>

#include <stdexcept>

#include <string>

using namespace std;

int main(int argc, char* argv[]) {

    try {

        if (argc == 5) {

            string file, query;

            for (int i = 1; i < argc; i += 2) {

                string option = argv[i];

                if (option == "--file" && file.empty()) {

                    file = argv[i + 1];

                } else if (option == "--query" && query.empty()) {

                    query = argv[i + 1];

                } else {

                    throw runtime_error("Expected --file path --query request");
                }
            }

            if (file.empty() || query.empty()) {

                throw runtime_error("Missing file or query");
            }

            runSetFile(file, query);

        } else if (argc == 2) {

            string task = argv[1];

            if (task == "1") {

                runAsteroids();

            } else if (task == "3") {

                runPartition();

            } else if (task == "4") {

                runSubarray();

            } else if (task == "5") {

                runTree();

            } else if (task == "6") {

                runHashTask();

            } else if (task == "6.1") {

                runBenchmark();

            } else if (task == "7") {

                runCache();

            } else {

                throw runtime_error("Unknown task");
            }

        } else {

            cout << "Tasks: ./laba2 1|3|4|5|6|6.1|7\n"

                      << "Set: ./laba2 --file data.txt --query \"SETADD 10\"\n";
        }

        return 0;

    } catch (const exception& e) {

        cerr << "Error: " << e.what() << '\n';

        return 1;
    }
}
