#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 2;
        int last = nums[nums.size() - 1];
        while (left <= right) {
            int mid = left + ( (right - left) >> 1);
            if (nums[mid] > nums[right]) {
                left = mid + 1;
            } else if (nums[mid] < nums[right]) {
                right = mid - 1;
            } else {
                right--;
            }
        }
        return nums[left];
    }
};