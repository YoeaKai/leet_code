from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        for i in range(n, 2 * n):
            # Store each y(i) into the respective x(i).
            nums[i - n] |= nums[i] << 10

        # 11111111 in binary.
        max_value = (2 << 9) - 1

        for i in range(n - 1, -1, -1):
            nums[2 * i + 1] = nums[i] >> 10
            nums[2 * i] = nums[i] & max_value

        return nums
