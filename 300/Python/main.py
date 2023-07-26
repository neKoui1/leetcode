# class Solution:
#     def lengthOfLIS(self, nums: List[int]) -> int:
#         n = len(nums)
#         dp = [ 1 for _ in range(n)]
#         for i in range(1, n):
#             for j in range(0, i):
#                 if nums[i] > nums[j]:
#                     dp[i] = max(dp[i], dp[j]+1)
        
#         return max(dp)

from bisect import *
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        dp = []
        n = len(nums)
        for _, v in enumerate(nums):
            index = bisect_left(dp, v)
            if index < len(dp):
                dp[index] = v
            else:
                dp.append(v)
        return len(dp)


