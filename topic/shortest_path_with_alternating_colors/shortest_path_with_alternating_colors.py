from typing import List
import sys


class Solution:
    def shortestAlternatingPaths(self, n: int, redEdges: List[List[int]], blueEdges: List[List[int]]) -> List[int]:
        graph = [[[], []] for _ in range(n)]

        for start, end in redEdges:
            graph[start][0].append(end)

        for start, end in blueEdges:
            graph[start][1].append(end)

        res = [[0, 0]] + [[sys.maxsize, sys.maxsize] for _ in range(n - 1)]
        queue = [[0, 0], [0, 1]]

        for idx, color in queue:
            for next in graph[idx][color]:
                if res[next][color] == sys.maxsize:
                    # 1 - color = another color
                    res[next][color] = res[idx][1 - color] + 1
                    queue.append([next, 1 - color])

        return [path if path < sys.maxsize else -1 for path in map(min, res)]
