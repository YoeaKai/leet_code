class Solution:
    # DP solution after dimensionality reduction.
    # def longestCommonSubsequence(self, text1: str, text2: str) -> int:
    #     dp = [0] * len(text2)
    #     for _, char1 in enumerate(text1):
    #         pre = dp[0]
    #         dp[0] += 0 if char1 != text2[0] else 1
    #         for j, char2 in enumerate(text2[1:], start=1):
    #             tmp = dp[j]
    #             dp[j] = pre + 1 if char1 == char2 else max(dp[j-1], dp[j])
    #             pre = tmp
    #     return dp[-1]

    # DP solution.
    def longestCommonSubsequence2(self, text1: str, text2: str) -> int:
        dp = [[0] * (len(text2) + 1) for _ in range(len(text1) + 1)]
        for i, char1 in enumerate(text1):
            for j, char2 in enumerate(text2):
                dp[i + 1][j + 1] = 1 + \
                    dp[i][j] if char1 == char2 else max(
                        dp[i][j + 1], dp[i + 1][j])
        return dp[-1][-1]
