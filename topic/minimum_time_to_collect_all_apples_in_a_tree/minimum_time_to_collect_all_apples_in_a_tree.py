from typing import List


class Solution:
    def minTime(self, n: int, edges: List[List[int]], hasApple: List[bool]) -> int:
        graph = [[] for _ in range(n)]
        for v1, v2 in edges:
            graph[v1].append(v2)
            graph[v2].append(v1)

        visited = set()

        def dfs(node):
            if node in visited:
                return 0

            visited.add(node)

            times = 0
            for child in graph[node]:
                times += dfs(child)
            if times > 0:
                return times + 2

            return 2 if hasApple[node] else 0

        return max(dfs(0) - 2, 0)
