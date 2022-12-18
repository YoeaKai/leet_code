from typing import List

class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []

        for idx, temperature in enumerate(temperatures):
            while stack and temperatures[stack[-1]] < temperature:
                cur = stack.pop()
                res[cur] = idx - cur
            stack.append(idx)

        return res
