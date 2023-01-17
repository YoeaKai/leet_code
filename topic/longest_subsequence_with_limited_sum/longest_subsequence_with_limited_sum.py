from typing import List
import bisect


class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        nums.sort()
        for i in range(1, len(nums)):
            nums[i] += nums[i - 1]

        res = []

        for query in queries:
            res.append(bisect.bisect_right(nums, query))

        return res
