from typing import List
from functools import reduce


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        # If you XOR same numbers it will return zero.
        # Refer: https://thepythonguru.com/python-builtin-functions/reduce/
        return reduce(lambda total, el: total ^ el, nums)
