from typing import List


class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        if not graph:
            return []

        res = []
        stack = [(0, [0])]

        while stack:
            node, path = stack.pop()

            if node == len(graph) - 1:
                res.append(path)

            for next_node in graph[node]:
                stack.append((next_node, path+[next_node]))

        return res
