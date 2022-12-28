from typing import List
import heapq


class Solution:
    def minStoneSum(self, piles: List[int], k: int) -> int:
        piles = [-pile for pile in piles]
        heapq.heapify(piles)

        for _ in range(k):
            heapq.heapreplace(piles, piles[0] // 2)

        return -sum(piles)
