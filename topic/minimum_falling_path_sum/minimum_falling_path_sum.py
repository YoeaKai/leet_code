from typing import List


class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        for i in range(1, len(matrix)):
            for j in range(len(matrix[i])):
                matrix[i][j] += min(matrix[i - 1][max(0, j - 1):j + 2])
        return min(matrix[-1])
