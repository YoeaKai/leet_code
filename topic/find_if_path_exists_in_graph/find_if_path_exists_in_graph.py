from typing import List
from collections import defaultdict


class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        graph = defaultdict(list)

        for v_1, v_2 in edges:
            graph[v_1].append(v_2)
            graph[v_2].append(v_1)

        seen = [False] * n

        def dfs(curr_node):
            if curr_node == destination:
                return True
            if not seen[curr_node]:
                seen[curr_node] = True
                for next_node in graph[curr_node]:
                    if dfs(next_node):
                        return True
            return False

        return dfs(source)
