class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        n = len(nums)
        for i in range(n-1):
            for j in range(i+1, n):
                if nums[i] + nums[j] == target:
                    return list([i, j])
        return list([])

a = Solution()
print(a.twoSum([2, 7, 11, 15], 9))
print(a.twoSum([3, 2, 4], 6))
print(a.twoSum([3, 3], 6))