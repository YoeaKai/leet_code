from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        exist = {}

        for num in nums:
            if num in exist:
                exist[num] += 1
            else:
                exist[num] = 1

        for num, count in exist.items():
            if count == 1:
                return num

        return 0
