/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:14:54
 * @link    github.com/taseikyo
 */

class Codec {
public:
    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        if (root == nullptr) {
            return ""; // 空树返回空字符串
        }

        string res;
        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {
            TreeNode* node = q.front();
            q.pop();

            if (node == nullptr) {
                res += "#,"; // 空节点用 # 表示，逗号作为分隔符
            } else {
                res += to_string(node->val) + ",";
                q.push(node->left);
                q.push(node->right);
            }
        }

        // 去掉末尾多余的逗号（可选）
        if (!res.empty()) {
            res.pop_back();
        }
        return res;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        if (data.empty()) {
            return nullptr;
        }

        // 1. 按逗号分割字符串，提取所有节点值
        vector<string> vals;
        string cur;
        for (char ch : data) {
            if (ch == ',') {
                vals.push_back(cur);
                cur.clear();
            } else {
                cur.push_back(ch);
            }
        }
        if (!cur.empty()) {
            vals.push_back(cur); // 处理最后一个值
        }

        // 2. 根节点
        if (vals.empty() || vals[0] == "#") {
            return nullptr;
        }
        TreeNode* root = new TreeNode(stoi(vals[0]));

        // 3. 层序构建：用队列存储待建立子节点的父节点
        queue<TreeNode*> q;
        q.push(root);
        int i = 1; // 当前处理的 vals 索引

        while (!q.empty() && i < vals.size()) {
            TreeNode* node = q.front();
            q.pop();

            // 处理左子节点
            if (i < vals.size()) {
                if (vals[i] != "#") {
                    node->left = new TreeNode(stoi(vals[i]));
                    q.push(node->left);
                }
                i++;
            }

            // 处理右子节点
            if (i < vals.size()) {
                if (vals[i] != "#") {
                    node->right = new TreeNode(stoi(vals[i]));
                    q.push(node->right);
                }
                i++;
            }
        }

        return root;
    }
};
