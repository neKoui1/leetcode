class Solution(object):
    """
    len(nums) mod 2  == 0
    if i mod 2 == 0 :
        nums[i] != nums[i+1]
    """
    def minDeletion(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        cnt, check = 0, True
        for i in range(n-1):
            if nums[i] == nums[i+1] and check:
                cnt += 1
            else:
                check = not check
        return cnt if (n-cnt)%2 == 0 else cnt+1

a = Solution()
print(a.minDeletion([1, 1, 2, 3, 5]))
print(a.minDeletion([1, 1, 2, 2, 3, 3]))