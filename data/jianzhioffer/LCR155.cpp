/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 15:49:14
 * @link    github.com/taseikyo
 */

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if (root == NULL) {
            return NULL;
        }
        auto stack = std::vector<Node*>{};
        auto dummy = new (Node);
        auto p = dummy;
        for (; stack.size() > 0 || root != NULL;) {
            if (root != NULL) {
                stack.push_back(root);
                root = root->left;
            } else {
                root = stack.back();
                stack.pop_back();
                p->right = root;
                root->left = p;
                p = root;
                root = root->right;
            }
        }
        p->right = dummy->right;
        dummy->right->left = p;

        return dummy->right;
    }
};
