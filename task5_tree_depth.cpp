#include "task5_tree_depth.h"
#include "lab1/BST.h"
#include "input.h"
#include <iostream>
using namespace std;
static int addWithDepth(BSTNode*& root, int value) {
    if (root == nullptr) {
        root = insertBST(root, value);
        return 1;
    }
    BSTNode* current = root;
    int depth = 1;
    while (current->key != value) {
        if (value < current->key) {
            if (current->left == nullptr) {
                current->left = insertBST(current->left, value);
                return depth + 1;
            }
            current = current->left;
        } else {
            if (current->right == nullptr) {
                current->right = insertBST(current->right, value);
                return depth + 1;
            }
            current = current->right;
        }
        ++depth;
    }
    return 0;
}
void runTree() {
    BSTNode* root = nullptr;
    bool first = true;
    try {
        int value = readInt();
        while (value != 0) {
            int depth = addWithDepth(root, value);
            if (depth != 0) {
                if (!first) {
                    cout << ' ';
                }
                cout << depth;
                first = false;
            }
            value = readInt();
        }
        cout << '\n';
    } catch (...) {
        freeBST(root);
        throw;
    }
    freeBST(root);
}
