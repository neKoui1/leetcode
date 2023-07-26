class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        if n < 2:
            return s
        start, maxLen = 0, 1
        dp = [ [False]*n for _ in range (n)]
        for i in range(n):
            dp[i][i] = True
            if i < n-1 and s[i] == s[i+1]:
                dp[i][i+1] = True
                start = i
                maxLen = 2
        
        for l in range(3, n+1):
            for i in range(n):
                j = i+l-1
                if j >= n:
                    break
                if s[i] == s[j] and dp[i+1][j-1]:
                    dp[i][j] = True
                    start = i
                    maxLen = l
        return s[start: start+maxLen]
