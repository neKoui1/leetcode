#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 2;
        int last = nums[nums.size() - 1];
        while (left <= right) {
            int mid = left + ( (right - left) >> 1 );
            if (nums[mid] < last) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        cout << "left = " << left << " right = " << right << endl;
        return nums[left];
    }
};

int main () {
    Solution s;
    vector<int> nums = {3, 4, 5, 1, 2};
    vector<int> nums2 = {4, 5, 6, 7, 0, 1, 2};
    vector<int> nums3 = {11, 13, 15, 17};
    cout << s.findMin(nums) << endl;
    cout << s.findMin(nums2) << endl;
    cout << s.findMin(nums3) << endl;

    return 0;
}