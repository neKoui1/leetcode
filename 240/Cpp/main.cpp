#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        for ( const auto& row : matrix ) {
            auto it = lower_bound(row.begin(), row.end(), target);
            if (it != row.end() && *it == target ) {
                return true;
            }
        }
        return false;
    }
};

int main () {
    vector< vector<int> > nums = {
        {1, 4, 7, 11, 15},
        {2, 5, 8, 12, 19},
        {3, 6, 9, 16, 22},
        {10, 13, 14, 17, 24},
        {18, 21, 23, 26, 30},
    };
    Solution s;
    cout << s.searchMatrix(nums, 5) << endl;
}