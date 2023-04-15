#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        int left = 0, right = nums.size() - 2;
        while (left <= right) {
            int mid = left + ( (right - left) >> 1 );
            if (nums[mid] > nums[mid+1]) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        return left;
    }

};

int main () {
    Solution s;
    vector<int> nums = {1, 2, 3, 1};
    cout << s.findPeakElement(nums) << endl;
    nums.erase(nums.begin(), nums.end());
    nums = {1, 2, 1, 3, 5, 6, 4};
    cout << s.findPeakElement(nums) << endl;

    return 0;
}