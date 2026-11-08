#include "task1_asteroids.h"

#include "lab1/Stack.h"

#include <iostream>

#include <stdexcept>

#include "input.h"

using namespace std;

void runAsteroids() {

    int n = readCount();

    StackNode* top = nullptr;

    StackNode* reversed = nullptr;

    try {

        for (int i = 0; i < n; ++i) {

            int size = readInt();

            char direction;

            if (!(cin >> direction) || size <= 0 || (direction != 'L' && direction != 'R')) {

                throw runtime_error("Expected positive size and L or R");
            }

            int current = direction == 'R' ? size : -size;

            while (current < 0 && top && top->data > 0) {

                current += pop(top);

                if (current > 0) {

                    push(top, current);

                    current = 0;
                }
            }

            if (current) {

                push(top, current);
            }
        }

        while (top) {

            push(reversed, pop(top));
        }

        cout << '[';

        bool first = true;

        while (reversed) {

            int x = pop(reversed);

            if (!first) {

                cout << ", ";
            }

            first = false;

            cout << '(' << (x > 0 ? x : -x) << ", " << (x > 0 ? 'R' : 'L') << ')';
        }

        cout << "]\n";

    } catch (...) {

        freeStack(top);

        freeStack(reversed);

        throw;
    }
}
