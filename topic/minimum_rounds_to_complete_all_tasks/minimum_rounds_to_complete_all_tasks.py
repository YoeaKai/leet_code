from typing import List
from collections import Counter


class Solution:
    def minimumRounds(self, tasks: List[int]) -> int:
        frequency = Counter(tasks).values()
        return -1 if 1 in frequency else sum((f + 2) // 3 for f in frequency)
