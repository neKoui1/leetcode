#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        for ( const auto& row : matrix ) {
            auto it = lower_bound(row.begin(), row.end(), target);
            if ( it != row.end() && *it == target ) {
                return true;
            }
        }
        return false;
    }

    // 右边界
    int BinarySearchRight(vector<int>& nums, int target) {
        int left = -1, right = nums.size();
        while (left + 1 < right) {
            int middle = left + ((right - left) >> 1);
            if(nums[middle] > target) {
                right = middle;
            } else {
                left = middle;
            }
        }
        return left;
    }

    // 左边界
    int BinarySearchLeft(vector<int>& nums, int target) {
        int left = -1, right = nums.size();
        while (left + 1 < right) {
            int middle = left + ((right - left) >> 1);
            if (nums[middle] < target) {
                left = middle;
            } else {
                right = middle;
            }
        }
        return right;
    }
};

int main () {
    vector<int> nums = {1, 2, 3, 3, 4, 5};
    Solution s;
    cout << s.BinarySearchLeft(nums, 3) << endl;
    cout << s.BinarySearchRight(nums, 3) << endl;
}