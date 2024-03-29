class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        d = dict()
        for i, num in enumerate(nums):
            if target - num in d:
                return [d[target-num], i]
            d[nums[i]] = i
        return []