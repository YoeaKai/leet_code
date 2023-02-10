from typing import List
from collections import deque


class Solution:
    def maxDistance(self, grid: List[List[int]]) -> int:
        distance = 0
        n = len(grid)

        queue = deque()
        for r in range(n):
            for c in range(n):
                if grid[r][c] == 1:
                    queue.append((r, c))

        if len(queue) == n * n:
            return -1

        row_variation = [0, -1, 0, 1]
        column_variation = [-1, 0, 1, 0]

        while queue:
            size = len(queue)
            distance += 1
            for _ in range(size):
                r, c = queue.popleft()
                for dir in range(4):
                    row = r + row_variation[dir]
                    column = c + column_variation[dir]
                    if 0 <= row < n and 0 <= column < n and grid[row][column] == 0:
                        grid[row][column] = 1
                        queue.append((row, column))

        return distance - 1
