#include "BST.h"

#include <iostream>

using namespace std;

BSTNode* insertBST(BSTNode* p, int k) {

    if (!p) {

        return new BSTNode{k, nullptr, nullptr};
    }

    if (k < p->key) {

        p->left = insertBST(p->left, k);

    } else if (k > p->key) {

        p->right = insertBST(p->right, k);
    }

    return p;
}

bool searchBST(BSTNode* p, int k) {

    while (p != nullptr) {

        if (p->key == k) {

            return true;
        }

        if (k < p->key) {

            p = p->left;

        } else {

            p = p->right;
        }
    }

    return false;
}

void printBSTInOrder(BSTNode* root) {

    if (root == nullptr) {

        return;
    }

    printBSTInOrder(root->left);

    cout << root->key << " ";

    printBSTInOrder(root->right);
}

void freeBST(BSTNode* root) {

    while (root) {

        if (root->left) {

            BSTNode* left = root->left;

            root->left = left->right;

            left->right = root;

            root = left;

        } else {

            BSTNode* next = root->right;

            delete root;

            root = next;
        }
    }
}
