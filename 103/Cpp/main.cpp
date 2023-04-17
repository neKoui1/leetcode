#include<iostream>
#include<vector>
#include<queue>
#include<algorithm>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int> > res;
        if (root == nullptr) {
            return res;
        }

        int cnt = 0;
        queue<TreeNode*> q;
        q.push(root);

        while( !q.empty() ) {
            auto length = q.size();
            vector<int> temp;
            cnt++;

            for ( int i = 0; i < length; i++ ) {
                auto node = q.front();
                q.pop();
                temp.push_back(node->val);
                if (node -> left != nullptr) {
                    q.push(node -> left);
                }
                if (node -> right != nullptr) {
                    q.push(node -> right);
                }
            }

            if ((cnt&1) == 0) {
                reverse(temp.begin(), temp.end());
            }
            res.push_back(temp);
        }
        return res;
    }
};