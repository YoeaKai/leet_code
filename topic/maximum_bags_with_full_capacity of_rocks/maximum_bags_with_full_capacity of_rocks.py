from typing import List


class Solution:
    def maximumBags(self, capacity: List[int], rocks: List[int], additionalRocks: int) -> int:
        remain = [cap - rock for cap, rock in zip(capacity, rocks)]
        remain.sort()
        full_bags = 0

        for curr in remain:
            if additionalRocks < curr:
                break
            additionalRocks -= curr
            full_bags += 1

        return full_bags
