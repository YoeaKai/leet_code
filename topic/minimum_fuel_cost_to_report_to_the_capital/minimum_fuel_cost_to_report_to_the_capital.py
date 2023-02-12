from typing import List
from collections import defaultdict
from math import ceil


class Solution:
    def minimumFuelCost(self, roads: List[List[int]], seats: int) -> int:
        graph = defaultdict(list)
        for c1, c2 in roads:
            graph[c1].append(c2)
            graph[c2].append(c1)

        self.res = 0

        def dfs(cur, prev, people=1):
            for next in graph[cur]:
                if next == prev:
                    continue
                people += dfs(next, cur)
            self.res += (int(ceil(people / seats)) if cur else 0)
            return people

        dfs(0, 0)

        return self.res
