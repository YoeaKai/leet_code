from typing import List


class Solution:
    def addToArrayForm(self, num: List[int], k: int) -> List[int]:
        for idx in range(len(num) - 1, -1, -1):
            k, num[idx] = divmod(num[idx] + k, 10)
        return [int(digit) for digit in str(k)] + num if k else num
