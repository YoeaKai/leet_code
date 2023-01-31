from typing import List


class Solution:
    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        scores = [(ages[i], scores[i]) for i in range(len(scores))]
        scores = sorted(scores, key=lambda x: (x[0], x[1]))
        dp = [scores[i][1] for i in range(len(scores))]
        res = dp[0]

        for i in range(1, len(scores)):
            for j in range(i):
                cur_score = dp[j] + scores[i][1]
                if scores[j][1] <= scores[i][1] and cur_score > dp[i]:
                    dp[i] = cur_score
            res = max(dp[i], res)

        return res
