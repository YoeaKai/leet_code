class Solution:
    def climbStairs(self, n: int) -> int:
        pre = sum = 1

        for _ in range(n-1):
            pre, sum = sum, sum + pre

        return sum
