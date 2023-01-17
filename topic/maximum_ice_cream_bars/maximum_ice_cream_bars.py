from typing import List


class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        return sum(1 for cost in sorted(costs) if (coins := coins-cost) >= 0)
