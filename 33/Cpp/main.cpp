#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        while (left <= right) {
            int mid = left + ( (right - left) >> 1);
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[0] <= nums[mid]) {
                if (target >= nums[0] && target < nums[mid]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                if (target > nums[mid] && target <= nums[nums.size() - 1]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }
        return -1;
    }
};

int main () {
    Solution s;
    vector<int> nums = {4,5,6,7,0,1,2};
    vector<int> nums2 = {4,5,6,7,0,1,2};
    vector<int> nums3 = {1};
    cout << s.search(nums, 0) << endl;
    cout << s.search(nums2, 3) << endl;
    cout << s.search(nums3, 0) << endl;
    vector<int> nums4 = {5, 1, 3};
    cout << s.search(nums4, 3) << endl;

    return 0;
}