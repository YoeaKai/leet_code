class Solution:
    def tribonacci(self, n: int) -> int:
        tmp_1, tmp_2, res = 1, 0, 0

        for _ in range(n):
            tmp_1, tmp_2, res = tmp_2, res, tmp_1 + tmp_2 + res

        return res
