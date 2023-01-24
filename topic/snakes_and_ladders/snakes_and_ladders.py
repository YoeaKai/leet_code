from typing import List
from collections import deque


class Solution:
    def snakesAndLadders(self, board: List[List[int]]) -> int:
        goal = len(board) * len(board)
        seen = set()
        queue = deque()
        queue.append((1, 0))

        while queue:
            label, step = queue.popleft()

            row, column = divmod(label-1, len(board))
            if row % 2 == 1:
                column = len(board)-1-column
            row = len(board)-1-row

            if board[row][column] != -1:
                label = board[row][column]

            if label == goal:
                return step

            for move_step in range(1, 7):
                next = label + move_step
                if next <= goal and next not in seen:
                    seen.add(next)
                    queue.append((next, step+1))

        return -1
