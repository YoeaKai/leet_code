class Solution:
    def numTilings(self, n: int) -> int:
        # dp[n] = dp[n-1] + dp[n-2] + 2*(dp[n-3]+...+dp[0])
        #       = dp[n-1] + dp[n-2] + dp[n-3] + dp[n-3] + 2*(dp[n-4]+...+dp[0])
        #       = dp[n-1] + dp[n-3] + (dp[n-2]+dp[n-3]+2*(dp[n-4]+...+dp[0]))
        #       = dp[n-1] + dp[n-3] + dp[n-1]
        #       = 2*dp[n-1] + dp[n-3]

        dp = [1, 2, 5] + [0] * n

        for i in range(3, n):
            dp[i] = (dp[i - 1] * 2 + dp[i - 3]) % 1000000007

        return dp[n - 1]
