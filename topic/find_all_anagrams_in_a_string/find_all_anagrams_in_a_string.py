from typing import List
from collections import Counter


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        res, counter, n, match = [], Counter(p), len(p), 0

        for i in range(len(s)):
            if s[i] in counter:
                if not counter[s[i]]:
                    match -= 1
                counter[s[i]] -= 1
                if not counter[s[i]]:
                    match += 1

            if i >= n and s[i-n] in counter:
                if not counter[s[i-n]]:
                    match -= 1
                counter[s[i-n]] += 1
                if not counter[s[i-n]]:
                    match += 1

            if match == len(counter):
                res.append(i-n+1)

        return res
