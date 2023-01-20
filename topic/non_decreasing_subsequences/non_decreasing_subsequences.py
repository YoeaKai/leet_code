from typing import List


class Solution:
    def findSubsequences(self, nums: List[int]) -> List[List[int]]:
        res = {()}

        for num in nums:
            res |= {sub + (num,) for sub in res if not sub or sub[-1] <= num}

        return [sub for sub in res if len(sub) >= 2]
