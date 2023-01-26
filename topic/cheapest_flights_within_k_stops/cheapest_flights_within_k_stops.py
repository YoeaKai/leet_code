from typing import List
from collections import defaultdict
import heapq


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        graph = defaultdict(dict)
        for start, end, price in flights:
            graph[start].append((end, price))

        visited = {}
        heap = [(0, 0, src)]

        while heap:
            cost, steps, node = heapq.heappop(heap)

            if node == dst and k >= steps - 1:
                return cost

            if node not in visited or visited[node] > steps:
                visited[node] = steps
                for next, price in graph[node]:
                    heapq.heappush(heap, (cost + price, steps + 1, next))

        return -1
