#include<iostream>
#include<vector>
#include<climits>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}
};
/**
 * @brief 
 * funciton1
 */
// class Solution {
// public:
//     vector<int> nums;
//     bool isValidBST(TreeNode* root) {
//         getVector(nums, root);
//         for (int i = 1 ; i < nums.size(); i++ ) {
//             if (nums[i] <= nums[i-1]) {
//                 return false;
//             }
//         }
//         return true;
//     }
//     void getVector(vector<int>& nums, TreeNode* node) {
//         if ( node == nullptr ) {
//             return;
//         }
//         getVector(nums, node->left);
//         nums.push_back(node->val);
//         getVector(nums, node->right);
//     }
// };


class Solution {
public:
    bool helper(TreeNode* node, long long lower, long long upper) {
        if ( node == nullptr ) {
            return true;
        }

        if ( (node -> val) <= lower || (node -> val) >= upper) {
            return false;
        }

        return helper(node->left, lower, node -> val) &&
            helper(node -> right, node -> val, upper);
    }
    bool isValidBST(TreeNode* root) {
        return helper(root, LLONG_MIN, LLONG_MAX);
    }
};