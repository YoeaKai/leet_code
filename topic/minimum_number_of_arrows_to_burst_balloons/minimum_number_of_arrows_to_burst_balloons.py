from typing import List

class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        points.sort(key = lambda x: x[1])

        arrows , curr_end = 1, points[0][1]

        for start, end in points:
            if curr_end < start:
                curr_end = end
                arrows  += 1

        return arrows
