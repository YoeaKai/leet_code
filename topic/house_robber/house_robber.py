from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        slip = 0

        for i in range(1, len(nums)):
            # If we steal nums[i], the maximum is nums[i] + slip.
            nums[i] += slip
            # If we don't steal nums[i], the maximum is between slip and steal it from the last run.
            slip = max(slip, nums[i-1])

        return max(slip, nums[-1])
