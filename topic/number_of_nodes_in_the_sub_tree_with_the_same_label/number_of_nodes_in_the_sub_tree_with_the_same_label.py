from typing import List
from collections import Counter, defaultdict


class Solution:
    def countSubTrees(self, n: int, edges: List[List[int]], labels: str) -> List[int]:
        def dfs(node: int, parent: int):
            count = Counter(labels[node])

            for child in graph.get(node, []):
                if child != parent:
                    count += dfs(child, node)

            res[node] = count[labels[node]]

            return count

        res = [0] * n

        graph = defaultdict(list)
        for a, b in edges:
            graph[a] += [b]
            graph[b] += [a]

        dfs(0, -1)

        return res
