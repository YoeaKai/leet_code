from typing import List


class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        left, right = int(max(weights)), sum(weights)

        while left < right:
            mid, count, cur = (left + right) // 2, 1, 0

            for weight in weights:
                if cur + weight > mid:
                    count += 1
                    cur = 0
                cur += weight

            if count > days:
                left = mid + 1
            else:
                right = mid

        return left
