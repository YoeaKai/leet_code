from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        res = [True]

        for i in range(1, len(s)+1):
            res += any(res[j] and s[j:i] in wordDict for j in range(i)),

        return res[-1]
