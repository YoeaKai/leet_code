from typing import List


class Solution:
    # Method 1: Bit Manipulation.
    def singleNumber(self, nums: List[int]) -> int:
        single = 0

        for i in range(32):
            count = 0
            for num in nums:
                if num & (1 << i) == (1 << i):
                    count += 1
            single |= (count % 3) << i

        return single if single < (1 << 31) else single - (1 << 32)

    # Method 2
    def singleNumber_2(self, nums: List[int]) -> int:
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
